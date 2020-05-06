package rest

import (
	"net/http"
	// "Food/util/pagination"
	"github.com/gin-gonic/gin"
	"Food/service"
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
	recipeService service.RecipeService
}

// NewRecipe godoc
func NewRecipe(recipeService service.RecipeService) RecipeResource {
	return &recipeResource{
		recipeService: recipeService,
	}
}

func (r *recipeResource) GetByCategory(c *gin.Context) {
	// page, size := pagination.GetPage(c)
	// r.recipeService
	c.JSON(http.StatusOK, nil)
}

func (r *recipeResource) GetCountByCategory(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
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

