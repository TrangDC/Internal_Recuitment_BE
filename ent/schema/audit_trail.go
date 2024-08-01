package schema

import (
	"entgo.io/ent/schema/edge"

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
		field.UUID("created_by", uuid.UUID{}).Optional().Annotations(),
		field.UUID("recordId", uuid.UUID{}).Annotations(),
		field.Enum("module").Values("hiring_teams", "job_positions", "hiring_jobs", "candidates", "skills", "users", "skill_types", "roles", "email_templates", "rec_teams").Annotations(),
		field.Enum("actionType").Optional().Values("create", "update", "delete").Default("create").Annotations(),
		field.String("note").Optional().MaxLen(500).Annotations(),
		field.Text("record_changes").Optional().Annotations(),
	}
}

// Edges of the AuditTrail make AuditTrailTypeId a foreign key to the AuditTrailType table
func (AuditTrail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_edge", User.Type).Ref("audit_edge").Unique().Field("created_by"),
	}
}

func (AuditTrail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
