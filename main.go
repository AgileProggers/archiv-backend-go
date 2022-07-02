package main

import (
	"github.com/AgileProggers/archiv-backend-go/controllers"
	"github.com/AgileProggers/archiv-backend-go/docs"
	"github.com/AgileProggers/archiv-backend-go/models"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})
	models.ConnectDatabase()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		vodsGroup := v1.Group("/vods")
		{
			vodsGroup.GET("/", controllers.GetVods)
			vodsGroup.POST("/", controllers.CreateVod)
			vodsGroup.GET("/:id", controllers.GetVodById)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))
	router.Run(":8080")
}
