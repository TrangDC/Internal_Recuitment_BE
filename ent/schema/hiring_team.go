package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HiringTeam holds the schema definition for the HiringTeam entity.
type HiringTeam struct {
	ent.Schema
}

// Fields of the HiringTeam.
func (HiringTeam) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
	}
}

// Edges of the HiringTeam.
func (HiringTeam) Edges() []ent.Edge {
	return nil
}

func (HiringTeam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
