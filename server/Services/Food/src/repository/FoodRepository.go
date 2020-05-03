package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// FoodRepository ...
type FoodRepository interface {
	FindAll() []entity.Food
	FindOne(id uint) entity.Food
}

type repository struct {
	connection *gorm.DB
}

func New(db *gorm.DB) FoodRepository {
	return &repository{
		connection: db,
	}
}

func (r *repository) FindAll() []entity.Food {
	var foods []entity.Food
	r.connection.Find(&foods)
	return foods
}

func (r *repository) FindOne(id uint) entity.Food {
	var food entity.Food
	r.connection.Where("id = ?", id).Find(&food)
	return food
}
