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
		field.String("name").MaxLen(256).NotEmpty().Annotations(entgql.OrderField("name")),
		field.Text("description").NotEmpty(),
		field.Int("amount").Default(0).Annotations(entgql.OrderField("amount")),
		field.Enum("status").Values("draft", "opened", "closed").Default("opened"),
		field.UUID("created_by", uuid.UUID{}).Optional().Annotations(),
		field.Enum("location").Values("ha_noi", "ho_chi_minh", "da_nang", "japan", "singapore"),
		field.Enum("salary_type").Values("range", "up_to", "negotiate", "minimum"),
		field.Int("salary_from").Default(0).Annotations(entgql.OrderField("salary_from")),
		field.Int("salary_to").Default(0).Annotations(entgql.OrderField("salary_to")),
		field.Enum("currency").Values("vnd", "usd", "jpy"),
		field.Time("last_apply_date").Optional().Annotations(entgql.OrderField("last_apply_date")),
		field.Int("priority").Default(4).Annotations(entgql.OrderField("priority")),
		field.UUID("hiring_team_id", uuid.UUID{}).Optional(),
		field.UUID("job_position_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the HiringJob
func (HiringJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_edge", User.Type).Ref("hiring_owner").Unique().Field("created_by"),
		edge.To("candidate_job_edges", CandidateJob.Type),
		edge.To("hiring_job_skill_edges", EntitySkill.Type),
		edge.From("hiring_team_edge", HiringTeam.Type).Ref("hiring_team_job_edges").Unique().Field("hiring_team_id"),
		edge.From("job_position_edge", JobPosition.Type).Ref("hiring_job_position_edges").Unique().Field("job_position_id"),
		edge.From("approval_users", User.Type).Ref("approval_jobs").Through("approval_steps", HiringJobStep.Type),
	}
}

func (HiringJob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		SlugMixin{},
	}
}
