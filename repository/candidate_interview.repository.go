package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/entitypermission"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/permission"
	"trec/ent/predicate"
	"trec/ent/recteam"
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/ent/user"
	"trec/middleware"
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
	UpdateCandidateInterviewStatus(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewStatusInput) (*ent.CandidateInterview, error)
	UpdateBulkCandidateInterviewStatus(ctx context.Context, predicates []predicate.CandidateInterview, status candidateinterview.Status) error
	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error)
	BuildQuery() *ent.CandidateInterviewQuery
	BuildBaseQuery() *ent.CandidateInterviewQuery
	BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error)
	BuildGetOne(ctx context.Context, query *ent.CandidateInterviewQuery) (*ent.CandidateInterview, error)
	GetDataForKeyword(ctx context.Context, record *ent.CandidateInterview, candidateJobRecord *ent.CandidateJob) (models.GroupModule, error)
	BuildStatusCountByCdJobID(ctx context.Context, candidateJobIDs []uuid.UUID) ([]models.CdInterviewCountByStatus, error)

	// third party
	CallbackInterviewSchedule(ctx context.Context, input models.MessageOutput) error

	// common function
	ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (string, error, error)
	ValidateCreateBulkInput(ctx context.Context, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateJob, error, error)
	ValidCandidateInterviewStatus(record *ent.CandidateInterview, status ent.CandidateInterviewStatusEditable) error
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
	return rps.client.CandidateInterview.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *candidateInterviewRepoImpl) BuildUpdate() *ent.CandidateInterviewUpdate {
	return rps.client.CandidateInterview.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildDelete() *ent.CandidateInterviewUpdate {
	return rps.client.CandidateInterview.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildQuery() *ent.CandidateInterviewQuery {
	return rps.client.CandidateInterview.Query().Where(candidateinterview.DeletedAtIsNil()).
		WithInterviewerEdges(
			func(query *ent.UserQuery) {
				query.Where(user.DeletedAtIsNil())
			},
		).WithCandidateJobEdge(
		func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.DeletedAtIsNil()).WithCandidateEdge().WithHiringJobEdge(
				func(query *ent.HiringJobQuery) {
					query.Where(hiringjob.DeletedAtIsNil()).WithHiringTeamEdge()
				},
			)
		},
	).WithCreatedByEdge()
}

func (rps *candidateInterviewRepoImpl) BuildBaseQuery() *ent.CandidateInterviewQuery {
	return rps.client.CandidateInterview.Query().Where(candidateinterview.DeletedAtIsNil())
}

