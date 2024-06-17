package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
)

// ID is the resolver for the id field.
func (r *candidateInterviewResolver) ID(ctx context.Context, obj *ent.CandidateInterview) (string, error) {
	return obj.ID.String(), nil
}

// CandidateJobID is the resolver for the candidate_job_id field.
func (r *candidateInterviewResolver) CandidateJobID(ctx context.Context, obj *ent.CandidateInterview) (string, error) {
	return obj.CandidateJobID.String(), nil
}

// Interviewer is the resolver for the interviewer field.
func (r *candidateInterviewResolver) Interviewer(ctx context.Context, obj *ent.CandidateInterview) ([]*ent.User, error) {
	return obj.Edges.InterviewerEdges, nil
}

// CandidateJob is the resolver for the candidate_job field.
func (r *candidateInterviewResolver) CandidateJob(ctx context.Context, obj *ent.CandidateInterview) (*ent.CandidateJob, error) {
	return obj.Edges.CandidateJobEdge, nil
}

// EditAble is the resolver for the edit_able field.
func (r *candidateInterviewResolver) EditAble(ctx context.Context, obj *ent.CandidateInterview) (bool, error) {
	if ent.CandidateInterviewStatusEditable.IsValid(ent.CandidateInterviewStatusEditable(obj.Edges.CandidateJobEdge.Status.String())) {
		return true, nil
	}
	return false, nil
}

// Owner is the resolver for the owner field.
func (r *candidateInterviewResolver) Owner(ctx context.Context, obj *ent.CandidateInterview) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}

// CandidateInterviewStatus is the resolver for the candidate_interview_status field.
func (r *candidateInterviewResolver) CandidateInterviewStatus(ctx context.Context, obj *ent.CandidateInterview) (ent.CandidateInterviewStatus, error) {
	return ent.CandidateInterviewStatus(obj.CandidateInterviewStatus), nil
}

// CandidateInterview returns graphql1.CandidateInterviewResolver implementation.
func (r *Resolver) CandidateInterview() graphql1.CandidateInterviewResolver {
	return &candidateInterviewResolver{r}
}

type candidateInterviewResolver struct{ *Resolver }
