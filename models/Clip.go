package models

import (
	"time"
)

type Clip struct {
	ID         string    `gorm:"index:unique" json:"id" binding:"required"`
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
	Vod        Vod       `json:"vod" gorm:"constraint:OnDelete:SET NULL;"`
	Creator    Creator   `json:"creator" binding:"required" gorm:"constraint:OnDelete:SET NULL;"`
	Game       Game      `json:"game" binding:"required" gorm:"constraint:OnDelete:SET NULL;"`
}
