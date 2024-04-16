package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/user"
	"trec/internal/util"
	"trec/repository"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type UserService interface {
	// query
	Selections(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error)
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

func (svc *userSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, orderBy *ent.UserOrder) (*ent.UserResponseGetAll, error) {
	var result *ent.UserResponseGetAll
	var edges []*ent.UserEdge
	var page int
	var perPage int
	query := svc.repoRegistry.User().BuildQuery()
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
