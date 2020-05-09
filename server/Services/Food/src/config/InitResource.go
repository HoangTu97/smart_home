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

	recipeRepository := repository.NewRecipe(DB)
	categoryRepository := repository.NewCategory(DB)
	ingredientRepository := repository.NewIngredient(DB)

	recipeService := service.NewRecipe(recipeRepository, recipeMapper)
	categoryService := service.NewCategory(categoryRepository, categoryMapper)
	ingredientService := service.NewIngredient(ingredientRepository, ingredientMapper)

	RecipeResource = rest.NewRecipe(recipeService, categoryService)
	CategoryResource = rest.NewCategory(categoryService)
	IngredientResource = rest.NewIngredient(ingredientService)
}
