package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *emailEventResolver) ID(ctx context.Context, obj *ent.EmailEvent) (string, error) {
	return obj.ID.String(), nil
}

// EmailEvent returns graphql1.EmailEventResolver implementation.
func (r *Resolver) EmailEvent() graphql1.EmailEventResolver { return &emailEventResolver{r} }

type emailEventResolver struct{ *Resolver }
