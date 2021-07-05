package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("summary"),
		field.String("description"),
		field.Strings("hashtags"),
		field.Strings("images").Optional(),
		field.String("homepage").Optional(),
		field.String("instagram").Optional(),
		field.String("facebook").Optional(),
		field.String("phone").Optional(),
		field.String("email").Optional(),
		field.Bool("recommended").Default(false),
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("likes", LikeClub.Type),
	}
}
