package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// IngredientRepository godoc
type IngredientRepository interface {
	FindAll() []entity.Ingredient
	FindOne(id uint) entity.Ingredient
}

type ingredientRepository struct {
	connection *gorm.DB
}

// NewIngredient godoc
func NewIngredient(db *gorm.DB) IngredientRepository {
	return &ingredientRepository{
		connection: db,
	}
}

func (r *ingredientRepository) FindAll() []entity.Ingredient {
	var ingredients []entity.Ingredient
	r.connection.Find(&ingredients)
	return ingredients
}

func (r *ingredientRepository) FindOne(id uint) entity.Ingredient {
	var ingredient entity.Ingredient
	r.connection.Where("id = ?", id).Find(&ingredient)
	return ingredient
}
