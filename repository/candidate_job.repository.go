package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidatejob"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateJobRepository interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJob, error)
	DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error
	UpdateCandidateJobAttachment(ctx context.Context, record *ent.CandidateJob, input *ent.UpdateCandidateAttachment) (*ent.CandidateJob, error)
	UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, status ent.CandidateJobStatus) (*ent.CandidateJob, error)
	// query
	GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error)
	BuildQuery() *ent.CandidateJobQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error)
	// common function
	ValidStatus(ctx context.Context, candidateId uuid.UUID, status *ent.CandidateJobStatus) error
}

type candidateJobRepoImpl struct {
	client *ent.Client
}

func NewCandidateJobRepository(client *ent.Client) CandidateJobRepository {
	return &candidateJobRepoImpl{
		client: client,
	}
}

// Base function
func (rps candidateJobRepoImpl) BuildCreate() *ent.CandidateJobCreate {
	return rps.client.CandidateJob.Create()
}

func (rps candidateJobRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateJobCreate) ([]*ent.CandidateJob, error) {
	return rps.client.CandidateJob.CreateBulk(input...).Save(ctx)
}

func (rps candidateJobRepoImpl) BuildUpdate() *ent.CandidateJobUpdate {
	return rps.client.CandidateJob.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateJobRepoImpl) BuildDelete() *ent.CandidateJobUpdate {
	return rps.client.CandidateJob.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps candidateJobRepoImpl) BuildQuery() *ent.CandidateJobQuery {
	return rps.client.CandidateJob.Query().WithAttachmentEdges(
		func(query *ent.AttachmentQuery) {
			query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobs))
		},
	)
}

func (rps candidateJobRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error) {
	return query.First(ctx)
}

func (rps candidateJobRepoImpl) BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error) {
	return query.All(ctx)
}

func (rps candidateJobRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error) {
	return query.Count(ctx)
}

func (rps candidateJobRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateJobQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps candidateJobRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.CandidateJob) *ent.CandidateJobUpdateOne {
	return model.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateJobRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateJobUpdateOne) (*ent.CandidateJob, error) {
	return update.Save(ctx)
}

// mutation
func (rps candidateJobRepoImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJob, error) {
	return rps.BuildCreate().
		SetHiringJobID(uuid.MustParse(input.HiringJobID)).
		SetUpdatedAt(time.Now().UTC()).
		SetCandidateID(uuid.MustParse(input.CandidateID)).
		SetStatus(candidatejob.Status(input.Status)).
		Save(ctx)
}

func (rps candidateJobRepoImpl) UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, status ent.CandidateJobStatus) (*ent.CandidateJob, error) {
	return rps.BuildUpdateOne(ctx, record).SetStatus(candidatejob.Status(status.String())).Save(ctx)
}

func (rps candidateJobRepoImpl) UpdateCandidateJobAttachment(ctx context.Context, record *ent.CandidateJob, input *ent.UpdateCandidateAttachment) (*ent.CandidateJob, error) {
	return rps.BuildUpdateOne(ctx, record).Save(ctx)
}

func (rps candidateJobRepoImpl) DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

// query
func (rps candidateJobRepoImpl) GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error) {
	return rps.BuildQuery().Where(candidatejob.IDEQ(candidateId)).First(ctx)
}

// common function
func (rps candidateJobRepoImpl) ValidStatus(ctx context.Context, candidateId uuid.UUID, status *ent.CandidateJobStatus) error {
	if !ent.CandidateJobStatusOpen.IsValid(ent.CandidateJobStatusOpen(status.String())) {
		return nil
	}
	query := rps.BuildQuery()
	if candidateId != uuid.Nil {
		query = query.Where(candidatejob.IDNEQ(candidateId))
	}
	openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(s)
	})
	query = query.Where(candidatejob.StatusIn(openStatus...))
	isExist, _ := rps.BuildExist(ctx, query)
	if isExist {
		return fmt.Errorf("model.candidates.validation.status_not_valid")
	}
	return nil
}
