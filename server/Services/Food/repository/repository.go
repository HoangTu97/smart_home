package repository

import (
	"Food/entity"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// DB db instance
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
	db.AutoMigrate(&entity.Category{}, &entity.Recipe{}, &entity.Ingredient{}, &entity.RecipeIngredients{})
	db.AutoMigrate(&entity.User{})
}

func initDB(db *gorm.DB) {
	count := 0
	db.Model(&entity.Category{}).Count(&count)
	if count == 0 {
		cate1 := entity.Category{Name: "Mexican Food", Image: "https://ak1.picdn.net/shutterstock/videos/19498861/thumb/1.jpg"}
		cate2 := entity.Category{Name: "Italian Food", Image: "https://images.unsplash.com/photo-1533777324565-a040eb52facd?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80"}
		db.Create(&cate1)
		db.Create(&cate2)

		// ingre1 := entity.Ingredient{Name: "Oil", Image: "https://ak7.picdn.net/shutterstock/videos/27252067/thumb/11.jpg", Description: "oil"}
		// ingre2 := entity.Ingredient{Name: "Salt", Image: "https://image.freepik.com/free-photo/sea-salt-wooden-bowl-isolated-white-background_29402-416.jpg", Description: "saltt"}
		// ingre3 := entity.Ingredient{Name: "Russet potatoes", Image: "http://www.valleyspuds.com/wp-content/uploads/Russet-Potatoes-cut.jpg", Description: "Russet potatoe"}
		// DB.Create(&ingre1)
		// DB.Create(&ingre2)
		// DB.Create(&ingre3)

		recipe1 := entity.Recipe{Name: "Oatmeal Cookies", Description: "abc", Image: "abc", Photos: "abc", Duration: 15, Categories: []entity.Category{cate1, cate2}}
		db.Create(&recipe1)

		// ingres1 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre1.ID, Quantity: 200}
		// ingres2 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre2.ID, Quantity: 5}
		// ingres3 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre3.ID, Quantity: 300}
		// DB.Create(&ingres1)
		// DB.Create(&ingres2)
		// DB.Create(&ingres3)
	}
}
