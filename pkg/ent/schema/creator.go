package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent/internal/annotations"
)

// Creator holds the schema definition for the User entity.
type Creator struct {
	ent.Schema
}

func (u Creator) Annotations() []schema.Annotation {
	return annotations.Generate(u)
}

// Fields of the User.
func (Creator) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the User.
func (Creator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clips", Clip.Type),
		edge.To("vods", Vod.Type),
	}
}