func (rps *candidateInterviewRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateInterviewQuery) (*ent.CandidateInterview, error) {
	return query.First(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildList(ctx context.Context, query *ent.CandidateInterviewQuery) ([]*ent.CandidateInterview, error) {
	return query.All(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildGetOne(ctx context.Context, query *ent.CandidateInterviewQuery) (*ent.CandidateInterview, error) {
	return query.First(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateInterviewQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateInterviewQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *candidateInterviewRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.CandidateInterview) *ent.CandidateInterviewUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateInterviewRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateInterviewUpdateOne) (*ent.CandidateInterview, error) {
	return update.Save(ctx)
}

// mutation
func (rps *candidateInterviewRepoImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, memberIds []uuid.UUID, status string) (*ent.CandidateInterview, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	createdById := payload.UserID
	create := rps.BuildCreate().SetTitle(strings.TrimSpace(input.Title)).
		AddInterviewerEdgeIDs(memberIds...).
		SetDescription(input.Description).
		SetCandidateJobID(uuid.MustParse(input.CandidateJobID)).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt).
		SetCreatedBy(createdById).
		SetCandidateJobStatus(candidateinterview.CandidateJobStatus(status)).
		SetLocation(input.Location).
		SetMeetingLink(input.MeetingLink)
	return create.Save(ctx)
}

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterview(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, record).SetTitle(strings.TrimSpace(input.Title)).
		SetDescription(input.Description).SetCandidateJobID(uuid.MustParse(input.CandidateJobID)).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt).
		AddInterviewerEdgeIDs(newMemberIds...).RemoveInterviewerEdgeIDs(removeMemberIds...).
		SetLocation(input.Location).
		SetMeetingLink(input.MeetingLink)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterviewSchedule(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewScheduleInput, newMemberIds, removeMemberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetInterviewDate(input.InterviewDate).
		SetStartFrom(input.StartFrom).
		SetEndAt(input.EndAt)
	if len(newMemberIds) > 0 && len(removeMemberIds) > 0 {
		update.AddInterviewerEdgeIDs(newMemberIds...).RemoveInterviewerEdgeIDs(removeMemberIds...)
	}
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *candidateInterviewRepoImpl) DeleteCandidateInterview(ctx context.Context, record *ent.CandidateInterview, memberIds []uuid.UUID) (*ent.CandidateInterview, error) {
	update := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveInterviewerEdgeIDs(memberIds...)
	return update.Save(ctx)
}

func (rps candidateInterviewRepoImpl) CreateBulkCandidateInterview(ctx context.Context, candidateJobs []*ent.CandidateJob, memberIds []uuid.UUID, input ent.NewCandidateInterview4CalendarInput) ([]*ent.CandidateInterview, error) {
	var createBulk []*ent.CandidateInterviewCreate
	var createInterviewers []*ent.CandidateInterviewerCreate
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	createdById := payload.UserID
	for _, record := range candidateJobs {
		createBulk = append(createBulk, rps.BuildCreate().SetTitle(input.Title).
			SetInterviewDate(input.InterviewDate).
			SetStartFrom(input.StartFrom).
			SetDescription(input.Description).
			SetEndAt(input.EndAt).
			SetCandidateJobID(record.ID).
			SetCreatedBy(createdById).
			SetID(uuid.New()).
			SetCandidateJobStatus(candidateinterview.CandidateJobStatus(record.Status.String())).
			SetLocation(input.Location).
			SetMeetingLink(input.MeetingLink))
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

func (rps *candidateInterviewRepoImpl) UpdateCandidateInterviewStatus(ctx context.Context, record *ent.CandidateInterview, input ent.UpdateCandidateInterviewStatusInput) (*ent.CandidateInterview, error) {
	return rps.BuildUpdateOne(ctx, record).SetStatus(candidateinterview.Status(input.Status.String())).Save(ctx)
}

func (rps *candidateInterviewRepoImpl) UpdateBulkCandidateInterviewStatus(ctx context.Context, predicates []predicate.CandidateInterview, status candidateinterview.Status) error {
	_, err := rps.BuildUpdate().Where(predicates...).SetStatus(status).Save(ctx)
	return err
}

// query
func (rps *candidateInterviewRepoImpl) GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterview, error) {
	query := rps.BuildQuery().Where(candidateinterview.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

func (rps candidateInterviewRepoImpl) GetDataForKeyword(ctx context.Context, record *ent.CandidateInterview, candidateJobRecord *ent.CandidateJob) (models.GroupModule, error) {
	var result models.GroupModule
	candidateQuery := rps.client.Candidate.Query().Where(candidate.DeletedAtIsNil(), candidate.IDEQ(candidateJobRecord.CandidateID)).
		WithReferenceUserEdge().WithCandidateSkillEdges(
		func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).WithSkillEdge(
				func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(
						func(stq *ent.SkillTypeQuery) {
							stq.Where(skilltype.DeletedAtIsNil())
						},
					)
				},
			)
		},
	)
	hiringjobQuery := rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil(), hiringjob.IDEQ(candidateJobRecord.HiringJobID)).WithHiringJobSkillEdges(
		func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).WithSkillEdge(
				func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(
						func(stq *ent.SkillTypeQuery) {
							stq.Where(skilltype.DeletedAtIsNil())
						},
					)
				},
			)
		},
	).WithOwnerEdge()
	candidateRecord, err := candidateQuery.First(ctx)
	if err != nil {
		return result, err
	}
	hiringJobRecord, err := hiringjobQuery.First(ctx)
	if err != nil {
		return result, err
	}
	hiringTeamRecord, err := rps.client.HiringTeam.Query().Where(hiringteam.DeletedAtIsNil(), hiringteam.IDEQ(hiringJobRecord.HiringTeamID)).WithUserEdges().First(ctx)
	if err != nil {
		return result, nil
	}
	return models.GroupModule{
		Candidate:    candidateRecord,
		HiringJob:    hiringJobRecord,
		HiringTeam:   hiringTeamRecord,
		CandidateJob: candidateJobRecord,
		Interview:    record,
	}, nil
}

