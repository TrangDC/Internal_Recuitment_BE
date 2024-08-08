package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateCertificateResolver) ID(ctx context.Context, obj *ent.CandidateCertificate) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// CandidateCertificate returns graphql1.CandidateCertificateResolver implementation.
func (r *Resolver) CandidateCertificate() graphql1.CandidateCertificateResolver {
	return &candidateCertificateResolver{r}
}

type candidateCertificateResolver struct{ *Resolver }
