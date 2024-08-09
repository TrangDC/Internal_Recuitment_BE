package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Job holds the schema definition for the Job entity.
type CandidateHistoryCall struct {
	ent.Schema
}

// Fields of the CandidateHistoryCall. NOTE : Part of the public API ( ultimately exposed to end job
func (CandidateHistoryCall) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(256).Optional(),
		field.UUID("candidate_id", uuid.UUID{}).Optional(),
		field.String("contact_to").MaxLen(256).Optional(),
		field.String("description").MaxLen(512).Optional(),
		field.Enum("type").Values("candidate", "others").Optional(),
		field.Time("date").Optional(),
		field.Time("start_time").Optional(),
		field.Time("end_time").Optional(),
		field.UUID("created_by_id", uuid.UUID{}).Optional(),
	}
}

// Edges of the CandidateHistoryCall
func (CandidateHistoryCall) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachment_edges", Attachment.Type),
		edge.From("candidate_edge", Candidate.Type).Ref("candidate_history_call_edges").Unique().Field("candidate_id"),
		edge.From("created_by_edge", User.Type).Ref("candidate_history_call_edges").Unique().Field("created_by_id"),
	}
}

func (CandidateHistoryCall) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
