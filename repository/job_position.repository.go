package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/jobposition"

	"github.com/google/uuid"
)

type JobPositionRepository interface {
	// mutation
	CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPosition, error)

	// query
	GetJobPosition(ctx context.Context, id uuid.UUID) (*ent.JobPosition, error)
	BuildQuery() *ent.JobPositionQuery

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
	return rps.client.JobPosition.Query().Where(jobposition.DeletedAtIsNil())
}

func (rps *jobPositionRepoImpl) BuildGet(ctx context.Context, query *ent.JobPositionQuery) (*ent.JobPosition, error) {
	return query.First(ctx)
}

func (rps *jobPositionRepoImpl) BuildExist(ctx context.Context, query *ent.JobPositionQuery) (bool, error) {
	return query.Exist(ctx)
}

// mutation
func (rps *jobPositionRepoImpl) CreateJobPosition(ctx context.Context, input ent.NewJobPositionInput) (*ent.JobPosition, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).Save(ctx)
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
