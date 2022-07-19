package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/AgileProggers/archiv-backend-go/pkg/database/internal/query"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/game"
	"github.com/AgileProggers/archiv-backend-go/pkg/ressources"
)

func Games() ([]*ent.Game, error) {
	return client.Game.Query().All(context.Background())
}

func GamesByQuery(params map[string][]string) ([]*ent.Game, error) {
	orderParams := params["order"]

	delete(params, "order")

	queryPredicate, err := query.BuildPredicate(game.Columns, params)

	if err != nil {
		return nil, fmt.Errorf("build query predicate: %v", err)
	}

	// TODO: maybe relations?
	buildQuery := client.Game.Query().Where(queryPredicate)

	if orderParams != nil {
		order := strings.Split(orderParams[0], ",")

		if len(order) != 2 {
			return nil, fmt.Errorf("invalid order params. Example: 'date,desc'")
		}

		column := strings.ToLower(order[0])
		direction := strings.ToLower(order[1])

		if query.ContainsColumn(game.Columns, column) {
			if direction == "asc" {
				buildQuery.Order(ent.Asc(column))
			} else {
				buildQuery.Order(ent.Desc(column))
			}
		}
	}

	return buildQuery.All(context.Background())
}

func GameById(id int) (*ent.Game, error) {
	return client.Game.Get(context.Background(), id)
}

// func Creategame(params map[string][]string) (*ent.Game, error) {
func CreateGame(game ressources.Game) (*ent.Game, error) {
	newGame := client.Game.
		Create().
		SetBoxArt(game.Boxart).
		SetName(game.Name)

	// TODO: add game, game and creator when its ready
	// game, err := database.vo(game.gameID)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot find game: %v", err)
	// }

	// newgame.Addgame(game)

	return newGame.Save(context.Background())
}

func PatchGame(id int) *ent.GameUpdateOne {
	return client.Game.UpdateOneID(id)
}

func DeleteGame(id int) error {
	return client.Game.DeleteOneID(id).
		Exec(context.Background())
}
