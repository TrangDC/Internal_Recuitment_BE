package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// CreatePreRequest is the resolver for the CreatePreRequest field.
func (r *mutationResolver) CreatePreRequest(ctx context.Context, input string) (string, error) {
	return "", nil
}

// CreateTeam is the resolver for the CreateTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, input ent.NewTeamInput) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().CreateTeam(ctx, input)
	if err != nil {
		return nil, err
	}
	return &ent.TeamResponse{
		Data: result,
	}, nil
}

// UpdateTeam is the resolver for the UpdateTeam field.
func (r *mutationResolver) UpdateTeam(ctx context.Context, id string, input ent.UpdateTeamInput) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().UpdateTeam(ctx, uuid.MustParse(id), input)
	if err != nil {
		return nil, err
	}
	return &ent.TeamResponse{
		Data: result,
	}, nil
}

// DeleteTeam is the resolver for the DeleteTeam field.
func (r *mutationResolver) DeleteTeam(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.Team().DeleteTeam(ctx, uuid.MustParse(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
