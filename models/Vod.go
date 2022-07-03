package models

import "time"

type Vod struct {
	ID         string    `gorm:"index:unique" json:"id" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Duration   int       `json:"duration" binding:"required"`
	Date       time.Time `json:"date" time_format:"2006-01-02T15:04:05.000Z" binding:"required"`
	Filename   string    `json:"filename" binding:"required"`
	Resolution string    `json:"resolution" binding:"required"`
	Fps        float32   `json:"fps" binding:"required"`
	Size       int       `json:"size" binding:"required"`
	Publish    bool      `json:"publish" binding:"required"`
}
