package repository

import (
	"Food/config"
	"Food/models"
)

func SavePost(post models.Post) models.Post {
	config.GetDB().Create(&post)
	return post
}
