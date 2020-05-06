package mapper

import (
	"Food/dto"
	"Food/entity"
	"strings"
)

// RecipeMapper godoc
type RecipeMapper interface {
	ToDTO(entity entity.Recipe) dto.RecipeDTO
	ToEntity(dto dto.RecipeDTO) entity.Recipe
	ToDTOS(entityList []entity.Recipe) []dto.RecipeDTO
	ToEntities(dtoList []dto.RecipeDTO) []entity.Recipe
}

type recipeMapper struct{}

// NewRecipe godoc
func NewRecipe() RecipeMapper {
	return &recipeMapper{}
}

func (mapper *recipeMapper) ToDTO(entity entity.Recipe) dto.RecipeDTO {
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

func (mapper *recipeMapper) ToEntity(dto dto.RecipeDTO) entity.Recipe {
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

func (mapper *recipeMapper) ToDTOS(entityList []entity.Recipe) []dto.RecipeDTO {
	dtos := []dto.RecipeDTO{}

	for _, entity := range entityList {
		dtos = append(dtos, mapper.ToDTO(entity))
	}

	return dtos
}

func (mapper *recipeMapper) ToEntities(dtoList []dto.RecipeDTO) []entity.Recipe {
	entities := []entity.Recipe{}

	for _, dto := range dtoList {
		entities = append(entities, mapper.ToEntity(dto))
	}

	return entities
}
