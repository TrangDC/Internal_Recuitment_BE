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

// CreateHiringTeam is the resolver for the CreateHiringTeam field.
func (r *mutationResolver) CreateHiringTeam(ctx context.Context, input ent.NewHiringTeamInput, note string) (*ent.HiringTeamResponse, error) {
	return r.serviceRegistry.HiringTeam().CreateHiringTeam(ctx, input, note)
}

// UpdateHiringTeam is the resolver for the UpdateHiringTeam field.
func (r *mutationResolver) UpdateHiringTeam(ctx context.Context, id string, input ent.UpdateHiringTeamInput, note string) (*ent.HiringTeamResponse, error) {
	return r.serviceRegistry.HiringTeam().UpdateHiringTeam(ctx, uuid.MustParse(id), input, note)
}

// DeleteHiringTeam is the resolver for the DeleteHiringTeam field.
func (r *mutationResolver) DeleteHiringTeam(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.HiringTeam().DeleteHiringTeam(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateJobPosition is the resolver for the CreateJobPosition field.
func (r *mutationResolver) CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPositionResponse, error) {
	return r.serviceRegistry.JobPosition().CreateJobPosition(ctx, input)
}

// UpdateJobPosition is the resolver for the UpdateJobPosition field.
func (r *mutationResolver) UpdateJobPosition(ctx context.Context, id string, input ent.UpdateJobPositionInput) (*ent.JobPositionResponse, error) {
	return r.serviceRegistry.JobPosition().UpdateJobPosition(ctx, uuid.MustParse(id), input)
}

// DeleteJobPosition is the resolver for the DeleteJobPosition field.
func (r *mutationResolver) DeleteJobPosition(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.JobPosition().DeleteJobPosition(ctx, uuid.MustParse(id))
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
func (r *mutationResolver) CreateCandidateJob(ctx context.Context, input ent.NewCandidateJobInput, note *string) (*ent.CandidateJobResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateJob().CreateCandidateJob(ctx, &input, inputNote)
}

// UpdateCandidateJobAttachment is the resolver for the UpdateCandidateJobAttachment field.
func (r *mutationResolver) UpdateCandidateJobAttachment(ctx context.Context, id string, input ent.UpdateCandidateAttachment, note *string) (*ent.CandidateJobResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateJob().UpdateCandidateJobAttachment(ctx, input, uuid.MustParse(id), inputNote)
}

// DeleteCandidateJob is the resolver for the DeleteCandidateJob field.
func (r *mutationResolver) DeleteCandidateJob(ctx context.Context, id string, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateJob().DeleteCandidateJob(ctx, uuid.MustParse(id), inputNote)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateCandidateJobStatus is the resolver for the UpdateCandidateJobStatus field.
func (r *mutationResolver) UpdateCandidateJobStatus(ctx context.Context, id string, input ent.UpdateCandidateJobStatus, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, input, uuid.MustParse(id), inputNote)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateJobFeedback is the resolver for the CreateCandidateJobFeedback field.
func (r *mutationResolver) CreateCandidateJobFeedback(ctx context.Context, input ent.NewCandidateJobFeedbackInput, note *string) (*ent.CandidateJobFeedbackResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateJobFeedback().CreateCandidateJobFeedback(ctx, &input, inputNote)
}

// UpdateCandidateJobFeedback is the resolver for the UpdateCandidateJobFeedback field.
func (r *mutationResolver) UpdateCandidateJobFeedback(ctx context.Context, id string, input ent.UpdateCandidateJobFeedbackInput, note *string) (*ent.CandidateJobFeedbackResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateJobFeedback().UpdateCandidateJobFeedback(ctx, uuid.MustParse(id), &input, inputNote)
}

// DeleteCandidateJobFeedback is the resolver for the DeleteCandidateJobFeedback field.
func (r *mutationResolver) DeleteCandidateJobFeedback(ctx context.Context, id string, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateJobFeedback().DeleteCandidateJobFeedback(ctx, uuid.MustParse(id), inputNote)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateInterview is the resolver for the CreateCandidateInterview field.
func (r *mutationResolver) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, note *string) (*ent.CandidateInterviewResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateInterview().CreateCandidateInterview(ctx, input, inputNote)
}

// UpdateCandidateInterview is the resolver for the UpdateCandidateInterview field.
func (r *mutationResolver) UpdateCandidateInterview(ctx context.Context, id string, input ent.UpdateCandidateInterviewInput, note *string) (*ent.CandidateInterviewResponse, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	return r.serviceRegistry.CandidateInterview().UpdateCandidateInterview(ctx, uuid.MustParse(id), input, inputNote)
}

// UpdateCandidateInterviewSchedule is the resolver for the UpdateCandidateInterviewSchedule field.
func (r *mutationResolver) UpdateCandidateInterviewSchedule(ctx context.Context, id string, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterviewResponse, error) {
	return r.serviceRegistry.CandidateInterview().UpdateCandidateInterviewSchedule(ctx, uuid.MustParse(id), input)
}

// DeleteCandidateInterview is the resolver for the DeleteCandidateInterview field.
func (r *mutationResolver) DeleteCandidateInterview(ctx context.Context, id string, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateInterview().DeleteCandidateInterview(ctx, uuid.MustParse(id), inputNote)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateCandidateInterview4Calendar is the resolver for the CreateCandidateInterview4Calendar field.
func (r *mutationResolver) CreateCandidateInterview4Calendar(ctx context.Context, input ent.NewCandidateInterview4CalendarInput, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateInterview().CreateCandidateInterview4Calendar(ctx, input, inputNote)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateCandidateInterviewStatus is the resolver for the UpdateCandidateInterviewStatus field.
func (r *mutationResolver) UpdateCandidateInterviewStatus(ctx context.Context, id string, input ent.UpdateCandidateInterviewStatusInput, note *string) (bool, error) {
	var inputNote string
	if note != nil {
		inputNote = *note
	}
	err := r.serviceRegistry.CandidateInterview().UpdateCandidateInterviewStatus(ctx, uuid.MustParse(id), input, inputNote)
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

// CreateSkill is the resolver for the CreateSkill field.
func (r *mutationResolver) CreateSkill(ctx context.Context, input ent.NewSkillInput, note string) (*ent.SkillResponse, error) {
	return r.serviceRegistry.Skill().CreateSkill(ctx, input, note)
}

// UpdateSkill is the resolver for the UpdateSkill field.
func (r *mutationResolver) UpdateSkill(ctx context.Context, id string, input ent.UpdateSkillInput, note string) (*ent.SkillResponse, error) {
	return r.serviceRegistry.Skill().UpdateSkill(ctx, uuid.MustParse(id), input, note)
}

// DeleteSkill is the resolver for the DeleteSkill field.
func (r *mutationResolver) DeleteSkill(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.Skill().DeleteSkill(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateSkillType is the resolver for the CreateSkillType field.
func (r *mutationResolver) CreateSkillType(ctx context.Context, input ent.NewSkillTypeInput, note string) (*ent.SkillTypeResponse, error) {
	return r.serviceRegistry.SkillType().CreateSkillType(ctx, input, note)
}

// UpdateSkillType is the resolver for the UpdateSkillType field.
func (r *mutationResolver) UpdateSkillType(ctx context.Context, id string, input ent.UpdateSkillTypeInput, note string) (*ent.SkillTypeResponse, error) {
	return r.serviceRegistry.SkillType().UpdateSkillType(ctx, uuid.MustParse(id), input, note)
}

// DeleteSkillType is the resolver for the DeleteSkillType field.
func (r *mutationResolver) DeleteSkillType(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.SkillType().DeleteSkillType(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateRole is the resolver for the CreateRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.NewRoleInput, note string) (*ent.RoleResponse, error) {
	return r.serviceRegistry.Role().CreateRole(ctx, input, note)
}

// UpdateRole is the resolver for the UpdateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id string, input ent.UpdateRoleInput, note string) (*ent.RoleResponse, error) {
	return r.serviceRegistry.Role().UpdateRole(ctx, uuid.MustParse(id), input, note)
}

// DeleteRole is the resolver for the DeleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.Role().DeleteRole(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateEmailTemplate is the resolver for the CreateEmailTemplate field.
func (r *mutationResolver) CreateEmailTemplate(ctx context.Context, input ent.NewEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error) {
	return r.serviceRegistry.EmailTemplate().CreateEmailTemplate(ctx, input, note)
}

// UpdateEmailTemplate is the resolver for the UpdateEmailTemplate field.
func (r *mutationResolver) UpdateEmailTemplate(ctx context.Context, id string, input ent.UpdateEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error) {
	return r.serviceRegistry.EmailTemplate().UpdateEmailTemplate(ctx, uuid.MustParse(id), input, note)
}

// UpdateEmailTemplateStatus is the resolver for the UpdateEmailTemplateStatus field.
func (r *mutationResolver) UpdateEmailTemplateStatus(ctx context.Context, id string, input ent.UpdateEmailTemplateStatusInput, note string) (bool, error) {
	err := r.serviceRegistry.EmailTemplate().UpdateEmailTemplateStatus(ctx, uuid.MustParse(id), input, note)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// DeleteEmailTemplate is the resolver for the DeleteEmailTemplate field.
func (r *mutationResolver) DeleteEmailTemplate(ctx context.Context, id string, note string) (bool, error) {
	err := r.serviceRegistry.EmailTemplate().DeleteEmailTemplate(ctx, uuid.MustParse(id), note)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ValidateCandidateInterview is the resolver for the ValidateCandidateInterview field.
func (r *mutationResolver) ValidateCandidateInterview(ctx context.Context, input ent.CandidateInterviewValidateInput) (*ent.CandidateInterviewResponseValidate, error) {
	return r.serviceRegistry.CandidateInterview().ValidateCandidateInterview(ctx, input)
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
