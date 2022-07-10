package database

import (
	"context"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/creator"
)

/*
func Creators(c *[]ent.Creator, query Creator) (err error) {
	result := database.Where(query).Find(c)
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
*/

func CreatorById(id int) (*ent.Creator, error) {
	return client.Creator.Get(context.Background(), id)
}

func CreateCreator() *ent.CreatorCreate {
	return client.Creator.Create()
}

func PatchCreator(id int) *ent.CreatorUpdateOne {
	return client.Creator.UpdateOneID(id)
}

func DeleteCreators(ids ...int) (int, error) {
	return client.Creator.Delete().
		Where(creator.IDIn(ids...)).
		Exec(context.Background())
}
