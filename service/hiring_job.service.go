package service

import (
	"trec/repository"

	"go.uber.org/zap"
)

type HiringJobService interface {
}

type hiringJobSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewHiringJobService(repoRegistry repository.Repository, logger *zap.Logger) HiringJobService {
	return &hiringJobSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}
