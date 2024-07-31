package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/hiringteam"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringTeamService interface {
	// mutation
	CreateHiringTeam(ctx context.Context, input ent.NewHiringTeamInput, note string) (*ent.HiringTeamResponse, error)
	UpdateHiringTeam(ctx context.Context, hiringTeamID uuid.UUID, input ent.UpdateHiringTeamInput, note string) (*ent.HiringTeamResponse, error)
	DeleteHiringTeam(ctx context.Context, hiringTeamID uuid.UUID, note string) error

	// query
	GetHiringTeam(ctx context.Context, hiringTeamID uuid.UUID) (*ent.HiringTeamResponse, error)
	GetHiringTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringTeamFreeWord,
		filter *ent.HiringTeamFilter, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringTeamFreeWord,
		filter *ent.HiringTeamFilter, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamSelectionResponseGetAll, error)
}

type hiringTeamSvcImpl struct {
	userSvcImpl               UserService
	hiringTeamApproverSvcImpl HiringTeamApproverService
	repoRegistry              repository.Repository
	dtoRegistry               dto.Dto
	logger                    *zap.Logger
}

func NewHiringTeamService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) HiringTeamService {
	return &hiringTeamSvcImpl{
		userSvcImpl:               NewUserService(repoRegistry, dtoRegistry, logger),
		hiringTeamApproverSvcImpl: NewHiringTeamApproverService(repoRegistry, logger),
		repoRegistry:              repoRegistry,
		dtoRegistry:               dtoRegistry,
		logger:                    logger,
	}
}

