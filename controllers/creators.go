package controllers

import (
	"net/http"
	"strconv"

	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gofiber/fiber/v2"
)

// GetCreators godoc
// @Summary Get all creators
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {array} models.Creator
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/ [get]
// @Param uuid query int false "The uuid of a creator"
// @Param name query string false "The name of a creator"
func GetCreators(c *fiber.Ctx) error {
	var creators []models.Creator
	var query models.Creator

	if err := c.QueryParser(&query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid params"})
	}

	if err := models.GetAllCreators(&creators, query); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "No creators found"})
	}

	return c.Status(http.StatusOK).JSON(creators)
}

// GetCreatorByID godoc
// @Summary Get creator by uuid
// @Tags Creators
// @Produce json
// @Success 200 {object} models.Creator
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/{uuid} [get]
// @Param uuid path int true "Unique Identifyer"
func GetCreatorByUUID(c *fiber.Ctx) error {
	var creator models.Creator
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := models.GetOneCreator(&creator, uuid); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Creator not found"})
	}

	return c.Status(http.StatusOK).JSON(creator)
}

// CreateCreator godoc
// @Summary Create creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /creators/ [post]
// @Param Body body models.Creator true "Creator obj"
func CreateCreator(c *fiber.Ctx) error {
	var newCreator models.Creator
	var creator models.Creator

	if err := c.BodyParser(&newCreator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect post body"})
	}

	if err := models.GetOneCreator(&creator, newCreator.UUID); err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Creator already exists"})
	}

	if err := models.AddNewCreator(&newCreator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Error while creating the model"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Created"})
}

// PatchCreator godoc
// @Summary Patch creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /creators/{uuid} [patch]
// @Param uuid path int true "Unique Identifier"
// @Param Body body models.Creator true "Creator obj"
func PatchCreator(c *fiber.Ctx) error {
	var newCreator models.Creator
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := c.BodyParser(&newCreator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect patch body"})
	}

	if err := models.PatchCreator(&newCreator, uuid); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Error while patching the model"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Updated"})
}

// DeleteCreator godoc
// @Summary Delete creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/{uuid} [delete]
// @Param uuid path string true "Unique Identifier"
func DeleteCreator(c *fiber.Ctx) error {
	var creator models.Creator
	uuid, err := strconv.Atoi(c.Params("uuid"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "UUID is not a number"})
	}

	if err := models.GetOneCreator(&creator, uuid); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Creator not found"})
	}

	if err := models.DeleteCreator(&creator, uuid); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Error while deleting the model"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Deleted"})
}
