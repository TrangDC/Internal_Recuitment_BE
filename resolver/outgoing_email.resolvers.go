package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *outgoingEmailResolver) ID(ctx context.Context, obj *ent.OutgoingEmail) (string, error) {
	return obj.ID.String(), nil
}

// RecipientType is the resolver for the recipient_type field.
func (r *outgoingEmailResolver) RecipientType(ctx context.Context, obj *ent.OutgoingEmail) (ent.OutgoingEmailRecipientType, error) {
	return ent.OutgoingEmailRecipientType(obj.RecipientType), nil
}

// Status is the resolver for the status field.
func (r *outgoingEmailResolver) Status(ctx context.Context, obj *ent.OutgoingEmail) (ent.OutgoingEmailStatus, error) {
	return ent.OutgoingEmailStatus(obj.Status), nil
}

// OutgoingEmail returns graphql1.OutgoingEmailResolver implementation.
func (r *Resolver) OutgoingEmail() graphql1.OutgoingEmailResolver { return &outgoingEmailResolver{r} }

type outgoingEmailResolver struct{ *Resolver }
