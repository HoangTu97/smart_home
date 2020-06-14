package mapper

import (
	"Food/dto"
	"Food/entity"
	"strings"
)

func ToRecipeDTO(entity entity.Recipe) dto.RecipeDTO {
	return dto.RecipeDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Name:        entity.Name,
		Image:       entity.Image,
		Description: entity.Description,
		Duration:    entity.Duration,
		Photos:      strings.Split(entity.Photos, ","),
	}
}

func ToRecipe(dto dto.RecipeDTO) entity.Recipe {
	return entity.Recipe{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Name:        dto.Name,
		Image:       dto.Image,
		Description: dto.Description,
		Duration:    dto.Duration,
		Photos:      strings.Join(dto.Photos, ","),
	}
}

func ToRecipeDTOS(entityList []entity.Recipe) []dto.RecipeDTO {
	dtos := make([]dto.RecipeDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToRecipeDTO(entity)
	}

	return dtos
}

func ToRecipes(dtoList []dto.RecipeDTO) []entity.Recipe {
	entities := make([]entity.Recipe, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToRecipe(dto)
	}

	return entities
}

func ToDTOSInterfaceFromEntitiesInterface(interfaces []interface{}) []interface{} {
	dtos := make([]interface{}, len(interfaces))

	for i, inter := range interfaces {
		dtos[i] = ToRecipeDTO(inter.(entity.Recipe))
	}

	return dtos
}
