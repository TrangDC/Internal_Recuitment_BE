package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateEducate struct {
	ent.Schema
}

// Fields of the CandidateEducate. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateEducate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.String("school_name").MaxLen(256).NotEmpty(),
		field.String("major").MaxLen(256).Optional(),
		field.String("gpa").MaxLen(256).Optional(),
		field.String("location").MaxLen(256).Optional(),
		field.String("description").MaxLen(512).Optional(),
		field.Time("start_date").Optional(),
		field.Time("end_date").Optional(),
		field.Int("order_id").Optional(),
		field.Bool("is_current").Default(false),
	}
}

// Edges of the CandidateEducate
func (CandidateEducate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment_edges", Attachment.Type),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_educate_edges").Unique().Field("candidate_id"),
	}
}

func (CandidateEducate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
