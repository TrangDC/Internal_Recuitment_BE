package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateResolver) ID(ctx context.Context, obj *ent.Candidate) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *candidateResolver) Status(ctx context.Context, obj *ent.Candidate) (ent.CandidateStatusEnum, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateStatus(ctx, obj.ID), nil
}

// Candidate returns graphql1.CandidateResolver implementation.
func (r *Resolver) Candidate() graphql1.CandidateResolver { return &candidateResolver{r} }

type candidateResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *candidateResolver) LastApplyDate(ctx context.Context, obj *ent.Candidate) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: LastApplyDate - last_apply_date"))
}
