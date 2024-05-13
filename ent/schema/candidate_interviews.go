package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateInterview struct {
	ent.Schema
}

// Fields of the CandidateInterview. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateInterview) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255).Annotations(entgql.OrderField("title")),
		field.Enum("candidate_job_status").Values("applied", "interviewing", "offering", "hired", "kiv", "offer_lost", "ex_staff").Default("applied"),
		field.UUID("candidate_job_id", uuid.UUID{}).Optional(),
		field.Time("interview_date").Optional(),
		field.Time("start_from").Optional(),
		field.Time("end_at").Optional(),
		field.Text("description"),
	}
}

// Edges of the CandidateInterview
func (CandidateInterview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_job_edge", CandidateJob.Type).Ref("candidate_job_interview").Unique().Field("candidate_job_id"),
		edge.To("attachment_edges", Attachment.Type),
		edge.From("interviewer_edges", User.Type).Ref("interview_edges").Through("user_interviewers", CandidateInterviewer.Type),
	}
}

func (CandidateInterview) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
