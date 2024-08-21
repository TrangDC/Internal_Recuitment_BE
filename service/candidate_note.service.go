package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidatejob"
	"trec/ent/candidatenote"
	"trec/internal/util"
	"trec/middleware"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateNoteService interface {
	CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNoteResponse, error)
	UpdateCandidateNote(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateNoteInput, note string) (*ent.CandidateNoteResponse, error)
	DeleteCandidateNote(ctx context.Context, id uuid.UUID, note string) error
	GetAllCandidateNotes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateNoteFilter, freeWord *ent.CandidateNoteFreeWord, orderBy *ent.CandidateNoteOrder) (*ent.CandidateNoteResponseGetAll, error)
	GetCandidateNote(ctx context.Context, id uuid.UUID) (*ent.CandidateNoteResponse, error)
}

type candidateNoteSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	dtoRegistry   dto.Dto
	logger        *zap.Logger
}

func NewCandidateNoteService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateNoteService {
	return &candidateNoteSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		dtoRegistry:   dtoRegistry,
		logger:        logger,
	}
}

func (svc *candidateNoteSvcImpl) CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNoteResponse, error) {
	var candidateNote *ent.CandidateNote
	err := svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		candidateNote, err = repoRegistry.CandidateNote().CreateCandidateNote(ctx, input)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateNote.ID, attachment.RelationTypeCandidateNotes, repoRegistry)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, candidateNote.ID)
	jsonStr, err := svc.dtoRegistry.CandidateNote().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidateNotes, jsonStr, audittrail.ActionTypeCreate, "")
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.CandidateNoteResponse{Data: result}, nil
}

func (svc *candidateNoteSvcImpl) UpdateCandidateNote(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateNoteInput, note string) (*ent.CandidateNoteResponse, error) {
	record, err := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err := repoRegistry.CandidateNote().UpdateCandidateNote(ctx, record, input)
		if err != nil {
			return err
		}
		return svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, id, input.Attachments, record.Edges.AttachmentEdges, attachment.RelationTypeCandidateNotes)
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, id)
	jsonStr, err := svc.dtoRegistry.CandidateNote().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, id, audittrail.ModuleCandidateNotes, jsonStr, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.CandidateNoteResponse{Data: result}, nil
}

func (svc *candidateNoteSvcImpl) DeleteCandidateNote(ctx context.Context, id uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	candidateJobQuery := svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(candidatejob.CandidateIDEQ(record.CandidateID))
	candidateJob, _ := svc.repoRegistry.CandidateJob().BuildGetOne(ctx, candidateJobQuery)
	if !svc.validPermissionDelete(payload, record, candidateJob.Edges.HiringJobEdge.Edges.HiringTeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err := repoRegistry.CandidateNote().DeleteCandidateNote(ctx, record)
		if err != nil {
			return err
		}
		return svc.attachmentSvc.RemoveAttachment(ctx, id, repoRegistry)
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonStr, err := svc.dtoRegistry.CandidateNote().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, id, audittrail.ModuleCandidateNotes, jsonStr, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return nil
}

func (svc candidateNoteSvcImpl) GetCandidateNote(ctx context.Context, id uuid.UUID) (*ent.CandidateNoteResponse, error) {
	result, err := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.CandidateNoteResponse{Data: result}, nil
}

func (svc *candidateNoteSvcImpl) GetAllCandidateNotes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateNoteFilter, freeWord *ent.CandidateNoteFreeWord, orderBy *ent.CandidateNoteOrder) (*ent.CandidateNoteResponseGetAll, error) {
	var (
		result        *ent.CandidateNoteResponseGetAll
		edges         []*ent.CandidateNoteEdge
		page, perPage int
	)
	query := svc.repoRegistry.CandidateNote().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateNote().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	orderFunc := ent.Desc(candidatenote.FieldCreatedAt)
	if orderBy != nil {
		orderFunc = ent.Desc(strings.ToLower(orderBy.Field.String()))
		if orderBy.Direction == ent.OrderDirectionAsc {
			orderFunc = ent.Asc(strings.ToLower(orderBy.Field.String()))
		}
	}
	query = query.Order(orderFunc)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	candidateNotes, err := svc.repoRegistry.CandidateNote().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateNotes, func(candidateNote *ent.CandidateNote, _ int) *ent.CandidateNoteEdge {
		return &ent.CandidateNoteEdge{
			Node:   candidateNote,
			Cursor: ent.Cursor{Value: candidateNote.ID.String()},
		}
	})
	result = &ent.CandidateNoteResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

// common
func (svc *candidateNoteSvcImpl) filter(query *ent.CandidateNoteQuery, filter *ent.CandidateNoteFilter) {
	if filter != nil {
		if filter.CandidateID != nil {
			query.Where(candidatenote.CandidateID(uuid.MustParse(*filter.CandidateID)))
		}
		if (filter.FromDate != nil && filter.ToDate != nil) && (!filter.FromDate.IsZero() && !filter.ToDate.IsZero()) {
			query.Where(candidatenote.CreatedAtGTE(*filter.FromDate), candidatenote.CreatedAtLTE(*filter.ToDate))
		}
	}
}

func (svc *candidateNoteSvcImpl) freeWord(query *ent.CandidateNoteQuery, freeWord *ent.CandidateNoteFreeWord) {
	if freeWord != nil && freeWord.Name != nil {
		query.Where(candidatenote.NameContainsFold(strings.TrimSpace(*freeWord.Name)))
	}
}

// permission
func (svc candidateNoteSvcImpl) validPermissionDelete(payload *middleware.Payload, record *ent.CandidateNote, hiringTeam *ent.HiringTeam) bool {
	if payload.ForTeam {
		// hiring team
		hiringTeamMemberIds := lo.Map(hiringTeam.Edges.HiringMemberEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		managerIds := lo.Map(hiringTeam.Edges.UserEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		if lo.Contains(hiringTeamMemberIds, payload.UserID) || lo.Contains(managerIds, payload.UserID) {
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

// Path: service/candidate_note.service.go
