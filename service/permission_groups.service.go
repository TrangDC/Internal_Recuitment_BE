package service

import (
	"context"
	"net/http"
	"trec/ent"
	"trec/internal/util"
	"trec/repository"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type PermissionGroupService interface {
	GetAllPermissionGroups(ctx context.Context) (*ent.PermissionGroupResponseGetAll, error)
}
type permissionGroupSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewPermissionGroupService(repoRegistry repository.Repository, logger *zap.Logger) PermissionGroupService {
	return &permissionGroupSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *permissionGroupSvcImpl) GetAllPermissionGroups(ctx context.Context) (*ent.PermissionGroupResponseGetAll, error) {
	result, err := svc.repoRegistry.PermissionGroup().GetAllPermissionGroups(ctx)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	edges := lo.Map(result, func(item *ent.PermissionGroup, index int) *ent.PermissionGroupEdge {
		return &ent.PermissionGroupEdge{
			Node: item,
		}
	})
	return &ent.PermissionGroupResponseGetAll{
		Edges: edges,
	}, nil
}
