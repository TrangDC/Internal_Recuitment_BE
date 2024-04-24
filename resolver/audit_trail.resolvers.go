package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *auditTrailResolver) ID(ctx context.Context, obj *ent.AuditTrail) (string, error) {
	return obj.ID.String(), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *auditTrailResolver) CreatedBy(ctx context.Context, obj *ent.AuditTrail) (string, error) {
	return obj.CreatedBy.String(), nil
}

// CreatedInfo is the resolver for the createdInfo field.
func (r *auditTrailResolver) CreatedInfo(ctx context.Context, obj *ent.AuditTrail) (*ent.User, error) {
	return obj.Edges.UserEdge, nil
}

// RecordID is the resolver for the recordId field.
func (r *auditTrailResolver) RecordID(ctx context.Context, obj *ent.AuditTrail) (string, error) {
	return obj.RecordId.String(), nil
}

// Module is the resolver for the module field.
func (r *auditTrailResolver) Module(ctx context.Context, obj *ent.AuditTrail) (ent.ProjectModule, error) {
	return ent.ProjectModule(obj.Module), nil
}

// ActionType is the resolver for the actionType field.
func (r *auditTrailResolver) ActionType(ctx context.Context, obj *ent.AuditTrail) (ent.AuditTrailAction, error) {
	return ent.AuditTrailAction(obj.ActionType), nil
}

// AuditTrail returns graphql1.AuditTrailResolver implementation.
func (r *Resolver) AuditTrail() graphql1.AuditTrailResolver { return &auditTrailResolver{r} }

type auditTrailResolver struct{ *Resolver }
