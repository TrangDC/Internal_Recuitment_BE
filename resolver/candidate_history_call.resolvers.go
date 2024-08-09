package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"trec/dto"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateHistoryCallResolver) ID(ctx context.Context, obj *ent.CandidateHistoryCall) (string, error) {
	return obj.ID.String(), nil
}

// CandidateID is the resolver for the candidate_id field.
func (r *candidateHistoryCallResolver) CandidateID(ctx context.Context, obj *ent.CandidateHistoryCall) (string, error) {
	return obj.CandidateID.String(), nil
}

// Type is the resolver for the type field.
func (r *candidateHistoryCallResolver) Type(ctx context.Context, obj *ent.CandidateHistoryCall) (ent.CandidateHistoryCallTypeEnum, error) {
	return ent.CandidateHistoryCallTypeEnum(obj.Type), nil
}

// Candidate is the resolver for the candidate field.
func (r *candidateHistoryCallResolver) Candidate(ctx context.Context, obj *ent.CandidateHistoryCall) (*ent.Candidate, error) {
	return obj.Edges.CandidateEdge, nil
}

// Edited is the resolver for the edited field.
func (r *candidateHistoryCallResolver) Edited(ctx context.Context, obj *ent.CandidateHistoryCall) (bool, error) {
	return dto.IsRecordEdited(obj.CreatedAt, obj.UpdatedAt), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *candidateHistoryCallResolver) CreatedBy(ctx context.Context, obj *ent.CandidateHistoryCall) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}

// CandidateHistoryCall returns graphql1.CandidateHistoryCallResolver implementation.
func (r *Resolver) CandidateHistoryCall() graphql1.CandidateHistoryCallResolver {
	return &candidateHistoryCallResolver{r}
}

type candidateHistoryCallResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *candidateHistoryCallResolver) Name(ctx context.Context, obj *ent.CandidateHistoryCall) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}
