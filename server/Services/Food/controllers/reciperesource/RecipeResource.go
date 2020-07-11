package reciperesource

import (
	"Food/dto/response"
	"Food/dto/response/recipe"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

func GetByCategory(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := service.FindOneCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByCateID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetCountByCategory(c *gin.Context) {
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := service.FindOneCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	result := service.CountRecipeByCateID(id)
	response.CreateSuccesResponse(c, &recipe.RecipeCountByCateResponseDTO{
		Value: result,
	})
}

func GetByCategoryName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	cateName := converter.MustString(c.Query("name"))

	cates, isExist := service.FindCateByName(cateName)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByCates(cates, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetByIngredient(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("ingredientId"))

	_, isExist := service.FindOneIngredientDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByIngredientID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetByIngredientName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	ingredientName := converter.MustString(c.Query("name"))

	ingredientIDs := service.FindIngredientIDsByName(ingredientName)
	if len(ingredientIDs) == 0 {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByIngredientIDIn(ingredientIDs, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetByRecipeName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	recipeName := converter.MustString(c.Query("name"))

	page := service.FindPageRecipeByName(recipeName, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := service.FindPageRecipe(pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetDetailByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	recipeData, isExist := service.FindOneRecipeWithCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, recipe.CreateRecipeDetailResponseDTO(recipeData))
}
