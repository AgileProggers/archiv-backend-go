package main

import (
	"fmt"
	"net/http"
	"time"

	docs "archiv-backend-go/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type message struct {
	Message string `json:"message"`
}
type vod struct {
	UUID       string    `json:"uuid" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Duration   int       `json:"duration" binding:"required"`
	Date       time.Time `json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `json:"filename" binding:"required"`
	Resolution string    `json:"resolution" binding:"required"`
	Fps        int       `json:"fps" binding:"required"`
	Size       int       `json:"size" binding:"required"`
	Publish    bool      `json:"publish" binding:"required"`
	Bitrate    int       `json:"bitrate" binding:"required"`
}

var vods = []vod{
	{
		UUID:       "lul",
		Title:      "lul",
		Duration:   5,
		Date:       time.Now(),
		Filename:   "lul",
		Resolution: "lul",
		Fps:        48,
		Size:       5,
		Publish:    true,
		Bitrate:    500,
	},
}

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} vod
// @Failure 404 {object} message
// @Router /vods/ [get]
func getVods(c *gin.Context) {
	if len(vods) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No vods found"})
		return
	}
	c.IndentedJSON(http.StatusOK, vods)
}

// GetVodByUUID godoc
// @Summary Get vod by uuid
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {object} vod
// @Failure 404 {object} message
// @Router /vods/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func getVodById(c *gin.Context) {
	uuid := c.Param("uuid")
	for i, v := range vods {
		if v.UUID == uuid {
			c.IndentedJSON(http.StatusOK, &vods[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Vod not found"})
}

// CreateVod godoc
// @Summary Create vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {object} message
// @Failure 400 {object} message
// @Router /vods/ [post]
// @Param Body body object true "Vod dict"
func createVod(c *gin.Context) {
	var newVod vod
	if err := c.BindJSON(&newVod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	vods = append(vods, newVod)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("%s created", newVod.UUID)})
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		vodsGroup := v1.Group("/vods")
		{
			vodsGroup.GET("/", getVods)
			vodsGroup.POST("/", createVod)
			vodsGroup.GET("/:uuid", getVodById)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))
	router.Run("localhost:8080")
}
