package models

type Creator struct {
	UUID  int    `gorm:"primaryKey;uniqueIndex;not null" json:"uuid"`
	Name  string `gorm:"not null" json:"name" binding:"required"`
	Clips []Clip `gorm:"foreignKey:CreatorID" json:"clips,omitempty"`
}