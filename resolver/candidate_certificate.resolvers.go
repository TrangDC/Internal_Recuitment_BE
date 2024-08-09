package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateCertificateResolver) ID(ctx context.Context, obj *ent.CandidateCertificate) (string, error) {
	return obj.ID.String(), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateCertificateResolver) Attachments(ctx context.Context, obj *ent.CandidateCertificate) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// CandidateCertificate returns graphql1.CandidateCertificateResolver implementation.
func (r *Resolver) CandidateCertificate() graphql1.CandidateCertificateResolver {
	return &candidateCertificateResolver{r}
}

type candidateCertificateResolver struct{ *Resolver }
