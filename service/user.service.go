package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/entitypermission"
	"trec/ent/hiringteam"
	"trec/ent/permission"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type UserService interface {
	// mutation
	CreateUser(ctx context.Context, input *ent.NewUserInput, note string) (*ent.UserResponse, error)
	UpdateUser(ctx context.Context, input *ent.UpdateUserInput, id uuid.UUID, note string) (*ent.UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID, note string) error
	UpdateUserStatus(ctx context.Context, input ent.UpdateUserStatusInput, id uuid.UUID, note string) (*ent.UserResponse, error)
	// query
	Selections(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserSelectionResponseGetAll, error)
	GetUser(ctx context.Context, id uuid.UUID) (*ent.UserResponse, error)
	GetUsers(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error)
	GetMe(ctx context.Context) (*ent.UserResponse, error)
	UpdateHiringTeam(ctx context.Context, teamName string, teamId uuid.UUID, userId []uuid.UUID, note string) error
	RemoveHiringTeam(ctx context.Context, teamId uuid.UUID, userId []uuid.UUID, note string) error
}

type userSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewUserService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) UserService {
	return &userSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *userSvcImpl) CreateUser(ctx context.Context, input *ent.NewUserInput, note string) (*ent.UserResponse, error) {
	var record *ent.User
	errString, err := svc.repoRegistry.User().ValidWorkEmail(ctx, uuid.Nil, input.WorkEmail)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	roleIds := lo.Map(input.RoleID, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	rolesPermissions, err := svc.repoRegistry.EntityPermission().BuildList(ctx,
		svc.repoRegistry.EntityPermission().BuildQuery().Where(entitypermission.EntityIDIn(roleIds...)))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	userPermissionInput := svc.dtoRegistry.User().NewUserEntityPermissionInput(rolesPermissions)
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.User().CreateUser(ctx, input, roleIds)
		if err != nil {
			return err
		}
		err := repoRegistry.EntityPermission().CreateAndUpdateEntityPermission(ctx, record.ID, userPermissionInput, nil, entitypermission.EntityTypeUser)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.User().GetUser(ctx, record.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagNotFound)
	}
	jsonString, err := svc.dtoRegistry.User().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleUsers, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.UserResponse{
		Data: result,
	}, nil
}

