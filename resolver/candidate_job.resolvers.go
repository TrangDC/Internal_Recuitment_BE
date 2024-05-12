package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"
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

// CandidateJob returns graphql1.CandidateJobResolver implementation.
func (r *Resolver) CandidateJob() graphql1.CandidateJobResolver { return &candidateJobResolver{r} }

type candidateJobResolver struct{ *Resolver }
