package repository

import (
	"Food/models"
)

func FindAllIngredient() []models.Ingredient {
	var ingredients []models.Ingredient
	GetDB().Find(&ingredients)
	return ingredients
}

func FindOneIngredient(id uint) (models.Ingredient, error) {
	var ingredient models.Ingredient
	result := GetDB().First(&ingredient, id)
	if result.Error != nil {
		return models.Ingredient{}, result.Error
	}
	return ingredient, nil
}

func FindIngredientByName(name string) []models.Ingredient {
	var ingredients []models.Ingredient
	GetDB().Where("name LIKE ?", "%" + name + "%").Find(&ingredients)
	return ingredients
}
