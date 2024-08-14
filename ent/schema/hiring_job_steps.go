package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// HiringJobStep holds the schema definition for the HiringJobStep entity.
type HiringJobStep struct {
	ent.Schema
}

// Fields of the HiringJobStep. NOTE : Part of the public API ( ultimately exposed to end job
func (HiringJobStep) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Annotations(entgql.OrderField("ID")),
		field.UUID("hiring_job_id", uuid.UUID{}).Optional(),
		field.Enum("type").Values("created", "opened", "closed").Default("created"),
		field.Time("created_at").Default(time.Now).Immutable().Annotations(entgql.OrderField("created_at")),
		field.Time("updated_at").Optional().Annotations(entgql.OrderField("updated_at")),
	}
}

// Edges of the HiringJobStep
func (HiringJobStep) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hiring_job_edge", HiringJob.Type).Ref("hiring_job_step").Unique().Field("hiring_job_id"),
	}
}
