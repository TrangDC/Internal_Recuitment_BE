package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/ent/jobposition"

	"github.com/google/uuid"
)

type JobPositionRepository interface {
	// mutation
	CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPosition, error)
	UpdateJobPosition(ctx context.Context, record *ent.JobPosition, input ent.UpdateJobPositionInput) (*ent.JobPosition, error)
	DeleteJobPosition(ctx context.Context, record *ent.JobPosition) (*ent.JobPosition, error)

	// query
	GetJobPosition(ctx context.Context, id uuid.UUID) (*ent.JobPosition, error)
	BuildQuery() *ent.JobPositionQuery
	BuildCount(ctx context.Context, query *ent.JobPositionQuery) (int, error)
	BuildList(ctx context.Context, query *ent.JobPositionQuery) ([]*ent.JobPosition, error)

	// common function
	ValidName(ctx context.Context, teamId uuid.UUID, name string) (error, error)
}

type jobPositionRepoImpl struct {
	client *ent.Client
}

func NewJobPositionRepository(client *ent.Client) JobPositionRepository {
	return &jobPositionRepoImpl{
		client: client,
	}
}

// base function
func (rps *jobPositionRepoImpl) BuildCreate() *ent.JobPositionCreate {
	return rps.client.JobPosition.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *jobPositionRepoImpl) BuildQuery() *ent.JobPositionQuery {
	return rps.client.JobPosition.Query().Where(jobposition.DeletedAtIsNil()).WithHiringJobPositionEdges(
		func(query *ent.HiringJobQuery) {
			query.Where(hiringjob.DeletedAtIsNil())
		},
	)
}

func (rps *jobPositionRepoImpl) BuildGet(ctx context.Context, query *ent.JobPositionQuery) (*ent.JobPosition, error) {
	return query.First(ctx)
}

func (rps *jobPositionRepoImpl) BuildList(ctx context.Context, query *ent.JobPositionQuery) ([]*ent.JobPosition, error) {
	return query.All(ctx)
}

func (rps *jobPositionRepoImpl) BuildCount(ctx context.Context, query *ent.JobPositionQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *jobPositionRepoImpl) BuildExist(ctx context.Context, query *ent.JobPositionQuery) (bool, error) {
	return query.Exist(ctx)
}
func (rps *jobPositionRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.JobPosition) *ent.JobPositionUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

// mutation
func (rps *jobPositionRepoImpl) CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPosition, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).Save(ctx)
}

func (rps *jobPositionRepoImpl) UpdateJobPosition(ctx context.Context, record *ent.JobPosition, input ent.UpdateJobPositionInput) (*ent.JobPosition, error) {
	return rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).Save(ctx)
}

func (rps *jobPositionRepoImpl) DeleteJobPosition(ctx context.Context, record *ent.JobPosition) (*ent.JobPosition, error) {
	update := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
	return update.Save(ctx)
}

// query
func (rps *jobPositionRepoImpl) GetJobPosition(ctx context.Context, id uuid.UUID) (*ent.JobPosition, error) {
	query := rps.BuildQuery().Where(jobposition.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *jobPositionRepoImpl) ValidName(ctx context.Context, jobPositionId uuid.UUID, name string) (error, error) {
	query := rps.BuildQuery().Where(jobposition.NameEqualFold(strings.TrimSpace(name)))
	if jobPositionId != uuid.Nil {
		query = query.Where(jobposition.IDNEQ(jobPositionId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.job_positions.validation.name_exist"), nil
	}
	return nil, nil
}

// Path: repository/job_position.repository.go
