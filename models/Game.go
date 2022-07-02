package models

type Game struct {
	ID     int    `gorm:"index:unique" json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Boxart string `json:"box_art" binding:"required"`
}
