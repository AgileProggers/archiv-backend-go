package database

import (
	"context"

	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/clip"
)

func Clips(c *[]ent.Clip, query ent.Clip) (err error) {
	// result := database.Where(query).Find(c)
	// if result.RowsAffected == 0 {
	// 	return errors.New("not found")
	// }
	return nil
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