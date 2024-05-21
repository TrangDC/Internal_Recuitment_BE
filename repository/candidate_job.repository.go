package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateJobRepository interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJob, error)
	DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error
	UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, status ent.CandidateJobStatus) (*ent.CandidateJob, error)
	UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error)
	// query
	GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error)
	BuildQuery() *ent.CandidateJobQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error)
	// common function
	ValidStatus(ctx context.Context, candidateId uuid.UUID, candidateJobId uuid.UUID, status *ent.CandidateJobStatus) error
	ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) error
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
	return rps.client.CandidateJob.Create().SetUpdatedAt(time.Now())
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
	return rps.client.CandidateJob.Query().Where(candidatejob.DeletedAtIsNil()).WithAttachmentEdges(
		func(query *ent.AttachmentQuery) {
			query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobs))
		},
	).WithCandidateEdge().WithHiringJobEdge(
		func(query *ent.HiringJobQuery) {
			query.WithTeamEdge().WithOwnerEdge()
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
	_, err := rps.client.Candidate.Update().Where(candidate.IDEQ(uuid.MustParse(input.CandidateID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = rps.client.HiringJob.Update().Where(hiringjob.IDEQ(uuid.MustParse(input.HiringJobID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
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

func (rps candidateJobRepoImpl) UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error) {
	return rps.BuildUpdateOne(ctx, record).RemoveAttachmentEdges().Save(ctx)
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
func (rps candidateJobRepoImpl) ValidStatus(ctx context.Context, candidateId uuid.UUID, candidateJobId uuid.UUID, status *ent.CandidateJobStatus) error {
	if !ent.CandidateJobStatusOpen.IsValid(ent.CandidateJobStatusOpen(status.String())) {
		return nil
	}
	query := rps.BuildQuery().Where(candidatejob.CandidateIDEQ(candidateId))
	if candidateJobId != uuid.Nil {
		query.Where(candidatejob.IDNEQ(candidateJobId))
	}
	openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(s)
	})
	query = query.Where(candidatejob.StatusIn(openStatus...))
	isExist, _ := rps.BuildExist(ctx, query)
	if isExist {
		return fmt.Errorf("model.candidate_job.validation.candidate_job_status_exist")
	}
	return nil
}

func (rps candidateJobRepoImpl) ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) error {
	candidateRecord, err := rps.client.Candidate.Query().Where(candidate.IDEQ(candidateId)).First(ctx)
	if err != nil {
		return err
	}
	if candidateRecord.IsBlacklist {
		return fmt.Errorf("model.candidate_job.validation.candidate_is_blacklist")
	}
	return nil
}
