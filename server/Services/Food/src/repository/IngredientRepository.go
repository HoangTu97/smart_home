package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// IngredientRepository godoc
type IngredientRepository interface {
	FindAll() []entity.Ingredient
	FindOne(id uint) (entity.Ingredient, error)
	FindByName(name string) []entity.Ingredient
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

func (r *ingredientRepository) FindOne(id uint) (entity.Ingredient, error) {
	var ingredient entity.Ingredient
	result := r.connection.First(&ingredient, id)
	if result.Error != nil {
		return entity.Ingredient{}, result.Error
	}
	return ingredient, nil
}

func (r *ingredientRepository) FindByName(name string) []entity.Ingredient {
	var ingredients []entity.Ingredient
	r.connection.Where("name LIKE ?", "%" + name + "%").Find(&ingredients)
	return ingredients
}
