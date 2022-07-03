package models

type Creator struct {
	UUID  int    `gorm:"primaryKey;not null" json:"uuid" binding:"required"`
	Name  string `gorm:"not null" json:"name" binding:"required"`
	Clips []Clip `gorm:"foreignKey:CreatorRefer" json:"clips"`
}
