package database

import (
	"errors"
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
	Clips      []Clip    `gorm:"foreignKey:Vod;association_foreignkey=UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clips"`
}

func GetAllVods(v *[]Vod, query Vod, o string) (err error) {
	if o == "" {
		o = "date desc"
	}
	result := database.Where(query).Where("publish = ?", true).Order(o).Preload("Clips").Find(v)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func AddNewVod(v *Vod) (err error) {
	if err = database.Create(v).Error; err != nil {
		return err
	}
	return nil
}

func GetOneVod(v *Vod, uuid string) (err error) {
	result := database.Where("uuid = ?", uuid).Where("publish = ?", true).Preload("Clips").Find(v)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func PatchVod(v *Vod, uuid string) (err error) {
	var vod Vod
	if err := GetOneVod(&vod, uuid); err != nil {
		return errors.New("vod not found")
	}
	if err := database.Where("uuid = ?", uuid).Updates(v).Error; err != nil {
		return errors.New("update failed")
	}
	return nil
}

func DeleteVod(v *Vod, uuid string) (err error) {
	database.Where("uuid = ?", uuid).Delete(v)
	return nil
}
