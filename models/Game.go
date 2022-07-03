package models

type Game struct {
	UUID   int    `gorm:"primaryKey;not null" json:"uuid" binding:"required"`
	Name   string `gorm:"not null" json:"name" binding:"required"`
	Boxart string `gorm:"not null" json:"box_art" binding:"required"`
	Clips  []Clip `gorm:"foreignKey:GameRefer" json:"clips"`
}
