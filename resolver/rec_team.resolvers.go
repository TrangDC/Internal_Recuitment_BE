package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *recTeamResolver) ID(ctx context.Context, obj *ent.RecTeam) (string, error) {
	return obj.ID.String(), nil
}

// LeaderID is the resolver for the leader_id field.
func (r *recTeamResolver) LeaderID(ctx context.Context, obj *ent.RecTeam) (string, error) {
	return obj.LeaderID.String(), nil
}

// Leader is the resolver for the leader field.
func (r *recTeamResolver) Leader(ctx context.Context, obj *ent.RecTeam) (*ent.User, error) {
	return obj.Edges.RecLeaderEdge, nil
}

// RecTeam returns graphql1.RecTeamResolver implementation.
func (r *Resolver) RecTeam() graphql1.RecTeamResolver { return &recTeamResolver{r} }

type recTeamResolver struct{ *Resolver }
