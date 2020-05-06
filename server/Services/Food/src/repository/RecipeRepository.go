package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// RecipeRepository godoc
type RecipeRepository interface {
	FindAll() []entity.Recipe
	FindOne(id uint) entity.Recipe
}

type recipeRepository struct {
	connection *gorm.DB
}

// NewRecipe godoc
func NewRecipe(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		connection: db,
	}
}

func (r *recipeRepository) FindAll() []entity.Recipe {
	var recipes []entity.Recipe
	r.connection.Find(&recipes)
	return recipes
}

func (r *recipeRepository) FindOne(id uint) entity.Recipe {
	var recipe entity.Recipe
	r.connection.Where("id = ?", id).Find(&recipe)
	return recipe
}
