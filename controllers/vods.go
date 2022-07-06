package controllers

import (
	"net/http"
	"strings"

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
// @Param uuid query string false "The uuid of a vod"
// @Param title query string false "The title of a vod"
// @Param duration query int false "The duration of a vod"
// @Param date query string false "The date of a vod"
// @Param filename query string false "The filename of a vod"
// @Param resolution query string false "The resolution of a vod"
// @Param fps query int false "The fps of a vod"
// @Param size query int false "The size of a vod"
// @Param order query string false "Set order direction divided by comma. Possible ordering values: 'date', 'duration', 'size'. Possible directions: 'asc', 'desc'. Example: 'date,desc'"
func GetVods(c *fiber.Ctx) error {
	var vod models.Vod
	var vods []models.Vod

	if err := c.QueryParser(&vod); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid params"})
	}

	// custom ordering query
	if orderParams := c.Query("order"); orderParams != "" {
		order := strings.Split(orderParams, ",")
		if len(order) != 2 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid order params. Example: 'date,desc'"})
		}
		if !stringInSlice(order[0], []string{"date", "duration", "size"}) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid first order param. 'date', 'duration' or 'size'"})
		}
		if !stringInSlice(order[1], []string{"asc", "desc"}) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid second order param. 'asc' or 'desc'"})
		}
		database.DB.Model((&vod)).Where(&vod).Where("publish = ?", true).Order(strings.Join(order, " ")).Preload("Clips").Find(&vods)

	} else {
		database.DB.Model((&vod)).Where(&vod).Where("publish = ?", true).Order("date desc").Preload("Clips").Find(&vods)
	}

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

	database.DB.Model((&vod)).Where("publish = ?", true).Where("uuid = ?", c.Params("uuid")).Preload("Clips").Find(&vod)

	if vod.UUID == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Vod not found"})
	}

	return c.Status(http.StatusOK).JSON(vod)
}

// CreateVod godoc
// @Summary Create vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /vods/ [post]
// @Param Body body models.Vod true "Vod dict"
func CreateVod(c *fiber.Ctx) error {
	var newVod models.Vod
	var vod models.Vod

	if err := c.BodyParser(&newVod); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while parsing the body"})
	}

	database.DB.Model(&vod).Find(&vod, "uuid = ?", newVod.UUID)
	if vod.UUID != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Vod already exists. Use PATCH to modify existing vods."})
	}

	if err := database.DB.Model(&vod).Create(&newVod).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "created"})
}
