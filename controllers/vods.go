package controllers

import (
	"fmt"
	"net/http"

	"github.com/AgileProggers/archiv-backend-go/models"
	"github.com/gin-gonic/gin"
)

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} models.Vod
// @Failure 404 {string} string
// @Router /vods/ [get]
func GetVods(c *gin.Context) {
	var vods []models.Vod

	result := models.DB.Model((&vods)).Where("publish = ?", true).Find(&vods)

	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No vods found"})
		return
	}

	c.IndentedJSON(http.StatusOK, vods)
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
func GetVodById(c *gin.Context) {
	var vods []models.Vod
	var v models.Vod

	result := models.DB.Model((&vods)).Where("id = ?", c.Param("id")).Where("publish = ?", true).Limit(1).Find(&v)

	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Vod not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, v)
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
func CreateVod(c *gin.Context) {
	var newVod models.Vod
	var vods []models.Vod

	models.DB.Find(&vods)

	if err := c.BindJSON(&newVod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Model(&vods).Where("id = ?", newVod.ID)

	if result.RowsAffected < 1 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Vod already exists. Use PATCH to modify existing vods."})
		return
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

	models.DB.Create(&v)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("%s created", newVod.ID)})
}
