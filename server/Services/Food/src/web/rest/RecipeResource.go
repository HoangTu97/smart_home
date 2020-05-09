package rest

import (
	"Food/dto/response"
	"Food/dto/response/recipe"
	"Food/service"
	"Food/util/converter"
	"Food/util/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecipeResource godoc
type RecipeResource interface {
	GetByCategory(c *gin.Context)
	GetCountByCategory(c *gin.Context)
	GetByIngredient(c *gin.Context)
	GetByIngredientName(c *gin.Context)
	GetByCategoryName(c *gin.Context)
	GetByRecipeName(c *gin.Context)
}

type recipeResource struct {
	recipeService   service.RecipeService
	categoryService service.CategoryService
}

// NewRecipe godoc
func NewRecipe(recipeService service.RecipeService, categoryService service.CategoryService) RecipeResource {
	return &recipeResource{
		recipeService:   recipeService,
		categoryService: categoryService,
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

func (r *recipeResource) GetByIngredient(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (r *recipeResource) GetByIngredientName(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (r *recipeResource) GetByCategoryName(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (r *recipeResource) GetByRecipeName(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
