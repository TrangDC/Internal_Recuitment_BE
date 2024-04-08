package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	graphql1 "trec/graphql"
)

// GetPreRequest is the resolver for the GetPreRequest field.
func (r *queryResolver) GetPreRequest(ctx context.Context) (string, error) {
	return "", nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
