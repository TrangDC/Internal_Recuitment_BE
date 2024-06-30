package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type EntityPermission struct {
	ent.Schema
}

// Fields of the EntityPermission. NOTE : Part of the public API ( ultimately exposed to end job
func (EntityPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("entity_id", uuid.UUID{}).Optional(),
		field.UUID("permission_id", uuid.UUID{}).Optional(),
		field.Bool("for_owner").Default(false),
		field.Bool("for_team").Default(false),
		field.Bool("for_all").Default(false),
		field.Enum("entity_type").Values("user", "role").Optional(),
	}
}

// Edges of the EntityPermission
func (EntityPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("permission_edges", Permission.Type).Ref("user_permission_edge").Unique().Field("permission_id"),
		edge.From("user_edge", User.Type).Ref("user_permission_edges").Unique().Field("entity_id"),
		edge.From("role_edge", Role.Type).Ref("role_permission_edges").Unique().Field("entity_id"),
	}
}

func (EntityPermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
