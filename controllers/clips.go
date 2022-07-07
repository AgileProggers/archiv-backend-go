package controllers

import (
	"net/http"
	"strings"

	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetClips godoc
// @Summary Get all clips
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {array} models.Clip
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /clips/ [get]
// @Param uuid query string false "The uuid of a clip"
// @Param title query string false "The title of a clip"
// @Param duration query int false "The duration of a clip"
// @Param date query string false "The date of a clip"
// @Param filename query string false "The filename of a clip"
// @Param resolution query string false "The resolution of a clip"
// @Param size query int false "The size of a clip"
// @Param viewcount query int false "The viewcount of a clip"
// @Param creator query int false "The creator id of a clip"
// @Param game query int false "The game id of a clip"
// @Param vod query string false "The vod id of a clip"
// @Param order query string false "Set order direction divided by comma. Possible ordering values: 'date', 'duration', 'size'. Possible directions: 'asc', 'desc'. Example: 'date,desc'"
func GetClips(c *fiber.Ctx) error {
	var clips []models.Clip
	var query models.Clip

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

	if err := models.GetAllClips(&clips, query, orderParams); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No clips found"})
	}

	return c.Status(http.StatusOK).JSON(clips)
}

// GetClipByID godoc
// @Summary Get clips by uuid
// @Tags Clips
// @Produce json
// @Success 200 {object} models.Clip
// @Failure 404 {string} string
// @Router /clips/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetClipByUUID(c *fiber.Ctx) error {
	var clip models.Clip

	if err := models.GetOneClip(&clip, c.Params("uuid")); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Clip not found"})
	}

	return c.Status(http.StatusOK).JSON(clip)
}

// CreateClip godoc
// @Summary Create clip
// @Tags Clips
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /clips/ [post]
// @Param Body body models.Clip true "Clip dict"
func CreateClip(c *fiber.Ctx) error {
	var newClip models.Clip
	var clip models.Clip

	if err := c.BodyParser(&newClip); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect post body"})
	}

	if err := models.GetOneClip(&clip, newClip.UUID); err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Clip already exists"})
	}

	if err := models.AddNewClip(&newClip); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created"})
}
