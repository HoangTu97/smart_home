package repository

import (
	"Food/util/page"
	"Food/util/pagination"
	"Food/entity"
	"github.com/jinzhu/gorm"
)

// RecipeRepository godoc
type RecipeRepository interface {
	FindAll() []entity.Recipe
	FindOne(id uint) (entity.Recipe, error)
	FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page
	FindPageByCates(cates []entity.Category, pageable pagination.Pageable) page.Page
	FindPageByName(name string, pageable pagination.Pageable) page.Page
	FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page
	FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page
	FindByName(name string) []entity.Recipe
	CountByCateID(cateID uint) int
}

type recipeRepository struct {
	connection *gorm.DB
}

// NewRecipe godoc
func NewRecipe(db *gorm.DB) RecipeRepository {
	return &recipeRepository{
		connection: db,
	}
}

func (r *recipeRepository) FindAll() []entity.Recipe {
	var recipes []entity.Recipe
	r.connection.Find(&recipes)
	return recipes
}

func (r *recipeRepository) FindOne(id uint) (entity.Recipe, error) {
	var recipe entity.Recipe
	result := r.connection.First(&recipe, id)
	if result.Error != nil {
		return entity.Recipe{}, result.Error
	}
	return recipe, nil
}

func (r *recipeRepository) FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Model(&entity.Category{ID: cateID}).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipeRepository) FindPageByCates(cates []entity.Category, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Model(&cates).Related(&recipes, "Recipes").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipeRepository) FindPageByName(name string, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Where("name LIKE ?", "%" + name + "%").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipeRepository) FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Where("id IN (?)", r.connection.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id = ?", ingredientID).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipeRepository) FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	var recipes []entity.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.connection.Where("id IN (?)", r.connection.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id IN (?)", ingredientIDs).SubQuery()).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipeRepository) FindByName(name string) []entity.Recipe {
	var recipes []entity.Recipe
	r.connection.Where("name LIKE ?", "%" + name + "%").Find(&recipes)
	return recipes
}

func (r *recipeRepository) CountByCateID(cateID uint) int {
	var result int
	r.connection.Table("cate_recipes").Where("category_id = ?", cateID).Count(&result)
	return result
}

func toInterfacesFromRecipes(recipes []entity.Recipe) []interface{} {
	content := make([]interface{}, len(recipes))
	for i, v := range recipes {
		content[i] = v
	}
	return content
}
