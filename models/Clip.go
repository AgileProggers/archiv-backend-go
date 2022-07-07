package models

import (
	"errors"
	"strings"
	"time"

	"github.com/AgileProggers/archiv-backend-go/database"
)

type Clip struct {
	UUID       string    `gorm:"primaryKey;uniqueIndex" json:"uuid"`
	Title      string    `gorm:"not null" json:"title" binding:"required"`
	Duration   int       `gorm:"not null" json:"duration" binding:"required"`
	Date       time.Time `gorm:"not null" json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `gorm:"not null" json:"filename" binding:"required"`
	Resolution string    `gorm:"not null" json:"resolution" binding:"required"`
	Size       int       `gorm:"not null" json:"size" binding:"required"`
	Viewcount  int       `gorm:"not null" json:"viewcount" binding:"required"`
	Creator    int       `gorm:"colum:creator" json:"creator"`
	Game       int       `gorm:"colum:game" json:"game"`
	Vod        string    `gorm:"colum:vod" json:"vod"`
}

func GetAllClips(c *[]Clip, o string) (err error) {
	if o == "" {
		o = "date desc"
	}
	result := database.DB.Order(o).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func AddNewClip(c *Clip) (err error) {
	var creator Creator
	var game Game
	var vod Vod
	omits := []string{}

	if err := GetOneCreator(&creator, c.Creator); err != nil {
		omits = append(omits, "Creator")
	}
	if err := GetOneGame(&game, c.Game); err != nil {
		omits = append(omits, "Game")
	}
	if err := GetOneVod(&vod, c.Vod); err != nil {
		omits = append(omits, "Vod")
	}
	if err = database.DB.Omit(strings.Join(omits, ",")).Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneClip(c *Clip, uuid string) (err error) {
	result := database.DB.Where("uuid = ?", uuid).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func PutOneClip(c *Clip) (err error) {
	database.DB.Save(c)
	return nil
}

func DeleteClip(c *Clip, uuid string) (err error) {
	database.DB.Where("uuid = ?", uuid).Delete(c)
	return nil
}
