package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/user"
	"trec/models"

	"github.com/google/uuid"
)

type CandidateInterviewRepository interface {
	CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID, status string) (*ent.CandidateInterview, error)
	UpdateCandidateInterview(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error)
	UpdateCandidateInterviewSchedule(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterview, error)
	DeleteCandidateInterview(ctx context.Context, record *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error)

	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error)
	BuildQuery() *ent.CandidateInterviewQuery
	BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error)

	// common function
	ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (string, error, error)
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
func (rps *candidateInterviewRepoImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID, status string) (*ent.CandidateInterview, error) {
	create := rps.BuildCreate().SetTitle(strings.TrimSpace(input.Title)).
		AddInterviewerEdgeIDs(memberIds...).
		SetDescription(input.Description).
		SetCandidateJobID(uuid.MustParse(input.CandidateJobID)).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt).
		SetCandidateJobStatus(candidateinterview.CandidateJobStatus(status))
	return create.Save(ctx)
}

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterview(ctx context.Context, model *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).SetTitle(strings.TrimSpace(input.Title)).
		SetDescription(input.Description).SetCandidateJobID(uuid.MustParse(input.CandidateJobID)).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt).
		AddInterviewerEdgeIDs(newMemberIds...).RemoveInterviewerEdgeIDs(removeMemberIds...)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterviewSchedule(ctx context.Context, model *ent.CandidateInterview, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) DeleteCandidateInterview(ctx context.Context, model *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveInterviewerEdgeIDs(memberIds...)
	return update.Save(ctx)
}

// query
func (rps *candidateInterviewRepoImpl) GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error) {
	query := rps.BuildQuery().Where(candidateinterview.IDEQ(id)).WithInterviewerEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	).WithCandidateJobEdge(
		func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.DeletedAtIsNil())
		},
	)
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *candidateInterviewRepoImpl) ValidTitle(ctx context.Context, candidateJobId uuid.UUID, candidateInterviewId uuid.UUID, title string, status string) (error, error) {
	query := rps.BuildQuery().Where(candidateinterview.TitleEqualFold(title),
		candidateinterview.CandidateJobStatusEQ(candidateinterview.CandidateJobStatus(status)),
		candidateinterview.CandidateJobID(candidateJobId))
	if candidateInterviewId != uuid.Nil {
		query = query.Where(candidateinterview.IDNEQ(candidateInterviewId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, fmt.Errorf("model.candidate_interviews.validation.title_exist")
	}
	return nil, nil
}

func (rps *candidateInterviewRepoImpl) ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (string, error, error) {
	record, err := rps.client.CandidateJob.Query().Where(candidatejob.IDEQ(input.CandidateJobId), candidatejob.DeletedAtIsNil()).WithHiringJobEdge(
		func(query *ent.HiringJobQuery) {
			query.Where(hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusOpened))
		},
	).WithCandidateEdge().First(ctx)
	if err != nil {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_job_not_found"), nil
	}
	if record.Edges.CandidateEdge.IsBlacklist {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_is_blacklist"), nil
	}
	if record.Edges.HiringJobEdge == nil {
		return "", fmt.Errorf("model.candidate_interviews.validation.job_close"), nil
	}
	if ent.CandidateJobStatusEnded.IsValid(ent.CandidateJobStatusEnded(record.Status.String())) || record.Status == candidatejob.StatusOffering {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_job_status_ended"), nil
	}
	stringError, err := rps.ValidTitle(ctx, input.CandidateJobId, candidateInterviewId, input.Title, record.Status.String())
	if err != nil || stringError != nil {
		return "", stringError, err
	}
	stringError, err = rps.ValidateSchedule(ctx, candidateInterviewId, input)
	if err != nil || stringError != nil {
		return "", stringError, err
	}
	return record.Status.String(), nil, err
}

func (rps *candidateInterviewRepoImpl) ValidateSchedule(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (error, error) {
	if input.InterviewDate.Before(time.Now()) {
		return fmt.Errorf("model.candidate_interviews.validation.interview_date_invalid"), nil
	}
	if input.StartFrom.After(*input.EndAt) {
		return fmt.Errorf("model.candidate_interviews.validation.start_from_end_at_invalid"), nil
	}
	query := rps.client.CandidateInterview.Query().
		Where(candidateinterview.IDNEQ(candidateInterviewId), candidateinterview.CandidateJobID(input.CandidateJobId),
			candidateinterview.InterviewDateEQ(*input.InterviewDate))
	query.Where(candidateinterview.Or(
		candidateinterview.And(candidateinterview.StartFromGTE(*input.StartFrom), candidateinterview.EndAtLTE(*input.EndAt)), // [start, [start, end] ,end]
		candidateinterview.StartFromIn(*input.StartFrom, *input.EndAt),
		candidateinterview.EndAtIn(*input.StartFrom, *input.EndAt)),
		candidateinterview.And(candidateinterview.StartFromLTE(*input.StartFrom), candidateinterview.EndAtGTE(*input.EndAt)))
	exist, err := query.Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return fmt.Errorf("model.candidate_interviews.validation.schedule_exist"), nil
	}
	return nil, nil
}
