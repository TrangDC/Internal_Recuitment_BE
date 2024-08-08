package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateCertificate struct {
	ent.Schema
}

// Fields of the CandidateCertificate. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateCertificate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.String("name").MaxLen(256).NotEmpty(),
		field.String("score").MaxLen(256).Optional(),
		field.Time("achieved_date").Optional(),
	}
}

// Edges of the CandidateCertificate
func (CandidateCertificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment_edges", Attachment.Type),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_certificate_edges").Unique().Field("candidate_id"),
	}
}

func (CandidateCertificate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
