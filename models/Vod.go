package models

import (
	"time"
)

type Vod struct {
	UUID       string    `gorm:"primaryKey;uniqueIndex" json:"uuid"`
	Title      string    `gorm:"not null" json:"title" binding:"required"`
	Duration   int       `gorm:"not null" json:"duration" binding:"required"`
	Date       time.Time `gorm:"not null" json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `gorm:"not null" json:"filename" binding:"required"`
	Resolution string    `gorm:"not null" json:"resolution" binding:"required"`
	Fps        float32   `gorm:"not null" json:"fps" binding:"required"`
	Size       int       `gorm:"not null" json:"size" binding:"required"`
	Publish    bool      `gorm:"not null" json:"publish" binding:"required"`
	Clips      []Clip    `gorm:"foreignKey:Vod" json:"clips"`
}
