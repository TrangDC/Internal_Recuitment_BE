package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Role struct {
	ent.Schema
}

// Fields of the Role. NOTE : Part of the public API ( ultimately exposed to end job
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entgql.OrderField("name")),
		field.String("description").Annotations(entgql.OrderField("DESCRIPTION")),
	}
}

// Edges of the Role
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role_permission_edges", EntityPermission.Type),
		edge.From("user_edges", User.Type).Ref("role_edges").Through("user_roles", UserRole.Type),
		edge.To("email_template_edges", EmailTemplate.Type).Through("email_template_roles", EmailRoleAttribute.Type),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
