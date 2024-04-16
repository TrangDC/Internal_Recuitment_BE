package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
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

// GetAllUsers is the resolver for the GetAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.UserOrder, freeWord *ent.UserFreeWord, filter *ent.UserFilter) (*ent.UserResponseGetAll, error) {
	result, err := r.serviceRegistry.User().GetUsers(ctx, pagination, orderBy)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetHiringJob is the resolver for the GetHiringJob field.
func (r *queryResolver) GetHiringJob(ctx context.Context, id string) (*ent.HiringJobResponse, error) {
	panic(fmt.Errorf("not implemented: GetHiringJob - GetHiringJob"))
}

// GetAllHiringJobs is the resolver for the GetAllHiringJobs field.
func (r *queryResolver) GetAllHiringJobs(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringJobFilter, freeWord *ent.HiringJobFreeWord, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error) {
	panic(fmt.Errorf("not implemented: GetAllHiringJobs - GetAllHiringJobs"))
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
