package rest

import (
	"Food/dto/response"
	"Food/dto/response/recipe"
	"Food/service"
	"Food/util/converter"
	"Food/util/pagination"
	"github.com/gin-gonic/gin"
)

// RecipeResource godoc
type RecipeResource interface {
	GetByCategory(c *gin.Context)
	GetCountByCategory(c *gin.Context)
	GetByCategoryName(c *gin.Context)
	GetByIngredient(c *gin.Context)
	GetByIngredientName(c *gin.Context)
	GetByRecipeName(c *gin.Context)
	GetAll(c *gin.Context)
	GetDetailByID(c *gin.Context)
}

type recipeResource struct {
	recipeService     service.RecipeService
	categoryService   service.CategoryService
	ingredientService service.IngredientService
}

// NewRecipe godoc
func NewRecipe(recipeService service.RecipeService, categoryService service.CategoryService, ingredientService service.IngredientService) RecipeResource {
	return &recipeResource{
		recipeService:     recipeService,
		categoryService:   categoryService,
		ingredientService: ingredientService,
	}
}

func (r *recipeResource) GetByCategory(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := r.categoryService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := r.recipeService.FindPageByCateID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetCountByCategory(c *gin.Context) {
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := r.categoryService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	result := r.recipeService.CountByCateID(id)
	response.CreateSuccesResponse(c, &recipe.RecipeCountByCateResponseDTO{
		Value: result,
	})
}

func (r *recipeResource) GetByCategoryName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	cateName := converter.MustString(c.Query("name"))

	cates, isExist := r.categoryService.FindByName(cateName)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := r.recipeService.FindPageByCates(cates, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetByIngredient(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("ingredientId"))

	_, isExist := r.ingredientService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := r.recipeService.FindPageByIngredientID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetByIngredientName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	ingredientName := converter.MustString(c.Query("name"))

	ingredientIDs := r.ingredientService.FindIDByName(ingredientName)
	if len(ingredientIDs) == 0 {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := r.recipeService.FindPageByIngredientIDIn(ingredientIDs, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetByRecipeName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	recipeName := converter.MustString(c.Query("name"))

	page := r.recipeService.FindPageByName(recipeName, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := r.recipeService.FindPage(pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipeResource) GetDetailByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	recipeData, isExist := r.recipeService.FindOneWithCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, recipe.CreateRecipeDetailResponseDTO(recipeData))
}
