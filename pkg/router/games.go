package router

import (
	"github.com/Gebes/there/v2"
)

// GetGames godoc
// @Summary Get all games
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {array} database.Game
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /games/ [get]
// @Param uuid query int false "The uuid of a game"
// @Param name query string false "The name of a game"
// @Param box_art query string false "The box_art of a game"
func GetGames(request there.HttpRequest) there.HttpResponse {
	var games []int //database.Game
	// var query database.Game

	// err := request.Body.BindJson(&query)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind query: %v", err))
	// }
	// err = bindingValidator.Struct(query)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("validation error: %v", err))
	// }

	// if err := database.GetAllGames(&games, query); err != nil {
	// 	return there.Error(there.StatusNotFound, "No games found")
	// }

	return there.Json(there.StatusOK, games)
}

// GetGameByUUID godoc
// @Summary Get game by uuid
// @Tags Games
// @Produce json
// @Success 200 {object} database.Game
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /games/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetGameByUUID(request there.HttpRequest) there.HttpResponse {
	var game int // database.Game
	// uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	// }

	// if err := database.GetOneGame(&game, uuid); err != nil {
	// 	return there.Error(there.StatusNotFound, "Game not found")
	// }

	return there.Json(there.StatusOK, game)
}

// CreateGame godoc
// @Summary Create game
// @Tags Games
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /games/ [post]
// @Param Body body database.Game true "Game obj"
func CreateGame(request there.HttpRequest) there.HttpResponse {
	// var newGame database.Game
	// var game database.Game

	// err := request.Body.BindJson(&newGame)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	// }

	// if err := database.GetOneGame(&game, newGame.UUID); err == nil {
	// 	return there.Error(there.StatusBadRequest, "Game already exists")
	// }

	// if err := database.AddNewGame(&newGame); err != nil {
	// 	return there.Error(there.StatusUnprocessableEntity, "Error while creating the model")
	// }

	return there.Message(there.StatusCreated, "Created")
}

// PatchGame godoc
// @Summary Patch game
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /games/{uuid} [patch]
// @Param uuid path int true "Unique Identifier"
// @Param Body body database.Game true "Game obj"
func PatchGame(request there.HttpRequest) there.HttpResponse {
	// var newGame database.Game
	// uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	// }

	// err = request.Body.BindJson(&newGame)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	// }

	// if err := database.PatchGame(&newGame, uuid); err != nil {
	// 	return there.Error(there.StatusUnprocessableEntity, "Error while patching the model")
	// }

	return there.Message(there.StatusOK, "Updated")
}

// DeleteGame godoc
// @Summary Delete game
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /games/{uuid} [delete]
// @Param uuid path string true "Unique Identifier"
func DeleteGame(request there.HttpRequest) there.HttpResponse {
	// var game database.Game
	// uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	// }

	// if err := database.GetOneGame(&game, uuid); err != nil {
	// 	return there.Error(there.StatusNotFound, "Game not found")
	// }

	// if err := database.DeleteGame(&game, uuid); err != nil {
	// 	return there.Error(there.StatusBadRequest, "Error while deleting the model")
	// }

	return there.Message(there.StatusOK, "Deleted")
}
