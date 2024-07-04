package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *roleResolver) ID(ctx context.Context, obj *ent.Role) (string, error) {
	return obj.ID.String(), nil
}

// EntityPermissions is the resolver for the entity_permissions field.
func (r *roleResolver) EntityPermissions(ctx context.Context, obj *ent.Role) ([]*ent.EntityPermission, error) {
	return obj.Edges.RolePermissionEdges, nil
}

// Role returns graphql1.RoleResolver implementation.
func (r *Resolver) Role() graphql1.RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }
