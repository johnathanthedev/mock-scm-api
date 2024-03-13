package users_service

import (
	"errors"

	models "scm-api/api/models"
	"scm-api/db"

	"github.com/google/uuid"
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

func GetUserIdByUsername(username string) (uuid.UUID, error) {
	var user models.User
	result := db.GetDB().Where("username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return uuid.Nil, errors.New("user not found")
		}

		return uuid.Nil, result.Error
	}

	return user.ID, nil
}
