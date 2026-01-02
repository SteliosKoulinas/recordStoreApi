package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email"    gorm:"unique"`
	Password string `json:"-" gorm:"not null"` // Exclude password from JSON responses
}
