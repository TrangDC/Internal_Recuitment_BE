package service

import (
	"trec/ent"
	"trec/internal/azuread"
	"trec/internal/azurestorage"
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
	HiringJob() HiringJobService
	AuditTrail() AuditTrailService
	Candidate() CandidateService
	CandidateJob() CandidateJobService
	CandidateJobFeedback() CandidateJobFeedbackService
	CandidateInterview() CandidateInterviewService
	Attachment() AttachmentService
	ExportData() ExportDataService
	ImportData() ImportDataService
	CandiateJobStep() CandidateJobStepService
}

// serviceImpl is the implementation of Service.
type serviceImpl struct {
	authService                 AuthService
	storageService              StorageService
	userService                 UserService
	teamService                 TeamService
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
}

// NewService creates a new Service.
func NewService(azureADOAuthClient azuread.AzureADOAuth, azureStorage azurestorage.AzureStorage, i18n models.I18n, entClient *ent.Client, logger *zap.Logger) Service {
	repoRegistry := repository.NewRepository(entClient)

	return &serviceImpl{
		authService:                 NewAuthService(azureADOAuthClient, logger),
		storageService:              NewStorageService(azureStorage, logger),
		userService:                 NewUserService(repoRegistry, logger),
		teamService:                 NewTeamService(repoRegistry, logger),
		hiringJobService:            NewHiringJobService(repoRegistry, logger),
		auditTrailService:           NewAuditTrailService(repoRegistry, logger),
		candidateService:            NewCandidateService(repoRegistry, logger),
		candidateJobService:         NewCandidateJobService(repoRegistry, logger),
		candidateJobFeedbackService: NewCandidateJobFeedbackService(repoRegistry, logger),
		candidateInterviewService:   NewCandidateInterviewService(repoRegistry, logger),
		attachmentService:           NewAttachmentService(repoRegistry, logger),
		exportDataService:           NewExportDataService(repoRegistry, i18n, logger),
		importDataService:           NewImportDataService(repoRegistry, logger),
		candidateJobStepService:     NewCandidateJobStepService(repoRegistry, logger),
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

// CandiateJobStep returns the CandidateJobStepService.
func (i serviceImpl) CandiateJobStep() CandidateJobStepService {
	return i.candidateJobStepService
}
