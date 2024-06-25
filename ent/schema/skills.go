package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Skill struct {
	ent.Schema
}

func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
		field.Text("description").MaxLen(255).Optional().Annotations(entgql.OrderField("description")),
		field.UUID("skill_type_id", uuid.UUID{}).Optional(),
	}
}

func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("skill_type_edge", SkillType.Type).Ref("skill_edges").Unique().Field("skill_type_id"),
	}
}

func (Skill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
