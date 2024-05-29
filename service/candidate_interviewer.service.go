package service

import (
	"context"
	"trec/ent"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CandidateInterviewerService interface {
	CreateCandidateInterviewer(ctx context.Context, memberIds []uuid.UUID, candidateInterviews []*ent.CandidateInterview, repoRegistry repository.Repository) ([]*ent.CandidateInterviewer, error)
}

type candidateInterviewerStepImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewCandidateInterviewerService(repoRegistry repository.Repository, logger *zap.Logger) CandidateInterviewerService {
	return &candidateInterviewerStepImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *candidateInterviewerStepImpl) CreateCandidateInterviewer(ctx context.Context, memberIds []uuid.UUID, candidateInterviews []*ent.CandidateInterview, repoRegistry repository.Repository) ([]*ent.CandidateInterviewer, error) {
	return repoRegistry.CandidateInterviewer().CreateBulkCandidateInterview(ctx, memberIds, candidateInterviews)
}
