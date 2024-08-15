package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/hiringjobstep"
	"trec/middleware"

	"github.com/google/uuid"
)

type HiringJobStepRepository interface {
	// mutation
	CreateHiringJobStep(ctx context.Context, step hiringjobstep.Type, hiringJobId uuid.UUID) error
	DeleteHiringJobStep(ctx context.Context, hiringJobId uuid.UUID) error
	BuildQuery() *ent.HiringJobStepQuery
}

type hiringJobStepRepoImpl struct {
	client *ent.Client
}

func NewHiringJobStepRepository(client *ent.Client) HiringJobStepRepository {
	return &hiringJobStepRepoImpl{
		client: client,
	}
}

// Base function
func (rps hiringJobStepRepoImpl) BuildCreate() *ent.HiringJobStepCreate {
	return rps.client.HiringJobStep.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps hiringJobStepRepoImpl) BuildQuery() *ent.HiringJobStepQuery {
	return rps.client.HiringJobStep.Query()
}

// mutation
func (rps hiringJobStepRepoImpl) CreateHiringJobStep(ctx context.Context, step hiringjobstep.Type, hiringJobId uuid.UUID) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	_, err := rps.BuildCreate().SetType(step).SetHiringJobID(hiringJobId).SetCreatedByEdgeID(payload.UserID).Save(ctx)
	return err
}

func (rps hiringJobStepRepoImpl) DeleteHiringJobStep(ctx context.Context, hiringJobId uuid.UUID) error {
	_, err := rps.client.HiringJobStep.Delete().Where(hiringjobstep.HiringJobID(hiringJobId)).Exec(ctx)
	return err
}
