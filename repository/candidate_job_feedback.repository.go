package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/hiringjob"
	"trec/middleware"

	"github.com/google/uuid"
)

type CandidateJobFeedbackRepository interface {
	// mutation
	CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput) (*ent.CandidateJobFeedback, error)
	UpdateCandidateJobFeedback(ctx context.Context, model *ent.CandidateJobFeedback, input *ent.UpdateCandidateJobFeedbackInput) (*ent.CandidateJobFeedback, error)
	// query
	GetCandidateJobFeedback(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJobFeedback, error)
	BuildQuery() *ent.CandidateJobFeedbackQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobFeedbackQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobFeedbackQuery) ([]*ent.CandidateJobFeedback, error)

	ValidJob(ctx context.Context, candidateJobId uuid.UUID) error
	ValidCandidate(ctx context.Context, candidateJobId uuid.UUID) error
}

type CandidateJobFeedbackRepoImpl struct {
	client *ent.Client
}

func NewCandidateJobFeedbackRepository(client *ent.Client) CandidateJobFeedbackRepository {
	return &CandidateJobFeedbackRepoImpl{
		client: client,
	}
}

// Base function
func (rps CandidateJobFeedbackRepoImpl) BuildCreate() *ent.CandidateJobFeedbackCreate {
	return rps.client.CandidateJobFeedback.Create()
}

func (rps CandidateJobFeedbackRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateJobFeedbackCreate) ([]*ent.CandidateJobFeedback, error) {
	return rps.client.CandidateJobFeedback.CreateBulk(input...).Save(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) BuildUpdate() *ent.CandidateJobFeedbackUpdate {
	return rps.client.CandidateJobFeedback.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps CandidateJobFeedbackRepoImpl) BuildDelete() *ent.CandidateJobFeedbackUpdate {
	return rps.client.CandidateJobFeedback.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps CandidateJobFeedbackRepoImpl) BuildQuery() *ent.CandidateJobFeedbackQuery {
	return rps.client.CandidateJobFeedback.Query().Where(candidatejobfeedback.DeletedAtIsNil()).WithAttachmentEdges(
		func(query *ent.AttachmentQuery) {
			query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobFeedbacks))
		},
	)
}

func (rps CandidateJobFeedbackRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateJobFeedbackQuery) (*ent.CandidateJobFeedback, error) {
	return query.First(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) BuildList(ctx context.Context, query *ent.CandidateJobFeedbackQuery) ([]*ent.CandidateJobFeedback, error) {
	return query.All(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateJobFeedbackQuery) (int, error) {
	return query.Count(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateJobFeedbackQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.CandidateJobFeedback) *ent.CandidateJobFeedbackUpdateOne {
	return model.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps CandidateJobFeedbackRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateJobFeedbackUpdateOne) (*ent.CandidateJobFeedback, error) {
	return update.Save(ctx)
}

// mutation
func (rps CandidateJobFeedbackRepoImpl) CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput) (*ent.CandidateJobFeedback, error) {
	id := ctx.Value(middleware.Key{})
	return rps.BuildCreate().
		SetCandidateJobID(uuid.MustParse(input.CandidateJobID)).
		SetUpdatedAt(time.Now().UTC()).
		SetCreatedBy(id.(uuid.UUID)).
		Save(ctx)
}

func (rps CandidateJobFeedbackRepoImpl) UpdateCandidateJobFeedback(ctx context.Context, model *ent.CandidateJobFeedback, input *ent.UpdateCandidateJobFeedbackInput) (*ent.CandidateJobFeedback, error) {
	return rps.BuildUpdateOne(ctx, model).SetFeedback(input.Feedback).Save(ctx)
}

// query
func (rps CandidateJobFeedbackRepoImpl) GetCandidateJobFeedback(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJobFeedback, error) {
	return rps.BuildQuery().Where(candidatejobfeedback.IDEQ(candidateId)).First(ctx)
}

// common
func (rps CandidateJobFeedbackRepoImpl) ValidJob(ctx context.Context, candidateJobId uuid.UUID) error {
	_, err := rps.client.CandidateJob.Query().Where(candidatejob.IDEQ(candidateJobId), candidatejob.HasHiringJobWith(
		hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusOpened),
	)).First(ctx)
	if err != nil {
		return fmt.Errorf("model.candidate_job_feedbacks.validation.job_close")
	}
	return err
}

func (rps CandidateJobFeedbackRepoImpl) ValidCandidate(ctx context.Context, candidateJobId uuid.UUID) error {
	_, err := rps.client.Candidate.Query().Where(candidate.IsBlacklist(false), candidate.HasCandidateJobEdgesWith(
		candidatejob.DeletedAtIsNil(), candidatejob.IDEQ(candidateJobId),
	)).First(ctx)
	if err != nil {
		return fmt.Errorf("model.candidate_job_feedbacks.validation.candidate_is_blacklist")
	}
	return err
}
