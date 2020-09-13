package routers

import (
	CategoryResource "Food/controllers/categoryresource"
	RecipeResource "Food/controllers/reciperesource"
	UserResource "Food/controllers/userresource"
	ImageResource "Food/controllers/imageresource"

	"github.com/gin-gonic/gin"
)

func InitRouterApi(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		publicRoutes := apiRoutes.Group("/public")
		{
			publicCategoryRoutes := publicRoutes.Group("/category")
			{
				publicCategoryRoutes.GET("/getAll", CategoryResource.GetAll)
			}

			publicRecipeRoutes := publicRoutes.Group("/recipe")
			{
				publicRecipeRoutes.GET("/getAll", RecipeResource.GetAll)
				publicRecipeRoutes.GET("/detail/:id", RecipeResource.GetDetailByID)
				publicRecipeRoutes.GET("/getByCategory/:categoryId", RecipeResource.GetByCategory)
				publicRecipeRoutes.GET("/countByCategory/:categoryId", RecipeResource.GetCountByCategory)
				publicRecipeRoutes.GET("/searchByRecipeName", RecipeResource.GetByRecipeName)
			}

			publicUserRoutes := publicRoutes.Group("/user")
			{
				publicUserRoutes.POST("/register", UserResource.Register)
				publicUserRoutes.POST("/login", UserResource.Login)
			}

			publicImageRoutes := publicRoutes.Group("/image")
			{
				publicImageRoutes.POST("/upload", ImageResource.Upload)
				publicImageRoutes.GET("/:id", ImageResource.FileDisplay)
				publicImageRoutes.GET("/:id/download", ImageResource.Download)
			}
		}

		// privateRoutes := apiRoutes.Group("/private")
		// {
			
		// }
	}
}