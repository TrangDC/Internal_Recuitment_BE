package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *entityPermissionResolver) ID(ctx context.Context, obj *ent.EntityPermission) (string, error) {
	return obj.ID.String(), nil
}

// Permission is the resolver for the permission field.
func (r *entityPermissionResolver) Permission(ctx context.Context, obj *ent.EntityPermission) (*ent.Permission, error) {
	return obj.Edges.PermissionEdges, nil
}

// EntityPermission returns graphql1.EntityPermissionResolver implementation.
func (r *Resolver) EntityPermission() graphql1.EntityPermissionResolver {
	return &entityPermissionResolver{r}
}

type entityPermissionResolver struct{ *Resolver }
