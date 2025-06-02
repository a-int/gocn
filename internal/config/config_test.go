package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name        string
		dbURLEnv    string
		wantDBURL   string
		expectError bool
	}{
		{
			name:      "DATABASE_URL is set",
			dbURLEnv:  "postgres://testuser:testpass@testhost:5432/testdb?sslmode=require",
			wantDBURL: "postgres://testuser:testpass@testhost:5432/testdb?sslmode=require",
		},
		{
			name:      "DATABASE_URL is not set (uses default)",
			dbURLEnv:  "", // Empty string to simulate unset variable
			wantDBURL: "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get current DATABASE_URL to restore it later
			originalDBURL := os.Getenv("DATABASE_URL")
			defer os.Setenv("DATABASE_URL", originalDBURL) // Restore original value

			if tt.dbURLEnv != "" {
				os.Setenv("DATABASE_URL", tt.dbURLEnv)
			} else {
				os.Unsetenv("DATABASE_URL") // Explicitly unset for the test case
			}

			cfg, err := LoadConfig()

			if (err != nil) != tt.expectError {
				t.Errorf("LoadConfig() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && cfg.DatabaseURL != tt.wantDBURL {
				t.Errorf("LoadConfig() DatabaseURL = %v, want %v", cfg.DatabaseURL, tt.wantDBURL)
			}
		})
	}
}
