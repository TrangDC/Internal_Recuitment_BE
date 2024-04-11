package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// GetPreRequest is the resolver for the GetPreRequest field.
func (r *queryResolver) GetPreRequest(ctx context.Context) (string, error) {
	return "", nil
}

// GetTeam is the resolver for the GetTeam field.
func (r *queryResolver) GetTeam(ctx context.Context, id string) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().GetTeam(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return &ent.TeamResponse{
		Data: result,
	}, nil
}

// GetAllTeams is the resolver for the GetAllTeams field.
func (r *queryResolver) GetAllTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.TeamFilter, freeWord *ent.TeamFreeWord, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error) {
	result, err := r.serviceRegistry.Team().GetTeams(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
