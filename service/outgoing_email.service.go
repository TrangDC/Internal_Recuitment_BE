package service

import (
	"context"
	"trec/ent"
	"trec/models"
	"trec/repository"

	"go.uber.org/zap"
)

type OutgoingEmailService interface {
	CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput) ([]*ent.OutgoingEmail, error)
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

func (svc outgoingEmailSvcImpl) CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput) ([]*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CreateBulkOutgoingEmail(ctx, input)
}

func (svc outgoingEmailSvcImpl) CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CallbackOutgoingEmail(ctx, input)
}
