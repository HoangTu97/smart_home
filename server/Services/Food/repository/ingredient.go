package repository

import (
	"Food/entity"
)

func FindAllIngredient() []entity.Ingredient {
	var ingredients []entity.Ingredient
	GetDB().Find(&ingredients)
	return ingredients
}

func FindOneIngredient(id uint) (entity.Ingredient, error) {
	var ingredient entity.Ingredient
	result := GetDB().First(&ingredient, id)
	if result.Error != nil {
		return entity.Ingredient{}, result.Error
	}
	return ingredient, nil
}

func FindIngredientByName(name string) []entity.Ingredient {
	var ingredients []entity.Ingredient
	GetDB().Where("name LIKE ?", "%" + name + "%").Find(&ingredients)
	return ingredients
}
