package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateAwardResolver) ID(ctx context.Context, obj *ent.CandidateAward) (string, error) {
	return obj.ID.String(), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateAwardResolver) Attachments(ctx context.Context, obj *ent.CandidateAward) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// CandidateAward returns graphql1.CandidateAwardResolver implementation.
func (r *Resolver) CandidateAward() graphql1.CandidateAwardResolver {
	return &candidateAwardResolver{r}
}

type candidateAwardResolver struct{ *Resolver }
