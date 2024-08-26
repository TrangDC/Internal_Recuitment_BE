package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringTeamApprover holds the schema definition for the HiringTeamApprover entity.
type HiringTeamApprover struct {
	ent.Schema
}

// Fields of the HiringTeamApprover.
func (HiringTeamApprover) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("hiring_team_id", uuid.UUID{}),
		field.Int("order_id").Positive(),
	}
}

// Edges of the HiringTeamApprover.
func (HiringTeamApprover) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_edge", User.Type).Unique().Required().Field("user_id"),
		edge.To("hiring_team_edge", HiringTeam.Type).Unique().Required().Field("hiring_team_id"),
	}
}

func (HiringTeamApprover) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
