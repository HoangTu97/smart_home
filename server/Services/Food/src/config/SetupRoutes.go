package config

import (
	sawggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	gin.SetMode(ServerSetting.RunMode)
	gin.ForceConsoleColor()

	server := gin.New()
	server.Use(gin.Recovery())

	apiRoutes := server.Group("/api")
	{
		publicRoutes := apiRoutes.Group("/public")
		{
			publicFoodRoutes := publicRoutes.Group("/food")
			{
				publicFoodRoutes.GET("/list", FoodResource.GetListFoods)
				publicFoodRoutes.GET("/detail/:id", FoodResource.GetFoodDetail)
			}
		}
	}

	url := ginSwagger.URL("http://localhost:" + ServerSetting.HTTPPort + "/swagger/doc.json") // The url pointing to API definition
	server.GET("/swagger/*any", ginSwagger.WrapHandler(sawggerFiles.Handler, url))

	server.Run(":" + ServerSetting.HTTPPort)
}
