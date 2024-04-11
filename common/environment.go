package common

import "os"

func GetEnv(key string, defaultValue string) string {
	// Get the value of the environment variable key
	// If the environment variable key does not exist, return the defaultValue
	// If the environment variable key exists, return the value of the environment variable key
	// Example: GetEnv("PORT", "8080")
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
