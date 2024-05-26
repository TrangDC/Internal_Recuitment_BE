package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateJobStepResolver) ID(ctx context.Context, obj *ent.CandidateJobStep) (string, error) {
	return obj.ID.String(), nil
}

// CandidateJobID is the resolver for the candidate_job_id field.
func (r *candidateJobStepResolver) CandidateJobID(ctx context.Context, obj *ent.CandidateJobStep) (string, error) {
	return obj.CandidateJobID.String(), nil
}

// CandidateJobStatus is the resolver for the candidate_job_status field.
func (r *candidateJobStepResolver) CandidateJobStatus(ctx context.Context, obj *ent.CandidateJobStep) (ent.CandidateJobStatus, error) {
	return ent.CandidateJobStatus(obj.CandidateJobStatus), nil
}

// CandidateJobStep returns graphql1.CandidateJobStepResolver implementation.
func (r *Resolver) CandidateJobStep() graphql1.CandidateJobStepResolver {
	return &candidateJobStepResolver{r}
}

type candidateJobStepResolver struct{ *Resolver }
