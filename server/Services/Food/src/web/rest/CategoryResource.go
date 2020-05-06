package rest

import (
	"Food/dto/response"
	"Food/dto/response/category"
	"Food/service"
	"Food/util/converter"
	"github.com/gin-gonic/gin"
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
	categoryDTO, isExist := r.categoryService.FindOne(id)
	if !isExist {
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
	id := converter.MustUint(c.Param("id"))
	categoryDTO, isExist := r.categoryService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, &category.CategoryNameResponseDTO{
		Name:  categoryDTO.Name,
	})
}