func (svc *hiringTeamSvcImpl) CreateHiringTeam(ctx context.Context, input ent.NewHiringTeamInput, note string) (*ent.HiringTeamResponse, error) {
	var (
		result    *ent.HiringTeam
		memberIds []uuid.UUID
	)
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	if !payload.ForAll {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	memberIds = lo.Map(input.Members, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	errString, err := svc.repoRegistry.HiringTeam().ValidInput(ctx, uuid.Nil, input.Name, memberIds, len(input.Approvers))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringTeam().CreateHiringTeam(ctx, input, memberIds)
		if err != nil {
			return err
		}
		err := svc.userSvcImpl.UpdateHiringTeam(ctx, result.Name, result.ID, memberIds, note)
		if err != nil {
			return err
		}
		return svc.hiringTeamApproverSvcImpl.HiringTeamApproverMutation(ctx, input.Approvers, result.ID, make([]*ent.HiringTeamApprover, 0))
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.HiringTeam().GetHiringTeam(ctx, result.ID)
	jsonString, err := svc.dtoRegistry.HiringTeam().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleHiringTeams, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.HiringTeamResponse{
		Data: result,
	}, nil
}

func (svc *hiringTeamSvcImpl) UpdateHiringTeam(ctx context.Context, hiringTeamID uuid.UUID, input ent.UpdateHiringTeamInput, note string) (*ent.HiringTeamResponse, error) {
	var (
		memberIds []uuid.UUID
		result    *ent.HiringTeam
	)
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.HiringTeam().GetHiringTeam(ctx, hiringTeamID)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, record) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if len(input.Members) != 0 {
		memberIds = lo.Map(input.Members, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
	}
	errString, err := svc.repoRegistry.HiringTeam().ValidInput(ctx, hiringTeamID, input.Name, memberIds, len(input.Approvers))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	newMemberIds, removeMemberIds := svc.updateMembers(record, memberIds)
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringTeam().UpdateHiringTeam(ctx, record, input, newMemberIds, removeMemberIds)
		if err != nil {
			return err
		}
		err = svc.userSvcImpl.UpdateHiringTeam(ctx, result.Name, result.ID, newMemberIds, note)
		if err != nil {
			return err
		}
		return svc.hiringTeamApproverSvcImpl.HiringTeamApproverMutation(ctx, input.Approvers, hiringTeamID, record.Edges.HiringTeamApprovers)
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.HiringTeam().GetHiringTeam(ctx, hiringTeamID)
	jsonString, err := svc.dtoRegistry.HiringTeam().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, hiringTeamID, audittrail.ModuleHiringTeams, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.HiringTeamResponse{
		Data: result,
	}, nil
}

func (svc *hiringTeamSvcImpl) DeleteHiringTeam(ctx context.Context, hiringTeamID uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	if !payload.ForAll {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	team, err := svc.repoRegistry.HiringTeam().GetHiringTeam(ctx, hiringTeamID)
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if len(team.Edges.HiringTeamJobEdges) != 0 {
		return util.WrapGQLError(ctx, "model.hiring_teams.validation.hiring_job_exist", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds := lo.Map(team.Edges.UserEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.HiringTeam().DeleteHiringTeam(ctx, team, memberIds)
		if err != nil {
			return err
		}
		err = repoRegistry.HiringTeam().DeleteRelationHiringTeam(ctx, hiringTeamID)
		if err != nil {
			return err
		}
		err = svc.userSvcImpl.RemoveHiringTeam(ctx, team.ID, memberIds, note)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.HiringTeam().AuditTrailDelete(team)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, hiringTeamID, audittrail.ModuleHiringTeams, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return nil
}

func (svc *hiringTeamSvcImpl) GetHiringTeam(ctx context.Context, hiringTeamID uuid.UUID) (*ent.HiringTeamResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.HiringTeam().BuildQuery()
	svc.validPermissionGet(payload, query)
	team, err := svc.repoRegistry.HiringTeam().BuildGetOne(ctx, query.Where(hiringteam.IDEQ(hiringTeamID)))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.HiringTeamResponse{
		Data: team,
	}, nil
}

func (svc *hiringTeamSvcImpl) GetHiringTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringTeamFreeWord, filter *ent.HiringTeamFilter, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var (
		result  *ent.HiringTeamResponseGetAll
		edges   []*ent.HiringTeamEdge
		page    int
		perPage int
		teams   []*ent.HiringTeam
		count   int
		err     error
	)
	query := svc.repoRegistry.HiringTeam().BuildQuery()
	svc.validPermissionGet(payload, query)
	teams, count, err = svc.getAllHiringTeams(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(teams, func(entity *ent.HiringTeam, index int) *ent.HiringTeamEdge {
		return &ent.HiringTeamEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	result = &ent.HiringTeamResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *hiringTeamSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringTeamFreeWord,
	filter *ent.HiringTeamFilter, orderBy ent.HiringTeamOrderBy) (*ent.HiringTeamSelectionResponseGetAll, error) {
	var (
		result  *ent.HiringTeamSelectionResponseGetAll
		edges   []*ent.HiringTeamSelectionEdge
		page    int
		perPage int
		teams   []*ent.HiringTeam
		count   int
	)
	query := svc.repoRegistry.HiringTeam().BuildBaseQuery()
	teams, count, err := svc.getAllHiringTeams(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		return nil, err
	}
	edges = lo.Map(teams, func(entity *ent.HiringTeam, index int) *ent.HiringTeamSelectionEdge {
		return &ent.HiringTeamSelectionEdge{
			Node: &ent.HiringTeamSelection{
				ID:   entity.ID.String(),
				Name: entity.Name,
			},
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	result = &ent.HiringTeamSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *hiringTeamSvcImpl) getAllHiringTeams(ctx context.Context, query *ent.HiringTeamQuery, pagination *ent.PaginationInput, freeWord *ent.HiringTeamFreeWord, filter *ent.HiringTeamFilter, orderBy ent.HiringTeamOrderBy) ([]*ent.HiringTeam, int, error) {
	var (
		page    int
		perPage int
		teams   []*ent.HiringTeam
		count   int
		err     error
	)
	svc.filter(ctx, query, filter)
	svc.freeWord(query, freeWord)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	if ent.HiringTeamOrderByAdditionalField.IsValid(ent.HiringTeamOrderByAdditionalField(orderBy.Field.String())) {
		count, teams, err = svc.getTeamListByAdditionOrder(ctx, query, page, perPage, orderBy)
		if err != nil {
			return nil, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	} else {
		count, teams, err = svc.getTeamsListByNormalOrder(ctx, query, page, perPage, orderBy)
		if err != nil {
			return nil, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	return teams, count, nil
}

// common function
func (svc *hiringTeamSvcImpl) freeWord(hiringTeamQuery *ent.HiringTeamQuery, input *ent.HiringTeamFreeWord) {
	if input != nil {
		if input.Name != nil {
			hiringTeamQuery.Where(hiringteam.NameContainsFold(*input.Name))
		}
	}
}

func (svc *hiringTeamSvcImpl) filter(ctx context.Context, hiringTeamQuery *ent.HiringTeamQuery, input *ent.HiringTeamFilter) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	if input != nil {
		if input.Name != nil {
			hiringTeamQuery.Where(hiringteam.NameEqualFold(*input.Name))
		}
		if input.ForOwner != nil {
			if *input.ForOwner {
				hiringTeamQuery.Where(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)))
			} else {
				hiringTeamQuery.Where(hiringteam.IDEQ(uuid.Nil))
			}
		}
		if input.ForHiringTeam != nil {
			if *input.ForHiringTeam {
				hiringTeamQuery.Where(hiringteam.Or(
					hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID)),
					hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)),
				))
			} else {
				hiringTeamQuery.Where(hiringteam.IDEQ(uuid.Nil))
			}
		}
		if input.ManagerIds != nil && len(input.ManagerIds) > 0 {
			managerIDs := lo.Map(input.ManagerIds, func(id *string, _ int) uuid.UUID {
				return uuid.MustParse(*id)
			})
			hiringTeamQuery.Where(hiringteam.HasUserEdgesWith(user.IDIn(managerIDs...)))
		}
	}
}

func (svc *hiringTeamSvcImpl) updateMembers(record *ent.HiringTeam, memberIds []uuid.UUID) ([]uuid.UUID, []uuid.UUID) {
	var newMemberIds []uuid.UUID
	var removeMemberIds []uuid.UUID
	currentMemberIds := lo.Map(record.Edges.UserEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	newMemberIds = lo.Filter(memberIds, func(memberId uuid.UUID, index int) bool {
		return !lo.Contains(currentMemberIds, memberId)
	})
	removeMemberIds = lo.Filter(currentMemberIds, func(memberId uuid.UUID, index int) bool {
		return !lo.Contains(memberIds, memberId)
	})
	return newMemberIds, removeMemberIds
}

func (svc hiringTeamSvcImpl) getTeamsListByNormalOrder(ctx context.Context, query *ent.HiringTeamQuery, page int, perPage int, orderBy ent.HiringTeamOrderBy) (int, []*ent.HiringTeam, error) {
	count, err := svc.repoRegistry.HiringTeam().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return 0, nil, err
	}
	if orderBy.Direction == ent.OrderDirectionAsc {
		query = query.Order(ent.Asc(strings.ToLower(orderBy.Field.String())))
	} else {
		query = query.Order(ent.Desc(strings.ToLower(orderBy.Field.String())))
	}
	if perPage != 0 && page != 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	teams, err := svc.repoRegistry.HiringTeam().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return 0, nil, err
	}
	return count, teams, nil
}

func (svc hiringTeamSvcImpl) getTeamListByAdditionOrder(ctx context.Context, query *ent.HiringTeamQuery, page int, perPage int, orderBy ent.HiringTeamOrderBy) (int, []*ent.HiringTeam, error) {
	teams, err := svc.repoRegistry.HiringTeam().BuildList(ctx, query.Order(ent.Desc(ent.HiringTeamOrderFieldCreatedAt.String())))
	if err != nil {
		svc.logger.Error(err.Error())
		return 0, nil, err
	}
	count := len(teams)
	switch orderBy.Field {
	case ent.HiringTeamOrderByFieldOpeningRequests:
		sort.Slice(teams, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return len(teams[i].Edges.HiringTeamJobEdges) < len(teams[j].Edges.HiringTeamJobEdges)
			} else {
				return len(teams[i].Edges.HiringTeamJobEdges) > len(teams[j].Edges.HiringTeamJobEdges)
			}
		})
	case ent.HiringTeamOrderByFieldNewestApplied:
		blankTeams := lo.Filter(teams, func(team *ent.HiringTeam, _ int) bool {
			return len(team.Edges.HiringTeamJobEdges) == 0
		})
		teams = lo.Filter(teams, func(team *ent.HiringTeam, _ int) bool {
			return len(team.Edges.HiringTeamJobEdges) != 0
		})
		sort.Slice(teams, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return (teams[i].Edges.HiringTeamJobEdges[0].LastApplyDate.After(teams[j].Edges.HiringTeamJobEdges[0].LastApplyDate))
			} else {
				return teams[i].Edges.HiringTeamJobEdges[0].LastApplyDate.Before(teams[j].Edges.HiringTeamJobEdges[0].LastApplyDate)
			}
		})
		teams = append(teams, blankTeams...)
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

// permission
func (svc *hiringTeamSvcImpl) validPermissionUpdate(payload *middleware.Payload, record *ent.HiringTeam) bool {
	if payload.ForAll {
		return true
	}
	currentManagerIds := lo.Map(record.Edges.UserEdges, func(user *ent.User, _ int) uuid.UUID {
		return user.ID
	})
	currentMemberIds := lo.Map(record.Edges.HiringMemberEdges, func(user *ent.User, _ int) uuid.UUID {
		return user.ID
	})
	if payload.ForOwner && payload.ForTeam && (lo.Contains(currentManagerIds, payload.UserID) || lo.Contains(currentMemberIds, payload.UserID)) {
		return true
	}
	if payload.ForOwner && lo.Contains(currentManagerIds, payload.UserID) {
		return true
	}
	if payload.ForTeam && lo.Contains(currentMemberIds, payload.UserID) {
		return true
	}
	return false
}

func (svc *hiringTeamSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.HiringTeamQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForOwner && payload.ForTeam {
		query.Where(hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))))
	}
	if payload.ForOwner {
		query.Where(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)))
	}
	if payload.ForTeam {
		query.Where(hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID)))
	}
}

// Path: service/team.service.go
