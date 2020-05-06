package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// IngredientsRepository godoc
type IngredientsRepository interface {
	FindAll() []entity.Ingredients
	FindOne(id uint) entity.Ingredients
}

type ingredientsRepository struct {
	connection *gorm.DB
}

// NewIngredients godoc
func NewIngredients(db *gorm.DB) IngredientsRepository {
	return &ingredientsRepository{
		connection: db,
	}
}

func (r *ingredientsRepository) FindAll() []entity.Ingredients {
	var ingredientsLst []entity.Ingredients
	r.connection.Find(&ingredientsLst)
	return ingredientsLst
}

func (r *ingredientsRepository) FindOne(id uint) entity.Ingredients {
	var ingredients entity.Ingredients
	r.connection.Where("id = ?", id).Find(&ingredients)
	return ingredients
}
