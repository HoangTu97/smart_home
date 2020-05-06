package config

import (
	"log"
	"github.com/joho/godotenv"
)

// LoadConfigFile godoc
func LoadConfigFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}