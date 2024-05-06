package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *attachmentResolver) ID(ctx context.Context, obj *ent.Attachment) (string, error) {
	return obj.ID.String(), nil
}

// DocumentID is the resolver for the document_id field.
func (r *attachmentResolver) DocumentID(ctx context.Context, obj *ent.Attachment) (string, error) {
	return obj.DocumentID.String(), nil
}

// Attachment returns graphql1.AttachmentResolver implementation.
func (r *Resolver) Attachment() graphql1.AttachmentResolver { return &attachmentResolver{r} }

type attachmentResolver struct{ *Resolver }
