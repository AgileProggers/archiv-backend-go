package main

import (
	"log"

	"github.com/AgileProggers/archiv-backend-go/controllers"
	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()
	database.ConnectDatabase()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := app.Group("/api/v1")
	{
		vodsGroup := v1.Group("/vods")
		{
			vodsGroup.Get("/", controllers.GetVods)
			vodsGroup.Post("/", controllers.CreateVod)
			vodsGroup.Get("/:id", controllers.GetVodById)
		}
	}
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
