package main

import (
	"log"

	"github.com/AgileProggers/archiv-backend-go/controllers"
	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/docs"
	"github.com/AgileProggers/archiv-backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func main() {
	// Create fiber
	app := fiber.New()
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Archiv API Metrics Page"}))

	// Connect DB
	dsn := "host=localhost user=test password=test dbname=test port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Vod{}, &models.Game{}, &models.Creator{}, &models.Clip{})
	database.DB = db

	// Create swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	app.Get("/swagger/*", swagger.New(swagger.Config{
		DefaultModelsExpandDepth: -1,
	}))

	// Create routes
	v1 := app.Group("/api/v1")
	{
		vodsGroup := v1.Group("/vods")
		{
			vodsGroup.Get("/", controllers.GetVods)
			vodsGroup.Post("/", controllers.CreateVod)
			vodsGroup.Get("/:uuid", controllers.GetVodByUUID)
			vodsGroup.Patch("/:uuid", controllers.PatchVod)
			vodsGroup.Delete("/:uuid", controllers.DeleteVod)
		}
		clipsGroup := v1.Group("/clips")
		{
			clipsGroup.Get("/", controllers.GetClips)
			clipsGroup.Post("/", controllers.CreateClip)
			clipsGroup.Get("/:uuid", controllers.GetClipByUUID)
			clipsGroup.Patch("/:uuid", controllers.PatchClip)
			clipsGroup.Delete("/:uuid", controllers.DeleteClip)
		}
		gamesGroup := v1.Group("/games")
		{
			gamesGroup.Get("/", controllers.GetGames)
			gamesGroup.Post("/", controllers.CreateGame)
			gamesGroup.Get("/:uuid", controllers.GetGameByUUID)
			gamesGroup.Patch("/:uuid", controllers.PatchGame)
			gamesGroup.Delete("/:uuid", controllers.DeleteGame)
		}
		creatorsGroup := v1.Group("/creators")
		{
			creatorsGroup.Get("/", controllers.GetCreators)
			creatorsGroup.Post("/", controllers.CreateCreator)
			creatorsGroup.Get("/:uuid", controllers.GetCreatorByUUID)
			creatorsGroup.Patch("/:uuid", controllers.PatchCreator)
			creatorsGroup.Delete("/:uuid", controllers.DeleteCreator)
		}
		yearsGroup := v1.Group("/years")
		{
			yearsGroup.Get("/", controllers.GetYears)
		}
	}

	// Run
	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
