package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type OutgoingEmail struct {
	ent.Schema
}

// Fields of the OutgoingEmail. NOTE : Part of the public API ( ultimately exposed to end job
func (OutgoingEmail) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("to", []string{}),
		field.JSON("cc", []string{}),
		field.JSON("bcc", []string{}),
		field.Text("subject").NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Text("signature"),
		field.UUID("candidate_id", uuid.UUID{}).Annotations().Optional(),
		field.Enum("recipient_type").Values(emailSendTos...),
		field.UUID("email_template_id", uuid.UUID{}).Optional().Annotations(),
		field.Enum("status").Values("pending", "sent", "failed").Default("pending"),
		field.Enum("event").
			Values("candidate_applied_to_kiv", "candidate_interviewing_to_kiv", "candidate_interviewing_to_offering", "created_interview", "updating_interview", "cancel_interview").
			Immutable(),
	}
}

// Edges of the OutgoingEmail
func (OutgoingEmail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_edge", Candidate.Type).Ref("outgoing_email_edges").Unique().Field("candidate_id"),
	}
}

func (OutgoingEmail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
