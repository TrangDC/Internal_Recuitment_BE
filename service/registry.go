package service

import (
	"trec/ent"
	"trec/internal/azuread"
	"trec/internal/azurestorage"
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
	Attachment() AttachmentService
}

// serviceImpl is the implementation of Service.
type serviceImpl struct {
	authService         AuthService
	storageService      StorageService
	userService         UserService
	teamService         TeamService
	hiringJobService    HiringJobService
	auditTrailService   AuditTrailService
	candidateService    CandidateService
	candidateJobService CandidateJobService
	attachmentService   AttachmentService
}

// NewService creates a new Service.
func NewService(azureADOAuthClient azuread.AzureADOAuth, azureStorage azurestorage.AzureStorage, entClient *ent.Client, logger *zap.Logger) Service {
	repoRegistry := repository.NewRepository(entClient)

	return &serviceImpl{
		authService:         NewAuthService(azureADOAuthClient, logger),
		storageService:      NewStorageService(azureStorage, logger),
		userService:         NewUserService(repoRegistry, logger),
		teamService:         NewTeamService(repoRegistry, logger),
		hiringJobService:    NewHiringJobService(repoRegistry, logger),
		auditTrailService:   NewAuditTrailService(repoRegistry, logger),
		candidateService:    NewCandidateService(repoRegistry, logger),
		candidateJobService: NewCandidateJobService(repoRegistry, logger),
		attachmentService:   NewAttachmentService(repoRegistry, logger),
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

// Attachment returns the AttachmentService.
func (i serviceImpl) Attachment() AttachmentService {
	return i.attachmentService
}
