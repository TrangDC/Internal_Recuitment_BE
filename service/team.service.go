package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/team"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type TeamService interface {
	formatTeamField(input string) string
	// mutation
	CreateTeam(ctx context.Context, input ent.NewTeamInput, note string) (*ent.TeamResponse, error)
	UpdateTeam(ctx context.Context, teamId uuid.UUID, input ent.UpdateTeamInput, note string) (*ent.TeamResponse, error)
	DeleteTeam(ctx context.Context, teamId uuid.UUID, note string) error
	// query
	GetTeam(ctx context.Context, teamId uuid.UUID) (*ent.TeamResponse, error)
	GetTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.TeamFreeWord,
		filter *ent.TeamFilter, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error)
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

func (svc *teamSvcImpl) formatTeamField(input string) string {
	switch input {
	case "Name":
		return "model.team.name"
	case "Members":
		return "model.team.members"
	}
	return ""
}

func (svc *teamSvcImpl) CreateTeam(ctx context.Context, input ent.NewTeamInput, note string) (*ent.TeamResponse, error) {
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
		return err
	})
	result, err := svc.repoRegistry.Team().GetTeam(ctx, team.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	record := models.AuditTrailData{
		Module: "module.team",
		Update: make([]interface{}, 0),
		Delete: make([]interface{}, 0),
	}
	record.Create = append(record.Create, models.AuditTrailCreateDelete{
		Field: svc.formatTeamField("Name"),
		Value: result.Name,
	})
	membersName := lo.Map(result.Edges.UserEdges, func(item *ent.User, index int) string {
		return item.Name
	})
	record.Create = append(record.Create, models.AuditTrailCreateDelete{
		Field: svc.formatTeamField("Members"),
		Value: membersName,
	})
	recordChanges, _ := json.Marshal([]interface{}{record})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleTeams, string(recordChanges), audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.TeamResponse{
		Data: result,
	}, nil
}

func (svc *teamSvcImpl) UpdateTeam(ctx context.Context, teamId uuid.UUID, input ent.UpdateTeamInput, note string) (*ent.TeamResponse, error) {
	var memberIds []uuid.UUID
	record, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.Team().ValidName(ctx, teamId, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	var result *ent.Team
	if len(input.Members) != 0 {
		memberIds = lo.Map(input.Members, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
	}
	newMemberIds, removeMemberIds := svc.updateMembers(record, memberIds)
	err = svc.repoRegistry.Team().ValidUserInAnotherTeam(ctx, teamId, memberIds)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Team().UpdateTeam(ctx, record, input, newMemberIds, removeMemberIds)
		return err
	})
	result, err = svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	recordChanges, err := svc.compareTeamUpdate(record, result)
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleTeams, recordChanges, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.TeamResponse{
		Data: result,
	}, nil
}

func (svc *teamSvcImpl) compareTeamUpdate(oldRecord *ent.Team, newRecord *ent.Team) (string, error) {
	result := models.AuditTrailData{
		Module: "module.team",
		Create: make([]interface{}, 0),
		Delete: make([]interface{}, 0),
	}
	if oldRecord.Name != newRecord.Name {
		result.Update = append(result.Update, models.AuditTrailUpdate{
			Field: svc.formatTeamField("Name"),
			Value: models.ValueChange{
				OldValue: oldRecord.Name,
				NewValue: newRecord.Name,
			},
		})
	}
	oldTeamMembers := lo.Map(oldRecord.Edges.UserEdges, func(item *ent.User, index int) string {
		return item.Name
	})
	newTeamMembers := lo.Map(newRecord.Edges.UserEdges, func(item *ent.User, index int) string {
		return item.Name
	})
	left, right := lo.Difference(oldTeamMembers, newTeamMembers)
	if len(left) != 0 || len(right) != 0 {
		result.Update = append(result.Update, models.AuditTrailUpdate{
			Field: svc.formatTeamField("Members"),
			Value: models.ValueChange{
				OldValue: oldTeamMembers,
				NewValue: newTeamMembers,
			},
		})
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (svc *teamSvcImpl) DeleteTeam(ctx context.Context, teamId uuid.UUID, note string) error {
	team, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	record := models.AuditTrailData{
		Module: "module.team",
        Create: make([]interface{}, 0),
        Update: make([]interface{}, 0),
	}
	record.Delete = append(record.Delete, models.AuditTrailCreateDelete{
        Field: svc.formatTeamField("Name"),
        Value: team.Name,
    })
	membersName := lo.Map(team.Edges.UserEdges, func(item *ent.User, index int) string {
        return item.Name
    })
	record.Delete = append(record.Delete, models.AuditTrailCreateDelete{
        Field: svc.formatTeamField("Members"),
        Value: membersName,
    })
	recordChanges, _ := json.Marshal([]interface{}{record})
	memberIds := lo.Map(team.Edges.UserEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Team().DeleteTeam(ctx, team, memberIds)
		return err
	})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, teamId, audittrail.ModuleTeams, string(recordChanges), audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return err
}

func (svc *teamSvcImpl) GetTeam(ctx context.Context, teamId uuid.UUID) (*ent.TeamResponse, error) {
	team, err := svc.repoRegistry.Team().GetTeam(ctx, teamId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.TeamResponse{
		Data: team,
	}, nil
}

func (svc *teamSvcImpl) GetTeams(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.TeamFreeWord, filter *ent.TeamFilter, orderBy *ent.TeamOrder) (*ent.TeamResponseGetAll, error) {
	var result *ent.TeamResponseGetAll
	var edges []*ent.TeamEdge
	var page int
	var perPage int
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
		page = *pagination.Page
		perPage = *pagination.PerPage
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
			Page:    page,
			PerPage: perPage,
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
