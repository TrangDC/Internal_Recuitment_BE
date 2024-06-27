package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type EntitySkill struct {
	ent.Schema
}

// Fields of the EntitySkill. NOTE : Part of the public API ( ultimately exposed to end job
func (EntitySkill) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("entity_type").Values("candidate", "hiring_job"),
		field.UUID("entity_id", uuid.UUID{}).Optional(),
		field.UUID("skill_id", uuid.UUID{}).Optional(),
		field.Int("order_id").Optional(),
	}
}

// Edges of the EntitySkill
func (EntitySkill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("skill_edge", Skill.Type).Ref("entity_skill_edges").Unique().Field("skill_id"),
		edge.From("hiring_job_edge", HiringJob.Type).Ref("hiring_job_skill_edges").Unique().Field("entity_id"),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_skill_edges").Unique().Field("entity_id"),
	}
}

func (EntitySkill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
