package controllers

import (
	"fmt"
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// GetClips godoc
// @Summary Get all clips
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {array} models.Clip
// @Failure 404 {string} string
// @Router /clips/ [get]
func GetClips(c *fiber.Ctx) error {
	var clip models.Clip
	var clips []models.Clip

	database.DB.Model((&clip)).Find(&clips)

	if len(clips) < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No clips found"})
	}

	return c.Status(http.StatusOK).JSON(clips)
}

// GetClipsByUUID godoc
// @Summary Get clips by uuid
// @Tags Clips
// @Produce json
// @Success 200 {object} models.Clip
// @Failure 404 {string} string
// @Router /clips/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetClipsByUUID(c *fiber.Ctx) error {
	var clip models.Clip

	result := database.DB.Model((&clip)).Where("uuid = ?", c.Params("uuid")).Limit(1).Find(&clip)

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Clip not found"})
	}

	return c.Status(http.StatusOK).JSON(clip)
}

// CreateClip godoc
// @Summary Create clip
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /clips/ [post]
// @Param Body body object true "Clip dict"
func CreateClip(c *fiber.Ctx) error {
	var newClip models.Clip
	var clip models.Clip

	if err := c.BodyParser(&newClip); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	database.DB.Model(&clip).Find(&clip, "uuid = ?", newClip.UUID)
	if clip.UUID != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Clip already exists. Use PATCH to modify existing clips."})
	}

	if err := database.DB.Model(&newClip).Create(&newClip).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("%s created", newClip.UUID)})
}
