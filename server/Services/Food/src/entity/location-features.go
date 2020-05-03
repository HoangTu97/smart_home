package entity

import (
	"github.com/jinzhu/gorm"
)

// LocationFeatures entity
type LocationFeatures struct {
	gorm.Model
	Location Location
}
