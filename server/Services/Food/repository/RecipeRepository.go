package repository

import (
	"Food/entity"
	"Food/pkg/page"
	"Food/pkg/pagination"
)

func FindAllRecipe() []entity.Recipe {
	var recipes []entity.Recipe
	db.Find(&recipes)
	return recipes
}

func FindOneRecipe(id uint) (entity.Recipe, error) {
	var recipe entity.Recipe
	result := db.First(&recipe, id)
	if result.Error != nil {
		return entity.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindOneRecipePreloadCate(id uint) (entity.Recipe, error) {
	var recipe entity.Recipe
	result := db.Preload("Categories").First(&recipe, id)
	if result.Error != nil {
		return entity.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindPageRecipeByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Model(&entity.Category{ID: cateID}).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByCates(cates []entity.Category, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Model(&cates).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByName(name string, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Where("name LIKE ?", "%" + name + "%").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Where("id IN (?)", db.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id = ?", ingredientID).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Where("id IN (?)", db.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id IN (?)", ingredientIDs).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipe(pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      db.Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

	return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindRecipeByName(name string) []entity.Recipe {
	var recipes []entity.Recipe
	db.Where("name LIKE ?", "%" + name + "%").Find(&recipes)
	return recipes
}

func CountRecipeByCateID(cateID uint) int {
	var result int
	db.Table("cate_recipes").Where("category_id = ?", cateID).Count(&result)
	return result
}

func toInterfacesFromRecipes(recipes []entity.Recipe) []interface{} {
	content := make([]interface{}, len(recipes))
	for i, v := range recipes {
		content[i] = v
	}
	return content
}
