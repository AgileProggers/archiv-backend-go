package models

type Creator struct {
	ID   int    `gorm:"index:unique" json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
