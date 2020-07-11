package models

import (
	"github.com/jinzhu/gorm"
)

// Interactions entity
type Interactions struct {
	gorm.Model
	UserID     uint
	LocationID uint
	Value      int32
}
