package controllers

import (
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// Get Games godoc
// @Summary Get all games
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {array} models.Game
// @Failure 404 {string} string
// @Router /games/ [get]
func GetGames(c *fiber.Ctx) error {
	game := new(models.Game)
	var games []models.Game

	result := database.DB.Model((&game)).Find(&games)

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No games found"})
	}

	return c.Status(http.StatusOK).JSON(games)
}

// GetGameByID godoc
// @Summary Get game by uuid
// @Tags Games
// @Produce json
// @Success 200 {object} models.Game
// @Failure 404 {string} string
// @Router /games/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetGameByUUID(c *fiber.Ctx) error {
	game := new(models.Game)
	var g models.Game

	result := database.DB.Model((&game)).Find(&g, "uuid = ?", c.Params("uuid"))

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Game not found"})
	}

	return c.Status(http.StatusOK).JSON(g)
}

// CreateGame godoc
// @Summary Create game
// @Tags Games
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /games/ [post]
// @Param Body body models.Game true "Game dict"
func CreateGame(c *fiber.Ctx) error {
	var newGame models.Game
	var game models.Game

	if err := c.BodyParser(&newGame); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while parsing the body"})
	}

	database.DB.Model(&game).Find(&game, "uuid = ?", newGame.UUID)
	if game.UUID > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Game already exists. Use PATCH to modify existing games."})
	}

	if err := database.DB.Model(&game).Create(&newGame).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "created"})
}
