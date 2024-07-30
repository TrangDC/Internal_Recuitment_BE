package service

import (
	"context"
	"github.com/samber/lo"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/jobposition"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type JobPositionService interface {
	// mutation
	CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPositionResponse, error)
	UpdateJobPosition(ctx context.Context, jobPositionId uuid.UUID, input ent.UpdateJobPositionInput) (*ent.JobPositionResponse, error)
	DeleteJobPosition(ctx context.Context, jobPositionId uuid.UUID) error

	// query
	GetJobPosition(ctx context.Context, jobPositionId uuid.UUID) (*ent.JobPositionResponse, error)
	GetJobPositions(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.JobPositionFreeWord,
		filter *ent.JobPositionFilter, orderBy *ent.JobPositionOrder) (*ent.JobPositionResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.JobPositionFreeWord,
		filter *ent.JobPositionFilter, orderBy *ent.JobPositionOrder) (*ent.JobPositionSelectionResponseGetAll, error)
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

func (svc *jobPositionSvcImpl) UpdateJobPosition(ctx context.Context, jobPositionId uuid.UUID, input ent.UpdateJobPositionInput) (*ent.JobPositionResponse, error) {
	var result *ent.JobPosition
	record, err := svc.repoRegistry.JobPosition().GetJobPosition(ctx, jobPositionId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.JobPosition().ValidName(ctx, jobPositionId, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.JobPosition().UpdateJobPosition(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.JobPosition().GetJobPosition(ctx, jobPositionId)
	return &ent.JobPositionResponse{
		Data: result,
	}, nil
}

func (svc *jobPositionSvcImpl) DeleteJobPosition(ctx context.Context, jobPositionId uuid.UUID) error {
	jobPositionRecord, err := svc.repoRegistry.JobPosition().GetJobPosition(ctx, jobPositionId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.JobPosition().DeleteJobPosition(ctx, jobPositionRecord)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

// query
func (svc *jobPositionSvcImpl) GetJobPosition(ctx context.Context, jobPositionId uuid.UUID) (*ent.JobPositionResponse, error) {
	result, err := svc.repoRegistry.JobPosition().GetJobPosition(ctx, jobPositionId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.JobPositionResponse{
		Data: result,
	}, nil
}

func (svc *jobPositionSvcImpl) GetJobPositions(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.JobPositionFreeWord,
	filter *ent.JobPositionFilter, orderBy *ent.JobPositionOrder) (*ent.JobPositionResponseGetAll, error) {
	var edges []*ent.JobPositionEdge
	jobPositions, count, page, perPage, err := svc.getAllJobPosition(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(jobPositions, func(entity *ent.JobPosition, index int) *ent.JobPositionEdge {
		return &ent.JobPositionEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.JobPositionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc *jobPositionSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.JobPositionFreeWord,
	filter *ent.JobPositionFilter, orderBy *ent.JobPositionOrder) (*ent.JobPositionSelectionResponseGetAll, error) {
	var edges []*ent.JobPositionSelectionEdge
	jobPositions, count, page, perPage, err := svc.getAllJobPosition(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(jobPositions, func(entity *ent.JobPosition, index int) *ent.JobPositionSelectionEdge {
		return &ent.JobPositionSelectionEdge{
			Node: &ent.JobPositionSelection{
				ID:   entity.ID.String(),
				Name: entity.Name,
			},
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.JobPositionSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc *jobPositionSvcImpl) getAllJobPosition(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.JobPositionFreeWord,
	filter *ent.JobPositionFilter, orderBy *ent.JobPositionOrder) ([]*ent.JobPosition, int, int, int, error) {
	var page int
	var perPage int
	query := svc.repoRegistry.JobPosition().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.JobPosition().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(jobposition.FieldCreatedAt)
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
	jobPositions, err := svc.repoRegistry.JobPosition().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return jobPositions, count, page, perPage, nil
}

// common function
func (svc *jobPositionSvcImpl) freeWord(query *ent.JobPositionQuery, input *ent.JobPositionFreeWord) {
	var jobPositions []predicate.JobPosition
	if input != nil {
		if input.Name != nil {
			jobPositions = append(jobPositions, jobposition.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
	if len(jobPositions) > 0 {
		query.Where(jobposition.Or(jobPositions...))
	}
}

func (svc *jobPositionSvcImpl) filter(query *ent.JobPositionQuery, input *ent.JobPositionFilter) {
	if input != nil {
		if input.Name != nil {
			query.Where(jobposition.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
	}
}

// Path: service/job_position.service.go
