package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate. NOTE : Part of the public API ( ultimately exposed to end job
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Annotations(entgql.OrderField("name")),
		field.String("email").MaxLen(255).NotEmpty().Unique().Annotations(entgql.OrderField("email")),
		field.String("phone").MaxLen(255).NotEmpty(),
		field.Time("dob").Optional(),
		field.Bool("is_blacklist").Default(false),
		field.Time("last_apply_date").Optional().Annotations(entgql.OrderField("last_apply_date")),
		field.Enum("reference_type").Values("eb", "rec", "hiring_platform", "reference", "headhunt").Default("eb"),
		field.String("reference_value").MaxLen(255).Optional(),
		field.UUID("reference_uid", uuid.UUID{}).Optional(),
		field.Time("recruit_time").Optional().Annotations(entgql.OrderField("recruit_time")),
		field.String("description").MaxLen(255).Optional(),
		field.String("country").MaxLen(255).Optional(),
		field.UUID("attachment_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the Candidate
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("candidate_job_edges", CandidateJob.Type),
		edge.From("reference_user_edge", User.Type).Ref("candidate_reference_edges").Unique().Field("reference_uid"),
	}
}

func (Candidate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
