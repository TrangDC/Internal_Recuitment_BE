package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *hiringTeamApproverResolver) ID(ctx context.Context, obj *ent.HiringTeamApprover) (string, error) {
	return obj.ID.String(), nil
}

// UserID is the resolver for the user_id field.
func (r *hiringTeamApproverResolver) UserID(ctx context.Context, obj *ent.HiringTeamApprover) (string, error) {
	return obj.UserID.String(), nil
}

// User is the resolver for the user field.
func (r *hiringTeamApproverResolver) User(ctx context.Context, obj *ent.HiringTeamApprover) (*ent.User, error) {
	return obj.Edges.UserEdgeOrErr()
}

// HiringTeamID is the resolver for the hiring_team_id field.
func (r *hiringTeamApproverResolver) HiringTeamID(ctx context.Context, obj *ent.HiringTeamApprover) (string, error) {
	return obj.HiringTeamID.String(), nil
}

// HiringTeamApprover returns graphql1.HiringTeamApproverResolver implementation.
func (r *Resolver) HiringTeamApprover() graphql1.HiringTeamApproverResolver {
	return &hiringTeamApproverResolver{r}
}

type hiringTeamApproverResolver struct{ *Resolver }
