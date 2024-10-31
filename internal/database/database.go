package database

import (
	"fmt"
	"time"

	"gitlab.com/xrs-cloud/lockbox/core/internal/config"
	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
	"gitlab.com/xrs-cloud/lockbox/core/internal/secrets"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase initializes the database connection using the provided configuration.
// It connects to the PostgreSQL database, configures connection pool settings, and performs schema migrations.
func InitDatabase(dbConfig config.DatabaseConfig) *gorm.DB {
	// Construct the DSN (Data Source Name) using the provided database configuration
	// This string contains the necessary information to connect to the database.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DatabaseName,
		dbConfig.Port,
		dbConfig.SSLMode,
	)
	global.Logger.Debugf("Connecting to database with %s", dsn)

	// Open a connection to the PostgreSQL database using GORM.
	// The connection uses the provided DSN, and GORM is configured with silent logging mode to suppress unnecessary logs.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		// If the connection fails, log the error and exit the application.
		global.Logger.Fatalf("Failed to connect to database: %v", err)
	}

	// Retrieve the underlying *sql.DB object to configure low-level database connection settings.
	sqlDB, err := db.DB()
	if err != nil {
		// If retrieval of the underlying *sql.DB object fails, log the error and exit.
		global.Logger.Fatalf("Failed to get database object: %v", err)
	}

	// Configure the database connection pool settings:
	// - SetMaxIdleConns: Limits the number of idle connections in the pool.
	// - SetMaxOpenConns: Limits the maximum number of open connections to the database.
	// - SetConnMaxLifetime: Specifies the maximum amount of time (in seconds) a connection may be reused.
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLife) * time.Second)

	// Automatically migrate the database schema based on the Secret model.
	// This ensures that the schema in the database stays up-to-date with the application's data models.
	db.AutoMigrate(&secrets.Secret{})

	// Return the initialized *gorm.DB object for use in the application.
	return db
}
