package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// CreatePreRequest is the resolver for the CreatePreRequest field.
func (r *mutationResolver) CreatePreRequest(ctx context.Context, input string) (string, error) {
	return "", nil
}

// CreateTeam is the resolver for the CreateTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, input ent.NewTeamInput, note string) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().CreateTeam(ctx, input, note)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTeam is the resolver for the UpdateTeam field.
func (r *mutationResolver) UpdateTeam(ctx context.Context, id string, input ent.UpdateTeamInput, note string) (*ent.TeamResponse, error) {
	result, err := r.serviceRegistry.Team().UpdateTeam(ctx, uuid.MustParse(id), input, note)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteTeam is the resolver for the DeleteTeam field.
func (r *mutationResolver) DeleteTeam(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.Team().DeleteTeam(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateHiringJob is the resolver for the CreateHiringJob field.
func (r *mutationResolver) CreateHiringJob(ctx context.Context, input ent.NewHiringJobInput) (*ent.HiringJobResponse, error) {
	result, err := r.serviceRegistry.HiringJob().CreateHiringJob(ctx, &input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateHiringJob is the resolver for the UpdateHiringJob field.
func (r *mutationResolver) UpdateHiringJob(ctx context.Context, id string, input ent.UpdateHiringJobInput) (*ent.HiringJobResponse, error) {
	result, err := r.serviceRegistry.HiringJob().UpdateHiringJob(ctx, &input, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteHiringJob is the resolver for the DeleteHiringJob field.
func (r *mutationResolver) DeleteHiringJob(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.HiringJob().DeleteHiringJob(ctx, uuid.MustParse(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateHiringJobStatus is the resolver for the UpdateHiringJobStatus field.
func (r *mutationResolver) UpdateHiringJobStatus(ctx context.Context, id string, status ent.HiringJobStatus) (*ent.HiringJobResponse, error) {
	result, err := r.serviceRegistry.HiringJob().UpdateHiringJobStatus(ctx, status, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCandidate is the resolver for the CreateCandidate field.
func (r *mutationResolver) CreateCandidate(ctx context.Context, input ent.NewCandidateInput) (*ent.CandidateResponse, error) {
	result, err := r.serviceRegistry.Candidate().CreateCandidate(ctx, &input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCandidate is the resolver for the UpdateCandidate field.
func (r *mutationResolver) UpdateCandidate(ctx context.Context, id string, input ent.UpdateCandidateInput) (*ent.CandidateResponse, error) {
	result, err := r.serviceRegistry.Candidate().UpdateCandidate(ctx, &input, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCandidate is the resolver for the DeleteCandidate field.
func (r *mutationResolver) DeleteCandidate(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.Candidate().DeleteCandidate(ctx, uuid.MustParse(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetBlackListCandidate is the resolver for the SetBlackListCandidate field.
func (r *mutationResolver) SetBlackListCandidate(ctx context.Context, id string, isBlackList bool) (bool, error) {
	err := r.serviceRegistry.Candidate().SetBlackListCandidate(ctx, uuid.MustParse(id), isBlackList)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateJob is the resolver for the CreateCandidateJob field.
func (r *mutationResolver) CreateCandidateJob(ctx context.Context, input ent.NewCandidateJobInput) (*ent.CandidateJobResponse, error) {
	result, err := r.serviceRegistry.CandidateJob().CreateCandidateJob(ctx, &input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCandidateJobAttachment is the resolver for the UpdateCandidateJobAttachment field.
func (r *mutationResolver) UpdateCandidateJobAttachment(ctx context.Context, id string, input ent.UpdateCandidateAttachment) (*ent.CandidateJobResponse, error) {
	result, err := r.serviceRegistry.CandidateJob().UpdateCandidateJobAttachment(ctx, input, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCandidateJob is the resolver for the DeleteCandidateJob field.
func (r *mutationResolver) DeleteCandidateJob(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.CandidateJob().DeleteCandidateJob(ctx, uuid.MustParse(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateCandidateJobStatus is the resolver for the UpdateCandidateJobStatus field.
func (r *mutationResolver) UpdateCandidateJobStatus(ctx context.Context, id string, status ent.CandidateJobStatus) (*ent.CandidateJobResponse, error) {
	result, err := r.serviceRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, status, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
