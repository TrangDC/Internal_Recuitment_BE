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

// Status is the resolver for the status field.
func (r *hiringJobStepResolver) Status(ctx context.Context, obj *ent.HiringJobStep) (ent.HiringJobStepStatusEnum, error) {
	return ent.HiringJobStepStatusEnum(obj.Status), nil
}

// Approver is the resolver for the approver field.
func (r *hiringJobStepResolver) Approver(ctx context.Context, obj *ent.HiringJobStep) (*ent.User, error) {
	return obj.Edges.ApprovalUserOrErr()
}

// HiringJobStep returns graphql1.HiringJobStepResolver implementation.
func (r *Resolver) HiringJobStep() graphql1.HiringJobStepResolver { return &hiringJobStepResolver{r} }

type hiringJobStepResolver struct{ *Resolver }
