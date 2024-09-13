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

type EmailEventService interface {
	Selections(ctx context.Context) (*ent.EmailEventSelectionResponseGetAll, error)
}

type emailEventSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewEmailEventService(repoRegistry repository.Repository, logger *zap.Logger) EmailEventService {
	return &emailEventSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *emailEventSvcImpl) Selections(ctx context.Context) (*ent.EmailEventSelectionResponseGetAll, error) {
	records, err := svc.repoRegistry.EmailEvent().BuildList(ctx, svc.repoRegistry.EmailEvent().BuildBaseQuery())
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges := lo.Map(records, func(record *ent.EmailEvent, _ int) *ent.EmailEventSelectionEdge {
		return &ent.EmailEventSelectionEdge{
			Node: record,
		}
	})
	return &ent.EmailEventSelectionResponseGetAll{Edges: edges}, nil
}
