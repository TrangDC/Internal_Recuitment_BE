package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/samber/lo"
)

// ID is the resolver for the id field.
func (r *emailTemplateResolver) ID(ctx context.Context, obj *ent.EmailTemplate) (string, error) {
	return obj.ID.String(), nil
}

// Event is the resolver for the event field.
func (r *emailTemplateResolver) Event(ctx context.Context, obj *ent.EmailTemplate) (*ent.EmailEvent, error) {
	return obj.Edges.EventEdgeOrErr()
}

// SendTo is the resolver for the send_to field.
func (r *emailTemplateResolver) SendTo(ctx context.Context, obj *ent.EmailTemplate) ([]ent.EmailTemplateSendTo, error) {
	result := lo.Map(obj.SendTo, func(s string, index int) ent.EmailTemplateSendTo {
		return ent.EmailTemplateSendTo(s)
	})
	return result, nil
}

// Status is the resolver for the status field.
func (r *emailTemplateResolver) Status(ctx context.Context, obj *ent.EmailTemplate) (ent.EmailTemplateStatus, error) {
	return ent.EmailTemplateStatus(obj.Status), nil
}

// Roles is the resolver for the roles field.
func (r *emailTemplateResolver) Roles(ctx context.Context, obj *ent.EmailTemplate) ([]*ent.Role, error) {
	return obj.Edges.RoleEdges, nil
}

// EmailTemplate returns graphql1.EmailTemplateResolver implementation.
func (r *Resolver) EmailTemplate() graphql1.EmailTemplateResolver { return &emailTemplateResolver{r} }

type emailTemplateResolver struct{ *Resolver }
