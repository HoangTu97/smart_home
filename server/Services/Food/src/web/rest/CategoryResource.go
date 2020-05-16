package rest

import (
	"Food/util/pagination"
	"github.com/gin-gonic/gin"
	"Food/dto/response"
	"Food/dto/response/category"
	"Food/service"
	"Food/util/converter"
)

// CategoryResource godoc
type CategoryResource interface {
	GetAll(c *gin.Context)
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

func (r *categoryResource) GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := r.categoryService.FindPage(pageable)

	response.CreateSuccesResponse(c, category.CreateCategoryListResponseDTOFromPage(page))
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
