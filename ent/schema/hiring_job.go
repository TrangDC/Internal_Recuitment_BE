package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type HiringJob struct {
	ent.Schema
}

// Fields of the HiringJob. NOTE : Part of the public API ( ultimately exposed to end job
func (HiringJob) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("NAME")),
		field.String("description").MaxLen(255).NotEmpty(),
		field.Int("amount").Default(0).Positive().Annotations(entgql.OrderField("AMOUNT")),
		field.Enum("status").Values("draft", "opened", "closed").Default("opened"),
		field.UUID("created_by", uuid.UUID{}).Optional().Annotations(),
		field.UUID("team_id", uuid.UUID{}).Optional().Annotations(),
		field.Enum("location").Values("ha_noi", "ho_chi_minh", "da_nang", "japan"),
		field.Enum("salary_type").Values("range", "up_to", "negotiate", "minimum"),
		field.Int("salary_from").Default(0).Positive().Annotations(entgql.OrderField("SALARY_FROM")),
		field.Int("salary_to").Default(0).Positive().Annotations(entgql.OrderField("SALARY_TO")),
		field.Enum("currency").Values("vnd", "usd", "jpy"),
	}
}

// Edges of the HiringJob
func (HiringJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_edge", User.Type).Ref("hiring_owner").Unique().Field("created_by"),
		edge.From("team_edge", Team.Type).Ref("hiring_team").Unique().Field("team_id"),
	}
}

func (HiringJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
