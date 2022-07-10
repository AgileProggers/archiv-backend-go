package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/internal/annotations"
)

// Clip holds the schema definition for the User entity.
type Clip struct {
	ent.Schema
}

func (u Clip) Annotations() []schema.Annotation {
	return annotations.Generate(u)
}

// Fields of the User.
func (Clip) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Int("duration").Positive(),
		field.Time("date"),
		field.String("filename").NotEmpty(),
		field.String("resolution"),
		field.Int("size"),
		field.Int("view_count"),
	}
}

// Edges of the User.
func (Clip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Creator.Type).
			Ref("clips").
			Unique(),
		edge.From("vod", Vod.Type).
			Ref("clips"),
		edge.To("game", Game.Type),
	}
}
