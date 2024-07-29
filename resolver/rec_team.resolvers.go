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
func (r *recTeamResolver) ID(ctx context.Context, obj *ent.RecTeam) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Leader is the resolver for the leader field.
func (r *recTeamResolver) Leader(ctx context.Context, obj *ent.RecTeam) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: Leader - leader"))
}

// RecTeam returns graphql1.RecTeamResolver implementation.
func (r *Resolver) RecTeam() graphql1.RecTeamResolver { return &recTeamResolver{r} }

type recTeamResolver struct{ *Resolver }
