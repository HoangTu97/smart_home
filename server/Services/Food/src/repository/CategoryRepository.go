package repository

import (
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// CategoryRepository godoc
type CategoryRepository interface {
	Save(category entity.Category) entity.Category
	FindOne(id uint) entity.Category
	FindAll() []entity.Category
	Delete(category entity.Category)
}

type categoryRepository struct {
	connection *gorm.DB
}

// NewCategory godoc
func NewCategory(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		connection: db,
	}
}

func (r *categoryRepository) Save(category entity.Category) entity.Category {
	r.connection.Save(&category)
	return category
}

func (r *categoryRepository) FindOne(id uint) entity.Category {
	var category entity.Category
	r.connection.First(&category, id)
	return category
}

func (r *categoryRepository) FindAll() []entity.Category {
	var categories []entity.Category
	r.connection.Find(&categories)
	return categories
}

func (r *categoryRepository) Delete(category entity.Category) {
	r.connection.Delete(&category)
}
