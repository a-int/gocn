package models

import "gorm.io/gorm"

// User represents a row in the "users" table
type User struct {
	// gorm.Model includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name string
	Age  uint
}
