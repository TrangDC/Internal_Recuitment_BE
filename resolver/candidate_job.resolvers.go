package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/hiringjob"
	graphql1 "trec/graphql"

	"github.com/samber/lo"
)

// ID is the resolver for the id field.
func (r *candidateJobResolver) ID(ctx context.Context, obj *ent.CandidateJob) (string, error) {
	return obj.ID.String(), nil
}

// CandidateID is the resolver for the candidate_id field.
func (r *candidateJobResolver) CandidateID(ctx context.Context, obj *ent.CandidateJob) (string, error) {
	return obj.CandidateID.String(), nil
}

// HiringJobID is the resolver for the hiring_job_id field.
func (r *candidateJobResolver) HiringJobID(ctx context.Context, obj *ent.CandidateJob) (string, error) {
	return obj.HiringJobID.String(), nil
}

// Status is the resolver for the status field.
func (r *candidateJobResolver) Status(ctx context.Context, obj *ent.CandidateJob) (ent.CandidateJobStatus, error) {
	return ent.CandidateJobStatus(obj.Status), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateJobResolver) Attachments(ctx context.Context, obj *ent.CandidateJob) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// Candidate is the resolver for the candidate field.
func (r *candidateJobResolver) Candidate(ctx context.Context, obj *ent.CandidateJob) (*ent.Candidate, error) {
	return obj.Edges.CandidateEdge, nil
}

// HiringJob is the resolver for the hiring_job field.
func (r *candidateJobResolver) HiringJob(ctx context.Context, obj *ent.CandidateJob) (*ent.HiringJob, error) {
	return obj.Edges.HiringJobEdge, nil
}

// Owner is the resolver for the owner field.
func (r *candidateJobResolver) Owner(ctx context.Context, obj *ent.CandidateJob) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}

// FailedReason is the resolver for the failed_reason field.
func (r *candidateJobResolver) FailedReason(ctx context.Context, obj *ent.CandidateJob) ([]ent.CandidateJobFailedReason, error) {
	result := lo.Map(obj.FailedReason, func(s string, index int) ent.CandidateJobFailedReason {
		return ent.CandidateJobFailedReason(s)
	})
	return result, nil
}

// IsAbleToDelete is the resolver for the is_able_to_delete field.
func (r *candidateJobResolver) IsAbleToDelete(ctx context.Context, obj *ent.CandidateJob) (bool, error) {
	return (obj.Edges.HiringJobEdge.Status != hiringjob.StatusOpened &&
			ent.CandidateJobStatusEnded.IsValid(ent.CandidateJobStatusEnded(obj.Status))),
		nil
}

// InterviewFeature is the resolver for the interview_feature field.
func (r *candidateJobResolver) InterviewFeature(ctx context.Context, obj *ent.CandidateJob) (int, error) {
	interviewFeature := lo.Filter(obj.Edges.CandidateJobInterview, func(cji *ent.CandidateInterview, index int) bool {
		return cji.InterviewDate.After(time.Now().UTC())
	})
	return len(interviewFeature), nil
}

// Steps is the resolver for the steps field.
func (r *candidateJobResolver) Steps(ctx context.Context, obj *ent.CandidateJob) ([]*ent.CandidateJobStep, error) {
	return obj.Edges.CandidateJobStep, nil
}

// CandidateJob returns graphql1.CandidateJobResolver implementation.
func (r *Resolver) CandidateJob() graphql1.CandidateJobResolver { return &candidateJobResolver{r} }

type candidateJobResolver struct{ *Resolver }
