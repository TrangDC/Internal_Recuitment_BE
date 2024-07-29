package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *hiringTeamResolver) ID(ctx context.Context, obj *ent.HiringTeam) (string, error) {
	return obj.ID.String(), nil
}

// Managers is the resolver for the managers field.
func (r *hiringTeamResolver) Managers(ctx context.Context, obj *ent.HiringTeam) ([]*ent.User, error) {
	return obj.Edges.UserEdges, nil
}

// OpeningRequests is the resolver for the opening_requests field.
func (r *hiringTeamResolver) OpeningRequests(ctx context.Context, obj *ent.HiringTeam) (int, error) {
	panic(fmt.Errorf("not implemented: OpeningRequests - opening_requests"))
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *hiringTeamResolver) IsAbleToDelete(ctx context.Context, obj *ent.HiringTeam) (bool, error) {
	panic(fmt.Errorf("not implemented: IsAbleToDelete - is_able_to_delete"))
}

// HiringTeam returns graphql1.HiringTeamResolver implementation.
func (r *Resolver) HiringTeam() graphql1.HiringTeamResolver { return &hiringTeamResolver{r} }

type hiringTeamResolver struct{ *Resolver }
