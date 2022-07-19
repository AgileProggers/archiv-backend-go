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

func Vods() ([]*ent.Vod, error) {
	return client.Vod.Query().All(context.Background())
}

func VodsByQuery(params map[string][]string) ([]*ent.Vod, error) {
	orderParams := params["order"]

	delete(params, "order")

	queryPredicate, err := query.BuildPredicate(vod.Columns, params)

	if err != nil {
		return nil, fmt.Errorf("build query predicate: %v", err)
	}

	buildQuery := client.Vod.Query().Where(queryPredicate).WithClips().WithGame()

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

func VodById(id int) (*ent.Vod, error) {
	return client.Vod.Get(context.Background(), id)
}

// func Createvod(params map[string][]string) (*ent.Vod, error) {
func Createvod(vod ressources.Vod) (*ent.Vod, error) {
	newVod := client.Vod.
		Create().
		SetDate(vod.Date).
		SetDuration(vod.Duration).
		SetFilename(vod.Filename).
		SetFps(float64(vod.Fps)).
		SetPublish(vod.Publish).
		SetResolution(vod.Resolution).
		SetSize(vod.Size).
		SetTitle(vod.Title)

	// TODO: add vod, game and creator when its ready
	// vod, err := database.vo(vod.vodID)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot find vod: %v", err)
	// }

	// newvod.Addvod(vod)

	return newVod.Save(context.Background())
}

func Patchvod(id int) *ent.VodUpdateOne {
	return client.Vod.UpdateOneID(id)
}

func Deletevod(id int) error {
	return client.Vod.DeleteOneID(id).
		Exec(context.Background())
}
