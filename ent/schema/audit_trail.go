package schema

import (
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AuditTrail holds the schema definition for the jobTitle entity.
type AuditTrail struct {
	ent.Schema
}

// Fields of the AuditTrail. NOTE : Part of the public API ( ultimately exposed to end users
func (AuditTrail) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.UUID("created_by", uuid.UUID{}).Optional().Annotations(),
		field.UUID("recordId", uuid.UUID{}).Annotations(),
		field.Enum("module").Values("teams").Annotations(),
		field.Enum("actionType").Optional().Values("create", "update", "delete").Default("create").Annotations(),
		field.String("note").Optional().MaxLen(500).Annotations(),
		field.Text("record_changes").Optional().Annotations(),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Optional().Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("DELETED_AT")),
	}
}

// Edges of the AuditTrail make AuditTrailTypeId a foreign key to the AuditTrailType table
func (AuditTrail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("created_by_edge", User.Type).Ref("audit_edge").Unique().Field("created_by"),
	}
}
