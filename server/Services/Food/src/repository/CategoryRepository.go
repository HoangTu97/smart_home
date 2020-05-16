package repository

import (
	"Food/util/page"
	"Food/util/pagination"
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// CategoryRepository godoc
type CategoryRepository interface {
	Save(category entity.Category) (entity.Category, error)
	FindOne(id uint) (entity.Category, error)
	FindOneByName(name string) ([]entity.Category, error)
	FindAll() []entity.Category
	FindPage(pageable pagination.Pageable) page.Page
	Delete(category entity.Category) error
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

func (r *categoryRepository) Save(category entity.Category) (entity.Category, error) {
	result := r.connection.Save(&category)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func (r *categoryRepository) FindOne(id uint) (entity.Category, error) {
	var category entity.Category
	result := r.connection.First(&category, id)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func (r *categoryRepository) FindOneByName(name string) ([]entity.Category, error) {
	var categories []entity.Category
	result := r.connection.Where("name LIKE ?", "%" + name + "%").Find(&categories)
	if result.Error != nil {
		return []entity.Category{}, result.Error
	}
	return categories, nil
}

func (r *categoryRepository) FindAll() []entity.Category {
	var categories []entity.Category
	r.connection.Find(&categories)
	return categories
}

func (r *categoryRepository) FindPage(pageable pagination.Pageable) page.Page {
	var categories []entity.Category

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Preload("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &categories)

	return page.From(toInterfacesFromCategories(categories), paginator.TotalRecord)
}

func (r *categoryRepository) Delete(category entity.Category) error {
	result := r.connection.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func toInterfacesFromCategories(categories []entity.Category) []interface{} {
	content := make([]interface{}, len(categories))
	for i, v := range categories {
		content[i] = v
	}
	return content
}
