package service

import (
	"context"
	"trec/ent"
	"trec/ent/attachment"
	"trec/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AttachmentService interface {
	CreateAttachment(ctx context.Context, input []*ent.NewAttachmentInput, relationId uuid.UUID, relationType attachment.RelationType, repoRegistry repository.Repository) ([]*ent.Attachment, error)
	GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error)
	GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error)
}

type attachmentSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewAttachmentService(repoRegistry repository.Repository, logger *zap.Logger) AttachmentService {
	return &attachmentSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *attachmentSvcImpl) CreateAttachment(ctx context.Context, input []*ent.NewAttachmentInput, relationId uuid.UUID, relationType attachment.RelationType, repoRegistry repository.Repository) ([]*ent.Attachment, error) {
	attachments, err := repoRegistry.Attachment().CreateAttachment(ctx, input, relationId, relationType)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}
	return attachments, nil
}

func (svc *attachmentSvcImpl) GetAttachment(ctx context.Context, attachmentId uuid.UUID) (*ent.Attachment, error) {
	attachment, err := svc.repoRegistry.Attachment().GetAttachment(ctx, attachmentId)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}
	return attachment, nil
}

func (svc *attachmentSvcImpl) GetAttachments(ctx context.Context, relationId uuid.UUID, relationType attachment.RelationType) ([]*ent.Attachment, error) {
	attachments, err := svc.repoRegistry.Attachment().GetAttachments(ctx, relationId, relationType)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}
	return attachments, nil
}
