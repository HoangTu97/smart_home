package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Setup initializes the database instance
func Setup() {
	// db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name))
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	DB = db
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}

// GetDB get connection
func GetDB() *gorm.DB {
	return DB
}
