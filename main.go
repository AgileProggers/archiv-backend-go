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

type vod struct {
	ID         string    `gorm:"index:unique" json:"id" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Duration   int       `json:"duration" binding:"required"`
	Date       time.Time `json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `json:"filename" binding:"required"`
	Resolution string    `json:"resolution" binding:"required"`
	Fps        float32   `json:"fps" binding:"required"`
	Size       int       `json:"size" binding:"required"`
	Publish    bool      `json:"publish" binding:"required"`
	Bitrate    float32   `json:"bitrate"`
}

type clip struct {
	ID         string    `gorm:"index:unique"json:"id" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Duration   int       `json:"duration" binding:"required"`
	Date       time.Time `json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `json:"filename" binding:"required"`
	Resolution string    `json:"resolution" binding:"required"`
	Fps        int       `json:"fps" binding:"required"`
	Size       int       `json:"size" binding:"required"`
	Publish    bool      `json:"publish" binding:"required"`
	Bitrate    float32   `json:"bitrate"`
	Viewcount  int       `json:"view_count" binding:"required"`
	Vod        vod       `json:"vod" gorm:"constraint:OnDelete:SET NULL;"`
	Creator    creator   `json:"creator" binding:"required" gorm:"constraint:OnDelete:SET NULL;"`
	Game       game      `json:"game" binding:"required" gorm:"constraint:OnDelete:SET NULL;"`
}

type creator struct {
	ID   int    `gorm:"index:unique"json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type game struct {
	ID     int    `gorm:"index:unique"json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Boxart string `json:"box_art" binding:"required"`
}

var vods = []vod{}
var clips = []clip{}
var creators = []creator{}
var games = []game{}

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} vod
// @Failure 404 {string} string
// @Router /vods/ [get]
func getVods(c *gin.Context) {
	if len(vods) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No vods found"})
		return
	}
	for i, v := range vods {
		vods[i].Bitrate = float32(v.Size * 8 / v.Duration)
	}
	c.IndentedJSON(http.StatusOK, vods)
}

// GetVodByID godoc
// @Summary Get vod by id
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {object} vod
// @Failure 404 {string} string
// @Router /vods/{id} [get]
// @Param id path string true "Unique Identifier"
func getVodById(c *gin.Context) {
	id := c.Param("id")
	for _, v := range vods {
		if v.ID == id {
			v.Bitrate = float32(v.Size * 8 / v.Duration)
			c.IndentedJSON(http.StatusOK, v)
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
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /vods/ [post]
// @Param Body body object true "Vod dict"
func createVod(c *gin.Context) {
	var newVod vod
	if err := c.BindJSON(&newVod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	vods = append(vods, newVod)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("%s created", newVod.ID)})
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
			vodsGroup.GET("/:id", getVodById)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)))
	router.Run("localhost:8080")
}
