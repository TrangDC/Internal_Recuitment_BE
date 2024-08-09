package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type AuditTrailService interface {
	// query
	GetAuditTrail(ctx context.Context, id uuid.UUID) (*ent.AuditTrailResponse, error)
	GetAuditTrails(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.AuditTrailFreeWord, filter *ent.AuditTrailFilter, orderBy *ent.AuditTrailOrder) (*ent.AuditTrailResponseGetAll, error)
}

type auditTrailSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewAuditTrailService(repoRegistry repository.Repository, logger *zap.Logger) AuditTrailService {
	return &auditTrailSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *auditTrailSvcImpl) GetAuditTrail(ctx context.Context, id uuid.UUID) (*ent.AuditTrailResponse, error) {
	result, err := svc.repoRegistry.AuditTrail().GetAuditTrail(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.AuditTrailResponse{
		Data: result,
	}, nil
}

func (svc *auditTrailSvcImpl) GetAuditTrails(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.AuditTrailFreeWord, filter *ent.AuditTrailFilter, orderBy *ent.AuditTrailOrder) (*ent.AuditTrailResponseGetAll, error) {
	var result *ent.AuditTrailResponseGetAll
	var edges []*ent.AuditTrailEdge
	var page int
	var perPage int
	query := svc.repoRegistry.AuditTrail().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.AuditTrail().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(audittrail.FieldCreatedAt)
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
	auditTrails, err := svc.repoRegistry.AuditTrail().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(auditTrails, func(auditTrail *ent.AuditTrail, index int) *ent.AuditTrailEdge {
		return &ent.AuditTrailEdge{
			Node: auditTrail,
			Cursor: ent.Cursor{
				Value: auditTrail.ID.String(),
			},
		}
	})
	result = &ent.AuditTrailResponseGetAll{
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
func (svc *auditTrailSvcImpl) freeWord(auditTrailQuery *ent.AuditTrailQuery, input *ent.AuditTrailFreeWord) {
	if input != nil {
		if input.RecordChange != nil {
			auditTrailQuery.Where(audittrail.RecordChangesContainsFold(strings.TrimSpace(*input.RecordChange)))
		}
	}
}

func (svc *auditTrailSvcImpl) filter(auditTrailQuery *ent.AuditTrailQuery, input *ent.AuditTrailFilter) {
	if input != nil {
		if input.RecordID != nil {
			auditTrailQuery.Where(audittrail.RecordIdEQ(uuid.MustParse(*input.RecordID)))
		}
		if input.Module != nil {
			auditTrailQuery.Where(audittrail.ModuleEQ(audittrail.Module(*input.Module)))
		}
		if input.ActionType != nil {
			auditTrailQuery.Where(audittrail.ActionTypeEQ(audittrail.ActionType(*input.ActionType)))
		}
		if input.FromDate != nil {
			auditTrailQuery.Where(audittrail.CreatedAtGTE(*input.FromDate))
		}
		if input.ToDate != nil {
			auditTrailQuery.Where(audittrail.CreatedAtLTE(*input.ToDate))
		}
	}
}
