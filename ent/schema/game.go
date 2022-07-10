package schema

import "entgo.io/ent"

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return nil
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return nil
}
