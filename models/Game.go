package models

type Game struct {
	UUID   int    `gorm:"primaryKey;uniqueIndex;not null" json:"uuid"`
	Name   string `gorm:"not null" json:"name" binding:"required"`
	Boxart string `gorm:"not null" json:"box_art" binding:"required"`
	Clips  []Clip `gorm:"foreignKey:Game;association_foreignkey=UUID" json:"-"`
}
