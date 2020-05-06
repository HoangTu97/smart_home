package config

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //
)

// DB godoc
var DB *gorm.DB

// SetupDB godoc
func SetupDB() {
	var err error
	DB, err = gorm.Open(DatabaseSetting.Type, DatabaseSetting.Name)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&entity.Category{}, &entity.Recipe{}, &entity.Ingredient{}, &entity.Ingredients{})
}

// CloseDB godoc
func CloseDB() {
	defer DB.Close()
}