func (svc *userSvcImpl) DeleteUser(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagNotFound)
	}
	roleIds := lo.Map(record.Edges.RoleEdges, func(member *ent.Role, index int) uuid.UUID {
		return member.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.User().DeleteUser(ctx, record, roleIds)
		if err != nil {
			return err
		}
		err = repoRegistry.EntityPermission().DeleteAllEntityPermission(ctx, record.ID)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.User().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleUsers, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *userSvcImpl) UpdateUser(ctx context.Context, input *ent.UpdateUserInput, id uuid.UUID, note string) (*ent.UserResponse, error) {
	var result *ent.User
	var roleIds []uuid.UUID
	record, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.User().ValidWorkEmail(ctx, id, input.WorkEmail)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if len(input.RoleID) != 0 {
		roleIds = lo.Map(input.RoleID, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
	}
	if len(record.Edges.HiringTeamEdges) != 0 {
		currentTeamId := record.Edges.HiringTeamEdges[0].ID.String()
		input.HiringTeamID = &currentTeamId
	}
	newRoleIds, removeRoleIds := svc.updateRoles(record, roleIds)
	rolesPermissions, err := svc.repoRegistry.EntityPermission().BuildList(ctx,
		svc.repoRegistry.EntityPermission().BuildQuery().Where(entitypermission.EntityIDIn(roleIds...)))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	userPermissionInput := svc.dtoRegistry.User().NewUserEntityPermissionInput(rolesPermissions)
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.User().UpdateUser(ctx, record, input, newRoleIds, removeRoleIds)
		if err != nil {
			return err
		}
		err = repoRegistry.EntityPermission().CreateAndUpdateEntityPermission(ctx, record.ID, userPermissionInput, record.Edges.UserPermissionEdges, entitypermission.EntityTypeUser)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.User().GetUser(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.User().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleUsers, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.UserResponse{
		Data: result,
	}, nil
}

func (svc *userSvcImpl) UpdateUserStatus(ctx context.Context, input ent.UpdateUserStatusInput, id uuid.UUID, note string) (*ent.UserResponse, error) {
	var result *ent.User
	record, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.User().UpdateUserStatus(ctx, record, user.Status(input.Status))
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.User().GetUser(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.User().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleUsers, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.UserResponse{
		Data: result,
	}, nil
}

// query
func (svc *userSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserSelectionResponseGetAll, error) {
	var result *ent.UserSelectionResponseGetAll
	var edges []*ent.UserSelectionEdge
	var page int
	var perPage int
	query := svc.repoRegistry.User().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.User().BuildCount(ctx, query)
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
		query = query.Order(ent.Desc(user.FieldCreatedAt))
	}
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(*pagination.PerPage).Offset((*pagination.Page - 1) * *pagination.PerPage)
	}
	users, err := svc.repoRegistry.User().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(users, func(user *ent.User, index int) *ent.UserSelectionEdge {
		return &ent.UserSelectionEdge{
			Node: &ent.UserSelection{
				ID:        user.ID.String(),
				Name:      user.Name,
				WorkEmail: user.WorkEmail,
			},
			Cursor: ent.Cursor{
				Value: user.ID.String(),
			},
		}
	})
	result = &ent.UserSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *userSvcImpl) GetUser(ctx context.Context, id uuid.UUID) (*ent.UserResponse, error) {
	result, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.UserResponse{
		Data: result,
	}, nil
}

func (svc *userSvcImpl) GetMe(ctx context.Context) (*ent.UserResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	result, err := svc.repoRegistry.User().GetUser(ctx, payload.UserID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.UserResponse{
		Data: result,
	}, nil
}

func (svc *userSvcImpl) GetUsers(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error) {
	var result *ent.UserResponseGetAll
	var edges []*ent.UserEdge
	var page int
	var perPage int
	query := svc.repoRegistry.User().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.User().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(user.FieldCreatedAt)
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
	users, err := svc.repoRegistry.User().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(users, func(user *ent.User, index int) *ent.UserEdge {
		return &ent.UserEdge{
			Node: user,
			Cursor: ent.Cursor{
				Value: user.ID.String(),
			},
		}
	})
	result = &ent.UserResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *userSvcImpl) UpdateHiringTeam(ctx context.Context, teamName string, teamId uuid.UUID, userId []uuid.UUID, note string) error {
	users, err := svc.repoRegistry.User().BuildList(ctx, svc.repoRegistry.User().BuildQuery().Where(user.IDIn(userId...)))
	if err != nil {
		return err
	}
	err = svc.repoRegistry.User().UpdateUserHiringTeam(ctx, userId, teamId)
	if err != nil {
		return err
	}
	recordAuditTrails := lo.Map(users, func(user *ent.User, index int) models.UserTeamAuditTrail {
		return models.UserTeamAuditTrail{
			RecordId:   user.ID,
			JsonString: svc.dtoRegistry.User().AuditTrailUpdateHiringTeam(user, teamName),
		}
	})
	err = svc.repoRegistry.AuditTrail().CreateBulkUserTeamAt(ctx, recordAuditTrails, note)
	return err
}

func (svc *userSvcImpl) RemoveHiringTeam(ctx context.Context, teamId uuid.UUID, userId []uuid.UUID, note string) error {
	users, err := svc.repoRegistry.User().BuildList(ctx, svc.repoRegistry.User().BuildQuery().Where(user.IDIn(userId...)))
	if err != nil {
		return err
	}
	err = svc.repoRegistry.User().DeleteUserHiringTeam(ctx, userId)
	if err != nil {
		return err
	}
	recordAuditTrails := lo.Map(users, func(user *ent.User, index int) models.UserTeamAuditTrail {
		return models.UserTeamAuditTrail{
			RecordId:   user.ID,
			JsonString: svc.dtoRegistry.User().AuditTrailUpdateHiringTeam(user, ""),
		}
	})
	err = svc.repoRegistry.AuditTrail().CreateBulkUserTeamAt(ctx, recordAuditTrails, note)
	return err
}

// common function
func (svc *userSvcImpl) freeWord(userQuery *ent.UserQuery, input *ent.UserFreeWord) {
	predicate := []predicate.User{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, user.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
		if input.WorkEmail != nil {
			predicate = append(predicate, user.WorkEmailContainsFold(strings.TrimSpace(*input.WorkEmail)))
		}
	}
	if len(predicate) > 0 {
		userQuery.Where(user.Or(predicate...))
	}
}

func (svc *userSvcImpl) filter(userQuery *ent.UserQuery, input *ent.UserFilter) {
	if input != nil {
		if input.Name != nil {
			userQuery.Where(user.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.Ids != nil {
			ids := lo.Map(input.Ids, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			userQuery.Where(user.IDIn(ids...))
		}
		if input.IgnoreIds != nil {
			ids := lo.Map(input.IgnoreIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			userQuery.Where(user.IDNotIn(ids...))
		}
		if input.IsAbleToInterviewer != nil {
			if *input.IsAbleToInterviewer {
				userQuery.Where(user.HasUserPermissionEdgesWith(
					entitypermission.HasPermissionEdgesWith(
						permission.OperationNameEQ(models.BeInterviewer),
					),
				))
			} else {
				userQuery.Where(user.Not(user.HasUserPermissionEdgesWith(
					entitypermission.HasPermissionEdgesWith(
						permission.OperationNameEQ(models.BeInterviewer),
					),
				)))
			}
		}
		if input.RoleID != nil {
			roles := lo.Map(input.RoleID, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			userQuery.Where(user.HasRoleEdgesWith(role.IDIn(roles...)))
		}
		if input.HiringTeamID != nil {
			teamIds := lo.Map(input.HiringTeamID, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			userQuery.Where(user.Or(user.HiringTeamIDIn(teamIds...), user.HasHiringTeamEdgesWith(hiringteam.IDIn(teamIds...))))
		}
		if input.IsAbleToLeaderRecTeam != nil {
			if *input.IsAbleToLeaderRecTeam {
				userQuery.Where(user.Not(user.HasLeadRecTeams()))
			} else {
				userQuery.Where(user.HasLeadRecTeams())
			}
		}
		if input.IsAbleToManagerHiringTeam != nil {
			if *input.IsAbleToManagerHiringTeam {
				userQuery.Where(user.Not(user.HasHiringTeamEdges()))
			} else {
				userQuery.Where(user.HasHiringTeamEdges())
			}
		}
	}
}

func (svc *userSvcImpl) updateRoles(record *ent.User, roleIds []uuid.UUID) ([]uuid.UUID, []uuid.UUID) {
	var newMemberIds []uuid.UUID
	var removeMemberIds []uuid.UUID
	currentMemberIds := lo.Map(record.Edges.RoleEdges, func(user *ent.Role, index int) uuid.UUID {
		return user.ID
	})
	newMemberIds = lo.Filter(roleIds, func(roleId uuid.UUID, index int) bool {
		return !lo.Contains(currentMemberIds, roleId)
	})
	removeMemberIds = lo.Filter(currentMemberIds, func(roleId uuid.UUID, index int) bool {
		return !lo.Contains(roleIds, roleId)
	})
	return newMemberIds, removeMemberIds
}

// Path: service/user.service.go
