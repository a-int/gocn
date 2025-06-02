package user

import (
	"fmt"
	"log"

	"github.com/a-int/gocn/internal/models"
	"gorm.io/gorm"
)

// FindAll fetches all users from the database.
// It takes a *gorm.DB instance and returns a slice of User or an error.
func FindAll(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		log.Printf("Error fetching users: %v", result.Error)
		return nil, result.Error
	}

	return users, nil
}

// PrintUsers prints the details of a slice of users.
func PrintUsers(users []models.User) {
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
