package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type EmailTemplate struct {
	ent.Schema
}

// Fields of the EmailTemplate. NOTE : Part of the public API ( ultimately exposed to end job
func (EmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("event").Values("candidate_applied_to_kiv", "candidate_interviewing_to_kiv", "candidate_interviewing_to_offering", "created_candidate", "updating_interview", "cancel_interview"),
		field.JSON("send_to", []string{"interviewer", "job_request", "team_manager", "team_member", "role", "candidate"}),
		field.JSON("cc", []string{}),
		field.JSON("bcc", []string{}),
		field.String("subject").MaxLen(255).NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Text("signature").Optional(),
		field.Enum("status").Values("active", "inactive").Default("active"),
	}
}

// Edges of the EmailTemplate
func (EmailTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role_edges", Role.Type).Ref("email_template_edges").Through("role_email_templates", EmailRoleAttribute.Type),
	}
}

func (EmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
