package router

import (
	"context"
	"fmt"
	"github.com/Gebes/there/v2"
	"github.com/go-playground/validator/v10"
)

var (
	app              *there.Router
	bindingValidator = validator.New()
)

func Listen() error {
	bindingValidator.SetTagName("binding")
	// Create fiber
	app = there.NewRouter()

	// Create routes
	v1 := app.Group("/api/v1")

	v1.Group("/vods").
		Get("/", GetVods).
		Post("/", CreateVod).
		Get("/:uuid", GetVodByUUID).
		Patch("/:uuid", PatchVod).
		Delete("/:uuid", DeleteVod)

	v1.Group("/clips").
		Get("/", GetClips).
		Post("/", CreateClip).
		Get("/:uuid", GetClipByUUID).
		Patch("/:uuid", PatchClip).
		Delete("/:uuid", DeleteClip)
	v1.Group("/games").
		Get("/", GetGames).
		Post("/", CreateGame).
		Get("/:uuid", GetGameByUUID).
		Patch("/:uuid", PatchGame).
		Delete("/:uuid", DeleteGame)

	v1.Group("/creators").
		Get("/", GetCreators).
		Post("/", CreateCreator).
		Get("/:uuid", GetCreatorByUUID).
		Patch("/:uuid", PatchCreator).
		Delete("/:uuid", DeleteCreator)

	v1.Group("/years").
		Get("/", GetYears)

	// Run
	err := app.Listen(8080)
	if err != nil {
		return fmt.Errorf("listen on port 8080: %v", err)
	}
	return nil
}

func Shutdown() error {
	return app.Server.Shutdown(context.Background())
}
