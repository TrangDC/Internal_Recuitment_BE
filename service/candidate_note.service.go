package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidatenote"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateNoteService interface {
	CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNoteResponse, error)
	UpdateCandidateNote(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateNoteInput, note string) (*ent.CandidateNoteResponse, error)
	DeleteCandidateNote(ctx context.Context, id uuid.UUID, note string) error
	GetCandidateNotes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateNoteFilter, freeWord *ent.CandidateNoteFreeWord, orderBy *ent.CandidateNoteOrder) (*ent.CandidateNoteResponseGetAll, error)
}

type candidateNoteSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	logger        *zap.Logger
}

func NewCandidateNoteService(repoRegistry repository.Repository, logger *zap.Logger) CandidateNoteService {
	return &candidateNoteSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
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
	return &ent.CandidateNoteResponse{Data: result}, nil
}

func (svc *candidateNoteSvcImpl) DeleteCandidateNote(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.CandidateNote().GetCandidateNote(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
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
	return nil
}

func (svc *candidateNoteSvcImpl) GetCandidateNotes(ctx context.Context, pagination *ent.PaginationInput, filter *ent.CandidateNoteFilter, freeWord *ent.CandidateNoteFreeWord, orderBy *ent.CandidateNoteOrder) (*ent.CandidateNoteResponseGetAll, error) {
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

func (svc *candidateNoteSvcImpl) filter(query *ent.CandidateNoteQuery, filter *ent.CandidateNoteFilter) {
	if filter != nil {
		if filter.CandidateID != nil {
			query.Where(candidatenote.CandidateID(uuid.MustParse(*filter.CandidateID)))
		}
		if (filter.FromDate != nil && filter.ToDate != nil) && (!filter.FromDate.IsZero() && !filter.ToDate.IsZero()) {
			query.Where(
				candidatenote.CreatedAtGTE(*filter.FromDate),
				candidatenote.CreatedAtLTE(*filter.ToDate),
			)
		}
	}
}

func (svc *candidateNoteSvcImpl) freeWord(query *ent.CandidateNoteQuery, freeWord *ent.CandidateNoteFreeWord) {
	if freeWord != nil && freeWord.Name != nil {
		query.Where(candidatenote.NameContainsFold(strings.TrimSpace(*freeWord.Name)))
	}
}
