package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateResolver) ID(ctx context.Context, obj *ent.Candidate) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *candidateResolver) Status(ctx context.Context, obj *ent.Candidate) (ent.CandidateStatusEnum, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateStatus(ctx, obj.ID), nil
}

// Candidate returns graphql1.CandidateResolver implementation.
func (r *Resolver) Candidate() graphql1.CandidateResolver { return &candidateResolver{r} }

type candidateResolver struct{ *Resolver }
