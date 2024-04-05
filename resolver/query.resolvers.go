package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// GetPreRequest is the resolver for the GetPreRequest field.
func (r *queryResolver) GetPreRequest(ctx context.Context, id string) (*ent.PreResponse, error) {
	result, err := r.serviceRegistry.Pre().PreFunction(ctx)
	if err != nil {
		return nil, err
	}
	data := ent.Pre{
		ID:           id,
		StringOutput: result,
	}
	return &ent.PreResponse{
		Data: &data,
	}, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
