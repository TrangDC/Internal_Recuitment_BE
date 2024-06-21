package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SkillType struct {
	ent.Schema
}

func (SkillType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
		field.Text("description").MaxLen(255).Optional().Annotations(entgql.OrderField("description")),
	}
}

func (SkillType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
