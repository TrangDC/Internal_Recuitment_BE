package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	"trec/ent/hiringjob"
	graphql1 "trec/graphql"

	"github.com/samber/lo"
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
	total := lo.Filter(obj.Edges.TeamJobEdges, func(record *ent.HiringJob, index int) bool {
		return record.Status == hiringjob.StatusOpened
	})
	return len(total), nil
}

// Team returns graphql1.TeamResolver implementation.
func (r *Resolver) Team() graphql1.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
