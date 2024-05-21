package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// CreateAttachmentSasurl is the resolver for the CreateAttachmentSASURL field.
func (r *mutationResolver) CreateAttachmentSasurl(ctx context.Context, input ent.AttachmentInput) (*ent.AttachmentResponse, error) {
	return r.serviceRegistry.Storage().CreateAttachmentSASURL(ctx, input)
}

// UpdateUser is the resolver for the UpdateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input ent.UpdateUserInput, note string) (*ent.UserResponse, error) {
	return r.serviceRegistry.User().UpdateUser(ctx, &input, uuid.MustParse(id), note)
}

// UpdateUserStatus is the resolver for the UpdateUserStatus field.
func (r *mutationResolver) UpdateUserStatus(ctx context.Context, id string, input ent.UpdateUserStatusInput, note string) (*ent.UserResponse, error) {
	return r.serviceRegistry.User().UpdateUserStatus(ctx, input, uuid.MustParse(id), note)
}

// CreateTeam is the resolver for the CreateTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, input ent.NewTeamInput, note string) (*ent.TeamResponse, error) {
	return r.serviceRegistry.Team().CreateTeam(ctx, input, note)
}

// UpdateTeam is the resolver for the UpdateTeam field.
func (r *mutationResolver) UpdateTeam(ctx context.Context, id string, input ent.UpdateTeamInput, note string) (*ent.TeamResponse, error) {
	return r.serviceRegistry.Team().UpdateTeam(ctx, uuid.MustParse(id), input, note)
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
func (r *mutationResolver) CreateHiringJob(ctx context.Context, input ent.NewHiringJobInput, note string) (*ent.HiringJobResponse, error) {
	return r.serviceRegistry.HiringJob().CreateHiringJob(ctx, &input, note)
}

// UpdateHiringJob is the resolver for the UpdateHiringJob field.
func (r *mutationResolver) UpdateHiringJob(ctx context.Context, id string, input ent.UpdateHiringJobInput, note string) (*ent.HiringJobResponse, error) {
	return r.serviceRegistry.HiringJob().UpdateHiringJob(ctx, &input, uuid.MustParse(id), note)
}

// DeleteHiringJob is the resolver for the DeleteHiringJob field.
func (r *mutationResolver) DeleteHiringJob(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.HiringJob().DeleteHiringJob(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateHiringJobStatus is the resolver for the UpdateHiringJobStatus field.
func (r *mutationResolver) UpdateHiringJobStatus(ctx context.Context, id string, status ent.HiringJobStatus, note string) (*ent.HiringJobResponse, error) {
	return r.serviceRegistry.HiringJob().UpdateHiringJobStatus(ctx, status, uuid.MustParse(id), note)
}

// CreateCandidate is the resolver for the CreateCandidate field.
func (r *mutationResolver) CreateCandidate(ctx context.Context, input ent.NewCandidateInput, note string) (*ent.CandidateResponse, error) {
	return r.serviceRegistry.Candidate().CreateCandidate(ctx, &input, note)
}

// UpdateCandidate is the resolver for the UpdateCandidate field.
func (r *mutationResolver) UpdateCandidate(ctx context.Context, id string, input ent.UpdateCandidateInput, note string) (*ent.CandidateResponse, error) {
	return r.serviceRegistry.Candidate().UpdateCandidate(ctx, &input, uuid.MustParse(id), note)
}

// DeleteCandidate is the resolver for the DeleteCandidate field.
func (r *mutationResolver) DeleteCandidate(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.Candidate().DeleteCandidate(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetBlackListCandidate is the resolver for the SetBlackListCandidate field.
func (r *mutationResolver) SetBlackListCandidate(ctx context.Context, id string, isBlackList bool, note string) (bool, error) {
	err := r.serviceRegistry.Candidate().SetBlackListCandidate(ctx, uuid.MustParse(id), isBlackList, note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateJob is the resolver for the CreateCandidateJob field.
func (r *mutationResolver) CreateCandidateJob(ctx context.Context, input ent.NewCandidateJobInput) (*ent.CandidateJobResponse, error) {
	return r.serviceRegistry.CandidateJob().CreateCandidateJob(ctx, &input)
}

// UpdateCandidateJobAttachment is the resolver for the UpdateCandidateJobAttachment field.
func (r *mutationResolver) UpdateCandidateJobAttachment(ctx context.Context, id string, input ent.UpdateCandidateAttachment) (*ent.CandidateJobResponse, error) {
	return r.serviceRegistry.CandidateJob().UpdateCandidateJobAttachment(ctx, input, uuid.MustParse(id))
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
	return r.serviceRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, status, uuid.MustParse(id))
}

// CreateCandidateJobFeedback is the resolver for the CreateCandidateJobFeedback field.
func (r *mutationResolver) CreateCandidateJobFeedback(ctx context.Context, input ent.NewCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error) {
	return r.serviceRegistry.CandidateJobFeedback().CreateCandidateJobFeedback(ctx, &input)
}

// UpdateCandidateJobFeedback is the resolver for the UpdateCandidateJobFeedback field.
func (r *mutationResolver) UpdateCandidateJobFeedback(ctx context.Context, id string, input ent.UpdateCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error) {
	return r.serviceRegistry.CandidateJobFeedback().UpdateCandidateJobFeedback(ctx, uuid.MustParse(id), &input)
}

// CreateCandidateInterview is the resolver for the CreateCandidateInterview field.
func (r *mutationResolver) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput) (*ent.CandidateInterviewResponse, error) {
	return r.serviceRegistry.CandidateInterview().CreateCandidateInterview(ctx, input)
}

// UpdateCandidateInterview is the resolver for the UpdateCandidateInterview field.
func (r *mutationResolver) UpdateCandidateInterview(ctx context.Context, id string, input ent.UpdateCandidateInterviewInput) (*ent.CandidateInterviewResponse, error) {
	return r.serviceRegistry.CandidateInterview().UpdateCandidateInterview(ctx, uuid.MustParse(id), input)
}

// UpdateCandidateInterviewSchedule is the resolver for the UpdateCandidateInterviewSchedule field.
func (r *mutationResolver) UpdateCandidateInterviewSchedule(ctx context.Context, id string, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterviewResponse, error) {
	return r.serviceRegistry.CandidateInterview().UpdateCandidateInterviewSchedule(ctx, uuid.MustParse(id), input)
}

// DeleteCandidateInterview is the resolver for the DeleteCandidateInterview field.
func (r *mutationResolver) DeleteCandidateInterview(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.CandidateInterview().DeleteCandidateInterview(ctx, uuid.MustParse(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateInterview4Calendar is the resolver for the CreateCandidateInterview4Calendar field.
func (r *mutationResolver) CreateCandidateInterview4Calendar(ctx context.Context, input ent.NewCandidateInterview4CalendarInput) (bool, error) {
	err := r.serviceRegistry.CandidateInterview().CreateCandidateInterview4Calendar(ctx, input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ImportCandidate is the resolver for the ImportCandidate field.
func (r *mutationResolver) ImportCandidate(ctx context.Context, file graphql.Upload) (bool, error) {
	err := r.serviceRegistry.ImportData().ImportCandidate(ctx, file)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
