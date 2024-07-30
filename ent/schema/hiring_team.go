package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HiringTeam holds the schema definition for the HiringTeam entity.
type HiringTeam struct {
	ent.Schema
}

// Fields of the HiringTeam.
func (HiringTeam) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(256).NotEmpty().Annotations(entgql.OrderField("name")),
	}
}

// Edges of the HiringTeam.
func (HiringTeam) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_edges", User.Type).Ref("hiring_team_edges").Through("user_hiring_teams", HiringTeamManager.Type).
			Comment("The uniqueness of the user is enforced on the edge schema"),
		edge.To("hiring_team_job_edges", HiringJob.Type),
		edge.To("hiring_member_edges", User.Type),
		edge.To("approvers_users", User.Type).Through("hiring_team_approvers", HiringTeamApprover.Type),
	}
}

func (HiringTeam) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
