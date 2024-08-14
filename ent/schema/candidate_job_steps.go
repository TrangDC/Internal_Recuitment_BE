package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateJobStep struct {
	ent.Schema
}

// Fields of the CandidateJobStep. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateJobStep) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("candidate_job_status").Values("applied", "interviewing", "offering", "hired", "failed_cv", "failed_interview", "offer_lost", "ex_staff").Default("applied"),
		field.UUID("candidate_job_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the CandidateJobStep
func (CandidateJobStep) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_job_edge", CandidateJob.Type).Ref("candidate_job_step").Unique().Field("candidate_job_id"),
	}
}

func (CandidateJobStep) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
