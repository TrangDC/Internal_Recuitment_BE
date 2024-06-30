package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// GetTeam is the resolver for the GetTeam field.
func (r *queryResolver) GetTeam(ctx context.Context, id string) (*ent.TeamResponse, error) {
	return r.serviceRegistry.Team().GetTeam(ctx, uuid.MustParse(id))
}

// GetAllTeams is the resolver for the GetAllTeams field.
func (r *queryResolver) GetAllTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.TeamFilter, freeWord *ent.TeamFreeWord, orderBy ent.TeamOrderBy) (*ent.TeamResponseGetAll, error) {
	return r.serviceRegistry.Team().GetTeams(ctx, pagination, freeWord, filter, orderBy)
}

// GetUser is the resolver for the GetUser field.UserSelectionResponseGetAll
func (r *queryResolver) GetUser(ctx context.Context, id string) (*ent.UserResponse, error) {
	return r.serviceRegistry.User().GetUser(ctx, uuid.MustParse(id))
}

// GetAllUsers is the resolver for the GetAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error) {
	return r.serviceRegistry.User().GetUsers(ctx, pagination, filter, freeWord, orderBy)
}

// GetMe is the resolver for the GetMe field.
func (r *queryResolver) GetMe(ctx context.Context) (*ent.UserResponse, error) {
	return r.serviceRegistry.User().GetMe(ctx)
}

// GetHiringJob is the resolver for the GetHiringJob field.
func (r *queryResolver) GetHiringJob(ctx context.Context, id string) (*ent.HiringJobResponse, error) {
	return r.serviceRegistry.HiringJob().GetHiringJob(ctx, uuid.MustParse(id))
}

// GetAllHiringJobs is the resolver for the GetAllHiringJobs field.
func (r *queryResolver) GetAllHiringJobs(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringJobFilter, freeWord *ent.HiringJobFreeWord, orderBy ent.HiringJobOrderBy) (*ent.HiringJobResponseGetAll, error) {
	return r.serviceRegistry.HiringJob().GetHiringJobs(ctx, pagination, freeWord, filter, orderBy)
}

// GetAuditTrail is the resolver for the GetAuditTrail field.
func (r *queryResolver) GetAuditTrail(ctx context.Context, id string) (*ent.AuditTrailResponse, error) {
	return r.serviceRegistry.AuditTrail().GetAuditTrail(ctx, uuid.MustParse(id))
}

// GetAllAuditTrails is the resolver for the GetAllAuditTrails field.
func (r *queryResolver) GetAllAuditTrails(ctx context.Context, pagination *ent.PaginationInput, filter *ent.AuditTrailFilter, freeWord *ent.AuditTrailFreeWord, orderBy *ent.AuditTrailOrder) (*ent.AuditTrailResponseGetAll, error) {
	return r.serviceRegistry.AuditTrail().GetAuditTrails(ctx, pagination, freeWord, filter, orderBy)
}

// GetCandidate is the resolver for the GetCandidate field.
func (r *queryResolver) GetCandidate(ctx context.Context, id string) (*ent.CandidateResponse, error) {
	return r.serviceRegistry.Candidate().GetCandidate(ctx, uuid.MustParse(id))
}

// GetAllCandidates is the resolver for the GetAllCandidates field.
func (r *queryResolver) GetAllCandidates(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateFilter, freeWord *ent.CandidateFreeWord, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error) {
	return r.serviceRegistry.Candidate().GetCandidates(ctx, pagination, freeWord, filter, orderBy)
}

// GetCandidateJob is the resolver for the GetCandidateJob field.
func (r *queryResolver) GetCandidateJob(ctx context.Context, id string) (*ent.CandidateJobResponse, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateJob(ctx, uuid.MustParse(id))
}

// GetAllCandidateJobs is the resolver for the GetAllCandidateJobs field.
func (r *queryResolver) GetAllCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateJobFilter, freeWord *ent.CandidateJobFreeWord, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateJobs(ctx, pagination, freeWord, filter, orderBy)
}

// GetCandidateJobGroupByStatus is the resolver for the GetCandidateJobGroupByStatus field.
func (r *queryResolver) GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (*ent.CandidateJobGroupByStatusResponse, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateJobGroupByStatus(ctx, pagination, filter, freeWord, orderBy)
}

// GetCandidateJobGroupByInterview is the resolver for the GetCandidateJobGroupByInterview field.
func (r *queryResolver) GetCandidateJobGroupByInterview(ctx context.Context, id string) (*ent.CandidateJobGroupByInterviewResponse, error) {
	return r.serviceRegistry.CandidateJob().GetCandidateJobGroupByInterview(ctx, uuid.MustParse(id))
}

// GetCandidateJobFeedback is the resolver for the GetCandidateJobFeedback field.
func (r *queryResolver) GetCandidateJobFeedback(ctx context.Context, id string) (*ent.CandidateJobFeedbackResponse, error) {
	return r.serviceRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, uuid.MustParse(id))
}

// GetAllCandidateJobFeedbacks is the resolver for the GetAllCandidateJobFeedbacks field.
func (r *queryResolver) GetAllCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateJobFeedbackFilter, freeWord *ent.CandidateJobFeedbackFreeWord, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error) {
	return r.serviceRegistry.CandidateJobFeedback().GetCandidateJobFeedbacks(ctx, pagination, freeWord, &filter, orderBy)
}

// GetCandidateInterview is the resolver for the GetCandidateInterview field.
func (r *queryResolver) GetCandidateInterview(ctx context.Context, id string) (*ent.CandidateInterviewResponse, error) {
	return r.serviceRegistry.CandidateInterview().GetCandidateInterview(ctx, uuid.MustParse(id))
}

