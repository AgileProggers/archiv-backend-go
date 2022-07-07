package controllers

import (
	"net/http"
	"strings"

	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} models.Vod
// @Failure 400 {string} string
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
	var vods []models.Vod
	var query models.Vod

	if err := c.QueryParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid params"})
	}

	// custom ordering query
	orderParams := ""
	if orderParams = c.Query("order"); orderParams != "" {
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
		orderParams = strings.Replace(orderParams, ",", " ", -1)
	}

	if err := models.GetAllVods(&vods, query, orderParams); err != nil {
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

	if err := models.GetOneVod(&vod, c.Params("uuid")); err != nil {
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
// @Failure 500 {string} string
// @Router /vods/ [post]
// @Param Body body models.Vod true "Vod dict"
func CreateVod(c *fiber.Ctx) error {
	var newVod models.Vod
	var vod models.Vod

	if err := c.BodyParser(&newVod); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect post body"})
	}

	if err := models.GetOneVod(&vod, newVod.UUID); err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Vod already exists"})
	}

	if err := models.AddNewVod(&newVod); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created"})
}
