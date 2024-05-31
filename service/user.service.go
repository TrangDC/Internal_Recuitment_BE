package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/team"
	"trec/ent/user"
	"trec/internal/util"
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
}

type userSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewUserService(repoRegistry repository.Repository, logger *zap.Logger) UserService {
	return &userSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *userSvcImpl) CreateUser(ctx context.Context, input *ent.NewUserInput, note string) (*ent.UserResponse, error) {
	var record *ent.User
	err := svc.repoRegistry.User().ValidWorkEmail(ctx, uuid.Nil, input.WorkEmail)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.User().CreateUser(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.UserResponse{
		Data: record,
	}, nil
}

func (svc *userSvcImpl) DeleteUser(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.User().DeleteUser(ctx, record)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *userSvcImpl) UpdateUser(ctx context.Context, input *ent.UpdateUserInput, id uuid.UUID, note string) (*ent.UserResponse, error) {
	var result *ent.User
	record, err := svc.repoRegistry.User().GetUser(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.User().ValidWorkEmail(ctx, id, input.WorkEmail)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.User().UpdateUser(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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
		result, err = repoRegistry.User().UpdateUserStatus(ctx, record, user.Status(input.Status))
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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

// common function
func (svc *userSvcImpl) freeWord(userQuery *ent.UserQuery, input *ent.UserFreeWord) {
	if input != nil {
		if input.Name != nil {
			userQuery.Where(user.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
		if input.WorkEmail != nil {
			userQuery.Where(user.WorkEmailContainsFold(strings.TrimSpace(*input.WorkEmail)))
		}
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
		if input.NotInTeam != nil {
			predicate := user.HasTeamEdgesWith(team.DeletedAtIsNil())
			if *input.NotInTeam {
				predicate = user.Not(user.HasTeamEdgesWith(team.DeletedAtIsNil()))
			}
			userQuery.Where(predicate)
		}
	}
}
