package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
