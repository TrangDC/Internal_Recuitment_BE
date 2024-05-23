package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateJob struct {
	ent.Schema
}

// Fields of the CandidateJob. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateJob) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("hiring_job_id", uuid.UUID{}).Optional(),
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.UUID("created_by", uuid.UUID{}).Optional(),
		field.Enum("status").Values("applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff").Default("applied"),
		field.JSON("failed_reason", []string{"poor_professionalism",
			"poor_fit_and_engagement", "over_expectations", "over_qualification", "language_deficiency",
			"weak_technical_skills", "poor_interpersonal_skills", "poor_problem_solving_skills", "poor_management_skills",
			"candidate_withdrawal", "others"}).Optional(),
	}
}

// Edges of the CandidateJob
func (CandidateJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment_edges", Attachment.Type),
		edge.From("hiring_job_edge", HiringJob.Type).Ref("candidate_job_edges").Unique().Field("hiring_job_id"),
		edge.To("candidate_job_feedback", CandidateJobFeedback.Type),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_job_edges").Unique().Field("candidate_id"),
		edge.To("candidate_job_interview", CandidateInterview.Type),
		edge.From("created_by_edge", User.Type).Ref("candidate_job_edges").Unique().Field("created_by"),
	}
}

func (CandidateJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
