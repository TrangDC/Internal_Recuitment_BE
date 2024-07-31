package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *jobPositionResolver) ID(ctx context.Context, obj *ent.JobPosition) (string, error) {
	return obj.ID.String(), nil
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *jobPositionResolver) IsAbleToDelete(ctx context.Context, obj *ent.JobPosition) (bool, error) {
	if len(obj.Edges.HiringJobPositionEdges) > 0 {
		return false, nil
	}
	return true, nil
}

// JobPosition returns graphql1.JobPositionResolver implementation.
func (r *Resolver) JobPosition() graphql1.JobPositionResolver { return &jobPositionResolver{r} }

type jobPositionResolver struct{ *Resolver }
