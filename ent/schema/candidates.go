package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate. NOTE : Part of the public API ( ultimately exposed to end job
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
		field.String("email").MaxLen(255).NotEmpty().Unique().Annotations(entgql.OrderField("email")),
		field.String("phone").MaxLen(255).NotEmpty(),
		field.Time("dob").Optional(),
		field.Bool("is_blacklist").Default(false),
		field.Time("last_apply_date").Optional().Annotations(entgql.OrderField("last_apply_date")),
	}
}

// Edges of the Candidate
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("candidate_job_edges", CandidateJob.Type),
	}
}

func (Candidate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
