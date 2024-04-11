package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/team"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type TeamService interface {
	// mutation
	CreateTeam(ctx context.Context, input ent.NewTeamInput) (*ent.Team, error)
	UpdateTeam(ctx context.Context, teamId uuid.UUID, input ent.UpdateTeamInput) (*ent.Team, error)
	DeleteTeam(ctx context.Context, teamId uuid.UUID) error
	// query
	GetTeam(ctx context.Context, teamId uuid.UUID) (*ent.Team, error)
	GetTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.TeamFreeWord, filter *ent.TeamFilter, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error)
	//
}

type teamSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewTeamService(repoRegistry repository.Repository, logger *zap.Logger) TeamService {
	return &teamSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *teamSvcImpl) CreateTeam(ctx context.Context, input ent.NewTeamInput) (*ent.Team, error) {
	var team *ent.Team
	var memberIds []uuid.UUID
	err := svc.repoRegistry.Team().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds = lo.Map(input.Members, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	err = svc.repoRegistry.Team().ValidUserInAnotherTeam(ctx, uuid.Nil, memberIds)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		team, err = repoRegistry.Team().CreateTeam(ctx, input, memberIds)
		if err != nil {
			return err
		}
		return nil
	})
	result, err := svc.repoRegistry.Team().GetTeam(ctx, team.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return result, nil
}

func (svc *teamSvcImpl) UpdateTeam(ctx context.Context, teamId uuid.UUID, input ent.UpdateTeamInput) (*ent.Team, error) {
	var memberIds []uuid.UUID
	team, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.Team().ValidName(ctx, teamId, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds = lo.Map(input.Members, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	newMemberIds, removeMemberids := svc.updateMembers(team, memberIds)
	err = svc.repoRegistry.Team().ValidUserInAnotherTeam(ctx, teamId, memberIds)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		team, err = repoRegistry.Team().UpdateTeam(ctx, team, input, newMemberIds, removeMemberids)
		if err != nil {
			return err
		}
		return nil
	})
	result, err := svc.repoRegistry.Team().GetTeam(ctx, team.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return result, nil
}

func (svc *teamSvcImpl) DeleteTeam(ctx context.Context, teamId uuid.UUID) error {
	team, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(team.Edges.UserEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Team().DeleteTeam(ctx, team, memberIds)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (svc *teamSvcImpl) GetTeam(ctx context.Context, teamId uuid.UUID) (*ent.Team, error) {
	team, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return team, nil
}

func (svc *teamSvcImpl) GetTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.TeamFreeWord, filter *ent.TeamFilter, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error) {
	var result *ent.TeamResponseGetAll
	var edges []*ent.TeamEdge
	query := svc.repoRegistry.Team().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.Team().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if orderBy != nil {
		if orderBy.Direction == ent.OrderDirectionAsc {
			query = query.Order(ent.Asc(strings.ToLower(orderBy.Field.String())))
		} else {
			query = query.Order(ent.Desc(strings.ToLower(orderBy.Field.String())))
		}
	} else {
		query = query.Order(ent.Desc(team.FieldCreatedAt))
	}
	if pagination != nil {
		query = query.Limit(*pagination.PerPage).Offset((*pagination.Page - 1) * *pagination.PerPage)
	}
	teams, err := svc.repoRegistry.Team().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(teams, func(team *ent.Team, index int) *ent.TeamEdge {
		return &ent.TeamEdge{
			Node: team,
			Cursor: ent.Cursor{
				Value: team.ID.String(),
			},
		}
	})
	result = &ent.TeamResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    *pagination.Page,
			PerPage: *pagination.PerPage,
		},
	}
	return result, nil
}

// common function
func (svc *teamSvcImpl) freeWord(teamQuery *ent.TeamQuery, input *ent.TeamFreeWord) {
	if input != nil {
		if input.Name != nil {
			teamQuery.Where(team.NameContainsFold(*input.Name))
		}
	}
}

func (svc *teamSvcImpl) filter(teamQuery *ent.TeamQuery, input *ent.TeamFilter) {
	if input != nil {
		if input.Name != nil {
			teamQuery.Where(team.NameEqualFold(*input.Name))
		}
	}
}

func (svc *teamSvcImpl) updateMembers(record *ent.Team, memberIds []uuid.UUID) ([]uuid.UUID, []uuid.UUID) {
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

// Path: service/team.service.go
