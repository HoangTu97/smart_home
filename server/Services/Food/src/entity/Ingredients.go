package entity

import (
	"github.com/jinzhu/gorm"
)

// Ingredients entity
type Ingredients struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	// Type     string `gorm:"type:""`
	Quantity uint32

	FoodID uint
}
