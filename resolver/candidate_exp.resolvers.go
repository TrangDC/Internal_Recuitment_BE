package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateExpResolver) ID(ctx context.Context, obj *ent.CandidateExp) (string, error) {
	return obj.ID.String(), nil
}

// CandidateExp returns graphql1.CandidateExpResolver implementation.
func (r *Resolver) CandidateExp() graphql1.CandidateExpResolver { return &candidateExpResolver{r} }

type candidateExpResolver struct{ *Resolver }
