package repository

import (
	"context"
	"time"
	"trec/ent"
)

type HiringJobRepository interface {
}

type hiringJobRepoImpl struct {
	client *ent.Client
}

func NewHiringJobRepository(client *ent.Client) HiringJobRepository {
	return &hiringJobRepoImpl{
		client: client,
	}
}

// Base function
func (rps *hiringJobRepoImpl) BuildCreate() *ent.HiringJobCreate {
	return rps.client.HiringJob.Create()
}

func (rps *hiringJobRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.HiringJobCreate) ([]*ent.HiringJob, error) {
	return rps.client.HiringJob.CreateBulk(input...).Save(ctx)
}

func (rps *hiringJobRepoImpl) BuildUpdate() *ent.HiringJobUpdate {
	return rps.client.HiringJob.Update().SetUpdatedAt(time.Now())
}

func (rps *hiringJobRepoImpl) BuildDelete() *ent.HiringJobUpdate {
	return rps.client.HiringJob.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *hiringJobRepoImpl) BuildQuery() *ent.HiringJobQuery {
	return rps.client.HiringJob.Query()
}

func (rps *hiringJobRepoImpl) BuildGet(ctx context.Context, query *ent.HiringJobQuery) (*ent.HiringJob, error) {
	return query.First(ctx)
}

func (rps *hiringJobRepoImpl) BuildList(ctx context.Context, query *ent.HiringJobQuery) ([]*ent.HiringJob, error) {
	return query.All(ctx)
}

func (rps *hiringJobRepoImpl) BuildCount(ctx context.Context, query *ent.HiringJobQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *hiringJobRepoImpl) BuildExist(ctx context.Context, query *ent.HiringJobQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *hiringJobRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.HiringJob) *ent.HiringJobUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *hiringJobRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.HiringJobUpdateOne) (*ent.HiringJob, error) {
	return update.Save(ctx)
}
