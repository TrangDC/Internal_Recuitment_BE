package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment. NOTE : Part of the public API ( ultimately exposed to end job
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("document_id", uuid.UUID{}).Unique(),
		field.String("document_name").MaxLen(255),
		field.Enum("relation_type").Values("candidate_jobs", "candidate_job_feedbacks", "candidates"),
		field.UUID("relation_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the Attachment
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_job_edge", CandidateJob.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_job_feedback_edge", CandidateJobFeedback.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_interview_edge", CandidateInterview.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_edge", Candidate.Type).Ref("attachment_edges").Unique().Field("relation_id"),
	}
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
