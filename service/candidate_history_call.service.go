package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidatejob"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/middleware"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateHistoryCallService interface {
	// mutation
	CreateCandidateHistoryCall(ctx context.Context, input ent.NewCandidateHistoryCallInput, note string) (*ent.CandidateHistoryCallResponse, error)
	UpdateCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID, input ent.UpdateCandidateHistoryCallInput, note string) (*ent.CandidateHistoryCallResponse, error)
	DeleteCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID, note string) error
	// query
	GetCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID) (*ent.CandidateHistoryCallResponse, error)
	GetCandidateHistoryCalls(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateHistoryCallFreeWord,
		filter *ent.CandidateHistoryCallFilter, orderBy *ent.CandidateHistoryCallOrder) (*ent.CandidateHistoryCallResponseGetAll, error)
}

type candidateHistoryCallSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	dtoRegistry   dto.Dto
	logger        *zap.Logger
}

func NewCandidateHistoryCallService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateHistoryCallService {
	return &candidateHistoryCallSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		dtoRegistry:   dtoRegistry,
		logger:        logger,
	}
}

func (svc *candidateHistoryCallSvcImpl) CreateCandidateHistoryCall(ctx context.Context, input ent.NewCandidateHistoryCallInput, note string) (*ent.CandidateHistoryCallResponse, error) {
	var result *ent.CandidateHistoryCall
	var err error
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.CandidateHistoryCall().CreateCandidateHistoryCall(ctx, input)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, result.ID, attachment.RelationTypeCandidateHistoryCalls, repoRegistry)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, result.ID)
	jsonString, err := svc.dtoRegistry.CandidateHistoryCall().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidateHistoryCalls, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateHistoryCallResponse{
		Data: result,
	}, nil
}

func (svc *candidateHistoryCallSvcImpl) UpdateCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID, input ent.UpdateCandidateHistoryCallInput, note string) (*ent.CandidateHistoryCallResponse, error) {
	var result *ent.CandidateHistoryCall
	record, err := svc.repoRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, candidateHistoryCallId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateHistoryCall().UpdateCandidateHistoryCall(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, record.ID, input.Attachments, record.Edges.AttachmentEdges, attachment.RelationTypeCandidateHistoryCalls)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, candidateHistoryCallId)
	jsonString, err := svc.dtoRegistry.CandidateHistoryCall().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidateHistoryCalls, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateHistoryCallResponse{
		Data: result,
	}, nil
}

func (svc *candidateHistoryCallSvcImpl) DeleteCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, candidateHistoryCallId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	candidateJobQuery := svc.repoRegistry.CandidateJob().BuildBaseQuery().
		Where(candidatejob.CandidateIDEQ(record.CandidateID)).
		WithHiringJobEdge(func(query *ent.HiringJobQuery) {
			query.WithHiringTeamEdge(func(query *ent.HiringTeamQuery) {
				query.WithHiringMemberEdges().WithUserEdges()
			})
		})
	candidateJob, _ := svc.repoRegistry.CandidateJob().BuildGetOne(ctx, candidateJobQuery)
	if !svc.validPermissionDelete(payload, record, candidateJob.Edges.HiringJobEdge.Edges.HiringTeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateHistoryCall().DeleteCandidateHistoryCall(ctx, record)
		if err != nil {
			return err
		}
		err = svc.attachmentSvc.RemoveAttachment(ctx, record.ID, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateHistoryCall().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleCandidateHistoryCalls, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateHistoryCallSvcImpl) GetCandidateHistoryCall(ctx context.Context, candidateHistoryCallId uuid.UUID) (*ent.CandidateHistoryCallResponse, error) {
	candidateHistoryCallRecord, err := svc.repoRegistry.CandidateHistoryCall().GetCandidateHistoryCall(ctx, candidateHistoryCallId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.CandidateHistoryCallResponse{
		Data: candidateHistoryCallRecord,
	}, nil
}

func (svc candidateHistoryCallSvcImpl) GetCandidateHistoryCalls(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateHistoryCallFreeWord,
	filter *ent.CandidateHistoryCallFilter, orderBy *ent.CandidateHistoryCallOrder) (*ent.CandidateHistoryCallResponseGetAll, error) {
	var edges []*ent.CandidateHistoryCallEdge
	candidateHistoryCalls, count, page, perPage, err := svc.getAllCandidateHistoryCall(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateHistoryCalls, func(entity *ent.CandidateHistoryCall, index int) *ent.CandidateHistoryCallEdge {
		return &ent.CandidateHistoryCallEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.CandidateHistoryCallResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc candidateHistoryCallSvcImpl) getAllCandidateHistoryCall(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateHistoryCallFreeWord,
	filter *ent.CandidateHistoryCallFilter, orderBy *ent.CandidateHistoryCallOrder) ([]*ent.CandidateHistoryCall, int, int, int, error) {
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateHistoryCall().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateHistoryCall().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidatehistorycall.FieldCreatedAt)
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
	candidateHistoryCalls, err := svc.repoRegistry.CandidateHistoryCall().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return candidateHistoryCalls, count, page, perPage, nil
}

// common function
func (svc *candidateHistoryCallSvcImpl) freeWord(query *ent.CandidateHistoryCallQuery, input *ent.CandidateHistoryCallFreeWord) {
	predicate := []predicate.CandidateHistoryCall{}
	if input != nil {
		if input.Description != nil {
			predicate = append(predicate, candidatehistorycall.DescriptionContainsFold(strings.TrimSpace(*input.Description)))
		}
		if input.Name != nil {
			predicate = append(predicate, candidatehistorycall.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
	if len(predicate) > 0 {
		query.Where(candidatehistorycall.Or(predicate...))
	}
}

func (svc *candidateHistoryCallSvcImpl) filter(query *ent.CandidateHistoryCallQuery, input *ent.CandidateHistoryCallFilter) {
	if input != nil {
		if (input.FromDate != nil && input.ToDate != nil) && (!input.FromDate.IsZero() && !input.ToDate.IsZero()) {
			query.Where(candidatehistorycall.DateGTE(*input.FromDate), candidatehistorycall.DateLTE(*input.ToDate))
		}
		if (input.StartTime != nil && input.EndTime != nil) && (!input.StartTime.IsZero() && !input.EndTime.IsZero()) {
			query.Where(candidatehistorycall.StartTimeGTE(*input.StartTime), candidatehistorycall.EndTimeLTE(*input.EndTime))
		}
		if input.CandidateID != nil {
			query.Where(candidatehistorycall.CandidateIDEQ(uuid.MustParse(*input.CandidateID)))
		}
		if input.Type != nil {
			query.Where(candidatehistorycall.TypeEQ(candidatehistorycall.Type(*input.Type)))
		}
	}
}

// permission
func (svc candidateHistoryCallSvcImpl) validPermissionDelete(payload *middleware.Payload, record *ent.CandidateHistoryCall, hiringTeam *ent.HiringTeam) bool {
	if payload.ForTeam {
		// hiring team
		memberIds := lo.Map(hiringTeam.Edges.HiringMemberEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		managerIds := lo.Map(hiringTeam.Edges.UserEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		if lo.Contains(memberIds, payload.UserID) || lo.Contains(managerIds, payload.UserID) {
			return true
		}
		// TODO: rec team
	}
	if payload.ForAll {
		return true
	}
	if payload.ForOwner && record.CreatedByID == payload.UserID {
		return true
	}
	return false
}

// Path: service/skill_types.service.go
