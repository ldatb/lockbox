package log

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// SanitizeHook is a custom Logrus hook that sanitizes log messages before they are logged.
// This ensures that sensitive information such as passwords, API keys, and tokens are not logged in plaintext.
type SanitizeHook struct{}

// Levels defines the log levels to which the hook is applied.
// This hook will be applied to all log levels, from debug to fatal.
func (hook *SanitizeHook) Levels() []logrus.Level {
	// Apply this hook to all log levels (Debug, Info, Warn, Error, Fatal, Panic).
	return logrus.AllLevels
}

// Fire is the method that is triggered before the log entry is written.
// It sanitizes the log message by escaping special characters and masking sensitive data.
func (hook *SanitizeHook) Fire(entry *logrus.Entry) error {
	// Sanitize log message by removing or escaping special characters (e.g., newlines, tabs).
	// This helps prevent logs from being manipulated or misformatted by malicious or unexpected input.
	sanitizedMessage := entry.Message
	sanitizedMessage = strings.ReplaceAll(sanitizedMessage, "\n", "\\n") // Escape newlines
	sanitizedMessage = strings.ReplaceAll(sanitizedMessage, "\r", "\\r") // Escape carriage returns
	sanitizedMessage = strings.ReplaceAll(sanitizedMessage, "\t", "\\t") // Escape tabs

	// Define regex patterns to detect sensitive information (e.g., passwords, API keys, tokens).
	// These patterns are case-insensitive to capture common sensitive data formats.
	sensitivePatterns := []string{
		`(?i)(password\s*=\s*)([^\s]+)`,    // Matches patterns like "password=somepassword"
		`(?i)(token\s*=\s*)([^\s]+)`,       // Matches patterns like "token=xyz"
		`(?i)(api[-_]?key\s*=\s*)([^\s]+)`, // Matches patterns like "apiKey=abc123" or "api_key=abc123"
		`(?i)(master\s*=\s*)([^\s]+)`,      // Matches patterns like "master=masterkey"
		`(?i)(key\s*=\s*)([^\s]+)`,         // Matches patterns like "key=somekey"
		`(?i)(secret\s*=\s*)([^\s]+)`,      // Matches patterns like "secret=mysecret"
	}

	// Iterate through the list of sensitive patterns and replace any matches with "*****".
	for _, pattern := range sensitivePatterns {
		// Compile the regular expression
		re := regexp.MustCompile(pattern)
		// Replace the sensitive value (second group) with "*****"
		sanitizedMessage = re.ReplaceAllString(sanitizedMessage, "$1*****")
	}

	// Update the log entry's message with the sanitized version.
	entry.Message = sanitizedMessage
	return nil
}
