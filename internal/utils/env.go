package utils

import "os"

// GetEnvOrFallback tries to get a value from the env vars,
// if the value does not exist, return a fallback.
func GetEnvOrFallback(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
