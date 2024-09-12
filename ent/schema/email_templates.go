package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type EmailTemplate struct {
	ent.Schema
}

// Fields of the EmailTemplate. NOTE : Part of the public API ( ultimately exposed to end job
func (EmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("event").Values("candidate_applied_to_kiv", "candidate_interviewing_to_kiv", "candidate_interviewing_to_offering", "created_interview", "updating_interview", "cancel_interview"),
		field.JSON("send_to", emailSendTos),
		field.JSON("cc", []string{}),
		field.JSON("bcc", []string{}),
		field.String("subject").MaxLen(256).NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Text("signature").Optional(),
		field.Enum("status").Values("active", "inactive").Default("active"),
		field.UUID("event_id", uuid.UUID{}),
	}
}

// Edges of the EmailTemplate
func (EmailTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role_edges", Role.Type).Ref("email_template_edges").Through("role_email_templates", EmailRoleAttribute.Type),
		edge.From("event_edge", EmailEvent.Type).Ref("template_edges").Unique().Field("event_id").Required(),
	}
}

func (EmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
