package database

// import (
// 	"errors"
// )

// type Game struct {
// 	UUID   int    `gorm:"primaryKey;uniqueIndex;not null" json:"uuid"`
// 	Name   string `gorm:"not null" json:"name" binding:"required"`
// 	Boxart string `gorm:"not null" json:"box_art" binding:"required"`
// 	Clips  []Clip `gorm:"foreignKey:Game;association_foreignkey=UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
// }

// func GetAllGames(g *[]Game, query Game) (err error) {
// 	result := database.Where(query).Find(g)
// 	if result.RowsAffected == 0 {
// 		return errors.New("not found")
// 	}
// 	return nil
// }

// func AddNewGame(g *Game) (err error) {
// 	if err = database.Create(g).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetOneGame(g *Game, uuid int) (err error) {
// 	result := database.Where("uuid = ?", uuid).Find(g)
// 	if result.RowsAffected == 0 {
// 		return errors.New("not found")
// 	}
// 	return nil
// }

// func PatchGame(g *Game, uuid int) (err error) {
// 	var game Game
// 	if err := GetOneGame(&game, uuid); err != nil {
// 		return errors.New("game not found")
// 	}
// 	if err := database.Where("uuid = ?", uuid).Updates(g).Error; err != nil {
// 		return errors.New("update failed")
// 	}
// 	return nil
// }

// func DeleteGame(g *Game, uuid int) (err error) {
// 	database.Where("uuid = ?", uuid).Delete(g)
// 	return nil
// }
