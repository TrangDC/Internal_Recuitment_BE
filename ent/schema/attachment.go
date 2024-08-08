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
		field.String("document_name").MaxLen(256),
		field.Enum("relation_type").Values("candidate_jobs", "candidate_job_feedbacks", "candidates", "candidate_educates", "candidate_awards", "candidate_certificates", "candidate_notes"),
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
		edge.From("candidate_educate_edge", CandidateEducate.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_award_edge", CandidateAward.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_certificate_edge", CandidateCertificate.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_history_call_edge", CandidateHistoryCall.Type).Ref("attachment_edges").Unique().Field("relation_id"),
		edge.From("candidate_note_edge", CandidateNote.Type).Ref("attachment_edges").Unique().Field("relation_id"),
	}
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
