package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RecTeam holds the schema definition for the RecTeam entity.
type RecTeam struct {
	ent.Schema
}

// Fields of the RecTeam. NOTE : Part of the public API ( ultimately exposed to end team
func (RecTeam) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(256).NotEmpty().Annotations(entgql.OrderField("name")),
		field.String("description").MaxLen(512).Optional(),
		field.UUID("leader_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the RecTeam
func (RecTeam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rec_member_edges", User.Type),
		edge.To("rec_team_job_edges", HiringJob.Type),
		edge.From("rec_leader_edge", User.Type).Ref("leader_rec_edge").Unique().Field("leader_id"),
	}
}

func (RecTeam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
