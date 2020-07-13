package repository

import (
	"Food/config"
	"Food/models"
)

func SaveUser(user models.User) models.User {
	config.GetDB().Create(&user)
	return user
}

func FindOneUserByName(name string) (models.User, error) {
	user := models.User{}
	result := config.GetDB().Where("Name = ?", name).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}