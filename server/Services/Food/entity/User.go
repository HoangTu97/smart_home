package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User entity
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string `gorm:"type:varchar(255)"`
	Address  string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)" json:"Password"`

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

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("ID", uuid.NewV4())
}
