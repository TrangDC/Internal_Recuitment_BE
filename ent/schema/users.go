package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User. NOTE : Part of the public API ( ultimately exposed to end users
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(256).NotEmpty().Annotations(entgql.OrderField("name")),
		field.String("work_email").MaxLen(256).Annotations(entgql.OrderField("work_email")),
		field.Enum("status").Values("active", "inactive").Default("active"),
		field.String("oid").Unique().MaxLen(256),
		field.UUID("rec_team_id", uuid.UUID{}).Optional(),
		field.String("location").MaxLen(256).Optional(),
		field.UUID("hiring_team_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("audit_edge", AuditTrail.Type),
		edge.To("hiring_owner", HiringJob.Type),
		edge.To("candidate_job_feedback", CandidateJobFeedback.Type),
		edge.To("interview_edges", CandidateInterview.Type).Through("interview_users", CandidateInterviewer.Type),
		edge.To("candidate_job_edges", CandidateJob.Type),
		edge.To("candidate_interview_edges", CandidateInterview.Type),
		edge.To("candidate_reference_edges", Candidate.Type),
		edge.To("user_permission_edges", EntityPermission.Type),
		edge.To("role_edges", Role.Type).Through("role_users", UserRole.Type),
		edge.To("hiring_team_edges", HiringTeam.Type).Through("hiring_team_users", HiringTeamManager.Type),
		edge.From("member_of_hiring_team_edges", HiringTeam.Type).Ref("hiring_member_edges").Unique().Field("hiring_team_id"),
		edge.From("approvers_hiring_teams", HiringTeam.Type).Ref("approvers_users").Through("hiring_team_approvers", HiringTeamApprover.Type),
		edge.To("leader_rec_edge", RecTeam.Type).Unique(),
		edge.From("rec_teams", RecTeam.Type).Ref("rec_member_edges").Unique().Field("rec_team_id"),
		edge.To("candidate_note_edges", CandidateNote.Type),
		edge.To("candidate_history_call_edges", CandidateHistoryCall.Type),
		edge.To("approval_jobs", HiringJob.Type).Through("approval_steps", HiringJobStep.Type),
		edge.To("hiring_job_rec_edges", HiringJob.Type),
		edge.To("candidate_job_rec_edges", CandidateJob.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
