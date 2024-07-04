package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *permissionGroupResolver) ID(ctx context.Context, obj *ent.PermissionGroup) (string, error) {
	return obj.ID.String(), nil
}

// GroupType is the resolver for the group_type field.
func (r *permissionGroupResolver) GroupType(ctx context.Context, obj *ent.PermissionGroup) (ent.PermissionGroupType, error) {
	return ent.PermissionGroupType(obj.GroupType), nil
}

// Permissions is the resolver for the permissions field.
func (r *permissionGroupResolver) Permissions(ctx context.Context, obj *ent.PermissionGroup) ([]*ent.Permission, error) {
	return obj.Edges.PermissionEdges, nil
}

// PermissionGroup returns graphql1.PermissionGroupResolver implementation.
func (r *Resolver) PermissionGroup() graphql1.PermissionGroupResolver {
	return &permissionGroupResolver{r}
}

type permissionGroupResolver struct{ *Resolver }
