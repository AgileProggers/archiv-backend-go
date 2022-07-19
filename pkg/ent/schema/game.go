package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int("game_id").Positive(),
		field.String("name").Unique(),
		field.String("box_art"),
	}

}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clip", Clip.Type).
			Ref("game").
			Unique(),
		edge.From("vod", Vod.Type).
			Ref("game").
			Unique(),
	}
}
