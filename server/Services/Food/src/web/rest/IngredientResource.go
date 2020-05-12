package rest

import (
	"Food/dto/response"
	"Food/dto/response/ingredient"
	"Food/service"
	"Food/util/converter"
	"github.com/gin-gonic/gin"
)

// IngredientResource godoc
type IngredientResource interface {
	GetNameByID(c *gin.Context)
	GetImageByID(c *gin.Context)
	GetByRecipeName(c *gin.Context)
	GetByRecipeID(c *gin.Context)
}

type ingredientResource struct {
	ingredientService        service.IngredientService
	recipeService            service.RecipeService
	recipeIngredientsService service.RecipeIngredientsService
}

// NewIngredient godoc
func NewIngredient(ingredientService service.IngredientService, recipeService service.RecipeService, recipeIngredientsService service.RecipeIngredientsService) IngredientResource {
	return &ingredientResource{
		ingredientService:        ingredientService,
		recipeService:            recipeService,
		recipeIngredientsService: recipeIngredientsService,
	}
}

func (r *ingredientResource) GetNameByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := r.ingredientService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &ingredient.IngredientNameResponseDTO{Name: ingredientDTO.Name})
}

func (r *ingredientResource) GetImageByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := r.ingredientService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &ingredient.IngredientImageResponseDTO{Image: ingredientDTO.Image})
}

func (r *ingredientResource) GetByRecipeName(c *gin.Context) {
	recipeName := converter.MustString(c.Query("recipeName"))
	recipeIDs := r.recipeService.FindIDsByName(recipeName)

	if len(recipeIDs) == 0 {
		response.CreateSuccesResponse(c, ingredient.CreateEmptyIngredientListResponseDTO())
		return
	}

	recipeIngredients := r.recipeIngredientsService.FindByRecipeIDs(recipeIDs)
	response.CreateSuccesResponse(c, ingredient.CreateIngredientListResponseDTO(recipeIngredients))
}

func (r *ingredientResource) GetByRecipeID(c *gin.Context) {
	recipeID := converter.MustUint(c.Query("recipeId"))
	_, isExist := r.recipeService.FindOne(recipeID)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	recipeIngredients := r.recipeIngredientsService.FindByRecipeID(recipeID)
	response.CreateSuccesResponse(c, ingredient.CreateIngredientListResponseDTO(recipeIngredients))
}
