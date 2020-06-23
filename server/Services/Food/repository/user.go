package repository

import (
	"Food/entity"
)

func SaveUser(user entity.User) entity.User {
	GetDB().Create(&user)
	return user
}

func FindOneUserByName(name string) (entity.User, error) {
	user := entity.User{}
	result := GetDB().Where("Name = ?", name).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}