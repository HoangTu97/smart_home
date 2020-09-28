package dto

import (
	"Food/helpers/security"
	"time"
)

// UserDTO godoc
type UserDTO struct {
	ID       uint
	UserID   string
	Name     string
	Password string
	Roles    []security.UserRole
	Address  string

	// features
	Age        uint8
	Gender     string
	Occupation string
	Long       float32
	Lat        float32
	ZipCode    uint16

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
