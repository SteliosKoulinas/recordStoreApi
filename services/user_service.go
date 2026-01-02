package services

import (
	"fmt"

	"github.com/SteliosKoulinas/recordStoreApi/db"
	"github.com/SteliosKoulinas/recordStoreApi/models"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService) GetAll() ([]models.Users, error) {
	var users []models.Users
	result := db.DB.Find(&users)
	return users, result.Error
}

func (s *UserService) Create(user *models.Users) error {
	// Check if user already exists
	var existing models.Users
	result := db.DB.
		Where("email = ?", user.Email).
		First(&existing)

	if result.Error == nil {
		// record exists
		return fmt.Errorf("user already exists")
	}

	if result.Error != gorm.ErrRecordNotFound {
		// some other db error
		return result.Error
	}
	return db.DB.Create(user).Error
}
