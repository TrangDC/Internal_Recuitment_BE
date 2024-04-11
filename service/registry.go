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
}

// serviceImpl is the implementation of Service.
type serviceImpl struct {
	authService    AuthService
	storageService StorageService
	UserService    UserService
	TeamService    TeamService
}

// NewService creates a new Service.
func NewService(azureADOAuthClient azuread.AzureADOAuth, azureStorage azurestorage.AzureStorage, entClient *ent.Client, logger *zap.Logger) Service {
	repoRegistry := repository.NewRepository(entClient)

	return &serviceImpl{
		authService:    NewAuthService(azureADOAuthClient, logger),
		storageService: NewStorageService(azureStorage, logger),
		UserService:    NewUserService(repoRegistry, logger),
		TeamService:    NewTeamService(repoRegistry, logger),
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
	return i.UserService
}

// Team returns the TeamService.
func (i serviceImpl) Team() TeamService {
	return i.TeamService
}
