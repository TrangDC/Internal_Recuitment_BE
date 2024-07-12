package schema

import (
	"entgo.io/ent"
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
		field.UUID("email_template_id", uuid.UUID{}).Optional().Annotations(),
		field.Enum("status").Values("pending", "sent", "failed").Default("pending"),
	}
}

// Edges of the OutgoingEmail
func (OutgoingEmail) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (OutgoingEmail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
