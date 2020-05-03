package dto

import (
	"time"
)

// FoodDTO ...
type FoodDTO struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Name       string
	Categories string
}
