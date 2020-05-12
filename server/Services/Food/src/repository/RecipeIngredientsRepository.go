package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// RecipeIngredientsRepository godoc
type RecipeIngredientsRepository interface {
	FindAll() []entity.RecipeIngredients
	FindByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients
	FindByRecipeID(recipeID uint) []entity.RecipeIngredients
	FindOne(id uint) entity.RecipeIngredients
}

type recipeIngredientsRepository struct {
	connection *gorm.DB
}

// NewRecipeIngredients godoc
func NewRecipeIngredients(db *gorm.DB) RecipeIngredientsRepository {
	return &recipeIngredientsRepository{
		connection: db,
	}
}

func (r *recipeIngredientsRepository) FindAll() []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	r.connection.Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func (r *recipeIngredientsRepository) FindOne(id uint) entity.RecipeIngredients {
	var recipeIngredient entity.RecipeIngredients
	r.connection.First(&recipeIngredient, id)
	return recipeIngredient
}

func (r *recipeIngredientsRepository) FindByRecipeIDs(recipeIDs []uint) []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	r.connection.Where("recipe_id IN (?)", recipeIDs).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func (r *recipeIngredientsRepository) FindByRecipeID(recipeID uint) []entity.RecipeIngredients {
	var recipeIngredientsLst []entity.RecipeIngredients
	r.connection.Where("recipe_id = ?", recipeID).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}
