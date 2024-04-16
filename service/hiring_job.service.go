package service

import (
	"context"
	"net/http"
	"trec/ent"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type HiringJobService interface {
	// mutation
	CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJobResponse, error)
	UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID) (*ent.HiringJobResponse, error)
	DeleteHiringJob(ctx context.Context, id uuid.UUID) error
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

func (svc *hiringJobSvcImpl) CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJobResponse, error) {
	var hiringJob *ent.HiringJob
	err := svc.repoRegistry.HiringJob().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		hiringJob, err = repoRegistry.HiringJob().CreateHiringJob(ctx, input)
		return nil
	})
	result, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, hiringJob.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID) (*ent.HiringJobResponse, error) {
	hiringJob, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.HiringJob().ValidName(ctx, id, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		hiringJob, err = repoRegistry.HiringJob().UpdateHiringJob(ctx, hiringJob, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) GetHiringJob(ctx context.Context, id uuid.UUID) (*ent.HiringJobResponse, error) {
	result, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) DeleteHiringJob(ctx context.Context, id uuid.UUID) error {
	hiringJob, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.HiringJob().DeleteHiringJob(ctx, hiringJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}
