package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TeamManager holds the schema definition for the TeamManager entity.
type TeamManager struct {
	ent.Schema
}

// Fields of the TeamManager. NOTE : Part of the public API ( ultimately exposed to end teamManagers
func (TeamManager) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entgql.OrderField("ID")),
		field.UUID("team_id", uuid.UUID{}).Annotations(),
		field.UUID("user_id", uuid.UUID{}).Annotations(),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Default(time.Now).Optional().Annotations(entgql.OrderField("UPDATED_AT")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("DELETED_AT")),
	}
}

// Edges of the TeamManager
func (TeamManager) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_edge", User.Type).Unique().Required().Field("user_id"),
		edge.To("team_edge", Team.Type).Unique().Required().Field("team_id"),
	}
}
