package service

import (
	"Food/entity"
	"Food/repository"
)

func FindRecipeIngredientsByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients {
	return repository.FindRecipeIngredientsByRecipeIDs(recipeIDs)
}

func FindRecipeIngredientsByRecipeID(recipeID uint) []entity.RecipeIngredients {
	return repository.FindRecipeIngredientsByRecipeID(recipeID)
}
