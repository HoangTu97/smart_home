package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// RecipeIngredientsRepository godoc
type RecipeIngredientsRepository interface {
	FindAll() []entity.RecipeIngredients
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
	var RecipeIngredientsLst []entity.RecipeIngredients
	r.connection.Find(&RecipeIngredientsLst)
	return RecipeIngredientsLst
}

func (r *recipeIngredientsRepository) FindOne(id uint) entity.RecipeIngredients {
	var RecipeIngredients entity.RecipeIngredients
	r.connection.Where("id = ?", id).Find(&RecipeIngredients)
	return RecipeIngredients
}
