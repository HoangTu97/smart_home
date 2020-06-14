package mapper

import (
	"Food/dto"
	"Food/entity"
)

func ToRecipeIngredientsDTO(entity entity.RecipeIngredients) dto.RecipeIngredientsDTO {
	return dto.RecipeIngredientsDTO{
		ID:           entity.ID,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
		RecipeID:     entity.RecipeID,
		IngredientID: entity.IngredientID,
	}
}

func ToRecipeIngredient(dto dto.RecipeIngredientsDTO) entity.RecipeIngredients {
	return entity.RecipeIngredients{
		ID:           dto.ID,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    dto.DeletedAt,
		RecipeID:     dto.RecipeID,
		IngredientID: dto.IngredientID,
	}
}

func ToRecipeIngredientsDTOS(entityList []entity.RecipeIngredients) []dto.RecipeIngredientsDTO {
	dtos := make([]dto.RecipeIngredientsDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = ToRecipeIngredientsDTO(v)
	}

	return dtos
}

func ToRecipeIngredients(dtoList []dto.RecipeIngredientsDTO) []entity.RecipeIngredients {
	entities := make([]entity.RecipeIngredients, len(dtoList))

	for i, v := range dtoList {
		entities[i] = ToRecipeIngredient(v)
	}

	return entities
}
