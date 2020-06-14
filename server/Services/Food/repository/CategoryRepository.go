package repository

import (
	"Food/entity"
	"Food/pkg/page"
	"Food/pkg/pagination"
)

func SaveCate(category entity.Category) (entity.Category, error) {
	result := db.Save(&category)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func FindOneCate(id uint) (entity.Category, error) {
	var category entity.Category
	result := db.First(&category, id)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func FindOneCateByName(name string) ([]entity.Category, error) {
	var categories []entity.Category
	result := db.Where("name LIKE ?", "%" + name + "%").Find(&categories)
	if result.Error != nil {
		return []entity.Category{}, result.Error
	}
	return categories, nil
}

func FindAllCate() []entity.Category {
	var categories []entity.Category
	db.Find(&categories)
	return categories
}

func FindPageCate(pageable pagination.Pageable) page.Page {
	var categories []entity.Category

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Preload("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &categories)

	return page.From(toInterfacesFromCategories(categories), paginator.TotalRecord)
}

func DeleteCate(category entity.Category) error {
	result := db.Delete(&category)
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
