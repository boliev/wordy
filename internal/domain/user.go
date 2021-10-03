package domain

import "gorm.io/gorm"

// User domain model
type User struct {
	gorm.Model
	Email    string
	Password string
}
