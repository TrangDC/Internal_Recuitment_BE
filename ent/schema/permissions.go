package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission. NOTE : Part of the public API ( ultimately exposed to end job
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Annotations(entgql.OrderField("title")),
		field.UUID("group_id", uuid.UUID{}).Optional(),
		field.Bool("for_owner").Default(false),
		field.Bool("for_team").Default(false),
		field.Bool("for_all").Default(false),
		field.String("operation_name").Optional(),
		field.UUID("parent_id", uuid.UUID{}).Optional(),
		field.Int("order_id"),
	}
}

// Edges of the Permission
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group_permission_edge", PermissionGroup.Type).Ref("permission_edges").Unique().Field("group_id"),
		edge.To("user_permission_edge", EntityPermission.Type),
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
