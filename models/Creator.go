package models

import (
	"errors"

	"github.com/AgileProggers/archiv-backend-go/database"
)

type Creator struct {
	UUID  int    `gorm:"primaryKey;uniqueIndex;not null" json:"uuid"`
	Name  string `gorm:"not null" json:"name" binding:"required"`
	Clips []Clip `gorm:"foreignKey:Creator;association_foreignkey=UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clips,omitempty"`
}

func GetAllCreators(c *[]Creator, query Creator) (err error) {
	result := database.DB.Where(query).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func AddNewCreator(c *Creator) (err error) {
	if err = database.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}

func GetOneCreator(c *Creator, uuid int) (err error) {
	result := database.DB.Where("uuid = ?", uuid).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func PatchCreator(c *Creator, uuid int) (err error) {
	var creator Creator
	if err := GetOneCreator(&creator, uuid); err != nil {
		return errors.New("creator not found")
	}
	if err := database.DB.Where("uuid = ?", uuid).Updates(c).Error; err != nil {
		return errors.New("update failed")
	}
	return nil
}

func DeleteCreator(c *Creator, uuid int) (err error) {
	database.DB.Where("uuid = ?", uuid).Delete(c)
	return nil
}
