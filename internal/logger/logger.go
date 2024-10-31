package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/xrs-cloud/lockbox/core/internal/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogFormatter is a custom log formatter for logrus.
// It inherits from logrus.TextFormatter and overrides the Format method to define a specific log format.
type LogFormatter struct {
	logrus.TextFormatter
}

// Format formats the log entry in a standard format: timestamp - LEVEL - message.
// It ensures all log messages follow a consistent structure, with timestamps in ISO8601 format
// and log levels in uppercase.
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Create the formatted log entry: timestamp - LEVEL - msg
	log := fmt.Sprintf(
		"%s %s - %s\n",
		entry.Time.Format(time.RFC3339),       // ISO8601 timestamp format
		strings.ToUpper(entry.Level.String()), // Convert log level to uppercase
		entry.Message,                         // Log message
	)
	return []byte(log), nil
}

// InitLogger initializes and configures the logger with sanitization and log rotation.
// It uses logrus for logging and lumberjack for log rotation. The logger is configured
// according to the provided logging configuration (level, file path, etc.).
func InitLogger(loggingConfig config.LoggingConfig) *logrus.Logger {
	// Create a new instance of logrus.Logger
	logger := logrus.New()

	// Set log level based on the provided configuration.
	// Valid log levels are: "debug", "info", "warn", "error", "fatal", "panic".
	level, err := logrus.ParseLevel(strings.ToLower(loggingConfig.Level))
	if err != nil {
		// If the log level is invalid, log the error and exit the application.
		log.Fatalf("Invalid log level: %v", err)
	}
	logger.SetLevel(level)

	// Configure log rotation using lumberjack.
	// This rotates log files when they reach a certain size, and manages log retention by
	// limiting the number of backups and the age of log files.
	logRotation := &lumberjack.Logger{
		Filename:   loggingConfig.FilePath, // Path to the log file
		MaxSize:    10,                     // Maximum size of each log file in MB before rotation
		MaxBackups: 3,                      // Maximum number of old log files to keep
		MaxAge:     28,                     // Maximum number of days to retain old log files
		Compress:   true,                   // Compress old log files to save space
	}

	// Configure the logger to write log output to both stdout and the rotated log file.
	multiWriter := io.MultiWriter(os.Stdout, logRotation)
	logger.SetOutput(multiWriter)

	// Apply the custom log formatter to ensure logs follow a consistent format.
	logger.SetFormatter(&LogFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp: true, // Ensure logs include full timestamps
		},
	})

	// Register a custom sanitize hook to sanitize sensitive information in logs.
	// This ensures sensitive data (e.g., passwords) is removed from logs before being output.
	logger.AddHook(&SanitizeHook{})

	// Return the initialized logger.
	return logger
}
