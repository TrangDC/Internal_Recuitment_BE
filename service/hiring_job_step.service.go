package service

import (
	"context"
	"trec/ent/hiringjobstep"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HiringJobStepService interface {
	CreateHiringJobStep(ctx context.Context, step hiringjobstep.Status, hiringJobId uuid.UUID, repoRegistry repository.Repository) error
}

type hiringJobStepImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewHiringJobStepService(repoRegistry repository.Repository, logger *zap.Logger) HiringJobStepService {
	return &hiringJobStepImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *hiringJobStepImpl) CreateHiringJobStep(ctx context.Context, step hiringjobstep.Status, hiringJobId uuid.UUID, repoRegistry repository.Repository) error {
	return repoRegistry.HiringJobStep().CreateHiringJobStep(ctx, step, hiringJobId)
}
