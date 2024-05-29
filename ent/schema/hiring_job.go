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
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
		field.Text("description").NotEmpty(),
		field.Int("amount").Default(0).Annotations(entgql.OrderField("amount")),
		field.Enum("status").Values("draft", "opened", "closed").Default("opened"),
		field.UUID("created_by", uuid.UUID{}).Optional().Annotations(),
		field.UUID("team_id", uuid.UUID{}).Optional().Annotations(),
		field.Enum("location").Values("ha_noi", "ho_chi_minh", "da_nang", "japan"),
		field.Enum("salary_type").Values("range", "up_to", "negotiate", "minimum"),
		field.Int("salary_from").Default(0).Annotations(entgql.OrderField("salary_from")),
		field.Int("salary_to").Default(0).Annotations(entgql.OrderField("salary_to")),
		field.Enum("currency").Values("vnd", "usd", "jpy"),
		field.Time("last_apply_date").Optional().Annotations(entgql.OrderField("last_apply_date")),
		field.Int("priority").Default(4).Annotations(entgql.OrderField("priority")),
	}
}

// Edges of the HiringJob
func (HiringJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_edge", User.Type).Ref("hiring_owner").Unique().Field("created_by"),
		edge.From("team_edge", Team.Type).Ref("team_job_edges").Unique().Field("team_id"),
		edge.To("candidate_job_edges", CandidateJob.Type),
	}
}

func (HiringJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
