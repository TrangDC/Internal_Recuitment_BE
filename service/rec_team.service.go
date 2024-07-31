package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/predicate"
	"trec/ent/recteam"
	"trec/ent/user"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type RecTeamService interface {
	// mutation
	CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput, note string) (*ent.RecTeamResponse, error)
	UpdateRecTeam(ctx context.Context, id string, input ent.UpdateRecTeamInput, note string) (*ent.RecTeamResponse, error)
	DeleteRecTeam(ctx context.Context, id uuid.UUID, note string) error

	// query
	GetRecTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord,
		filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord,
		filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamSelectionResponseGetAll, error)
}

type recTeamSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewRecTeamService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) RecTeamService {
	return &recTeamSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *recTeamSvcImpl) CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput, note string) (*ent.RecTeamResponse, error) {
	var result *ent.RecTeam
	errString, err := svc.repoRegistry.RecTeam().ValidInput(ctx, uuid.Nil, input.Name, uuid.MustParse(input.LeaderID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.RecTeam().CreateRecTeam(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	results, _ := svc.repoRegistry.RecTeam().GetRecTeam(ctx, result.ID)
	return &ent.RecTeamResponse{
		Data: results,
	}, nil
}

func (svc *recTeamSvcImpl) DeleteRecTeam(ctx context.Context, id uuid.UUID, note string) error {
	recTeam, err := svc.repoRegistry.RecTeam().GetRecTeam(ctx, id)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(recTeam.Edges.RecMemberEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.RecTeam().DeleteRecTeam(ctx, recTeam, memberIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *recTeamSvcImpl) UpdateRecTeam(ctx context.Context, recTeamId string, input ent.UpdateRecTeamInput, note string) (*ent.RecTeamResponse, error) {
	var result *ent.RecTeam
	record, err := svc.repoRegistry.RecTeam().GetRecTeam(ctx, uuid.MustParse(recTeamId))
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.RecTeam().ValidInput(ctx, uuid.MustParse(recTeamId), input.Name, uuid.MustParse(input.LeaderID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.RecTeam().UpdateRecTeam(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.RecTeam().GetRecTeam(ctx, uuid.MustParse(recTeamId))
	return &ent.RecTeamResponse{
		Data: result,
	}, nil
}

// query
func (svc *recTeamSvcImpl) GetRecTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord, filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamResponseGetAll, error) {
	var (
		result   *ent.RecTeamResponseGetAll
		edges    []*ent.RecTeamEdge
		page     int
		perPage  int
		recTeams []*ent.RecTeam
		count    int
		err      error
	)
	query := svc.repoRegistry.RecTeam().BuildQuery()
	recTeams, count, page, perPage, err = svc.getAllRecTeams(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(recTeams, func(entity *ent.RecTeam, index int) *ent.RecTeamEdge {
		return &ent.RecTeamEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	result = &ent.RecTeamResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}

	return result, nil
}

func (svc *recTeamSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord, filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamSelectionResponseGetAll, error) {
	var (
		result   *ent.RecTeamSelectionResponseGetAll
		edges    []*ent.RecTeamSelectionEdge
		page     int
		perPage  int
		recTeams []*ent.RecTeam
		count    int
		err      error
	)
	query := svc.repoRegistry.RecTeam().BuildBaseQuery()
	recTeams, count, page, perPage, err = svc.getAllRecTeams(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(recTeams, func(entity *ent.RecTeam, index int) *ent.RecTeamSelectionEdge {
		return &ent.RecTeamSelectionEdge{
			Node: &ent.RecTeamSelection{
				ID:   entity.ID.String(),
				Name: entity.Name,
			},
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	result = &ent.RecTeamSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Page:    page,
			PerPage: perPage,
			Total:   count,
		},
	}

	return result, nil
}

func (svc *recTeamSvcImpl) getAllRecTeams(ctx context.Context, query *ent.RecTeamQuery, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord, filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) ([]*ent.RecTeam, int, int, int, error) {
	var (
		page     int
		perPage  int
		recTeams []*ent.RecTeam
		count    int
		err      error
	)
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err = svc.repoRegistry.RecTeam().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(recteam.FieldCreatedAt)
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
	recTeams, err = svc.repoRegistry.RecTeam().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return recTeams, count, page, perPage, nil
}

// common function
func (svc *recTeamSvcImpl) freeWord(recTeamQuery *ent.RecTeamQuery, input *ent.RecTeamFreeWord) {
	var recTeams []predicate.RecTeam
	if input != nil {
		if input.Name != nil {
			recTeams = append(recTeams, recteam.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
		if input.Description != nil {
			recTeams = append(recTeams, recteam.DescriptionContainsFold(strings.TrimSpace(*input.Description)))
		}
	}
	if len(recTeams) > 0 {
		recTeamQuery.Where(recteam.Or(recTeams...))
	}
}

func (svc *recTeamSvcImpl) filter(recTeamQuery *ent.RecTeamQuery, input *ent.RecTeamFilter) {
	if input != nil {
		if input.Name != nil {
			recTeamQuery.Where(recteam.NameContainsFold(*input.Name))
		}
		if input.LeaderIds != nil && len(input.LeaderIds) > 0 {
			leaderIDs := lo.Map(input.LeaderIds, func(id *string, _ int) uuid.UUID {
				return uuid.MustParse(*id)
			})
			recTeamQuery.Where(recteam.HasRecLeaderEdgeWith(user.IDIn(leaderIDs...)))
		}
	}
}

// Path: service/rec_team.service.go
