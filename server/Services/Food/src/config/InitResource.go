package config

import (
	"Food/repository"
	"Food/service"
	"Food/service/mapper"
	"Food/web/rest"
)

var (
	RecipeResource     rest.RecipeResource
	CategoryResource   rest.CategoryResource
	IngredientResource rest.IngredientResource
)

// InitResource godoc
func InitResource() {
	recipeMapper := mapper.NewRecipe()
	categoryMapper := mapper.NewCategory()
	ingredientMapper := mapper.NewIngredient()
	recipeIngredientsMapper := mapper.NewRecipeIngredients()

	recipeRepository := repository.NewRecipe(DB)
	categoryRepository := repository.NewCategory(DB)
	ingredientRepository := repository.NewIngredient(DB)
	recipeIngredientsRepository := repository.NewRecipeIngredients(DB)

	recipeService := service.NewRecipe(recipeRepository, recipeMapper)
	categoryService := service.NewCategory(categoryRepository, categoryMapper)
	ingredientService := service.NewIngredient(ingredientRepository, ingredientMapper)
	recipeIngredientsService := service.NewRecipeIngredients(recipeIngredientsRepository, recipeIngredientsMapper)

	RecipeResource = rest.NewRecipe(recipeService, categoryService, ingredientService)
	CategoryResource = rest.NewCategory(categoryService)
	IngredientResource = rest.NewIngredient(ingredientService, recipeService, recipeIngredientsService)
}
