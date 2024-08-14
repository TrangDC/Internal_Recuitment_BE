package service

import (
	"context"
	"github.com/google/uuid"
	"trec/ent"
	"trec/models"
	"trec/repository"

	"go.uber.org/zap"
)

type OutgoingEmailService interface {
	CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput, candidateId uuid.UUID) ([]*ent.OutgoingEmail, error)
	CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error)
}

type outgoingEmailSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewOutgoingEmailService(repoRegistry repository.Repository, logger *zap.Logger) OutgoingEmailService {
	return &outgoingEmailSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc outgoingEmailSvcImpl) CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput, candidateId uuid.UUID) ([]*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CreateBulkOutgoingEmail(ctx, input, candidateId)
}

func (svc outgoingEmailSvcImpl) CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CallbackOutgoingEmail(ctx, input)
}
