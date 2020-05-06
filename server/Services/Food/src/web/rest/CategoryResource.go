package rest

import (
	"Food/dto/response"
	"Food/dto/response/category"
	"Food/service"
	"Food/util/converter"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CategoryResource godoc
type CategoryResource interface {
	GetByID(c *gin.Context)
	GetNameByID(c *gin.Context)
}

type categoryResource struct {
	categoryService service.CategoryService
}

// NewCategory godoc
func NewCategory(categoryService service.CategoryService) CategoryResource {
	return &categoryResource{
		categoryService: categoryService,
	}
}

func (r *categoryResource) GetByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	categoryDTO := r.categoryService.FindOne(id)
	if categoryDTO.ID == 0 {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &category.CategoryDetailResponseDTO{
		ID:    categoryDTO.ID,
		Name:  categoryDTO.Name,
		Image: categoryDTO.Image,
	})
}

func (r *categoryResource) GetNameByID(c *gin.Context) {
	c.JSON(http.StatusOK, &category.CategoryNameResponseDTO{})
}
