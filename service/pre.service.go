package service

import (
	"context"
	"trec/repository"

	"go.uber.org/zap"
)

type PreService interface {
	PreFunction(ctx context.Context) (string, error)
}

type jobTitleSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewPreService(repoRegistry repository.Repository, logger *zap.Logger) PreService {
	return &jobTitleSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc jobTitleSvcImpl) PreFunction(ctx context.Context) (string, error) {
	result, _ := svc.repoRegistry.JobTitle().PreFunction(ctx)
	svc.logger.Info("PreService.PreFunction", zap.String("[Job_Title]:", result))
	return "Success - Service", nil
}
