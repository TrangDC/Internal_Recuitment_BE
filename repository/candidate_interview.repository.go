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
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateInterviewRepository interface {
	CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID, status string) (*ent.CandidateInterview, error)
	UpdateCandidateInterview(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error)
	UpdateCandidateInterviewSchedule(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewScheduleInput, newMemberIds, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error)
	DeleteCandidateInterview(ctx context.Context, record *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error)
	CreateBulkCandidateInterview(ctx context.Context, candidateJobs []*ent.CandidateJob, memberIds []uuid.UUID, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateInterview, error)
	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error)
	BuildQuery() *ent.CandidateInterviewQuery
	BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error)

	// common function
	ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (string, error, error)
	ValidateCreateBulkInput(ctx context.Context, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateJob, error, error)
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

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterviewSchedule(ctx context.Context, model *ent.CandidateInterview, input ent.UpdateCandidateInterviewScheduleInput, newMemberIds, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt)
	if len(newMemberIds) > 0 && len(removeMemberIds) > 0 {
		update.AddInterviewerEdgeIDs(newMemberIds...).RemoveInterviewerEdgeIDs(removeMemberIds...)
	}
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) DeleteCandidateInterview(ctx context.Context, model *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveInterviewerEdgeIDs(memberIds...)
	return update.Save(ctx)
}

func (rps candidateInterviewRepoImpl) CreateBulkCandidateInterview(ctx context.Context, candidateJobs []*ent.CandidateJob, memberIds []uuid.UUID, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateInterview, error) {
	var createBulk []*ent.CandidateInterviewCreate
	var createInterviewers []*ent.CandidateInterviewerCreate
	for _, record := range candidateJobs {
		createBulk = append(createBulk, rps.BuildCreate().SetTitle(input.Title).
			SetInterviewDate(input.InterviewDate).
			SetStartFrom(input.StartFrom).
			SetDescription(input.Description).
			SetEndAt(input.EndAt).
			SetCandidateJobID(record.ID).
			SetCandidateJobStatus(candidateinterview.CandidateJobStatus(record.Status.String())))
	}
	candidateInterviews, err := rps.client.CandidateInterview.CreateBulk(createBulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	for _, record := range candidateInterviews {
		createBulkThings := lo.Map(memberIds, func(item uuid.UUID, index int) *ent.CandidateInterviewerCreate {
			return rps.client.CandidateInterviewer.Create().SetCandidateInterviewID(record.ID).SetUserID(item)
		})
		createInterviewers = append(createInterviewers, createBulkThings...)
	}
	_, err = rps.client.CandidateInterviewer.CreateBulk(createInterviewers...).Save(ctx)
	return candidateInterviews, err
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
	stringError, err = rps.ValidateSchedule(ctx, candidateInterviewId, input.CandidateJobId, input.InterviewDate, input.StartFrom, input.EndAt)
	if err != nil || stringError != nil {
		return "", stringError, err
	}
	return record.Status.String(), nil, err
}

func (rps *candidateInterviewRepoImpl) ValidateCreateBulkInput(ctx context.Context, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateJob, error, error) {
	// validate job
	_, err := rps.client.HiringJob.Query().Where(
		hiringjob.IDEQ(uuid.MustParse(input.JobID)),
		hiringjob.DeletedAtIsNil(),
		hiringjob.StatusEQ(hiringjob.StatusOpened),
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("model.candidate_interviews.validation.job_close"), nil
	}
	// validate candidate
	candidateIds := lo.Uniq(lo.Map(input.CandidateID, func(item string, index int) uuid.UUID {
		return uuid.MustParse(item)
	}))
	candidates, err := rps.client.Candidate.Query().Where(candidate.IDIn(candidateIds...), candidate.IsBlacklist(false), candidate.DeletedAtIsNil()).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	if len(candidates) != len(candidateIds) {
		return nil, fmt.Errorf("model.candidate_interviews.validation.candidate_is_blacklist"), nil
	}
	// validate candidate job
	candidateJobs, err := rps.client.CandidateJob.Query().Where(
		candidatejob.CandidateIDIn(candidateIds...), candidatejob.DeletedAtIsNil(),
		candidatejob.HiringJobIDEQ(uuid.MustParse(input.JobID)), candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	if len(candidateJobs) != len(candidateIds) {
		return nil, fmt.Errorf("model.candidate_interviews.validation.candidate_job_not_found"), nil
	}
	// validate title
	for _, record := range candidateJobs {
		stringError, err := rps.ValidTitle(ctx, record.ID, uuid.Nil, input.Title, record.Status.String())
		if err != nil || stringError != nil {
			return nil, stringError, err
		}
		stringError, err = rps.ValidateSchedule(ctx, uuid.Nil, record.ID, &input.InterviewDate, &input.StartFrom, &input.EndAt)
		if err != nil || stringError != nil {
			return nil, stringError, err
		}
	}
	return candidateJobs, nil, nil
}

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

func (rps *candidateInterviewRepoImpl) ValidateSchedule(ctx context.Context, candidateInterviewId uuid.UUID, candidateJobId uuid.UUID, interviewDate, startFrom, endAt *time.Time) (error, error) {
	if startFrom.After(*endAt) {
		return fmt.Errorf("model.candidate_interviews.validation.start_from_end_at_invalid"), nil
	}
	query := rps.client.CandidateInterview.Query().
		Where(candidateinterview.CandidateJobID(candidateJobId), candidateinterview.DeletedAtIsNil(),
			candidateinterview.InterviewDateEQ(*interviewDate))
	if candidateInterviewId != uuid.Nil {
		query.Where(candidateinterview.IDNEQ(candidateInterviewId))
	}
	query.Where(candidateinterview.Or(
		candidateinterview.And(candidateinterview.StartFromGTE(*startFrom), candidateinterview.EndAtLTE(*endAt)), // [start, [start, end] ,end]
		candidateinterview.StartFromIn(*startFrom, *endAt),
		candidateinterview.EndAtIn(*startFrom, *endAt)),
		candidateinterview.And(candidateinterview.StartFromLTE(*startFrom), candidateinterview.EndAtGTE(*endAt)))
	exist, err := query.Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return fmt.Errorf("model.candidate_interviews.validation.schedule_exist"), nil
	}
	return nil, nil
}
