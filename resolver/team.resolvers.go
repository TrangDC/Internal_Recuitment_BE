package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *teamResolver) ID(ctx context.Context, obj *ent.Team) (string, error) {
	return obj.ID.String(), nil
}

// Members is the resolver for the members field.
func (r *teamResolver) Members(ctx context.Context, obj *ent.Team) ([]*ent.User, error) {
	return obj.Edges.UserEdges, nil
}

// OpeningRequests is the resolver for the opening_requests field.
func (r *teamResolver) OpeningRequests(ctx context.Context, obj *ent.Team) (int, error) {
	return 0, nil
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *teamResolver) IsAbleToDelete(ctx context.Context, obj *ent.Team) (bool, error) {
	return true, nil
}

// Team returns graphql1.TeamResolver implementation.
func (r *Resolver) Team() graphql1.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
