package service

import (
	"trec/config"
	"trec/dto"
	"trec/ent"
	"trec/internal/azuread"
	"trec/internal/azurestorage"
	"trec/internal/servicebus"
	"trec/models"
	"trec/repository"

	"go.uber.org/zap"
)

// Service is the interface for all services.
type Service interface {
	Auth() AuthService
	Storage() StorageService
	User() UserService
	JobPosition() JobPositionService
	HiringJob() HiringJobService
	AuditTrail() AuditTrailService
	Candidate() CandidateService
	CandidateJob() CandidateJobService
	CandidateJobFeedback() CandidateJobFeedbackService
	CandidateInterview() CandidateInterviewService
	Attachment() AttachmentService
	ExportData() ExportDataService
	ImportData() ImportDataService
	CandidateJobStep() CandidateJobStepService
	Skill() SkillService
	SkillType() SkillTypeService
	Role() RoleService
	PermissionGroup() PermissionGroupService
	EmailTemplate() EmailTemplateService
	Email() EmailService
	Report() ReportService
	OutgoingEmail() OutgoingEmailService
	HiringTeam() HiringTeamService
	RecTeam() RecTeamService
	CandidateHistoryCall() CandidateHistoryCallService
	CandidateNote() CandidateNoteService
	HiringJobStep() HiringJobStepService
	CandidateActivity() CandidateActivityService
	EmailEvent() EmailEventService
}

// serviceImpl is the implementation of Service.
type serviceImpl struct {
	authService                 AuthService
	storageService              StorageService
	userService                 UserService
	jobPositionService          JobPositionService
	hiringJobService            HiringJobService
	auditTrailService           AuditTrailService
	candidateService            CandidateService
	candidateJobService         CandidateJobService
	candidateJobFeedbackService CandidateJobFeedbackService
	candidateInterviewService   CandidateInterviewService
	attachmentService           AttachmentService
	exportDataService           ExportDataService
	importDataService           ImportDataService
	candidateJobStepService     CandidateJobStepService
	skillService                SkillService
	skillTypeService            SkillTypeService
	roleService                 RoleService
	permissionGroup             PermissionGroupService
	emailTemplate               EmailTemplateService
	emailService                EmailService
	reportService               ReportService
	outgoingEmailService        OutgoingEmailService
	hiringTeamService           HiringTeamService
	recTeamService              RecTeamService
	candidateHistoryCall        CandidateHistoryCallService
	candidateNoteService        CandidateNoteService
	hiringJobStepService        HiringJobStepService
	candidateActivity           CandidateActivityService
	emailEventService           EmailEventService
}

// NewService creates a new Service.
func NewService(azureADOAuthClient azuread.AzureADOAuth, azureStorage azurestorage.AzureStorage, serviceBusClient servicebus.ServiceBus,
	i18n models.I18n, entClient *ent.Client, logger *zap.Logger, configs *config.Configurations) Service {
	repoRegistry := repository.NewRepository(entClient)
	dtoRegistry := dto.NewDto()
	return &serviceImpl{
		authService:                 NewAuthService(azureADOAuthClient, logger),
		storageService:              NewStorageService(azureStorage, logger),
		userService:                 NewUserService(repoRegistry, dtoRegistry, logger),
		jobPositionService:          NewJobPositionService(repoRegistry, dtoRegistry, logger),
		hiringJobService:            NewHiringJobService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		auditTrailService:           NewAuditTrailService(repoRegistry, logger),
		candidateService:            NewCandidateService(repoRegistry, dtoRegistry, logger),
		candidateJobService:         NewCandidateJobService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		candidateJobFeedbackService: NewCandidateJobFeedbackService(repoRegistry, dtoRegistry, logger),
		candidateInterviewService:   NewCandidateInterviewService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		attachmentService:           NewAttachmentService(repoRegistry, logger),
		exportDataService:           NewExportDataService(repoRegistry, i18n, logger),
		importDataService:           NewImportDataService(repoRegistry, logger),
		candidateJobStepService:     NewCandidateJobStepService(repoRegistry, logger),
		skillService:                NewSkillService(repoRegistry, dtoRegistry, logger),
		skillTypeService:            NewSkillTypeService(repoRegistry, dtoRegistry, logger),
		roleService:                 NewRoleService(repoRegistry, dtoRegistry, logger),
		permissionGroup:             NewPermissionGroupService(repoRegistry, logger),
		emailTemplate:               NewEmailTemplateService(repoRegistry, dtoRegistry, logger),
		emailService:                NewEmailService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		reportService:               NewReportService(repoRegistry, logger),
		outgoingEmailService:        NewOutgoingEmailService(repoRegistry, logger),
		hiringTeamService:           NewHiringTeamService(repoRegistry, dtoRegistry, logger),
		recTeamService:              NewRecTeamService(repoRegistry, dtoRegistry, logger),
		candidateHistoryCall:        NewCandidateHistoryCallService(repoRegistry, dtoRegistry, logger),
		candidateNoteService:        NewCandidateNoteService(repoRegistry, dtoRegistry, logger),
		hiringJobStepService:        NewHiringJobStepService(repoRegistry, dtoRegistry, logger),
		candidateActivity:           NewCandidateActivityService(repoRegistry, logger),
		emailEventService:           NewEmailEventService(repoRegistry, logger),
	}
}

