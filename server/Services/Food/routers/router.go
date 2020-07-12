package routers

import (
	_ "Food/docs"
	CategoryResource "Food/controllers/categoryresource"
	RecipeResource "Food/controllers/reciperesource"
	UserResource "Food/controllers/userresource"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
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
		}

		// privateRoutes := apiRoutes.Group("/private")
		// {
			
		// }
	}

	return r
}
