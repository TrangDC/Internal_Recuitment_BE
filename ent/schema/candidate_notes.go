package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CandidateNotes holds the schema definition for the CandidateNotes entity.
type CandidateNote struct {
	ent.Schema
}

// Fields of the CandidateNotes.
func (CandidateNote) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("candidate_id", uuid.UUID{}),
		field.UUID("created_by_id", uuid.UUID{}),
		field.String("name").NotEmpty().MaxLen(256),
		field.Text("description").NotEmpty(),
	}
}

// Edges of the CandidateNotes.
func (CandidateNote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate_edge", Candidate.Type).Field("candidate_id").Ref("candidate_note_edges").Unique().Required(),
		edge.From("created_by_edge", User.Type).Field("created_by_id").Ref("candidate_note_edges").Unique().Required(),
		edge.To("attachment_edges", Attachment.Type),
	}
}

func (CandidateNote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}
