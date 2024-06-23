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
func (r *candidateJobFeedbackResolver) ID(ctx context.Context, obj *ent.CandidateJobFeedback) (string, error) {
	return obj.ID.String(), nil
}

// CreatedBy is the resolver for the created_by field.
func (r *candidateJobFeedbackResolver) CreatedBy(ctx context.Context, obj *ent.CandidateJobFeedback) (string, error) {
	return obj.CreatedBy.String(), nil
}

// CandidateJobID is the resolver for the candidate_job_id field.
func (r *candidateJobFeedbackResolver) CandidateJobID(ctx context.Context, obj *ent.CandidateJobFeedback) (string, error) {
	return obj.CandidateJobID.String(), nil
}

// CandidateJob is the resolver for the candidate_job field.
func (r *candidateJobFeedbackResolver) CandidateJob(ctx context.Context, obj *ent.CandidateJobFeedback) (*ent.CandidateJob, error) {
	return obj.Edges.CandidateJobEdge, nil
}

// Owner is the resolver for the owner field.
func (r *candidateJobFeedbackResolver) Owner(ctx context.Context, obj *ent.CandidateJobFeedback) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}

// Edited is the resolver for the edited field.
func (r *candidateJobFeedbackResolver) Edited(ctx context.Context, obj *ent.CandidateJobFeedback) (bool, error) {
	return dto.IsRecordEdited(obj.CreatedAt, obj.UpdatedAt), nil
}

// Attachments is the resolver for the attachments field.
func (r *candidateJobFeedbackResolver) Attachments(ctx context.Context, obj *ent.CandidateJobFeedback) ([]*ent.Attachment, error) {
	return obj.Edges.AttachmentEdges, nil
}

// CandidateJobFeedback returns graphql1.CandidateJobFeedbackResolver implementation.
func (r *Resolver) CandidateJobFeedback() graphql1.CandidateJobFeedbackResolver {
	return &candidateJobFeedbackResolver{r}
}

type candidateJobFeedbackResolver struct{ *Resolver }
