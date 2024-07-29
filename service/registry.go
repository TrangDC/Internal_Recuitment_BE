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
	Team() TeamService
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
}

// serviceImpl is the implementation of Service.
type serviceImpl struct {
	authService                 AuthService
	storageService              StorageService
	userService                 UserService
	teamService                 TeamService
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
		teamService:                 NewTeamService(repoRegistry, dtoRegistry, logger),
		jobPositionService:          NewJobPositionService(repoRegistry, dtoRegistry, logger),
		hiringJobService:            NewHiringJobService(repoRegistry, dtoRegistry, logger),
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

// Team returns the TeamService.
func (i serviceImpl) Team() TeamService {
	return i.teamService
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
