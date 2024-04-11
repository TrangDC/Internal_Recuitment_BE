package service

import (
	"trec/repository"

	"go.uber.org/zap"
)

type TeamMngService interface {
}

type teamMngSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewTeamMngService(repoRegistry repository.Repository, logger *zap.Logger) TeamMngService {
	return &teamMngSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}
