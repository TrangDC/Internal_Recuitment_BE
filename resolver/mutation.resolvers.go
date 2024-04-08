package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	graphql1 "trec/graphql"
)

// CreatePreRequest is the resolver for the CreatePreRequest field.
func (r *mutationResolver) CreatePreRequest(ctx context.Context, input string) (string, error) {
	return "", nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