func (rps candidateInterviewRepoImpl) BuildStatusCountByCdJobID(ctx context.Context, candidateJobIDs []uuid.UUID) ([]models.CdInterviewCountByStatus, error) {
	result := make([]models.CdInterviewCountByStatus, 0)
	err := rps.BuildBaseQuery().Select(candidateinterview.FieldStatus).Where(candidateinterview.CandidateJobIDIn(candidateJobIDs...)).GroupBy(candidateinterview.FieldStatus).Aggregate(ent.Count()).Scan(ctx, &result)
	return result, err
}

// third party
func (rps *candidateInterviewRepoImpl) CallbackInterviewSchedule(ctx context.Context, input models.MessageOutput) error {
	candidateInterview, err := rps.GetCandidateInterview(ctx, uuid.MustParse(input.ID))
	if err != nil {
		return err
	}
	_, err = candidateInterview.Update().SetStatus(candidateinterview.StatusInterviewing).Save(ctx)
	return err
}

// common function
func (rps *candidateInterviewRepoImpl) ValidateInput(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate) (string, error, error) {
	record, err := rps.client.CandidateJob.Query().Where(candidatejob.IDEQ(input.CandidateJobId), candidatejob.DeletedAtIsNil()).
		WithHiringJobEdge(func(query *ent.HiringJobQuery) {
			query.Where(hiringjob.DeletedAtIsNil(), hiringjob.StatusEQ(hiringjob.StatusOpened))
		}).
		WithCandidateEdge().
		First(ctx)
	if err != nil {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_job_not_found"), nil
	}
	if record.Edges.CandidateEdge.IsBlacklist {
		return "", fmt.Errorf("model.candidate_interviews.validation.candidate_is_blacklist"), nil
	}
	if record.Edges.HiringJobEdge == nil {
		return "", fmt.Errorf("model.candidate_interviews.validation.job_close"), nil
	}
	if record.Status != candidatejob.StatusInterviewing {
		return "", fmt.Errorf("model.candidate_interviews.validation.invalid_candidate_job_status"), nil
	}
	stringError, err := rps.ValidateSchedule(ctx, candidateInterviewId, input, record.CandidateID)
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
	candidateJobs, err := rps.client.CandidateJob.Query().
		Where(
			candidatejob.CandidateIDIn(candidateIds...), candidatejob.DeletedAtIsNil(),
			candidatejob.HiringJobIDEQ(uuid.MustParse(input.JobID)), candidatejob.StatusEQ(candidatejob.StatusInterviewing),
		).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	if len(candidateJobs) < len(candidateIds) {
		return nil, fmt.Errorf("model.candidate_interviews.validation.candidate_job_not_found"), nil
	}
	for _, record := range candidateJobs {
		if record.Status != candidatejob.StatusInterviewing {
			return nil, fmt.Errorf("model.candidate_interviews.validation.invalid_candidate_job_status"), nil
		}
		var inputValidate models.CandidateInterviewInputValidate
		jsonString, _ := json.Marshal(input)
		_ = json.Unmarshal(jsonString, &inputValidate)
		stringError, err := rps.ValidateSchedule(ctx, uuid.Nil, inputValidate, record.CandidateID)
		if err != nil || stringError != nil {
			return nil, stringError, err
		}
	}
	return candidateJobs, nil, nil
}

