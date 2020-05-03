package config

import (
	"Food/service"
	"Food/repository"
	"Food/web/rest"
	"Food/service/mapper"
)

var (
	FoodResource rest.FoodResource
)

func InitResource() {
	foodMapper := mapper.New()

	foodRepository := repository.New(DB)

	foodService := service.New(foodRepository, foodMapper)

	FoodResource = rest.New(foodService)
}