package entity

import (
	"github.com/jinzhu/gorm"
)

// User entity
type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`

	// features
	Age        uint8
	Gender     string
	Occupation string
	Long       float32 `gorm:"index:location"`
	Lat        float32 `gorm:"index:location"`
	ZipCode    uint16
}
