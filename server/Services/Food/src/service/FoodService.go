package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

// FoodService ...
type FoodService interface {
	FindAll() []dto.FoodDTO
	FindOne(id uint) dto.FoodDTO
}

type foodService struct {
	foodRepository repository.FoodRepository
	foodMapper     mapper.FoodMapper
}

// New ...
func New(foodRepository repository.FoodRepository, foodMapper mapper.FoodMapper) FoodService {
	return &foodService{
		foodRepository: foodRepository,
		foodMapper:     foodMapper,
	}
}

func (service *foodService) FindAll() []dto.FoodDTO {
	foods := service.foodRepository.FindAll()
	return service.foodMapper.ToDTOS(foods)
}

func (service *foodService) FindOne(id uint) dto.FoodDTO {
	food := service.foodRepository.FindOne(id)
	return service.foodMapper.ToDTO(food)
}
