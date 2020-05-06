package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Food/service"
)

// IngredientResource godoc
type IngredientResource interface {
	GetNameByID(c *gin.Context)
	GetImageByID(c *gin.Context)
	GetByRecipeName(c *gin.Context)
}

type ingredientResource struct {
	ingredientService service.IngredientService
}

// NewIngredient godoc
func NewIngredient(ingredientService service.IngredientService) IngredientResource {
	return &ingredientResource{
		ingredientService: ingredientService,
	}
}

func (r *ingredientResource) GetNameByID(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (r *ingredientResource) GetImageByID(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (r *ingredientResource) GetByRecipeName(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
