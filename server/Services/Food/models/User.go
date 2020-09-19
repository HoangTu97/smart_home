package models

import (
	"gorm.io/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

// User entity
type User struct {
	ID       uint      `gorm:"primary_key"`
	UserID   uuid.UUID `gorm:"type:uuid;index:user_id"`
	Name     string    `gorm:"type:varchar(255)"`
	Address  string    `gorm:"type:varchar(255)"`
	Password string    `gorm:"type:varchar(255)" json:"Password"`

	// features
	Age        uint8
	Gender     string
	Occupation string
	Long       float32 `gorm:"index:location"`
	Lat        float32 `gorm:"index:location"`
	ZipCode    uint16

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserID = uuid.NewV4()
	return
}
