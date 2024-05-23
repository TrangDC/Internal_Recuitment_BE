package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *userResolver) Status(ctx context.Context, obj *ent.User) (ent.UserStatus, error) {
	return ent.UserStatus(obj.Status), nil
}

// Team is the resolver for the team field.
func (r *userResolver) Team(ctx context.Context, obj *ent.User) (*ent.Team, error) {
	return obj.Edges.TeamEdges[0], nil
}

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
