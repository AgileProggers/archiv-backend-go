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
			return nil, fmt.Errorf("Invalid order params. Example: 'date,desc'")
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
	return client.Clip.
		Create().
		SetDate(clip.Date).
		SetDuration(clip.Duration).
		SetFilename(clip.Filename).
		SetResolution(clip.Resolution).
		SetSize(clip.Size).
		SetTitle(clip.Title).
		SetViewCount(clip.ViewCount).
		Save(context.Background())
}

func PatchClip(id int, clip ressources.Clip) (*ent.Clip, error) {
	originalClip, err := client.Clip.Get(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("Clip not found")
	}

	 fmt.Println(clip)
	 fmt.Println(originalClip)


	return client.Clip.UpdateOneID(id).
		SetDate(clip.Date).
		SetDuration(clip.Duration).
		SetFilename(clip.Filename).
		SetResolution(clip.Resolution).
		SetSize(clip.Size).
		SetTitle(clip.Title).
		SetViewCount(clip.ViewCount).
		Save(context.Background())
}

func DeleteClips(ids ...int) (int, error) {
	return client.Clip.Delete().
		Where(clip.IDIn(ids...)).
		Exec(context.Background())
}
