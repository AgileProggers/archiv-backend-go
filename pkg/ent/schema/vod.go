package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vod holds the schema definition for the Vod entity.
type Vod struct {
	ent.Schema
}

// Fields of the Vod.
func (Vod) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Int("duration"),
		field.Time("date"),
		field.String("filename"),
		field.String("resolution"),
		field.Float("fps"),
		field.Int("size"),
		field.Bool("publish"),
	}
}

// Edges of the Vod.
func (Vod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clips", Clip.Type),
		edge.To("game", Game.Type),
	}
}
