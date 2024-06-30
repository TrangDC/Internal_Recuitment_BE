package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type PermissionGroup struct {
	ent.Schema
}

// Fields of the PermissionGroup. NOTE : Part of the public API ( ultimately exposed to end job
func (PermissionGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Annotations(entgql.OrderField("title")),
		field.UUID("parent_id", uuid.UUID{}).Optional(),
		field.Enum("group_type").Values("function", "system").Default("function"),
		field.Int("order_id"),
	}
}

// Edges of the PermissionGroup
func (PermissionGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group_permission_parent", PermissionGroup.Type).Ref("group_permission_children").Unique().Field("parent_id"),
		edge.To("group_permission_children", PermissionGroup.Type),
		edge.To("permission_edges", Permission.Type),
	}
}

func (PermissionGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
