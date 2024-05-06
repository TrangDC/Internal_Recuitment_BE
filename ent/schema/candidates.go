package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate. NOTE : Part of the public API ( ultimately exposed to end job
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("email").MaxLen(255).NotEmpty().Unique(),
		field.String("phone").MaxLen(255).NotEmpty(),
		field.Time("dob"),
		field.Bool("is_blacklist").Default(false),
	}
}

// Edges of the Candidate
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Candidate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
