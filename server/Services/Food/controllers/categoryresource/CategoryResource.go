package categoryresource

import (
	"Food/dto/response"
	"Food/dto/response/category"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Category all
// @Summary GetAll
// @Tags PublicCategory
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=category.CategoryListResponseDTO{items=category.CategoryListItemResponseDTO{}}} "desc"
// @Router /api/public/category/getAll [get]
func GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := service.FindPageCate(pageable)

	response.CreateSuccesResponse(c, category.CreateCategoryListResponseDTOFromPage(page))
}

func GetByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	categoryDTO, isExist := service.FindOneCate(id)
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

func GetNameByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	categoryDTO, isExist := service.FindOneCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, &category.CategoryNameResponseDTO{
		Name:  categoryDTO.Name,
	})
}
