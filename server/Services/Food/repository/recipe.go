package repository

import (
	"Food/models"
	"Food/pkg/page"
	"Food/pkg/pagination"
)

func FindAllRecipe() []models.Recipe {
	var recipes []models.Recipe
	GetDB().Find(&recipes)
	return recipes
}

func FindOneRecipe(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := GetDB().First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindOneRecipePreloadCate(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := GetDB().Preload("Categories").First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindPageRecipeByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Model(&models.Category{ID: cateID}).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByCates(cates []models.Category, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Model(&cates).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByName(name string, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Where("name LIKE ?", "%" + name + "%").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Where("id IN (?)", GetDB().Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id = ?", ingredientID).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Where("id IN (?)", GetDB().Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id IN (?)", ingredientIDs).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipe(pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      GetDB().Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

	return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindRecipeByName(name string) []models.Recipe {
	var recipes []models.Recipe
	GetDB().Where("name LIKE ?", "%" + name + "%").Find(&recipes)
	return recipes
}

func CountRecipeByCateID(cateID uint) int {
	var result int
	GetDB().Table("cate_recipes").Where("category_id = ?", cateID).Count(&result)
	return result
}

func toInterfacesFromRecipes(recipes []models.Recipe) []interface{} {
	content := make([]interface{}, len(recipes))
	for i, v := range recipes {
		content[i] = v
	}
	return content
}
