package config

import (
	"fmt"
	"os"
)

// Config holds the application's configuration values.
type Config struct {
	DatabaseURL string
}

// LoadConfig loads configuration from environment variables.
// It returns a Config struct or an error if a required variable is missing.
func LoadConfig() (*Config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Using a default for now, but in a real app, you might return an error
		// or have a more sophisticated default/validation.
		databaseURL = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		fmt.Println("DATABASE_URL environment variable not set, using default connection string.")
	}

	return &Config{
		DatabaseURL: databaseURL,
	}, nil
}
