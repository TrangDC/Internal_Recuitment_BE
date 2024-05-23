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
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// CandidateJob returns graphql1.CandidateJobResolver implementation.
func (r *Resolver) CandidateJob() graphql1.CandidateJobResolver { return &candidateJobResolver{r} }

type candidateJobResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *candidateJobResolver) CreatedBy(ctx context.Context, obj *ent.CandidateJob) (*ent.User, error) {
	return obj.Edges.CreatedByEdge, nil
}
