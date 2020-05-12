package service

import (
	"Food/entity"
	"Food/repository"
	"Food/service/mapper"
)

// RecipeIngredientsService godoc
type RecipeIngredientsService interface {
	FindByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients
	FindByRecipeID(recipeID uint) []entity.RecipeIngredients
}

type recipeIngredientsService struct {
	recipeIngredientsRepository repository.RecipeIngredientsRepository
	recipeIngredientsMapper     mapper.RecipeIngredientsMapper
}

// NewRecipeIngredients godoc
func NewRecipeIngredients(recipeIngredientsRepository repository.RecipeIngredientsRepository, recipeIngredientsMapper mapper.RecipeIngredientsMapper) RecipeIngredientsService {
	return &recipeIngredientsService{
		recipeIngredientsRepository: recipeIngredientsRepository,
		recipeIngredientsMapper:     recipeIngredientsMapper,
	}
}

func (s *recipeIngredientsService) FindByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients {
	return s.recipeIngredientsRepository.FindByRecipeIDs(recipeIDs)
}

func (s *recipeIngredientsService) FindByRecipeID(recipeID uint) []entity.RecipeIngredients {
	return s.recipeIngredientsRepository.FindByRecipeID(recipeID)
}