// Auth returns the AuthService.
func (i serviceImpl) Auth() AuthService {
	return i.authService
}

// Storage returns the StorageService.
func (i serviceImpl) Storage() StorageService {
	return i.storageService
}

// User returns the UserService.
func (i serviceImpl) User() UserService {
	return i.userService
}

// JobPosition returns the JobPositionService.
func (i serviceImpl) JobPosition() JobPositionService {
	return i.jobPositionService
}

// HiringJob returns the HiringJobService.
func (i serviceImpl) HiringJob() HiringJobService {
	return i.hiringJobService
}

// AuditTrail returns the AuditTrailService.
func (i serviceImpl) AuditTrail() AuditTrailService {
	return i.auditTrailService
}

// Candidate returns the CandidateService.
func (i serviceImpl) Candidate() CandidateService {
	return i.candidateService
}

// CandidateJob returns the CandidateJobService.
func (i serviceImpl) CandidateJob() CandidateJobService {
	return i.candidateJobService
}

// CandidateJobFeedback returns the CandidateJobService.
func (i serviceImpl) CandidateJobFeedback() CandidateJobFeedbackService {
	return i.candidateJobFeedbackService
}

// CandidateInterview returns the CandidateJobService.
func (i serviceImpl) CandidateInterview() CandidateInterviewService {
	return i.candidateInterviewService
}

// Attachment returns the AttachmentService.
func (i serviceImpl) Attachment() AttachmentService {
	return i.attachmentService
}

// ExportData returns the ExportDataService.
func (i serviceImpl) ExportData() ExportDataService {
	return i.exportDataService
}

// ImportData returns the ImportDataService.
func (i serviceImpl) ImportData() ImportDataService {
	return i.importDataService
}

// CandidateJobStep returns the CandidateJobStepService.
func (i serviceImpl) CandidateJobStep() CandidateJobStepService {
	return i.candidateJobStepService
}

func (i serviceImpl) Skill() SkillService {
	return i.skillService
}

func (i serviceImpl) SkillType() SkillTypeService {
	return i.skillTypeService
}

func (i serviceImpl) Role() RoleService {
	return i.roleService
}

func (i serviceImpl) PermissionGroup() PermissionGroupService {
	return i.permissionGroup
}

func (i serviceImpl) EmailTemplate() EmailTemplateService {
	return i.emailTemplate
}

func (i serviceImpl) Email() EmailService {
	return i.emailService
}

func (i serviceImpl) Report() ReportService {
	return i.reportService
}

func (i serviceImpl) OutgoingEmail() OutgoingEmailService {
	return i.outgoingEmailService
}

func (i serviceImpl) HiringTeam() HiringTeamService {
	return i.hiringTeamService
}

func (i serviceImpl) RecTeam() RecTeamService {
	return i.recTeamService
}

func (i serviceImpl) CandidateHistoryCall() CandidateHistoryCallService {
	return i.candidateHistoryCall
}

func (i serviceImpl) CandidateNote() CandidateNoteService {
	return i.candidateNoteService
}

func (i serviceImpl) HiringJobStep() HiringJobStepService {
	return i.hiringJobStepService
}

func (i serviceImpl) CandidateActivity() CandidateActivityService {
	return i.candidateActivity
}

func (i serviceImpl) EmailEvent() EmailEventService {
	return i.emailEventService
}
