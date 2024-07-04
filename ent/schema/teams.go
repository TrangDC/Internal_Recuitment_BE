package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team. NOTE : Part of the public API ( ultimately exposed to end team
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
	}
}

// Edges of the Team
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_edges", User.Type).Ref("team_edges").Through("user_teams", TeamManager.Type).
			Comment("The uniqueness of the user is enforced on the edge schema"),
		edge.To("team_job_edges", HiringJob.Type),
		edge.To("member_edges", User.Type),
	}
}

func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
