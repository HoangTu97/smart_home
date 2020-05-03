package rest

import (
	"Food/util/logger"
	"strconv"
	"Food/dto/response"
	"Food/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// FoodResource ...
type FoodResource interface {
	GetListFoods(ctx *gin.Context)
	GetFoodDetail(ctx *gin.Context)
}

type resource struct {
	foodService service.FoodService
}

func New(foodService service.FoodService) FoodResource {
	return &resource{
		foodService: foodService,
	}
}

func (r *resource) GetListFoods(ctx *gin.Context) {
	foodDTOS := r.foodService.FindAll()

	responseDTOS := []response.FoodListItemResponseDTO{}
	for _, foodDTO := range foodDTOS {
		responseDTOS = append(responseDTOS, response.FoodListItemResponseDTO{
			ID:         foodDTO.ID,
			Name:       foodDTO.Name,
			Categories: strings.Split(foodDTO.Categories, ","),
		})
	}

	ctx.JSON(http.StatusOK, responseDTOS)
}

func (r *resource) GetFoodDetail(ctx *gin.Context) {
	logger.Info("GetFoodDetail : id:{}", ctx.Param("id"))
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return
	}

	foodDTO := r.foodService.FindOne(uint(id))

	responseDTO := response.FoodListItemResponseDTO{
		ID:         foodDTO.ID,
		Name:       foodDTO.Name,
		Categories: strings.Split(foodDTO.Categories, ","),
	}

	ctx.JSON(http.StatusOK, responseDTO)
}