// GetAllCandidateInterviews is the resolver for the GetAllCandidateInterviews field.
func (r *queryResolver) GetAllCandidateInterviews(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateInterviewFilter, freeWord *ent.CandidateInterviewFreeWord, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error) {
	return r.serviceRegistry.CandidateInterview().GetCandidateInterviews(ctx, pagination, freeWord, filter, orderBy)
}

// GetAllCandidateInterview4Calendar is the resolver for the GetAllCandidateInterview4Calendar field.
func (r *queryResolver) GetAllCandidateInterview4Calendar(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateInterviewCalendarFilter, freeWord *ent.CandidateInterviewFreeWord, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error) {
	return r.serviceRegistry.CandidateInterview().GetAllCandidateInterview4Calendar(ctx, pagination, freeWord, filter, orderBy)
}

// ExportSampleCandidate is the resolver for the ExportSampleCandidate field.
func (r *queryResolver) ExportSampleCandidate(ctx context.Context, lang ent.I18nLanguage) (*ent.Base64Response, error) {
	return r.serviceRegistry.ExportData().ExportSampleCandidates(ctx, lang)
}

// GetSkill is the resolver for the GetSkill field.
func (r *queryResolver) GetSkill(ctx context.Context, id string) (*ent.SkillResponse, error) {
	return r.serviceRegistry.Skill().GetSkill(ctx, uuid.MustParse(id))
}

// GetAllSkills is the resolver for the GetAllSkills field.
func (r *queryResolver) GetAllSkills(ctx context.Context, pagination *ent.PaginationInput, filter *ent.SkillFilter, freeWord *ent.SkillFreeWord, orderBy *ent.SkillOrder) (*ent.SkillResponseGetAll, error) {
	return r.serviceRegistry.Skill().GetSkills(ctx, pagination, freeWord, filter, orderBy)
}

// GetSkillType is the resolver for the GetSkillType field.
func (r *queryResolver) GetSkillType(ctx context.Context, id string) (*ent.SkillTypeResponse, error) {
	return r.serviceRegistry.SkillType().GetSkillType(ctx, uuid.MustParse(id))
}

// GetAllSkillTypes is the resolver for the GetAllSkillTypes field.
func (r *queryResolver) GetAllSkillTypes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.SkillTypeFilter, freeWord *ent.SkillTypeFreeWord, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeResponseGetAll, error) {
	return r.serviceRegistry.SkillType().GetSkillTypes(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionUsers is the resolver for the SelectionUsers field.
func (r *queryResolver) SelectionUsers(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserSelectionResponseGetAll, error) {
	return r.serviceRegistry.User().Selections(ctx, pagination, filter, freeWord, orderBy)
}

// SelectionTeams is the resolver for the SelectionTeams field.
func (r *queryResolver) SelectionTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.TeamFilter, freeWord *ent.TeamFreeWord, orderBy ent.TeamOrderBy) (*ent.TeamSelectionResponseGetAll, error) {
	return r.serviceRegistry.Team().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionHiringJobs is the resolver for the SelectionHiringJobs field.
func (r *queryResolver) SelectionHiringJobs(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringJobFilter, freeWord *ent.HiringJobFreeWord, orderBy ent.HiringJobOrderBy) (*ent.HiringJobSelectionResponseGetAll, error) {
	return r.serviceRegistry.HiringJob().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionCandidates is the resolver for the SelectionCandidates field.
func (r *queryResolver) SelectionCandidates(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateFilter, freeWord *ent.CandidateFreeWord, orderBy *ent.CandidateOrder) (*ent.CandidateSelectionResponseGetAll, error) {
	return r.serviceRegistry.Candidate().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionSkills is the resolver for the SelectionSkills field.
func (r *queryResolver) SelectionSkills(ctx context.Context, pagination *ent.PaginationInput, filter *ent.SkillFilter, freeWord *ent.SkillFreeWord, orderBy *ent.SkillOrder) (*ent.SkillSelectionResponseGetAll, error) {
	return r.serviceRegistry.Skill().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionSkillTypes is the resolver for the SelectionSkillTypes field.
func (r *queryResolver) SelectionSkillTypes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.SkillTypeFilter, freeWord *ent.SkillTypeFreeWord, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeSelectionResponseGetAll, error) {
	return r.serviceRegistry.SkillType().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionRole is the resolver for the SelectionRole field.
func (r *queryResolver) SelectionRole(ctx context.Context, pagination *ent.PaginationInput, filter *ent.RoleFilter, freeWord *ent.RoleFreeWord, orderBy *ent.RoleOrder) (*ent.RoleSelectionResponseGetAll, error) {
	return r.serviceRegistry.Role().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// GetRole is the resolver for the GetRole field.
func (r *queryResolver) GetRole(ctx context.Context, id string) (*ent.RoleResponse, error) {
	return r.serviceRegistry.Role().GetRole(ctx, uuid.MustParse(id))
}

// GetAllRoles is the resolver for the GetAllRoles field.
func (r *queryResolver) GetAllRoles(ctx context.Context, pagination *ent.PaginationInput, filter *ent.RoleFilter, freeWord *ent.RoleFreeWord, orderBy *ent.RoleOrder) (*ent.RoleResponseGetAll, error) {
	return r.serviceRegistry.Role().GetRoles(ctx, pagination, freeWord, filter, orderBy)
}

// GetAllPermissionGroups is the resolver for the GetAllPermissionGroups field.
func (r *queryResolver) GetAllPermissionGroups(ctx context.Context) (*ent.PermissionGroupResponseGetAll, error) {
	return r.serviceRegistry.PermissionGroup().GetAllPermissionGroups(ctx)
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
