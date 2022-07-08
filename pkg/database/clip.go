package database

import (
	"errors"
	"strings"
	"time"
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

func GetAllClips(c *[]Clip, query Clip, o string) (err error) {
	if o == "" {
		o = "date desc"
	}
	result := database.Where(query).Order(o).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func AddNewClip(c *Clip) (err error) {
	var creator Creator
	var game Game
	var vod Vod
	var omits []string

	if err := GetOneCreator(&creator, c.Creator); err != nil {
		omits = append(omits, "Creator")
	}
	if err := GetOneGame(&game, c.Game); err != nil {
		omits = append(omits, "Game")
	}
	if err := GetOneVod(&vod, c.Vod); err != nil {
		omits = append(omits, "Vod")
	}
	if err = database.Omit(strings.Join(omits, ",")).Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneClip(c *Clip, uuid string) (err error) {
	result := database.Where("uuid = ?", uuid).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func PatchClip(c *Clip, uuid string) (err error) {
	var clip Clip
	if err := GetOneClip(&clip, uuid); err != nil {
		return errors.New("clip not found")
	}
	if err := database.Where("uuid = ?", uuid).Updates(c).Error; err != nil {
		return errors.New("update failed")
	}
	return nil
}

func DeleteClip(c *Clip, uuid string) (err error) {
	database.Where("uuid = ?", uuid).Delete(c)
	return nil
}
