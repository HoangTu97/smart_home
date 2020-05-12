package mapper

import (
	"Food/dto"
	"Food/entity"
)

// RecipeIngredientsMapper godoc
type RecipeIngredientsMapper interface {
	ToDTO(entity entity.RecipeIngredients) dto.RecipeIngredientsDTO
	ToEntity(dto dto.RecipeIngredientsDTO) entity.RecipeIngredients
	ToDTOS(entityList []entity.RecipeIngredients) []dto.RecipeIngredientsDTO
	ToEntities(dtoList []dto.RecipeIngredientsDTO) []entity.RecipeIngredients
}

type recipeIngredientsMapper struct{}

// NewRecipeIngredients godoc
func NewRecipeIngredients() RecipeIngredientsMapper {
	return &recipeIngredientsMapper{}
}

func (mapper *recipeIngredientsMapper) ToDTO(entity entity.RecipeIngredients) dto.RecipeIngredientsDTO {
	return dto.RecipeIngredientsDTO{
		ID:           entity.ID,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
		RecipeID:     entity.RecipeID,
		IngredientID: entity.IngredientID,
	}
}

func (mapper *recipeIngredientsMapper) ToEntity(dto dto.RecipeIngredientsDTO) entity.RecipeIngredients {
	return entity.RecipeIngredients{
		ID:           dto.ID,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    dto.DeletedAt,
		RecipeID:     dto.RecipeID,
		IngredientID: dto.IngredientID,
	}
}

func (mapper *recipeIngredientsMapper) ToDTOS(entityList []entity.RecipeIngredients) []dto.RecipeIngredientsDTO {
	dtos := make([]dto.RecipeIngredientsDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = mapper.ToDTO(v)
	}

	return dtos
}

func (mapper *recipeIngredientsMapper) ToEntities(dtoList []dto.RecipeIngredientsDTO) []entity.RecipeIngredients {
	entities := make([]entity.RecipeIngredients, len(dtoList))

	for i, v := range dtoList {
		entities[i] = mapper.ToEntity(v)
	}

	return entities
}
