package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateJobFeedback struct {
	ent.Schema
}

// Fields of the CandidateJobFeedback. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateJobFeedback) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_job_id", uuid.UUID{}).Optional(),
		field.UUID("created_by", uuid.UUID{}).Optional(),
		field.Enum("candidate_job_status").Values("applied", "interviewing", "offering", "hired", "failed_cv", "failed_interview", "offer_lost", "ex_staff").Default("applied"),
		field.Text("feedback"),
	}
}

// Edges of the CandidateJobFeedback
func (CandidateJobFeedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("created_by_edge", User.Type).Ref("candidate_job_feedback").Unique().Field("created_by"),
		edge.From("candidate_job_edge", CandidateJob.Type).Ref("candidate_job_feedback").Unique().Field("candidate_job_id"),
		edge.To("attachment_edges", Attachment.Type),
	}
}

func (CandidateJobFeedback) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
