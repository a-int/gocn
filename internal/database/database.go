package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a database connection.
// It takes the databaseURL string and returns a *gorm.DB instance or logs a fatal error if the connection fails.
func Connect(databaseURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")
	return db
}

// Migrate runs auto-migrations for the provided models.
// It takes a *gorm.DB instance and a variadic list of model interfaces.
// It logs a fatal error if migration fails.
func Migrate(db *gorm.DB, models ...interface{}) {
	log.Println("Running database auto-migrations...")
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	log.Println("Database auto-migration completed successfully.")
}
