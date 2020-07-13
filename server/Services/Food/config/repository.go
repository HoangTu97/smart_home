package config

import (
	"Food/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// DB db instance
var DB *gorm.DB

// Setup initializes the database instance
func SetupDB() {
	// db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name))
	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	migrateDB(db)
	initDB(db)

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

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Category{}, &models.Recipe{}, &models.Ingredient{}, &models.RecipeIngredients{})
	db.AutoMigrate(&models.User{})
}

func initDB(db *gorm.DB) {
	count := 0
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		cate1 := models.Category{Name: "Mexican Food", Image: "https://ak1.picdn.net/shutterstock/videos/19498861/thumb/1.jpg"}
		cate2 := models.Category{Name: "Italian Food", Image: "https://images.unsplash.com/photo-1533777324565-a040eb52facd?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80"}
		db.Create(&cate1)
		db.Create(&cate2)

		recipe1 := models.Recipe{Name: "Oatmeal Cookies", Description: "abc", Image: "abc", Photos: "abc", Duration: 15, Categories: []models.Category{cate1, cate2}}
		db.Create(&recipe1)
	}
}
