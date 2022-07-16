package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/AgileProggers/archiv-backend-go/pkg/database/internal/query"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/vod"
	"github.com/AgileProggers/archiv-backend-go/pkg/ressources"
)

func Creators() ([]*ent.Creator, error) {
	return client.Creator.Query().All(context.Background())
}

func CreatorsByQuery(params map[string][]string) ([]*ent.Creator, error) {
	orderParams := params["order"]

	delete(params, "order")
	
	queryPredicate, err := query.BuildPredicate(vod.Columns, params)

	if err != nil {
		return nil, fmt.Errorf("build query predicate: %v", err)
	}

	buildQuery := client.Creator.Query().Where(queryPredicate).WithClips()

	if orderParams != nil {
		order := strings.Split(orderParams[0], ",")

		if len(order) != 2 {
			return nil, fmt.Errorf("invalid order params. Example: 'date,desc'")
		}

		column := strings.ToLower(order[0])
		direction := strings.ToLower(order[1])

		if query.ContainsColumn(vod.Columns, column) {
			if direction == "asc" {
				buildQuery.Order(ent.Asc(column))
			} else {
				buildQuery.Order(ent.Desc(column))
			}
		}
	}

	return buildQuery.All(context.Background())
}

func CreatorById(id int) (*ent.Creator, error) {
	return client.Creator.Get(context.Background(), id)
}

// func Createvod(params map[string][]string) (*ent.Creator, error) {
func CreateCreator(creator ressources.Creator) (*ent.Creator, error) {
	newVod := client.Creator.
		Create().
		SetName(creator.Name)

	
	// TODO: add vod, game and creator when its ready
	// vod, err := database.vo(vod.vodID)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot find vod: %v", err)
	// }

	// newvod.Addvod(vod)
	
	return newVod.Save(context.Background())
}

func PatchCreator(id int) (*ent.CreatorUpdateOne) {
	return client.Creator.UpdateOneID(id)
}

func DeleteCreator(id int) (error) {
	return client.Creator.DeleteOneID(id).
		Exec(context.Background())
}