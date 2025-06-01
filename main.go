package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User represents a row in the "users" table
type User struct {
	// gorm.Model includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name string
	Age  uint
}

func main() {
	// Database connection string
	// This should ideally come from environment variables in a real app
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		log.Println("DATABASE_URL environment variable not set, using default connection string:", dsn)
	}

	fmt.Println(dsn)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")

	// Auto-migrate the schema (create the users table if it doesn't exist)
	// Note: In production, you might use dedicated migration tools
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate user table: %v", err)
	}

	log.Println("User table checked/migrated successfully")

	// Find all users
	var users []User
	result := db.Find(&users)

	if result.Error != nil {
		log.Printf("Error fetching users: %v", result.Error)
		return
	}

	// Check if the table is empty
	if len(users) == 0 {
		fmt.Println("The 'users' table is currently empty.")
	} else {
		fmt.Printf("Found %d user(s) in the 'users' table:\n", len(users))
		for _, user := range users {
			fmt.Printf(" - ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
		}
	}
}
