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

	DB.LogMode(true)

	// Migrate the schema
	DB.AutoMigrate(&entity.Category{}, &entity.Recipe{}, &entity.Ingredient{}, &entity.RecipeIngredients{})
	count := 0
	DB.Model(entity.Category{}).Count(&count)
	if count == 0 {
		cate1 := entity.Category{Name: "Mexican Food", Image: "https://ak1.picdn.net/shutterstock/videos/19498861/thumb/1.jpg"}
		cate2 := entity.Category{Name: "Italian Food", Image: "https://images.unsplash.com/photo-1533777324565-a040eb52facd?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80"}
		DB.Create(&cate1)
		DB.Create(&cate2)

		ingre1 := entity.Ingredient{Name: "Oil", Image: "https://ak7.picdn.net/shutterstock/videos/27252067/thumb/11.jpg", Description: "oil"}
		ingre2 := entity.Ingredient{Name: "Salt", Image: "https://image.freepik.com/free-photo/sea-salt-wooden-bowl-isolated-white-background_29402-416.jpg", Description: "saltt"}
		ingre3 := entity.Ingredient{Name: "Russet potatoes", Image: "http://www.valleyspuds.com/wp-content/uploads/Russet-Potatoes-cut.jpg", Description: "Russet potatoe"}
		DB.Create(&ingre1)
		DB.Create(&ingre2)
		DB.Create(&ingre3)

		recipe1 := entity.Recipe{Name: "Oatmeal Cookies", Description: "abc", Image: "abc", Photos: "abc", Duration: 15, Categories: []entity.Category{cate1, cate2}}
		DB.Create(&recipe1)

		ingres1 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre1.ID, Quantity: 200}
		ingres2 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre2.ID, Quantity: 5}
		ingres3 := entity.RecipeIngredients{RecipeID: recipe1.ID, IngredientID: ingre3.ID, Quantity: 300}
		DB.Create(&ingres1)
		DB.Create(&ingres2)
		DB.Create(&ingres3)

		// recipe1.RecipeIngredients = []entity.RecipeIngredients{ingres1, ingres2, ingres3}
		// ingre1.RecipeIngredients = []entity.RecipeIngredients{ingres1}
		// ingre2.RecipeIngredients = []entity.RecipeIngredients{ingres2}
		// ingre3.RecipeIngredients = []entity.RecipeIngredients{ingres3}

	}
}

// CloseDB godoc
func CloseDB() {
	defer DB.Close()
}
