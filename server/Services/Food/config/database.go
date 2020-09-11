package config

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite

	"Food/helpers/converter"
	"Food/models"
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
	// initDB(db)

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
	db.AutoMigrate(
		&models.Category{},
		&models.Recipe{},
		&models.Ingredient{},
		&models.RecipeIngredients{},
		&models.User{},
		&models.UserRecipeInteraction{},
	)
}

func initDB(db *gorm.DB) {
	count := 0
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		// cate1 := models.Category{Name: "Mexican Food", Image: "https://ak1.picdn.net/shutterstock/videos/19498861/thumb/1.jpg"}
		// cate2 := models.Category{Name: "Italian Food", Image: "https://images.unsplash.com/photo-1533777324565-a040eb52facd?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80"}
		// db.Create(&cate1)
		// db.Create(&cate2)

		// recipe1 := models.Recipe{Name: "Oatmeal Cookies", Description: "abc", Image: "abc", Photos: "abc", Duration: 15, Categories: []models.Category{cate1, cate2}}
		// db.Create(&recipe1)

		csvfile, err := os.Open("data/RAW_recipes.csv")
		if err != nil {
			log.Fatalln("Couldn't open the csv file", err)
		}
		mapCsv := csvToMap(csvfile)
		for _, item := range mapCsv {
			recipe := models.Recipe{
				Name: item["name"],
				// Description: item["description"],
				Description: "",
				Image:       "",
				Photos:      "",
				Duration:    converter.MustUint32(item["minutes"]),
			}
			db.Create(&recipe)
			fmt.Printf("Recipe : %v\n", recipe)
		}

		csvfile2, err2 := os.Open("data/RAW_interactions.csv")
		if err2 != nil {
			log.Fatalln("Couldn't open the csv file", err2)
		}
		mapCsv2 := csvToMap(csvfile2)
		for _, item := range mapCsv2 {
			t, _ := time.Parse("2006-01-02", item["date"])
			interaction := models.UserRecipeInteraction{
				UserID:   converter.MustUint(item["user_id"]),
				RecipeID: converter.MustUint(item["recipe_id"]),
				Rating:   converter.MustInt(item["rating"]),
				// Review:    item["review"],
				Review:    "",
				CreatedAt: t,
			}
			db.Create(&interaction)
			fmt.Printf("Interaction : %v\n", interaction)
		}
	}
}

// csvToMap takes a reader and returns an array of dictionaries, using the header row as the keys
func csvToMap(reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}
