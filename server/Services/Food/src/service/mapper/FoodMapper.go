package mapper

import (
	"Food/dto"
	"Food/entity"
)

// FoodMapper ...
type FoodMapper interface {
	ToDTO(entity entity.Food) dto.FoodDTO
	ToEntity(dto dto.FoodDTO) entity.Food
	ToDTOS(entityList []entity.Food) []dto.FoodDTO
	ToEntities(dtoList []dto.FoodDTO) []entity.Food
}

type foodMapper struct{}

func New() FoodMapper {
	return &foodMapper{}
}

func (mapper *foodMapper) ToDTO(entity entity.Food) dto.FoodDTO {
	return dto.FoodDTO{
		ID:         entity.ID,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
		DeletedAt:  entity.DeletedAt,
		Name:       entity.Name,
		Categories: entity.Categories,
	}
}

func (mapper *foodMapper) ToEntity(dto dto.FoodDTO) entity.Food {
	return entity.Food{
		ID:         dto.ID,
		CreatedAt:  dto.CreatedAt,
		UpdatedAt:  dto.UpdatedAt,
		DeletedAt:  dto.DeletedAt,
		Name:       dto.Name,
		Categories: dto.Categories,
	}
}

func (mapper *foodMapper) ToDTOS(entityList []entity.Food) []dto.FoodDTO {
	dtos := []dto.FoodDTO{}

	for _, entity := range entityList {
		dtos = append(dtos, mapper.ToDTO(entity))
	}

	return dtos
}

func (mapper *foodMapper) ToEntities(dtoList []dto.FoodDTO) []entity.Food {
	entities := []entity.Food{}

	for _, dto := range dtoList {
		entities = append(entities, mapper.ToEntity(dto))
	}

	return entities
}
