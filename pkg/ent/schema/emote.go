package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Emote holds the schema definition for the Emote entity.
type Emote struct {
	ent.Schema
}

// Fields of the Emote.
func (Emote) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("url").
			Unique(),
	}
}

// Edges of the Emote.
func (Emote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("provider", Provider.Type).
			Ref("emotes").
			Unique(),
	}
}
