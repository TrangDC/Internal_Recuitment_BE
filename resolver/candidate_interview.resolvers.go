package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateInterviewResolver) ID(ctx context.Context, obj *ent.CandidateInterview) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// CandidateJobID is the resolver for the candidate_job_id field.
func (r *candidateInterviewResolver) CandidateJobID(ctx context.Context, obj *ent.CandidateInterview) (string, error) {
	panic(fmt.Errorf("not implemented: CandidateJobID - candidate_job_id"))
}

// Interviewer is the resolver for the interviewer field.
func (r *candidateInterviewResolver) Interviewer(ctx context.Context, obj *ent.CandidateInterview) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Interviewer - interviewer"))
}

// CandidateInterview returns graphql1.CandidateInterviewResolver implementation.
func (r *Resolver) CandidateInterview() graphql1.CandidateInterviewResolver {
	return &candidateInterviewResolver{r}
}

type candidateInterviewResolver struct{ *Resolver }
