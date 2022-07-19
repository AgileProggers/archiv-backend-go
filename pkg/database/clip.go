package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/AgileProggers/archiv-backend-go/pkg/database/internal/query"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
	"github.com/AgileProggers/archiv-backend-go/pkg/ressources"
)

func Clips() ([]*ent.Clip, error) {
	return client.Clip.Query().All(context.Background())
}

func ClipsByQuery(params map[string][]string) ([]*ent.Clip, error) {
	orderParams := params["order"]

	delete(params, "order")

	queryPredicate, err := query.BuildPredicate(clip.Columns, params)

	if err != nil {
		return nil, fmt.Errorf("build query predicate: %v", err)
	}

	buildQuery := client.Clip.Query().Where(queryPredicate).WithCreator().WithGame().WithVod()

	if orderParams != nil {
		order := strings.Split(orderParams[0], ",")

		if len(order) != 2 {
			return nil, fmt.Errorf("invalid order params. Example: 'date,desc'")
		}

		column := strings.ToLower(order[0])
		direction := strings.ToLower(order[1])

		if query.ContainsColumn(clip.Columns, column) {
			if direction == "asc" {
				buildQuery.Order(ent.Asc(column))
			} else {
				buildQuery.Order(ent.Desc(column))
			}
		}
	}

	return buildQuery.All(context.Background())
}

func ClipById(id int) (*ent.Clip, error) {
	return client.Clip.Get(context.Background(), id)
}

// func CreateClip(params map[string][]string) (*ent.Clip, error) {
func CreateClip(clip ressources.Clip) (*ent.Clip, error) {
	newClip := client.Clip.
		Create().
		SetDate(clip.Date).
		SetDuration(clip.Duration).
		SetFilename(clip.Filename).
		SetResolution(clip.Resolution).
		SetSize(clip.Size).
		SetTitle(clip.Title).
		SetViewCount(clip.ViewCount)

	// TODO: add vod, game and creator when its ready
	// vod, err := database.vo(clip.VodID)
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot find vod: %v", err)
	// }

	// newClip.AddVod(vod)

	return newClip.Save(context.Background())
}

func PatchClip(id int) *ent.ClipUpdateOne {
	return client.Clip.UpdateOneID(id)
}

func DeleteClip(id int) error {
	return client.Clip.DeleteOneID(id).
		Exec(context.Background())
}
