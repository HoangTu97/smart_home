package config

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //
)

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open(DatabaseSetting.Type, DatabaseSetting.Name)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&entity.Food{}, &entity.Ingredients{})
}

func CloseDB() {
	defer DB.Close()
}