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
	// query
	Selections(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseSelectionGetAll, error)
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

func (svc *userSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, filter *ent.UserFilter, freeWord *ent.UserFreeWord, orderBy *ent.UserOrder) (*ent.UserResponseSelectionGetAll, error) {
	var result *ent.UserResponseSelectionGetAll
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
	result = &ent.UserResponseSelectionGetAll{
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
