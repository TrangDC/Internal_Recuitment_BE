package service

import (
	"context"
	"trec/ent/candidatejob"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CandidateJobStepService interface {
	CreateCandidateJobStep(ctx context.Context, status candidatejob.Status, candidateJobId uuid.UUID, repoRegistry repository.Repository) error
}

type candidateJobStepImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewCandidateJobStepService(repoRegistry repository.Repository, logger *zap.Logger) CandidateJobStepService {
	return &attachmentSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *attachmentSvcImpl) CreateCandidateJobStep(ctx context.Context, status candidatejob.Status, candidateJobId uuid.UUID, repoRegistry repository.Repository) error {
	return repoRegistry.CandidateJobStep().CreateCandidateJobStep(ctx, status, candidateJobId)
}
