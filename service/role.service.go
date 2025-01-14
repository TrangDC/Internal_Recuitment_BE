package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/entitypermission"
	"trec/ent/predicate"
	"trec/ent/role"
	"trec/ent/userrole"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type RoleService interface {
	// mutation
	CreateRole(ctx context.Context, input ent.NewRoleInput, note string) (*ent.RoleResponse, error)
	UpdateRole(ctx context.Context, roleId uuid.UUID, input ent.UpdateRoleInput, note string) (*ent.RoleResponse, error)
	DeleteRole(ctx context.Context, roleId uuid.UUID, note string) error
	// query
	GetRole(ctx context.Context, roleId uuid.UUID) (*ent.RoleResponse, error)
	GetRoles(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RoleFreeWord,
		filter *ent.RoleFilter, orderBy *ent.RoleOrder) (*ent.RoleResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RoleFreeWord,
		filter *ent.RoleFilter, orderBy *ent.RoleOrder) (*ent.RoleSelectionResponseGetAll, error)
}

type roleSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewRoleService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) RoleService {
	return &roleSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

func (svc *roleSvcImpl) CreateRole(ctx context.Context, input ent.NewRoleInput, note string) (*ent.RoleResponse, error) {
	var result *ent.Role
	var record *ent.Role
	errString, err := svc.repoRegistry.Role().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.EntityPermission().ValidActionPermission(ctx, input.EntityPermissions)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.Role().CreateRole(ctx, input)
		if err != nil {
			return err
		}
		err = repoRegistry.EntityPermission().CreateAndUpdateEntityPermission(ctx, record.ID, input.EntityPermissions, nil, entitypermission.EntityTypeRole)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.Role().GetRole(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.Role().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleRoles, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.RoleResponse{
		Data: result,
	}, nil
}

func (svc *roleSvcImpl) UpdateRole(ctx context.Context, roleId uuid.UUID, input ent.UpdateRoleInput, note string) (*ent.RoleResponse, error) {
	var result *ent.Role
	record, err := svc.repoRegistry.Role().GetRole(ctx, roleId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.Role().ValidName(ctx, roleId, *input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.EntityPermission().ValidActionPermission(ctx, input.EntityPermissions)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	lo.ForEach(record.Edges.UserEdges, func(_ *ent.User, i int) {
		record.Edges.UserEdges[i].Edges.UserPermissionEdges, _ = svc.repoRegistry.EntityPermission().BuildList(
			ctx,
			svc.repoRegistry.EntityPermission().BuildQuery().Where(entitypermission.EntityID(record.Edges.UserEdges[i].ID)),
		)
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Role().UpdateRole(ctx, record, input)
		if err != nil {
			return err
		}
		err = repoRegistry.EntityPermission().CreateAndUpdateEntityPermission(ctx, record.ID, input.EntityPermissions, record.Edges.RolePermissionEdges, entitypermission.EntityTypeRole)
		if err != nil {
			return err
		}
		return svc.updateUserPermissions(ctx, repoRegistry, record.Edges.UserEdges, note)
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.Role().GetRole(ctx, roleId)
	jsonString, err := svc.dtoRegistry.Role().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, roleId, audittrail.ModuleRoles, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.RoleResponse{
		Data: result,
	}, nil
}

func (svc *roleSvcImpl) DeleteRole(ctx context.Context, roleId uuid.UUID, note string) error {
	roleRecord, err := svc.repoRegistry.Role().GetRole(ctx, roleId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	lo.ForEach(roleRecord.Edges.UserEdges, func(_ *ent.User, i int) {
		roleRecord.Edges.UserEdges[i].Edges.UserPermissionEdges, _ = svc.repoRegistry.EntityPermission().BuildList(
			ctx,
			svc.repoRegistry.EntityPermission().BuildQuery().Where(entitypermission.EntityID(roleRecord.Edges.UserEdges[i].ID)),
		)
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.Role().DeleteRole(ctx, roleRecord)
		if err != nil {
			return err
		}
		err = repoRegistry.EntityPermission().DeleteBulkEntityPermissionByEntityID(ctx, roleId)
		if err != nil {
			return err
		}
		return svc.updateUserPermissions(ctx, repoRegistry, roleRecord.Edges.UserEdges, note)
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.Role().AuditTrailDelete(roleRecord)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, roleId, audittrail.ModuleRoles, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *roleSvcImpl) GetRole(ctx context.Context, roleId uuid.UUID) (*ent.RoleResponse, error) {
	roleRecord, err := svc.repoRegistry.Role().GetRole(ctx, roleId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.RoleResponse{
		Data: roleRecord,
	}, nil
}

func (svc roleSvcImpl) GetRoles(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RoleFreeWord,
	filter *ent.RoleFilter, orderBy *ent.RoleOrder) (*ent.RoleResponseGetAll, error) {
	var edges []*ent.RoleEdge
	roles, count, page, perPage, err := svc.getAllRole(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(roles, func(entity *ent.Role, index int) *ent.RoleEdge {
		return &ent.RoleEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.RoleResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc roleSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RoleFreeWord,
	filter *ent.RoleFilter, orderBy *ent.RoleOrder) (*ent.RoleSelectionResponseGetAll, error) {
	var edges []*ent.RoleSelectionEdge
	roles, count, page, perPage, err := svc.getAllRole(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(roles, func(entity *ent.Role, index int) *ent.RoleSelectionEdge {
		return &ent.RoleSelectionEdge{
			Node: &ent.RoleSelection{
				ID:                entity.ID.String(),
				Name:              entity.Name,
				EntityPermissions: entity.Edges.RolePermissionEdges,
			},
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.RoleSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc roleSvcImpl) updateUserPermissions(ctx context.Context, repoRegistry repository.Repository, roleUsers []*ent.User, note string) error {
	for _, userRec := range roleUsers {
		// instead of just get the ids, get the whole records to avoid user role audit trail
		userRec.Edges.RoleEdges, _ = repoRegistry.Role().BuildList(ctx, repoRegistry.Role().BuildBaseQuery().Where(
			role.HasUserRolesWith(userrole.UserID(userRec.ID)),
		))
		roleIDs := lo.Map(userRec.Edges.RoleEdges, func(item *ent.Role, _ int) uuid.UUID {
			return item.ID
		})
		rolesPermissions, _ := repoRegistry.EntityPermission().BuildList(ctx,
			repoRegistry.EntityPermission().BuildQuery().Where(entitypermission.EntityIDIn(roleIDs...)))
		userPermissionInput := svc.dtoRegistry.User().NewUserEntityPermissionInput(rolesPermissions)
		err := repoRegistry.EntityPermission().CreateAndUpdateEntityPermission(ctx, userRec.ID, userPermissionInput, userRec.Edges.UserPermissionEdges, entitypermission.EntityTypeUser)
		if err != nil {
			return err
		}
		result, _ := repoRegistry.User().GetUser(ctx, userRec.ID)
		jsonString, err := svc.dtoRegistry.User().AuditTrailUpdate(userRec, result)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
		_, err = repoRegistry.AuditTrail().AuditTrailMutation(ctx, userRec.ID, audittrail.ModuleUsers, jsonString, audittrail.ActionTypeUpdate, note)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
	}
	return nil
}

func (svc roleSvcImpl) getAllRole(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.RoleFreeWord,
	filter *ent.RoleFilter, orderBy *ent.RoleOrder) ([]*ent.Role, int, int, int, error) {
	var page int
	var perPage int
	query := svc.repoRegistry.Role().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.Role().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(role.FieldCreatedAt)
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
	roles, err := svc.repoRegistry.Role().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return roles, count, page, perPage, nil
}

// common function
func (svc *roleSvcImpl) freeWord(query *ent.RoleQuery, input *ent.RoleFreeWord) {
	predicate := []predicate.Role{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, role.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
	if len(predicate) > 0 {
		query.Where(role.Or(predicate...))
	}
}

func (svc *roleSvcImpl) filter(query *ent.RoleQuery, input *ent.RoleFilter) {
	if input != nil {
		if input.Name != nil {
			query.Where(role.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
	}
}

// Path: service/role.service.go
