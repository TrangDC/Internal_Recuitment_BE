package repository

import (
	"context"
	"time"
	"trec/ent"
)

type TeamManagerRepository interface {
}

type teamManagerRepoImpl struct {
	client *ent.Client
}

func NewTeamManagerRepository(client *ent.Client) TeamManagerRepository {
	return &teamManagerRepoImpl{
		client: client,
	}
}

// Base function
func (rps *teamManagerRepoImpl) BuildCreate() *ent.TeamManagerCreate {
	return rps.client.TeamManager.Create()
}

func (rps *teamManagerRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.TeamManagerCreate) ([]*ent.TeamManager, error) {
	return rps.client.TeamManager.CreateBulk(input...).Save(ctx)
}

func (rps *teamManagerRepoImpl) BuildUpdate() *ent.TeamManagerUpdate {
	return rps.client.TeamManager.Update().SetUpdatedAt(time.Now())
}

func (rps *teamManagerRepoImpl) BuildDelete() *ent.TeamManagerUpdate {
	return rps.client.TeamManager.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *teamManagerRepoImpl) BuildQuery() *ent.TeamManagerQuery {
	return rps.client.TeamManager.Query()
}

func (rps *teamManagerRepoImpl) BuildGet(ctx context.Context, query *ent.TeamManagerQuery) (*ent.TeamManager, error) {
	return query.First(ctx)
}

func (rps *teamManagerRepoImpl) BuildList(ctx context.Context, query *ent.TeamManagerQuery) ([]*ent.TeamManager, error) {
	return query.All(ctx)
}

func (rps *teamManagerRepoImpl) BuildCount(ctx context.Context, query *ent.TeamManagerQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *teamManagerRepoImpl) BuildExist(ctx context.Context, query *ent.TeamManagerQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *teamManagerRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.TeamManager) *ent.TeamManagerUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *teamManagerRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.TeamManagerUpdateOne) (*ent.TeamManager, error) {
	return update.Save(ctx)
}
