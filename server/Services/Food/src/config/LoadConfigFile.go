package config

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadConfigFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}