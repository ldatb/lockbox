package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Global contains global variables that are accessible throughout the entire application.
// These variables provide centralized access to shared resources such as logging and database connections.
var (
	// Logger is a globally accessible instance of logrus.Logger.
	// It is used to log messages and errors throughout the application, providing a centralized logging mechanism.
	// Logrus allows for structured and leveled logging, making it easier to monitor application behavior.
	Logger *logrus.Logger

	// Database is a globally accessible instance of gorm.DB.
	// It represents the connection to the database and is used throughout the application to perform database operations.
	// The connection is initialized once during application startup and is reused by different components.
	Database *gorm.DB
)
