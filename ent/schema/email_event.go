package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailEvent holds the schema definition for the EmailEvent entity.
type EmailEvent struct {
	ent.Schema
}

// Fields of the EmailEvent. NOTE : Part of the public API ( ultimately exposed to end job )
func (EmailEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.Enum("module").Values("interview", "application", "job_request"),
		field.Enum("action").Values(
			"create", "update", "cancel",
			"cd_applied", "cd_interviewing", "cd_offering", "cd_failed_cv", "cd_failed_interview", "cd_offer_lost", "cd_hired",
			"close", "open", "reopen", "need_approval",
		),
		field.Text("name"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional(),
	}
}

func (EmailEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("template_edges", EmailTemplate.Type),
		edge.To("outgoing_email_edges", OutgoingEmail.Type),
	}
}
