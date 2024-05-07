package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User. NOTE : Part of the public API ( ultimately exposed to end users
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("work_email").MaxLen(255).Annotations(entgql.OrderField("WORK_EMAIL")),
		field.String("oid").Unique().MaxLen(255),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("audit_edge", AuditTrail.Type),
		edge.To("hiring_owner", HiringJob.Type),
		edge.To("team_edges", Team.Type).Through("team_users", TeamManager.Type),
		edge.To("candidate_job_feedback", CandidateJobFeedback.Type),
		edge.To("interview_edges", CandidateInterview.Type).Through("interview_users", CandidateInterviewer.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
