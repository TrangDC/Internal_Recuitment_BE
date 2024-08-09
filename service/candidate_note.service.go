package service

import (
	"context"
	"net/http"
	"trec/ent"
	"trec/ent/attachment"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CandidateNoteService interface {
	CreateCandidateNote(ctx context.Context, input ent.NewCandidateNoteInput) (*ent.CandidateNoteResponse, error)
	UpdateCandidateNote(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateNoteInput, note string) (*ent.CandidateNoteResponse, error)
	DeleteCandidateNote(ctx context.Context, id uuid.UUID, note string) error
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
