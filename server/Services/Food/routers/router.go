package routers

import (
	CategoryResource "Food/web/rest/categoryresource"
	RecipeResource "Food/web/rest/reciperesource"
	UserResource "Food/web/rest/userresource"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(include.CORS())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoutes := r.Group("/api")
	{
		publicRoutes := apiRoutes.Group("/public")
		{
			publicCategoryRoutes := publicRoutes.Group("/category")
			{
				// publicCategoryRoutes.GET("/detail/:id", CategoryResource.GetByID)
				// publicCategoryRoutes.GET("/getName/:id", CategoryResource.GetNameByID)
				publicCategoryRoutes.GET("/getAll", CategoryResource.GetAll)
			}

			publicRecipeRoutes := publicRoutes.Group("/recipe")
			{
				publicRecipeRoutes.GET("/getAll", RecipeResource.GetAll)
				publicRecipeRoutes.GET("/detail/:id", RecipeResource.GetDetailByID)
				publicRecipeRoutes.GET("/getByCategory/:categoryId", RecipeResource.GetByCategory)
				publicRecipeRoutes.GET("/countByCategory/:categoryId", RecipeResource.GetCountByCategory)
				// publicRecipeRoutes.GET("/getByIngredient/:ingredientId", RecipeResource.GetByIngredient)
				// publicRecipeRoutes.GET("/searchByIngredientName", RecipeResource.GetByIngredientName)
				// publicRecipeRoutes.GET("/searchByCategoryName", RecipeResource.GetByCategoryName)
				publicRecipeRoutes.GET("/searchByRecipeName", RecipeResource.GetByRecipeName)
			}

			// publicIngredientRoutes := publicRoutes.Group("/ingredient")
			// {
			// 	publicIngredientRoutes.GET("/getName/:id", IngredientResource.GetNameByID)
			// 	publicIngredientRoutes.GET("/getImage/:id", IngredientResource.GetImageByID)
			// 	publicIngredientRoutes.GET("/searchIngredients", IngredientResource.GetByRecipeName)
			// 	publicIngredientRoutes.GET("/searchIngredientsByRecipeId", IngredientResource.GetByRecipeID)
			// }

			publicUserRoutes := publicRoutes.Group("/user")
			{
				publicUserRoutes.POST("/register", UserResource.Register)
				publicUserRoutes.POST("/login", UserResource.Login)
			}
		}

		// privateRoutes := apiRoutes.Group("/private")
		// {
			
		// }
	}

	return r
}
