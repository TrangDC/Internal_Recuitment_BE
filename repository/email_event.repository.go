package repository

import (
	"context"
	"trec/ent"
	"trec/ent/emailevent"

	"github.com/google/uuid"
)

type EmailEventRepository interface {
	BuildBaseQuery() *ent.EmailEventQuery
	BuildList(ctx context.Context, query *ent.EmailEventQuery) ([]*ent.EmailEvent, error)
	GetEmailEvent(ctx context.Context, id uuid.UUID) (*ent.EmailEvent, error)
}

type emailEventRepoImpl struct {
	entClient *ent.Client
}

func NewEmailEventRepository(entClient *ent.Client) EmailEventRepository {
	return &emailEventRepoImpl{
		entClient: entClient,
	}
}

func (rps *emailEventRepoImpl) BuildBaseQuery() *ent.EmailEventQuery {
	return rps.entClient.EmailEvent.Query()
}

func (rps *emailEventRepoImpl) BuildList(ctx context.Context, query *ent.EmailEventQuery) ([]*ent.EmailEvent, error) {
	return query.All(ctx)
}

func (rps *emailEventRepoImpl) GetEmailEvent(ctx context.Context, id uuid.UUID) (*ent.EmailEvent, error) {
	return rps.BuildBaseQuery().Where(emailevent.ID(id)).Only(ctx)
}
