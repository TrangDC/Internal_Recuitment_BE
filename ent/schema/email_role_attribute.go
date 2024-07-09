package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailRoleAttribute holds the schema definition for the EmailRoleAttribute entity.
type EmailRoleAttribute struct {
	ent.Schema
}

// Fields of the EmailRoleAttribute. NOTE : Part of the public API ( ultimately exposed to end teamManagers
func (EmailRoleAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("email_template_id", uuid.UUID{}).Annotations(),
		field.UUID("role_id", uuid.UUID{}).Annotations(),
	}
}

// Edges of the EmailRoleAttribute
func (EmailRoleAttribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("email_template_edge", Team.Type).Unique().Required().Field("email_template_id"),
		edge.To("role_edge", Role.Type).Unique().Required().Field("role_id"),
	}
}

func (EmailRoleAttribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
