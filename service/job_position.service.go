package service

import (
	"context"
	"net/http"
	"trec/dto"
	"trec/ent"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type JobPositionService interface {
	// mutation
	CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPositionResponse, error)
}

type jobPositionSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewJobPositionService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) JobPositionService {
	return &jobPositionSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *jobPositionSvcImpl) CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPositionResponse, error) {
	var result *ent.JobPosition
	errString, err := svc.repoRegistry.JobPosition().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = svc.repoRegistry.JobPosition().CreateJobPosition(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.JobPosition().GetJobPosition(ctx, result.ID)
	return &ent.JobPositionResponse{
		Data: result,
	}, nil
}

// Path: service/job_position.service.go
