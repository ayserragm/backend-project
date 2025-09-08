package repository

import (
	"github.com/ayserragm/backend-project/internal/db"
	"github.com/ayserragm/backend-project/internal/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
