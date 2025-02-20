package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type JobPosition struct {
	ent.Schema
}

func (JobPosition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(256).NotEmpty().Annotations(entgql.OrderField("name")),
		field.Text("description").MaxLen(512).Optional().Annotations(entgql.OrderField("description")),
	}
}

func (JobPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hiring_job_position_edges", HiringJob.Type),
	}
}

func (JobPosition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
