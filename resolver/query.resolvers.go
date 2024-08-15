package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trec/ent"
	graphql1 "trec/graphql"

	"github.com/google/uuid"
)

// GetHiringTeam is the resolver for the GetHiringTeam field.
func (r *queryResolver) GetHiringTeam(ctx context.Context, id string) (*ent.HiringTeamResponse, error) {
	return r.serviceRegistry.HiringTeam().GetHiringTeam(ctx, uuid.MustParse(id))
}

// GetAllHiringTeams is the resolver for the GetAllHiringTeams field.
func (r *queryResolver) GetAllHiringTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringTeamFilter, freeWord *ent.HiringTeamFreeWord, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamResponseGetAll, error) {
	return r.serviceRegistry.HiringTeam().GetHiringTeams(ctx, pagination, freeWord, filter, orderBy)
}

// GetAllRecTeams is the resolver for the GetAllRecTeams field.
func (r *queryResolver) GetAllRecTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.RecTeamFilter, freeWord *ent.RecTeamFreeWord, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamResponseGetAll, error) {
	return r.serviceRegistry.RecTeam().GetRecTeams(ctx, pagination, freeWord, filter, orderBy)
}

// GetRecTeam is the resolver for the GetRecTeam field.
func (r *queryResolver) GetRecTeam(ctx context.Context, id string) (*ent.RecTeamResponse, error) {
	return r.serviceRegistry.RecTeam().GetRecTeam(ctx, uuid.MustParse(id))
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

// GetJobPosition is the resolver for the GetJobPosition field.
func (r *queryResolver) GetJobPosition(ctx context.Context, id string) (*ent.JobPositionResponse, error) {
	return r.serviceRegistry.JobPosition().GetJobPosition(ctx, uuid.MustParse(id))
}

// GetAllJobPositions is the resolver for the GetAllJobPositions field.
func (r *queryResolver) GetAllJobPositions(ctx context.Context, pagination *ent.PaginationInput, filter *ent.JobPositionFilter, freeWord *ent.JobPositionFreeWord, orderBy *ent.JobPositionOrder) (*ent.JobPositionResponseGetAll, error) {
	return r.serviceRegistry.JobPosition().GetJobPositions(ctx, pagination, freeWord, filter, orderBy)
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
func (r *queryResolver) GetAllCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateJobFilter, freeWord *ent.CandidateJobFreeWord, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
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

// ValidProcessingCandidateJobExistByCandidateID is the resolver for the ValidProcessingCandidateJobExistByCandidateID field.
func (r *queryResolver) ValidProcessingCandidateJobExistByCandidateID(ctx context.Context, candidateID string) (bool, error) {
	return r.serviceRegistry.CandidateJob().ValidProcessingCdJobExistByCdID(ctx, uuid.MustParse(candidateID))
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

// SelectionHiringTeams is the resolver for the SelectionHiringTeams field.
func (r *queryResolver) SelectionHiringTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.HiringTeamFilter, freeWord *ent.HiringTeamFreeWord, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamSelectionResponseGetAll, error) {
	return r.serviceRegistry.HiringTeam().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionRecTeams is the resolver for the SelectionRecTeams field.
func (r *queryResolver) SelectionRecTeams(ctx context.Context, pagination *ent.PaginationInput, filter *ent.RecTeamFilter, freeWord *ent.RecTeamFreeWord, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamSelectionResponseGetAll, error) {
	return r.serviceRegistry.RecTeam().Selections(ctx, pagination, freeWord, filter, orderBy)
}

// SelectionJobPositions is the resolver for the SelectionJobPositions field.
func (r *queryResolver) SelectionJobPositions(ctx context.Context, pagination *ent.PaginationInput, filter *ent.JobPositionFilter, freeWord *ent.JobPositionFreeWord, orderBy *ent.JobPositionOrder) (*ent.JobPositionSelectionResponseGetAll, error) {
	return r.serviceRegistry.JobPosition().Selections(ctx, pagination, freeWord, filter, orderBy)
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

// GetEmailTemplate is the resolver for the GetEmailTemplate field.
func (r *queryResolver) GetEmailTemplate(ctx context.Context, id string) (*ent.EmailTemplateResponse, error) {
	return r.serviceRegistry.EmailTemplate().GetEmailTemplate(ctx, uuid.MustParse(id))
}

// GetAllEmailTemplates is the resolver for the GetAllEmailTemplates field.
func (r *queryResolver) GetAllEmailTemplates(ctx context.Context, pagination *ent.PaginationInput, filter *ent.EmailTemplateFilter, freeWord *ent.EmailTemplateFreeWord, orderBy *ent.EmailTemplateOrder) (*ent.EmailTemplateResponseGetAll, error) {
	return r.serviceRegistry.EmailTemplate().GetEmailTemplates(ctx, pagination, freeWord, filter, *orderBy)
}

// GetAllEmailTemplateKeywords is the resolver for the GetAllEmailTemplateKeywords field.
func (r *queryResolver) GetAllEmailTemplateKeywords(ctx context.Context, filter ent.EmailTemplateKeywordFilter) (*ent.GetEmailTemplateKeywordResponse, error) {
	return r.serviceRegistry.EmailTemplate().GetAllEmailTemplateKeyword(filter)
}

// GetCandidateHistoryCall is the resolver for the GetCandidateHistoryCall field.
func (r *queryResolver) GetCandidateHistoryCall(ctx context.Context, id string) (*ent.CandidateHistoryCallResponse, error) {
	return r.serviceRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, uuid.MustParse(id))
}

// GetAllCandidateHistoryCalls is the resolver for the GetAllCandidateHistoryCalls field.
func (r *queryResolver) GetAllCandidateHistoryCalls(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateHistoryCallFilter, freeWord *ent.CandidateHistoryCallFreeWord, orderBy *ent.CandidateHistoryCallOrder) (*ent.CandidateHistoryCallResponseGetAll, error) {
	return r.serviceRegistry.CandidateHistoryCall().GetCandidateHistoryCalls(ctx, pagination, freeWord, filter, orderBy)
}

// GetAllPermissionGroups is the resolver for the GetAllPermissionGroups field.
func (r *queryResolver) GetAllPermissionGroups(ctx context.Context) (*ent.PermissionGroupResponseGetAll, error) {
	return r.serviceRegistry.PermissionGroup().GetAllPermissionGroups(ctx)
}

// ReportCandidateLcc is the resolver for the ReportCandidateLCC field.
func (r *queryResolver) ReportCandidateLcc(ctx context.Context) (*ent.ReportCandidateLCCResponse, error) {
	return r.serviceRegistry.Report().ReportCandidateLCC(ctx)
}

// ReportCandidateColumnChart is the resolver for the ReportCandidateColumnChart field.
func (r *queryResolver) ReportCandidateColumnChart(ctx context.Context, filter ent.ReportFilter) (*ent.ReportCandidateColumnChartResponse, error) {
	return r.serviceRegistry.Report().ReportCandidateColumnChart(ctx, filter)
}

// ReportApplication is the resolver for the ReportApplication field.
func (r *queryResolver) ReportApplication(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationResponse, error) {
	return r.serviceRegistry.Report().ReportApplication(ctx, filter)
}

// ReportApplicationReportTable is the resolver for the ReportApplicationReportTable field.
func (r *queryResolver) ReportApplicationReportTable(ctx context.Context, filter ent.ReportFilter) (*ent.ReportApplicationReportTableResponse, error) {
	return r.serviceRegistry.Report().ReportApplicationReportTable(ctx, filter)
}

// ReportCandidateConversionRateChart is the resolver for the ReportCandidateConversionRateChart field.
func (r *queryResolver) ReportCandidateConversionRateChart(ctx context.Context) (*ent.ReportCandidateConversionRateChartResponse, error) {
	return r.serviceRegistry.Report().ReportCandidateConversionRateChart(ctx)
}

// ReportCandidateConversionRateTable is the resolver for the ReportCandidateConversionRateTable field.
func (r *queryResolver) ReportCandidateConversionRateTable(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.ReportOrderBy) (*ent.ReportCandidateConversionRateTableResponse, error) {
	return r.serviceRegistry.Report().ReportCandidateConversionRateTable(ctx, pagination, orderBy)
}

// GetCandidateNote is the resolver for the GetCandidateNote field.
func (r *queryResolver) GetCandidateNote(ctx context.Context, id string) (*ent.CandidateNoteResponse, error) {
	return r.serviceRegistry.CandidateNote().GetCandidateNote(ctx, uuid.MustParse(id))
}

// GetAllCandidateNotes is the resolver for the GetAllCandidateNotes field.
func (r *queryResolver) GetAllCandidateNotes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateNoteFilter, freeWord *ent.CandidateNoteFreeWord, orderBy *ent.CandidateNoteOrder) (*ent.CandidateNoteResponseGetAll, error) {
	return r.serviceRegistry.CandidateNote().GetAllCandidateNotes(ctx, pagination, filter, freeWord, orderBy)
}

// GetCandidateActivities is the resolver for the GetCandidateActivities field.
func (r *queryResolver) GetCandidateActivities(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateActivityFilter, freeWord *ent.CandidateActivityFreeWord, orderBy ent.CandidateActivityOrder) (*ent.CandidateActivityResponse, error) {
	return r.serviceRegistry.CandidateActivity().GetAllCandidateActivities(ctx, pagination, filter, freeWord, orderBy)
}

// GetAllOutgoingEmails is the resolver for the GetAllOutgoingEmails field.
func (r *queryResolver) GetAllOutgoingEmails(ctx context.Context, pagination *ent.PaginationInput, filter ent.OutgoingEmailFilter, freeWord *ent.OutgoingEmailFreeWord, orderBy *ent.OutgoingEmailOrder) (*ent.OutgoingEmailResponseGetAll, error) {
	return r.serviceRegistry.OutgoingEmail().GetAllOutgoingEmails(ctx, pagination, freeWord, filter, orderBy)
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
