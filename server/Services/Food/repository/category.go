package repository

import (
	"Food/entity"
	"Food/pkg/page"
	"Food/pkg/pagination"
)

func SaveCate(category entity.Category) (entity.Category, error) {
	result := GetDB().Save(&category)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func FindOneCate(id uint) (entity.Category, error) {
	var category entity.Category

	// key := "CATE_" + converter.ToStr(id)
	// if gredis.Exists(key) {
	// 	data, err := gredis.Get(key)
	// 	if err == nil {
	// 		json.Unmarshal(data, &category)
	// 		return category, nil
	// 	}
	// 	logging.Info("FindOneCate", err)
	// }

	result := GetDB().First(&category, id)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}

	// gredis.Set(key, category, 3600)
	return category, nil
}

func FindOneCateByName(name string) ([]entity.Category, error) {
	var categories []entity.Category
	result := GetDB().Where("name LIKE ?", "%" + name + "%").Find(&categories)
	if result.Error != nil {
		return []entity.Category{}, result.Error
	}
	return categories, nil
}

func FindAllCate() []entity.Category {
	var categories []entity.Category
	GetDB().Find(&categories)
	return categories
}

func FindPageCate(pageable pagination.Pageable) page.Page {
	var categories []entity.Category

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Preload("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &categories)

	return page.From(toInterfacesFromCategories(categories), paginator.TotalRecord)
}

func DeleteCate(category entity.Category) error {
	result := GetDB().Delete(&category)
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
