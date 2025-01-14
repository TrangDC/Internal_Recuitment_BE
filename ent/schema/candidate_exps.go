package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateExp struct {
	ent.Schema
}

// Fields of the CandidateExp. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateExp) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.Text("position").NotEmpty(),
		field.Text("company").NotEmpty(),
		field.Text("location").Optional(),
		field.Text("description").Optional(),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		field.Int("order_id").Optional(),
		field.Bool("is_current").Default(false),
	}
}

// Edges of the CandidateExp
func (CandidateExp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_exp_edges").Unique().Field("candidate_id"),
	}
}

func (CandidateExp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
