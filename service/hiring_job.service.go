package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringJobService interface {
	// mutation
	CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJobResponse, error)
	UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID) (*ent.HiringJobResponse, error)
	UpdateHiringJobStatus(ctx context.Context, status ent.HiringJobStatus, id uuid.UUID) (*ent.HiringJobResponse, error)
	DeleteHiringJob(ctx context.Context, id uuid.UUID) error
	// query
	GetHiringJob(ctx context.Context, id uuid.UUID) (*ent.HiringJobResponse, error)
	GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord, filter *ent.HiringJobFilter, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error)
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
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
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

func (svc *hiringJobSvcImpl) UpdateHiringJobStatus(ctx context.Context, status ent.HiringJobStatus, id uuid.UUID) (*ent.HiringJobResponse, error) {
	hiringJob, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		hiringJob, err = repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, status)
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

func (svc *hiringJobSvcImpl) GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord, filter *ent.HiringJobFilter, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error) {
	var result *ent.HiringJobResponseGetAll
	var edges []*ent.HiringJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.HiringJob().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.HiringJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(hiringjob.FieldCreatedAt)
	if orderBy != nil {
		order = ent.Desc(strings.ToLower(orderBy.Field.String()))
		if orderBy.Direction == ent.OrderDirectionAsc {
			order = ent.Asc(strings.ToLower(orderBy.Field.String()))
		}
	}
	query = query.Order(order)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, index int) *ent.HiringJobEdge {
		return &ent.HiringJobEdge{
			Node: hiringJob,
			Cursor: ent.Cursor{
				Value: hiringJob.ID.String(),
			},
		}
	})
	result = &ent.HiringJobResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

// common function
func (svc *hiringJobSvcImpl) freeWord(hiringJobQuery *ent.HiringJobQuery, input *ent.HiringJobFreeWord) {
	if input != nil {
		if input.Name != nil {
			hiringJobQuery.Where(hiringjob.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
}

func (svc *hiringJobSvcImpl) filter(hiringJobQuery *ent.HiringJobQuery, input *ent.HiringJobFilter) {
	if input != nil {
		if input.Name != nil {
			hiringJobQuery.Where(hiringjob.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.TeamIds != nil {
			ids := lo.Map(input.TeamIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.TeamIDIn(ids...))
		}
		if input.Status != nil {
			hiringJobQuery.Where(hiringjob.StatusEQ(hiringjob.Status(*input.Status)))
		}
	}
}
