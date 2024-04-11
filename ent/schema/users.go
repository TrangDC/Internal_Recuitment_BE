package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User. NOTE : Part of the public API ( ultimately exposed to end users
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("work_email").MaxLen(255).Annotations(entgql.OrderField("WORK_EMAIL")),
		field.String("oid").Unique().MaxLen(255),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).Optional().Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("DELETED_AT")),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("audit_edge", AuditTrail.Type),
		edge.To("team_edges", Team.Type).Through("team_users", TeamManager.Type),
	}
}
