package mapper

import (
	"Food/dto"
	"Food/entity"
)

// IngredientMapper godoc
type IngredientMapper interface {
	ToDTO(entity entity.Ingredient) dto.IngredientDTO
	ToEntity(dto dto.IngredientDTO) entity.Ingredient
	ToDTOS(entityList []entity.Ingredient) []dto.IngredientDTO
	ToEntities(dtoList []dto.IngredientDTO) []entity.Ingredient
}

type ingredientMapper struct{}

// NewIngredient godoc
func NewIngredient() IngredientMapper {
	return &ingredientMapper{}
}

func (mapper *ingredientMapper) ToDTO(entity entity.Ingredient) dto.IngredientDTO {
	return dto.IngredientDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Name:        entity.Name,
		Image:       entity.Image,
		Description: entity.Description,
	}
}

func (mapper *ingredientMapper) ToEntity(dto dto.IngredientDTO) entity.Ingredient {
	return entity.Ingredient{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Name:        dto.Name,
		Image:       dto.Image,
		Description: dto.Description,
	}
}

func (mapper *ingredientMapper) ToDTOS(entityList []entity.Ingredient) []dto.IngredientDTO {
	dtos := []dto.IngredientDTO{}

	for _, entity := range entityList {
		dtos = append(dtos, mapper.ToDTO(entity))
	}

	return dtos
}

func (mapper *ingredientMapper) ToEntities(dtoList []dto.IngredientDTO) []entity.Ingredient {
	entities := []entity.Ingredient{}

	for _, dto := range dtoList {
		entities = append(entities, mapper.ToEntity(dto))
	}

	return entities
}
