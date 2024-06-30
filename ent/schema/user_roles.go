package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserRole holds the schema definition for the UserRole entity.
type UserRole struct {
	ent.Schema
}

// Fields of the UserRole. NOTE : Part of the public API ( ultimately exposed to end teamManagers
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("role_id", uuid.UUID{}).Annotations(),
		field.UUID("user_id", uuid.UUID{}).Annotations(),
	}
}

// Edges of the UserRole
func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_edge", User.Type).Unique().Required().Field("user_id"),
		edge.To("role_edge", Role.Type).Unique().Required().Field("role_id"),
	}
}

func (UserRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
