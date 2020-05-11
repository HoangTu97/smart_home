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
	dtos := make([]dto.IngredientDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = mapper.ToDTO(v)
	}

	return dtos
}

func (mapper *ingredientMapper) ToEntities(dtoList []dto.IngredientDTO) []entity.Ingredient {
	entities := make([]entity.Ingredient, len(dtoList))

	for i, v := range dtoList {
		entities[i] = mapper.ToEntity(v)
	}

	return entities
}
