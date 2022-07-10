package database

import (
	"context"
	"fmt"
	"github.com/AgileProggers/archiv-backend-go/pkg/database/internal/query"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
)

func Clips() ([]*ent.Clip, error) {
	return client.Clip.Query().All(context.Background())
}

func ClipsByQuery(params map[string][]string) ([]*ent.Clip, error) {
	queryPredicate, err := query.BuildPredicate(clip.Columns, params)
	if err != nil {
		return nil, fmt.Errorf("build query predicate: %v", err)
	}
	return client.Clip.Query().Where(queryPredicate).All(context.Background())
}

func ClipById(id int) (*ent.Clip, error) {
	return client.Clip.Get(context.Background(), id)
}

func CreateClip() *ent.ClipCreate {
	return client.Clip.Create()
}

func PatchClip(id int) *ent.ClipUpdateOne {
	return client.Clip.UpdateOneID(id)
}

func DeleteClips(ids ...int) (int, error) {
	return client.Clip.Delete().
		Where(clip.IDIn(ids...)).
		Exec(context.Background())
}
