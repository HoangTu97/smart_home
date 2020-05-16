package config

import (
	sawggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

// SetupRoutes godoc
func SetupRoutes() {
	gin.SetMode(ServerSetting.RunMode)
	gin.ForceConsoleColor()

	server := gin.New()
	server.Use(gin.Recovery())

	apiRoutes := server.Group("/api")
	{
		publicRoutes := apiRoutes.Group("/public")
		{
			publicCategoryRoutes := publicRoutes.Group("/category")
			{
				publicCategoryRoutes.GET("/detail/:id", CategoryResource.GetByID)
				publicCategoryRoutes.GET("/getName/:id", CategoryResource.GetNameByID)
				publicCategoryRoutes.GET("/getAll", CategoryResource.GetAll)
			}

			publicRecipeRoutes := publicRoutes.Group("/recipe")
			{
				publicRecipeRoutes.GET("/getAll", RecipeResource.GetAll)
				publicRecipeRoutes.GET("/detail/:id", RecipeResource.GetDetailByID)
				publicRecipeRoutes.GET("/getByCategory/:categoryId", RecipeResource.GetByCategory)
				publicRecipeRoutes.GET("/countByCategory/:categoryId", RecipeResource.GetCountByCategory)
				publicRecipeRoutes.GET("/getByIngredient/:ingredientId", RecipeResource.GetByIngredient)
				publicRecipeRoutes.GET("/searchByIngredientName", RecipeResource.GetByIngredientName)
				publicRecipeRoutes.GET("/searchByCategoryName", RecipeResource.GetByCategoryName)
				publicRecipeRoutes.GET("/searchByRecipeName", RecipeResource.GetByRecipeName)
			}

			publicIngredientRoutes := publicRoutes.Group("/ingredient")
			{
				publicIngredientRoutes.GET("/getName/:id", IngredientResource.GetNameByID)
				publicIngredientRoutes.GET("/getImage/:id", IngredientResource.GetImageByID)
				publicIngredientRoutes.GET("/searchIngredients", IngredientResource.GetByRecipeName)
				publicIngredientRoutes.GET("/searchIngredientsByRecipeId", IngredientResource.GetByRecipeID)
			}

			// getCategoryById(categoryId)
			// getCategoryName(categoryId)
			// getIngredientName(ingredientID)
			// getIngredientUrl(ingredientID)
			// getAllIngredients(idArray)
			// getRecipes(categoryId)
			// getRecipesByIngredient(ingredientId)
			// getNumberOfRecipes(categoryId)
			// getRecipesByIngredientName(ingredientName)
			// getRecipesByCategoryName(categoryName)
			// getRecipesByRecipeName(recipeName)
		}
	}

	url := ginSwagger.URL("http://localhost:" + ServerSetting.HTTPPort + "/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(sawggerFiles.Handler, url))

	server.Run(":" + ServerSetting.HTTPPort)
}
