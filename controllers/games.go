package controllers

import (
	"net/http"
	"strconv"

	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// Get Games godoc
// @Summary Get all games
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {array} models.Game
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /games/ [get]
// @Param uuid query int false "The uuid of a game"
// @Param name query string false "The name of a game"
// @Param box_art query string false "The box_art of a game"
func GetGames(c *fiber.Ctx) error {
	var games []models.Game
	var query models.Game

	if err := c.QueryParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid params"})
	}

	if err := models.GetAllGames(&games, query); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No games found"})
	}

	return c.Status(http.StatusOK).JSON(games)
}

// GetGameByID godoc
// @Summary Get game by uuid
// @Tags Games
// @Produce json
// @Success 200 {object} models.Game
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /games/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetGameByUUID(c *fiber.Ctx) error {
	var game models.Game
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := models.GetOneGame(&game, uuid); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Game not found"})
	}

	return c.Status(http.StatusOK).JSON(game)
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
// @Param Body body models.Game true "Game obj"
func CreateGame(c *fiber.Ctx) error {
	var newGame models.Game
	var game models.Game

	if err := c.BodyParser(&newGame); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect post body"})
	}

	if err := models.GetOneGame(&game, newGame.UUID); err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Game already exists"})
	}

	if err := models.AddNewGame(&newGame); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created"})
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
// @Param Body body models.Game true "Game obj"
func PatchGame(c *fiber.Ctx) error {
	var newGame models.Game
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := c.BodyParser(&newGame); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect patch body"})
	}

	if err := models.PatchGame(&newGame, uuid); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Error while patching the model"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated"})
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
func DeleteGame(c *fiber.Ctx) error {
	var game models.Game
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := models.GetOneGame(&game, uuid); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Game not found"})
	}

	if err := models.DeleteGame(&game, uuid); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Error while deleting the model"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted"})
}
