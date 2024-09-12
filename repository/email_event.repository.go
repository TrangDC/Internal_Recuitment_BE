package repository

import (
	"context"
	"trec/ent"
)

type EmailEventRepository interface {
	BuildBaseQuery() *ent.EmailEventQuery
	BuildList(ctx context.Context, query *ent.EmailEventQuery) ([]*ent.EmailEvent, error)
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
