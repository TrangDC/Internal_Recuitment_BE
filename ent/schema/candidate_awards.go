package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateAward struct {
	ent.Schema
}

// Fields of the CandidateAward. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateAward) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.String("name").MaxLen(256).NotEmpty(),
		field.Time("achieved_date").Optional(),
	}
}

// Edges of the CandidateAward
func (CandidateAward) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment_edges", Attachment.Type),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_award_edges").Unique().Field("candidate_id"),
	}
}

func (CandidateAward) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
