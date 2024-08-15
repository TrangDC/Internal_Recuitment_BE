package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *hiringJobStepResolver) ID(ctx context.Context, obj *ent.HiringJobStep) (string, error) {
	return obj.ID.String(), nil
}

// Type is the resolver for the type field.
func (r *hiringJobStepResolver) Type(ctx context.Context, obj *ent.HiringJobStep) (ent.HiringJobStepTypeEnum, error) {
	return ent.HiringJobStepTypeEnum(obj.Type), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *hiringJobStepResolver) CreatedBy(ctx context.Context, obj *ent.HiringJobStep) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}

// HiringJobStep returns graphql1.HiringJobStepResolver implementation.
func (r *Resolver) HiringJobStep() graphql1.HiringJobStepResolver { return &hiringJobStepResolver{r} }

type hiringJobStepResolver struct{ *Resolver }
