package main

import (
	"log"

	"github.com/a-int/gocn/internal/config"
	"github.com/a-int/gocn/internal/database"
	"github.com/a-int/gocn/internal/models"
	"github.com/a-int/gocn/internal/user"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	db := database.Connect(cfg.DatabaseURL)

	// Auto-migrate the schema (create the users table if it doesn't exist)
	// Note: In production, you might use dedicated migration tools
	database.Migrate(db, &models.User{})

	// Find all users
	users, err := user.FindAll(db)
	if err != nil {
		// user.FindAll already logs the error, so we can just return here
		return
	}

	// Print users
	user.PrintUsers(users)
}
