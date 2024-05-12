package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/user"

	"github.com/google/uuid"
)

type CandidateInterviewRepository interface {
	CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID) (*ent.CandidateInterview, error)
	UpdateCandidateInterview(ctx context.Context, model *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error)
	DeleteCandidateInterview(ctx context.Context, model *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error)

	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error)
	BuildQuery() *ent.CandidateInterviewQuery
	BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error)

	// common function
	ValidTitle(ctx context.Context, candidateJobId uuid.UUID, candidateInterviewId uuid.UUID, title string, status string) error
	ValidateStatus(ctx context.Context, candidateJobId uuid.UUID) (string, error)
	ValidJob(ctx context.Context, candidateJobId uuid.UUID) error
	ValidCandidate(ctx context.Context, candidateJobId uuid.UUID) error
	ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, title string, candidateJobId uuid.UUID) error
}

type candidateInterviewRepoImpl struct {
	client *ent.Client
}

func NewCandidateInterviewRepository(client *ent.Client) CandidateInterviewRepository {
	return &candidateInterviewRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *candidateInterviewRepoImpl) BuildCreate() *ent.CandidateInterviewCreate {
	return rps.client.CandidateInterview.Create().SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildUpdate() *ent.CandidateInterviewUpdate {
	return rps.client.CandidateInterview.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildDelete() *ent.CandidateInterviewUpdate {
	return rps.client.CandidateInterview.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildQuery() *ent.CandidateInterviewQuery {
	return rps.client.CandidateInterview.Query().Where(candidateinterview.DeletedAtIsNil()).WithInterviewerEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	)
}

func (rps *candidateInterviewRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateInterviewQuery) (*ent.CandidateInterview, error) {
	return query.First(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error) {
	return query.All(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateInterviewQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.CandidateInterview) *ent.CandidateInterviewUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateInterviewUpdateOne) (*ent.CandidateInterview, error) {
	return update.Save(ctx)
}

// mutation
func (rps *candidateInterviewRepoImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	create := rps.BuildCreate().SetTitle(strings.TrimSpace(input.Title)).AddUserInterviewerIDs(memberIds...)
	return create.Save(ctx)
}

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterview(ctx context.Context, model *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).SetTitle(strings.TrimSpace(input.Title)).
		AddUserInterviewerIDs(newMemberIds...).RemoveUserInterviewerIDs(removeMemberIds...)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) DeleteCandidateInterview(ctx context.Context, model *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveUserInterviewerIDs(memberIds...)
	return update.Save(ctx)
}

// query
func (rps *candidateInterviewRepoImpl) GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error) {
	query := rps.BuildQuery().Where(candidateinterview.IDEQ(id)).WithInterviewerEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	)
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *candidateInterviewRepoImpl) ValidTitle(ctx context.Context, candidateJobId uuid.UUID, candidateInterviewId uuid.UUID, title string, status string) error {
	query := rps.BuildQuery().Where(candidateinterview.TitleEqualFold(title),
		candidateinterview.CandidateJobStatusEQ(candidateinterview.CandidateJobStatus(status)),
		candidateinterview.CandidateJobID(candidateJobId))
	if candidateInterviewId != uuid.Nil {
		query = query.Where(candidateinterview.IDNEQ(candidateInterviewId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("model.candidate_interviews.validation.title_exist")
	}
	return nil
}

func (rps *candidateInterviewRepoImpl) ValidateStatus(ctx context.Context, candidateJobId uuid.UUID) (string, error) {
	record, err := rps.client.CandidateJob.Query().Where(candidatejob.IDEQ(candidateJobId), candidatejob.DeletedAtIsNil()).First(ctx)
	if err != nil {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_job_not_found")
	}
	if ent.CandidateJobStatusEnded.IsValid(ent.CandidateJobStatusEnded(record.Status.String())) {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_job_status_ended")
	}
	return record.Status.String(), nil
}

func (rps candidateInterviewRepoImpl) ValidJob(ctx context.Context, candidateJobId uuid.UUID) error {
	_, err := rps.client.CandidateJob.Query().Where(candidatejob.IDEQ(candidateJobId), candidatejob.HasHiringJobWith(
		hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusOpened),
	)).First(ctx)
	if err != nil {
		return fmt.Errorf("model.candidate_interviews.validation.job_close")
	}
	return err
}

func (rps candidateInterviewRepoImpl) ValidCandidate(ctx context.Context, candidateJobId uuid.UUID) error {
	_, err := rps.client.Candidate.Query().Where(candidate.IsBlacklist(false), candidate.HasCandidateJobEdgesWith(
		candidatejob.DeletedAtIsNil(), candidatejob.IDEQ(candidateJobId),
	)).First(ctx)
	if err != nil {
		return fmt.Errorf("model.candidate_interviews.validation.candidate_is_blacklist")
	}
	return err
}

func (rps *candidateInterviewRepoImpl) ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, title string, candidateJobId uuid.UUID) error {
	err := rps.ValidCandidate(ctx, candidateJobId)
	if err != nil {
		return err
	}
	err = rps.ValidJob(ctx, candidateJobId)
	if err != nil {
		return err
	}
	status, err := rps.ValidateStatus(ctx, candidateJobId)
	if err != nil {
		return err
	}
	err = rps.ValidTitle(ctx, candidateJobId, candidateInterviewId, title, status)
	if err != nil {
		return err
	}
	return nil
}
