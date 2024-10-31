package config

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"strconv"

	"github.com/alyu/configparser"
)

// Config contains all configurations for the application.
// It aggregates server, security, database, and logging settings.
type Config struct {
	// Server holds configurations related to the server's network settings, such as host and port.
	Server ServerConfig

	// Security holds configurations related to API key management and cryptographic operations.
	Security SecurityConfig

	// Database holds configurations related to database connections, credentials, and pooling.
	Database DatabaseConfig

	// Logging holds configurations related to application logging, including log level and file location.
	Logging LoggingConfig
}

// ServerConfig contains server-related configurations.
// This struct holds information for binding the server to a specific host and port.
type ServerConfig struct {
	// Host defines the server's host address (e.g., "0.0.0.0" for all interfaces).
	Host string

	// Port defines the port on which the server listens (e.g., "8080").
	Port string
}

// SecurityConfig contains security-related configurations.
// This struct manages cryptographic settings, API key parameters, and secure passphrases.
type SecurityConfig struct {
	// APIKeyLength defines the length of the generated API keys (e.g., 32 for a 32-character key).
	APIKeyLength int

	// APIKeyValidity defines the duration (in seconds) for which the API key remains valid.
	APIKeyValidity int
}

// DatabaseConfig contains database-related configurations.
// This struct manages database credentials, connection settings, and connection pool configurations.
type DatabaseConfig struct {
	// Host defines the database server's host address (e.g., "localhost").
	Host string

	// Port defines the port on which the database is running (e.g., "5432" for PostgreSQL).
	Port string

	// Username holds the username for authenticating to the database.
	Username string

	// Password holds the password for authenticating to the database.
	Password string

	// DatabaseName holds the name of the database to connect to.
	DatabaseName string

	// SSLMode defines the SSL mode to use when connecting to the database (e.g., "disable", "require").
	SSLMode string

	// MaxIdleConns defines the maximum number of idle connections in the connection pool.
	MaxIdleConns int

	// MaxOpenConns defines the maximum number of open connections in the connection pool.
	MaxOpenConns int

	// MaxConnLife defines the maximum lifetime (in seconds) of a connection before it is reused.
	MaxConnLife int
}

// LoggingConfig contains logging-related configurations.
// This struct manages the application's log level, file paths, and log rotation settings.
type LoggingConfig struct {
	// Level defines the log level (e.g., "info", "debug", "error").
	Level string

	// FilePath specifies the file path where log files will be written.
	FilePath string

	// MaxLogLength defines the maximum size (in characters) before log files are rotated.
	MaxLogLength int
}

// LoadConfig loads the configuration from a .conf file and environment variables.
// The MASTER_CRYPTO_PASS environment variable is required for production security, and if it is not set, a random key is generated with a warning.
func LoadConfig(filePath string) (*Config, error) {
	// Read the .conf file using configparser
	configFile, err := configparser.Read(filePath)
	if err != nil {
		return nil, err
	}

	// Load the server configuration section
	serverSection, err := configFile.Section("server")
	if err != nil {
		log.Fatalf("Error accessing 'server' section: %v", err)
	}

	// Load the security configuration section
	securitySection, err := configFile.Section("security")
	if err != nil {
		log.Fatalf("Error accessing 'security' section: %v", err)
	}

	// Load the master cryptographic passphrase from the environment
	// If it is not defined, generate a random one and save it
	envFieldName := "MASTER_CRYPTO_PASS"
	masterCryptoPass := os.Getenv(envFieldName)
	if masterCryptoPass == "" {
		// If the environment variable is not set, generate a random 64-character hexadecimal key (32 bytes) and print a warning
		log.Println("WARNING: MASTER_CRYPTO_PASS environment variable is not set. A random key has been generated.")
		log.Println("WARNING: Using a randomly generated key is NOT suitable for production environments. Please configure a consistent key via the MASTER_CRYPTO_PASS environment variable to ensure data can be encrypted and decrypted consistently across sessions.")
		masterCryptoPass = generateRandomKey(32)
		os.Setenv(envFieldName, masterCryptoPass)
	}

	// Load the database configuration section
	databaseSection, err := configFile.Section("database")
	if err != nil {
		log.Fatalf("Error accessing 'database' section: %v", err)
	}

	// Load the logging configuration section
	loggingSection, err := configFile.Section("logging")
	if err != nil {
		log.Fatalf("Error accessing 'logging' section: %v", err)
	}

	// Fill in the configuration values using defaults where applicable
	config := &Config{
		Server: ServerConfig{
			Host: getValueOrDefault(serverSection, "host", "0.0.0.0"),
			Port: getValueOrDefault(serverSection, "port", "8080"),
		},
		Security: SecurityConfig{
			APIKeyLength:   getValueOrDefaultAsInt(securitySection, "api_key_length", 32),
			APIKeyValidity: getValueOrDefaultAsInt(securitySection, "api_key_validity", 60), // 1 minute
		},
		Database: DatabaseConfig{
			Host:         getValueOrDefault(databaseSection, "host", "localhost"),
			Port:         getValueOrDefault(databaseSection, "port", "5432"),
			Username:     getValueOrDefault(databaseSection, "username", "lockboxuser"),
			Password:     getValueOrDefault(databaseSection, "password", "password"),
			DatabaseName: getValueOrDefault(databaseSection, "db_name", "lockboxdb"),
			SSLMode:      getValueOrDefault(databaseSection, "ssl_mode", "disable"),
			MaxIdleConns: getValueOrDefaultAsInt(databaseSection, "max_idle_conns", 10),
			MaxOpenConns: getValueOrDefaultAsInt(databaseSection, "max_open_conns", 100),
			MaxConnLife:  getValueOrDefaultAsInt(databaseSection, "max_conn_life", 60),
		},
		Logging: LoggingConfig{
			Level:        getValueOrDefault(loggingSection, "level", "info"),
			FilePath:     getValueOrDefault(loggingSection, "filepath", "lockbox.log"),
			MaxLogLength: getValueOrDefaultAsInt(loggingSection, "max_log_length", 1000),
		},
	}

	return config, nil
}

// getValueOrDefault retrieves a string value from the configuration section, or returns a default value if not found.
// This ensures the application has reasonable defaults even if some configuration parameters are missing.
func getValueOrDefault(section *configparser.Section, key, defaultValue string) string {
	if section == nil {
		return defaultValue
	}

	value := section.ValueOf(key)
	if value == "" {
		return defaultValue
	}

	return value
}

// getValueOrDefaultAsInt retrieves an integer value from the configuration section, or returns a default value if not found.
// This converts the string value from the config file into an integer, with error handling for invalid formats.
func getValueOrDefaultAsInt(section *configparser.Section, key string, defaultValue int) int {
	value := getValueOrDefault(section, key, strconv.Itoa(defaultValue))

	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return valueAsInt
}

// generateRandomKey generates a secure random key of the specified length (in bytes) and returns it as a hexadecimal string.
// This is used when the MASTER_CRYPTO_PASS environment variable is not set, generating a 64-character random key (32 bytes).
func generateRandomKey(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate random key: %v", err)
	}
	return hex.EncodeToString(bytes)
}
