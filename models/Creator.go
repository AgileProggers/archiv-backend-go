package models

import (
	"errors"

	"github.com/AgileProggers/archiv-backend-go/database"
)

type Creator struct {
	UUID  int    `gorm:"primaryKey;uniqueIndex;not null" json:"uuid"`
	Name  string `gorm:"not null" json:"name" binding:"required"`
	Clips []Clip `gorm:"foreignKey:Creator;association_foreignkey=UUID" json:"clips,omitempty"`
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

func PutOneCreator(c *Creator) (err error) {
	database.DB.Save(c)
	return nil
}

func DeleteCreator(c *Creator, uuid int) (err error) {
	database.DB.Where("uuid = ?", uuid).Delete(c)
	return nil
}
