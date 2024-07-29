package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringTeamManager holds the schema definition for the HiringTeamManager entity.
type HiringTeamManager struct {
	ent.Schema
}

// Fields of the HiringTeamManager.
func (HiringTeamManager) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("hiring_team_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the HiringTeamManager.
func (HiringTeamManager) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_edge", User.Type).Unique().Required().Field("user_id"),
		edge.To("hiring_team_edge", HiringTeam.Type).Unique().Required().Field("hiring_team_id"),
	}
}

func (HiringTeamManager) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
