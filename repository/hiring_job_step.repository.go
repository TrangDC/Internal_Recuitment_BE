package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/hiringjobstep"

	"github.com/google/uuid"
)

type HiringJobStepRepository interface {
	// base func
	BuildCreate() *ent.HiringJobStepCreate
	BuildQuery() *ent.HiringJobStepQuery
	// mutation
	CreateHiringJobStep(ctx context.Context, step hiringjobstep.Status, hiringJobId uuid.UUID) error
	DeleteHiringJobStep(ctx context.Context, hiringJobId uuid.UUID) error
	CreateBulkHiringJobSteps(ctx context.Context, creates []*ent.HiringJobStepCreate) ([]*ent.HiringJobStep, error)
	UpdateHiringJobStepStatus(ctx context.Context, record *ent.HiringJobStep, inputStatus ent.HiringJobStepStatusEnum) (*ent.HiringJobStep, error)
	UpdateHiringJobStepByRecLeader(ctx context.Context, recLeaderStep *ent.HiringJobStep, recLeaderID uuid.UUID) (*ent.HiringJobStep, error)
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
	currentTime := time.Now().UTC()
	return rps.client.HiringJobStep.Create().SetCreatedAt(currentTime).SetUpdatedAt(currentTime)
}

func (rps hiringJobStepRepoImpl) BuildQuery() *ent.HiringJobStepQuery {
	return rps.client.HiringJobStep.Query()
}

func (rps *hiringJobStepRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.HiringJobStep) *ent.HiringJobStepUpdateOne {
	return record.Update().SetUpdatedAt(time.Now().UTC())
}

// mutation
func (rps hiringJobStepRepoImpl) CreateHiringJobStep(ctx context.Context, step hiringjobstep.Status, hiringJobId uuid.UUID) error {
	_, err := rps.BuildCreate().SetStatus(step).SetHiringJobID(hiringJobId).Save(ctx)
	return err
}

func (rps hiringJobStepRepoImpl) DeleteHiringJobStep(ctx context.Context, hiringJobId uuid.UUID) error {
	_, err := rps.client.HiringJobStep.Delete().Where(hiringjobstep.HiringJobID(hiringJobId)).Exec(ctx)
	return err
}

func (rps hiringJobStepRepoImpl) CreateBulkHiringJobSteps(ctx context.Context, creates []*ent.HiringJobStepCreate) ([]*ent.HiringJobStep, error) {
	return rps.client.HiringJobStep.CreateBulk(creates...).Save(ctx)
}

func (rps hiringJobStepRepoImpl) UpdateHiringJobStepStatus(ctx context.Context, record *ent.HiringJobStep, inputStatus ent.HiringJobStepStatusEnum) (*ent.HiringJobStep, error) {
	return record.Update().SetUpdatedAt(time.Now().UTC()).SetStatus(hiringjobstep.Status(inputStatus)).Save(ctx)
}

func (rps hiringJobStepRepoImpl) UpdateHiringJobStepByRecLeader(ctx context.Context, recLeaderStep *ent.HiringJobStep, recLeaderID uuid.UUID) (*ent.HiringJobStep, error) {
	return rps.BuildUpdateOne(ctx, recLeaderStep).SetUserID(recLeaderID).Save(ctx)
}
