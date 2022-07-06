package models

import (
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
