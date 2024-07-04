package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *permissionResolver) ID(ctx context.Context, obj *ent.Permission) (string, error) {
	return obj.ID.String(), nil
}

// Permission returns graphql1.PermissionResolver implementation.
func (r *Resolver) Permission() graphql1.PermissionResolver { return &permissionResolver{r} }

type permissionResolver struct{ *Resolver }
