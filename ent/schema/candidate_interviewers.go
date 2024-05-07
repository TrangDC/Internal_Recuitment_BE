package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateInterviewer holds the schema definition for the CandidateInterviewer entity.
type CandidateInterviewer struct {
	ent.Schema
}

// Fields of the CandidateInterviewer. NOTE : Part of the public API ( ultimately exposed to end teamManagers
func (CandidateInterviewer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_interview_id", uuid.UUID{}).Annotations(),
		field.UUID("user_id", uuid.UUID{}).Annotations(),
	}
}

// Edges of the CandidateInterviewer
func (CandidateInterviewer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_edge", User.Type).Unique().Required().Field("user_id"),
		edge.To("interview_edge", CandidateInterview.Type).Unique().Required().Field("candidate_interview_id"),
	}
}

func (CandidateInterviewer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
