package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LikeClub holds the schema definition for the LikeClub entity.
type LikeClub struct {
	ent.Schema
}

// Fields of the LikeClub.
func (LikeClub) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Time("created_at"),
	}
}

// Edges of the LikeClub.
func (LikeClub) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("club", Club.Type).
			Ref("likes").
			Unique(),
	}
}
