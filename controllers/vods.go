package controllers

import (
	"fmt"
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} models.Vod
// @Failure 404 {string} string
// @Router /vods/ [get]
func GetVods(c *fiber.Ctx) error {
	var vods []models.Vod

	result := database.DB.Model((&vods)).Where("publish = ?", true).Find(&vods)

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No vods found"})
	}

	return c.Status(http.StatusOK).JSON(vods)
}

// GetVodByID godoc
// @Summary Get vod by id
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {object} models.Vod
// @Failure 404 {string} string
// @Router /vods/{id} [get]
// @Param id path string true "Unique Identifier"
func GetVodById(c *fiber.Ctx) error {
	var vods []models.Vod
	var v models.Vod

	result := database.DB.Model((&vods)).Where("id = ?", c.Params("id")).Where("publish = ?", true).Limit(1).Find(&v)

	if result.RowsAffected < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Vod not found"})
	}

	return c.Status(http.StatusOK).JSON(v)
}

// CreateVod godoc
// @Summary Create vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /vods/ [post]
// @Param Body body object true "Vod dict"
func CreateVod(c *fiber.Ctx) error {
	var newVod models.Vod
	var vods []models.Vod
	var exists models.Vod

	database.DB.Find(&vods)

	if err := c.BodyParser(&newVod); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if result := database.DB.Model(&vods).Where("id = ?", newVod.ID).Find(&exists); result.RowsAffected > 0 {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"message": "Vod already exists. Use PATCH to modify existing vods."})
	}

	v := models.Vod{
		ID:         newVod.ID,
		Title:      newVod.Title,
		Duration:   newVod.Duration,
		Date:       newVod.Date,
		Filename:   newVod.Filename,
		Resolution: newVod.Resolution,
		Fps:        newVod.Fps,
		Size:       newVod.Size,
		Publish:    newVod.Publish,
	}

	database.DB.Create(&v)
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("%s created", newVod.ID)})
}