func (rps *candidateInterviewRepoImpl) ValidateSchedule(ctx context.Context, candidateInterviewId uuid.UUID, input models.CandidateInterviewInputValidate, candidateID uuid.UUID) (error, error) {
	currentDate := time.Now().UTC()
	currentDate = time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC).Add(-(time.Hour * 12))
	if currentDate.After(input.InterviewDate) {
		return fmt.Errorf("model.candidate_interviews.validation.interview_date_is_past"), nil
	}
	// timeDifference := interviewDate.Sub(currentDate).Minutes()
	// localTimeZone := 0
	// if timeDifference < 60*12 {
	// 	localTimeZone = ((60 * 12) - int(timeDifference)) / 60
	// } else {
	// 	localTimeZone = (int(timeDifference) - (60 * 12)) / 60
	// }
	utcTime := time.Now().UTC()
	currentTime := time.Date(utcTime.Year(), utcTime.Month(), utcTime.Day(), utcTime.Hour(), utcTime.Minute(), 0, 0, time.UTC)
	if currentTime.After(input.StartFrom) {
		return fmt.Errorf("model.candidate_interviews.validation.start_from_is_past"), nil
	}
	if input.StartFrom.After(input.EndAt) {
		return fmt.Errorf("model.candidate_interviews.validation.start_from_end_at_invalid"), nil
	}
	interviewPredicates := []predicate.CandidateInterview{
		candidateinterview.DeletedAtIsNil(),
		candidateinterview.InterviewDateEQ(input.InterviewDate),
		candidateinterview.Or(
			candidateinterview.And(candidateinterview.StartFromLTE(input.StartFrom), candidateinterview.EndAtGTE(input.EndAt)), // outside
			candidateinterview.And(candidateinterview.StartFromGTE(input.StartFrom), candidateinterview.EndAtLTE(input.EndAt)), // inside
			candidateinterview.StartFromIn(input.StartFrom, input.EndAt),
			candidateinterview.EndAtIn(input.StartFrom, input.EndAt),
		),
	}
	if candidateInterviewId != uuid.Nil {
		interviewPredicates = append(interviewPredicates, candidateinterview.IDNEQ(candidateInterviewId))
	}
	// Valid candidate schedule
	otherCdJobIDs, err := rps.client.CandidateJob.Query().
		Where(candidatejob.DeletedAtIsNil(), candidatejob.CandidateID(candidateID), candidatejob.StatusEQ(candidatejob.StatusInterviewing)).
		IDs(ctx)
	if err != nil {
		return nil, err
	}
	busyCandidatePredicates := interviewPredicates
	busyCandidatePredicates = append(busyCandidatePredicates, candidateinterview.CandidateJobIDIn(otherCdJobIDs...))
	exist, err := rps.client.CandidateInterview.Query().Where(busyCandidatePredicates...).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return fmt.Errorf("model.candidate_interviews.validation.candidate_schedule_exist"), nil
	}
	// Valid interviewer schedule
	busyInterviewerNames := make([]string, 0)
	beInterviewerPermissionID, err := rps.client.Permission.Query().Where(permission.OperationNameEQ(models.BeInterviewer)).FirstID(ctx)
	if err != nil {
		return nil, err
	}
	for _, idStr := range input.Interviewer {
		userRec, err := rps.client.User.Query().Where(user.ID(uuid.MustParse(idStr))).
			WithHiringTeamEdges(func(query *ent.HiringTeamQuery) { query.Where(hiringteam.DeletedAtIsNil()) }).
			WithLeaderRecEdge(func(query *ent.RecTeamQuery) { query.Where(recteam.DeletedAtIsNil()) }).
			WithUserPermissionEdges(func(query *ent.EntityPermissionQuery) {
				query.Where(entitypermission.PermissionID(beInterviewerPermissionID))
			}).
			WithInterviewEdges(func(query *ent.CandidateInterviewQuery) { query.Where(interviewPredicates...) }).
			First(ctx)
		if err != nil {
			return nil, err
		}
		if len(userRec.Edges.HiringTeamEdges) == 0 && userRec.Edges.LeaderRecEdge != nil && len(userRec.Edges.UserPermissionEdges) == 0 {
			return fmt.Errorf("model.candidate_interviews.validation.interviewer_not_have_permission"), nil
		}
		if len(userRec.Edges.InterviewEdges) > 0 {
			busyInterviewerNames = append(busyInterviewerNames, userRec.Name)
		}
	}
	if len(busyInterviewerNames) > 0 {
		return fmt.Errorf("The following interviewers have another interview at this time and are not available: %s", strings.Join(busyInterviewerNames, ", ")), nil
	}
	return nil, nil
}

func (rps *candidateInterviewRepoImpl) ValidCandidateInterviewStatus(record *ent.CandidateInterview, status ent.CandidateInterviewStatusEditable) error {
	if ent.CandidateInterviewStatusEditable.IsValid(ent.CandidateInterviewStatusEditable(record.Status)) {
		return fmt.Errorf("model.candidate_interviews.validation.invalid_editable")
	}
	if status == ent.CandidateInterviewStatusEditableDone && record.Status == candidateinterview.StatusInvitedToInterview {
		return fmt.Errorf("model.candidate_interviews.validation.invalid_input")
	}
	return nil
}

// Path: repository/candidate_interviewer.repository.go
