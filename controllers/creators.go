package controllers

import (
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// GetCreators godoc
// @Summary Get all creators
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {array} models.Creator
// @Failure 404 {string} string
// @Router /creators/ [get]
func GetCreators(c *fiber.Ctx) error {
	var creators []models.Creator

	result := database.DB.Model((&creators)).Find(&creators)

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No creators found"})
	}

	return c.Status(http.StatusOK).JSON(creators)
}

// GetCreatorByID godoc
// @Summary Get creator by uuid
// @Tags Creators
// @Produce json
// @Success 200 {object} models.Creator
// @Failure 404 {string} string
// @Router /creators/{uuid} [get]
// @Param uuid path string true "Unique Identifyer"
func GetCreatorByUUID(c *fiber.Ctx) error {
	var creators []models.Creator
	var cr models.Creator

	result := database.DB.Model((&creators)).Preload("Clips").Find(&cr, "uuid = ?", c.Params("uuid"))

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Creator not found"})
	}

	return c.Status(http.StatusOK).JSON(cr)
}

// CreateCreator godoc
// @Summary Create creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /creators/ [post]
// @Param Body body models.Creator true "Creator dict"
func CreateCreator(c *fiber.Ctx) error {
	var newCreator models.Creator
	var creator models.Creator

	if err := c.BodyParser(&newCreator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	database.DB.Model(&creator).Find(&creator, "uuid = ?", newCreator.UUID)
	if creator.UUID > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Creator already exists. Use PATCH to modify existing creators."})
	}

	if err := database.DB.Model(&creator).Create(&newCreator).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "created"})
}
