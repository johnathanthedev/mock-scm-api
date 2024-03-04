package users_service

import (
	"errors"

	"scm-api/db"
	models "scm-api/internal/api/models"

	"gorm.io/gorm"
)

func UserExists(username string) bool {
	var existingUser models.User
	result := db.GetDB().Where("username = ?", username).First(&existingUser)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateUser(username string) (*models.User, error) {
	newUser := &models.User{
		Username: username,
	}

	if err := db.GetDB().Create(newUser).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}
