package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateNoteResolver) ID(ctx context.Context, obj *ent.CandidateNote) (string, error) {
	return obj.ID.String(), nil
}

// Candidate is the resolver for the candidate field.
func (r *candidateNoteResolver) Candidate(ctx context.Context, obj *ent.CandidateNote) (*ent.Candidate, error) {
	return obj.Edges.CandidateEdgeOrErr()
}

// CreatedBy is the resolver for the created_by field.
func (r *candidateNoteResolver) CreatedBy(ctx context.Context, obj *ent.CandidateNote) (*ent.User, error) {
	return obj.Edges.CreatedByEdgeOrErr()
}

// Attachments is the resolver for the attachments field.
func (r *candidateNoteResolver) Attachments(ctx context.Context, obj *ent.CandidateNote) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdgesOrErr()
}

// CandidateNote returns graphql1.CandidateNoteResolver implementation.
func (r *Resolver) CandidateNote() graphql1.CandidateNoteResolver { return &candidateNoteResolver{r} }

type candidateNoteResolver struct{ *Resolver }
