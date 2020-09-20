package repository

import (
	"Food/config"
	"Food/models"
)

func SaveComment(commnent models.Comment) models.Comment {
	config.GetDB().Create(&commnent)
	return commnent
}
