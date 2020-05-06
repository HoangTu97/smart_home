package service

import (
	"Food/repository"
	"Food/service/mapper"
)

// RecipeService godoc
type RecipeService interface {
}

type recipeService struct {
	recipeRepository repository.RecipeRepository
	recipeMapper     mapper.RecipeMapper
}

// NewRecipe godoc
func NewRecipe(recipeRepository repository.RecipeRepository, recipeMapper mapper.RecipeMapper) RecipeService {
	return &recipeService{
		recipeRepository: recipeRepository,
		recipeMapper:     recipeMapper,
	}
}
