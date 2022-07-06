package main

import (
	"log"

	"github.com/AgileProggers/archiv-backend-go/controllers"
	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
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
			vodsGroup.Get("/:uuid", controllers.GetVodByUUID)
		}
		clipsGroup := v1.Group("/clips")
		{
			clipsGroup.Get("/", controllers.GetClips)
			clipsGroup.Post("/", controllers.CreateClip)
			clipsGroup.Get("/:uuid", controllers.GetClipsByUUID)
		}
		gamesGroup := v1.Group("/games")
		{
			gamesGroup.Get("/", controllers.GetGames)
			gamesGroup.Post("/", controllers.CreateGame)
			gamesGroup.Get("/:uuid", controllers.GetGameByUUID)
		}
		creatorsGroup := v1.Group("/creators")
		{
			creatorsGroup.Get("/", controllers.GetCreators)
			creatorsGroup.Post("/", controllers.CreateCreator)
			creatorsGroup.Get("/:uuid", controllers.GetCreatorByUUID)
		}
	}
	v1.Get("/years", controllers.GetYears)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DefaultModelsExpandDepth: -1,
	}))
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Archiv API Metrics Page"}))
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
