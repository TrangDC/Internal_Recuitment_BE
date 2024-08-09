package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/predicate"
	"trec/ent/recteam"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
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
	GetRecTeam(ctx context.Context, id uuid.UUID) (*ent.RecTeamResponse, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord,
		filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamSelectionResponseGetAll, error)
}

type recTeamSvcImpl struct {
	userSvcImpl  UserService
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewRecTeamService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) RecTeamService {
	return &recTeamSvcImpl{
		userSvcImpl:  NewUserService(repoRegistry, dtoRegistry, logger),
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *recTeamSvcImpl) CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput, note string) (*ent.RecTeamResponse, error) {
	var record *ent.RecTeam
	errString, err, userRecord := svc.repoRegistry.RecTeam().ValidInput(ctx, uuid.Nil, input.Name, uuid.MustParse(input.LeaderID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = svc.userSvcImpl.RemoveRecTeam(ctx, []uuid.UUID{userRecord.ID}, note, repoRegistry)
		if err != nil {
			return err
		}
		record, err = repoRegistry.RecTeam().CreateRecTeam(ctx, input)
		if err != nil {
			return err
		}
		err = svc.userSvcImpl.AuditTrailCreateRecTeamWLeader(ctx, userRecord, input.Name, note, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.RecTeam().GetRecTeam(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.RecTeam().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleRecTeams, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.RecTeamResponse{
		Data: result,
	}, nil
}

func (svc *recTeamSvcImpl) DeleteRecTeam(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.RecTeam().GetRecTeam(ctx, id)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(record.Edges.RecMemberEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := repoRegistry.RecTeam().DeleteRecTeam(ctx, record, memberIds)
		if err != nil {
			return err
		}
		err = svc.userSvcImpl.SetRecTeam(ctx, "", uuid.Nil, memberIds, record.Name, note, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.RecTeam().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, id, audittrail.ModuleRecTeams, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *recTeamSvcImpl) UpdateRecTeam(ctx context.Context, recTeamId string, input ent.UpdateRecTeamInput, note string) (*ent.RecTeamResponse, error) {
	var result *ent.RecTeam
	var oldUserRec string
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.RecTeam().GetRecTeam(ctx, uuid.MustParse(recTeamId))
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, record) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	errString, err, userRecord := svc.repoRegistry.RecTeam().ValidInput(ctx, uuid.MustParse(recTeamId), input.Name, uuid.MustParse(input.LeaderID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if userRecord.Edges.RecTeams != nil {
		oldUserRec = userRecord.Edges.RecTeams.Name
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.RecTeam().UpdateRecTeam(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.userSvcImpl.SetRecTeam(ctx, record.Name, record.ID, []uuid.UUID{uuid.MustParse(input.LeaderID)}, oldUserRec, note, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.RecTeam().GetRecTeam(ctx, uuid.MustParse(recTeamId))
	jsonString, err := svc.dtoRegistry.RecTeam().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, uuid.MustParse(recTeamId), audittrail.ModuleRecTeams, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.RecTeamResponse{
		Data: result,
	}, nil
}

// query
func (svc *recTeamSvcImpl) GetRecTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord,
	filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
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
	svc.validPermissionGet(payload, query)
	recTeams, count, page, perPage, err = svc.getAllRecTeams(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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

func (svc *recTeamSvcImpl) GetRecTeam(ctx context.Context, id uuid.UUID) (*ent.RecTeamResponse, error) {
	recTeam, err := svc.repoRegistry.RecTeam().GetRecTeam(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.RecTeamResponse{
		Data: recTeam,
	}, nil
}

func (svc *recTeamSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RecTeamFreeWord,
	filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) (*ent.RecTeamSelectionResponseGetAll, error) {
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
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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

func (svc *recTeamSvcImpl) getAllRecTeams(ctx context.Context, query *ent.RecTeamQuery, pagination *ent.PaginationInput,
	freeWord *ent.RecTeamFreeWord, filter *ent.RecTeamFilter, orderBy *ent.RecTeamOrderBy) ([]*ent.RecTeam, int, int, int, error) {
	var (
		page     int
		perPage  int
		recTeams []*ent.RecTeam
		count    int
		err      error
	)
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	if orderBy == nil || orderBy.Field != ent.RecTeamOrderByFieldLeader {
		count, err = svc.repoRegistry.RecTeam().BuildCount(ctx, query)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, 0, 0, 0, err
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
			query = query.Limit(perPage).Offset((page - 1) * perPage)
		}
		recTeams, err = svc.repoRegistry.RecTeam().BuildList(ctx, query)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, 0, 0, 0, err
		}
	} else {
		count, recTeams, err = svc.getTeamsListByAdditionalOrder(ctx, query, page, perPage, orderBy)
		if err != nil {
			return nil, 0, 0, 0, err
		}
	}
	return recTeams, count, page, perPage, nil
}

// common function
func (svc *recTeamSvcImpl) getTeamsListByAdditionalOrder(ctx context.Context, query *ent.RecTeamQuery, page, perPage int, orderBy *ent.RecTeamOrderBy) (int, []*ent.RecTeam, error) {
	teams, err := svc.repoRegistry.RecTeam().BuildList(ctx, query.Order(ent.Desc(ent.RecTeamOrderByFieldCreatedAt.String())))
	if err != nil {
		svc.logger.Error(err.Error())
		return 0, nil, err
	}
	count := len(teams)
	switch orderBy.Field {
	case ent.RecTeamOrderByFieldLeader:
		sort.Slice(teams, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return teams[i].Edges.RecLeaderEdge.Name < teams[j].Edges.RecLeaderEdge.Name
			} else {
				return teams[i].Edges.RecLeaderEdge.Name > teams[j].Edges.RecLeaderEdge.Name
			}
		})
	}
	// Split slice by page and perPage
	if page != 0 && perPage != 0 {
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(teams) {
			return count, nil, nil
		}
		if start <= len(teams) && end > len(teams) {
			return count, teams[start:], nil
		}
		teams = teams[start:end]
	}
	return count, teams, nil
}

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

// permission
func (svc *recTeamSvcImpl) validPermissionUpdate(payload *middleware.Payload, record *ent.RecTeam) bool {
	if payload.ForAll {
		return true
	}
	currentLeaderId := record.LeaderID
	currentMemberIds := lo.Map(record.Edges.RecMemberEdges, func(user *ent.User, _ int) uuid.UUID {
		return user.ID
	})
	if payload.ForOwner && payload.UserID == currentLeaderId {
		return true
	}
	if payload.ForTeam && lo.Contains(currentMemberIds, payload.UserID) {
		return true
	}
	return false
}

func (svc *recTeamSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.RecTeamQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForOwner {
		query.Where(recteam.LeaderIDEQ(payload.UserID))
	}
	if payload.ForTeam {
		query.Where(recteam.HasRecMemberEdgesWith(user.IDEQ(payload.UserID)))
	}
}

// Path: service/rec_team.service.go
