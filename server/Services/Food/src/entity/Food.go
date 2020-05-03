package entity

import (
	"time"
)

// Food entity
type Food struct {
	ID          uint `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
	Name        string     `gorm:"type:varchar(255)"`
	Categories  string     `gorm:"type:varchar(255)"`
	Ingredients []Ingredients
}
