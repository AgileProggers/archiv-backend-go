package controllers

import (
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/database"
	"github.com/gofiber/fiber/v2"
)

func GetYears(c *fiber.Ctx) error {
	type Year struct {
		Year  int `json:"year"`
		Count int `json:"count"`
	}
	var years []Year

	database.DB.Raw("SELECT to_char(vods.date, 'yyyy') AS year, COUNT(to_char(vods.date, 'yyyy')) AS count FROM vods GROUP BY to_char(vods.date, 'yyyy') ORDER BY year DESC").Find(&years)
	if len(years) < 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "No years found"})
	}

	return c.Status(http.StatusOK).JSON(years)
}
