package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateEducateResolver) ID(ctx context.Context, obj *ent.CandidateEducate) (string, error) {
	return obj.ID.String(), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateEducateResolver) Attachments(ctx context.Context, obj *ent.CandidateEducate) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// CandidateEducate returns graphql1.CandidateEducateResolver implementation.
func (r *Resolver) CandidateEducate() graphql1.CandidateEducateResolver {
	return &candidateEducateResolver{r}
}

type candidateEducateResolver struct{ *Resolver }
