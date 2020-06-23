package dto

import "time"

// UserDTO godoc
type UserDTO struct {
	ID       string
	Name     string
	Address  string
	Password string

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
