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
	var vod models.Vod
	var vods []models.Vod

	database.DB.Model((&vod)).Where("publish = ?", true).Preload("Clips").Find(&vods)

	if len(vods) < 1 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No vods found"})
	}

	return c.Status(http.StatusOK).JSON(vods)
}

// GetVodByUUID godoc
// @Summary Get vod by uuid
// @Tags Vods
// @Produce json
// @Success 200 {object} models.Vod
// @Failure 404 {string} string
// @Router /vods/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetVodByUUID(c *fiber.Ctx) error {
	var vod models.Vod
	var v models.Vod

	database.DB.Model((&vod)).Where("uuid = ?", c.Params("uuid")).Where("publish = ?", true).Preload("Clips").First(&v)

	if v.UUID == "" {
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
	var vod models.Vod

	if err := c.BodyParser(&newVod); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	database.DB.Model(&vod).Find(&vod, "uuid = ?", newVod.UUID)
	if vod.UUID != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Vod already exists. Use PATCH to modify existing vods."})
	}

	if err := database.DB.Model(&vod).Create(&newVod).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": fmt.Sprintf("%s created", newVod.UUID)})
}
