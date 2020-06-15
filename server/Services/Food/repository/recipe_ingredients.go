package repository

import (
	"Food/entity"
)

func FindRecipeIngredientsAll() []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	GetDB().Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func FindOneRecipeIngredients(id uint) entity.RecipeIngredients {
	var recipeIngredient entity.RecipeIngredients
	GetDB().First(&recipeIngredient, id)
	return recipeIngredient
}

func FindRecipeIngredientsByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	GetDB().Where("recipe_id IN (?)", recipeIDs).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func FindRecipeIngredientsByRecipeID(recipeID uint) []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	GetDB().Where("recipe_id = ?", recipeID).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}
