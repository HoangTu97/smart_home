package entity

import (
	"github.com/jinzhu/gorm"
)

// UserFeatures entity
type UserFeatures struct {
	gorm.Model
	User       User
}
