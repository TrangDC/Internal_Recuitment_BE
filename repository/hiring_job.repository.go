package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/internal/util"

	"github.com/google/uuid"
)

type HiringJobRepository interface {
	// mutation
	CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJob, error)
	UpdateHiringJob(ctx context.Context, record *ent.HiringJob, input *ent.UpdateHiringJobInput) (*ent.HiringJob, error)
	UpdateHiringJobStatus(ctx context.Context, record *ent.HiringJob, status ent.HiringJobStatus) (*ent.HiringJob, error)
	DeleteHiringJob(ctx context.Context, record *ent.HiringJob) error
	// query
	GetHiringJob(ctx context.Context, hiringJobId uuid.UUID) (*ent.HiringJob, error)
	BuildQuery() *ent.HiringJobQuery
	BuildCount(ctx context.Context, query *ent.HiringJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.HiringJobQuery) ([]*ent.HiringJob, error)
	// common function
	ValidName(ctx context.Context, hiringJobId uuid.UUID, name string) error
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
	return query.WithOwnerEdge().WithTeamEdge().All(ctx)
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

// mutation
func (rps *hiringJobRepoImpl) CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJob, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetAmount(input.Amount).
		SetLocation(hiringjob.Location(input.Location)).
		SetSlug(util.SlugGeneration(input.Name)).
		SetSalaryType(hiringjob.SalaryType(input.SalaryType)).
		SetSalaryFrom(input.SalaryFrom).
		SetSalaryTo(input.SalaryTo).
		SetCurrency(hiringjob.Currency(input.Currency)).
		SetStatus(hiringjob.Status(input.Status)).
		SetTeamID(uuid.MustParse(input.TeamID)).
		SetCreatedBy(uuid.MustParse(input.CreatedBy)).
		Save(ctx)
}

func (rps *hiringJobRepoImpl) UpdateHiringJob(ctx context.Context, record *ent.HiringJob, input *ent.UpdateHiringJobInput) (*ent.HiringJob, error) {
	return rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetAmount(input.Amount).
		SetLocation(hiringjob.Location(input.Location)).
		SetSlug(util.SlugGeneration(input.Name)).
		SetSalaryType(hiringjob.SalaryType(input.SalaryType)).
		SetSalaryFrom(input.SalaryFrom).
		SetSalaryTo(input.SalaryTo).
		SetCurrency(hiringjob.Currency(input.Currency)).
		SetTeamID(uuid.MustParse(input.TeamID)).
		SetCreatedBy(uuid.MustParse(input.CreatedBy)).
		Save(ctx)
}

func (rps *hiringJobRepoImpl) UpdateHiringJobStatus(ctx context.Context, record *ent.HiringJob, status ent.HiringJobStatus) (*ent.HiringJob, error) {
	return rps.BuildUpdateOne(ctx, record).SetStatus(hiringjob.Status(status)).Save(ctx)
}

func (rps *hiringJobRepoImpl) DeleteHiringJob(ctx context.Context, record *ent.HiringJob) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).Save(ctx)
	return err
}

// query
func (rps *hiringJobRepoImpl) GetHiringJob(ctx context.Context, hiringJobId uuid.UUID) (*ent.HiringJob, error) {
	return rps.BuildQuery().WithOwnerEdge().WithTeamEdge().Where(hiringjob.IDEQ(hiringJobId)).First(ctx)
}

// common function
func (rps *hiringJobRepoImpl) ValidName(ctx context.Context, hiringJobId uuid.UUID, name string) error {
	query := rps.BuildQuery().Where(hiringjob.SlugEQ(util.SlugGeneration(name)))
	if hiringJobId != uuid.Nil {
		query = query.Where(hiringjob.IDNEQ(hiringJobId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if isExist {
		return fmt.Errorf("model.hiring_jobs.validation.name_exist")
	}
	return err
}
