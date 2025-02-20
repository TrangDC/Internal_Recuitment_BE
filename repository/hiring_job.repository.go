package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobstep"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/jobposition"
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/ent/user"
	"trec/internal/util"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type HiringJobRepository interface {
	// mutation
	CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput) (*ent.HiringJob, error)
	UpdateHiringJob(ctx context.Context, record *ent.HiringJob, input *ent.UpdateHiringJobInput) (*ent.HiringJob, error)
	UpdateHiringJobStatus(ctx context.Context, record *ent.HiringJob, status ent.HiringJobStatus) (*ent.HiringJob, error)
	DeleteHiringJob(ctx context.Context, record *ent.HiringJob) error
	DeleteRelationHiringJob(ctx context.Context, recordId uuid.UUID) error
	// query
	GetHiringJob(ctx context.Context, hiringJobId uuid.UUID) (*ent.HiringJob, error)
	BuildQuery() *ent.HiringJobQuery
	BuildBaseQuery() *ent.HiringJobQuery
	BuildCount(ctx context.Context, query *ent.HiringJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.HiringJobQuery) ([]*ent.HiringJob, error)
	BuildGetOne(ctx context.Context, query *ent.HiringJobQuery) (*ent.HiringJob, error)
	// common function
	ValidName(ctx context.Context, hiringJobId uuid.UUID, name string, hiringTeamID string) (error, error)
	ValidPriority(ctx context.Context, hiringJobId uuid.UUID, hiringTeamID uuid.UUID, priority int) (error, error)
	ValidStatus(recordStatus hiringjob.Status, inputStatus ent.HiringJobStatus, cdJobs []*ent.CandidateJob) error
	ValidStatusWhenUpdate(record *ent.HiringJob, input *ent.UpdateHiringJobInput, recTeamChange, hiringTeamChange bool) (error, error)
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
	return rps.client.HiringJob.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
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
	return rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil()).WithCandidateJobEdges(
		func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.DeletedAtIsNil())
		},
	).WithOwnerEdge().WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	).WithHiringJobSkillEdges(
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
	).WithJobPositionEdge(
		func(query *ent.JobPositionQuery) {
			query.Where(jobposition.DeletedAtIsNil())
		},
	).WithApprovalSteps(
		func(query *ent.HiringJobStepQuery) {
			query.WithApprovalUser().Order(ent.Asc(hiringjobstep.FieldOrderID))
		},
	).WithRecTeamEdge().WithRecInChargeEdge()
}

func (rps *hiringJobRepoImpl) BuildGetOne(ctx context.Context, query *ent.HiringJobQuery) (*ent.HiringJob, error) {
	return query.First(ctx)
}

func (rps hiringJobRepoImpl) BuildBaseQuery() *ent.HiringJobQuery {
	return rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil())
}

func (rps *hiringJobRepoImpl) BuildGet(ctx context.Context, query *ent.HiringJobQuery) (*ent.HiringJob, error) {
	return query.First(ctx)
}

func (rps *hiringJobRepoImpl) BuildList(ctx context.Context, query *ent.HiringJobQuery) ([]*ent.HiringJob, error) {
	return query.All(ctx)
}

func (rps *hiringJobRepoImpl) BuildCount(ctx context.Context, query *ent.HiringJobQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *hiringJobRepoImpl) BuildExist(ctx context.Context, query *ent.HiringJobQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *hiringJobRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.HiringJob) *ent.HiringJobUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
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
		SetHiringTeamID(uuid.MustParse(input.HiringTeamID)).
		SetCreatedBy(uuid.MustParse(input.CreatedBy)).
		SetPriority(input.Priority).
		SetJobPositionID(uuid.MustParse(input.JobPositionID)).
		SetRecTeamID(uuid.MustParse(input.RecTeamID)).
		SetStatus(hiringjob.StatusPendingApprovals).
		SetLevel(hiringjob.Level(input.Level)).
		SetNote(input.Note).
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
		SetCreatedBy(uuid.MustParse(input.CreatedBy)).
		SetHiringTeamID(uuid.MustParse(input.HiringTeamID)).
		SetRecTeamEdgeID(uuid.MustParse(input.RecTeamID)).
		SetRecInChargeID(uuid.MustParse(input.RecInChargeID)).
		SetPriority(input.Priority).
		SetJobPositionID(uuid.MustParse(input.JobPositionID)).
		SetLevel(hiringjob.Level(input.Level)).
		SetNote(input.Note).
		Save(ctx)
}

func (rps *hiringJobRepoImpl) UpdateHiringJobStatus(ctx context.Context, record *ent.HiringJob, status ent.HiringJobStatus) (*ent.HiringJob, error) {
	update := rps.BuildUpdateOne(ctx, record).SetStatus(hiringjob.Status(status))
	switch status {
	case ent.HiringJobStatusOpened:
		update.SetOpenedAt(time.Now().UTC())
	case ent.HiringJobStatusClosed:
		update.SetClosedAt(time.Now().UTC())
	}
	return update.Save(ctx)
}

func (rps *hiringJobRepoImpl) DeleteHiringJob(ctx context.Context, record *ent.HiringJob) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).Save(ctx)
	return err
}

func (rps *hiringJobRepoImpl) DeleteRelationHiringJob(ctx context.Context, recordId uuid.UUID) error {
	_, err := rps.client.EntitySkill.Delete().Where(entityskill.EntityIDEQ(recordId)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJob.Update().Where(candidatejob.HiringJobID(recordId)).SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).ClearCandidateJobStep().Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobStep.Delete().Where(candidatejobstep.HasCandidateJobEdgeWith(candidatejob.HiringJobID(recordId))).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateInterview.Update().Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.HiringJobID(recordId))).SetUpdatedAt(time.Now().UTC()).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

// query
func (rps *hiringJobRepoImpl) GetHiringJob(ctx context.Context, hiringJobId uuid.UUID) (*ent.HiringJob, error) {
	return rps.BuildQuery().Where(hiringjob.IDEQ(hiringJobId)).First(ctx)
}

// common function
func (rps *hiringJobRepoImpl) ValidName(ctx context.Context, hiringJobId uuid.UUID, name string, hiringTeamId string) (error, error) {
	query := rps.BuildQuery().Where(hiringjob.NameEqualFold(strings.TrimSpace(name)), hiringjob.HiringTeamID(uuid.MustParse(hiringTeamId)))
	if hiringJobId != uuid.Nil {
		query = query.Where(hiringjob.IDNEQ(hiringJobId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.hiring_jobs.validation.name_exist"), nil
	}
	return nil, nil
}

func (rps *hiringJobRepoImpl) ValidPriority(ctx context.Context, hiringJobId uuid.UUID, hiringTeamID uuid.UUID, priority int) (error, error) {
	query := rps.BuildQuery().Where(
		hiringjob.PriorityEQ(priority),
		hiringjob.HiringTeamID(hiringTeamID),
		hiringjob.StatusIn(hiringjob.StatusPendingApprovals, hiringjob.StatusOpened),
	)
	if hiringJobId != uuid.Nil {
		query = query.Where(hiringjob.IDNEQ(hiringJobId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		switch priority {
		case 1:
			return fmt.Errorf("model.hiring_jobs.validation.priority_ugent_exist"), nil
		case 2:
			return fmt.Errorf("model.hiring_jobs.validation.priority_high_exist"), nil
		}
	}
	return nil, nil
}

func (rps *hiringJobRepoImpl) ValidStatus(recordStatus hiringjob.Status, inputStatus ent.HiringJobStatus, cdJobs []*ent.CandidateJob) error {
	switch recordStatus {
	case hiringjob.StatusPendingApprovals:
		if inputStatus != ent.HiringJobStatusCancelled {
			return errors.New("model.hiring_jobs.validation.invalid_status_update")
		}
	case hiringjob.StatusOpened:
		switch inputStatus {
		case ent.HiringJobStatusClosed:
			cdJobHiredExists := false
			for _, cdJob := range cdJobs {
				switch cdJob.Status {
				case candidatejob.StatusApplied, candidatejob.StatusInterviewing, candidatejob.StatusOffering:
					return errors.New("model.hiring_jobs.validation.close_processing_job")
				case candidatejob.StatusHired:
					cdJobHiredExists = true
				}
			}
			if !cdJobHiredExists {
				return errors.New("model.hiring_jobs.validation.close_not_hired_job")
			}
		case ent.HiringJobStatusCancelled:
			if lo.SomeBy(cdJobs, func(item *ent.CandidateJob) bool {
				return !ent.CandidateJobStatusFailed.IsValid(ent.CandidateJobStatusFailed(item.Status))
			}) {
				return errors.New("model.hiring_jobs.validation.cancel_not_failed_job")
			}
		}
	case hiringjob.StatusClosed:
		if inputStatus != ent.HiringJobStatusOpened {
			return errors.New("model.hiring_jobs.validation.invalid_status_update")
		}
	case hiringjob.StatusCancelled:
		return errors.New("model.hiring_jobs.validation.invalid_status_update")
	}
	return nil
}

func (r *hiringJobRepoImpl) ValidStatusWhenUpdate(record *ent.HiringJob, input *ent.UpdateHiringJobInput, recTeamChange, hiringTeamChange bool) (error, error) {
	switch record.Status {
	case hiringjob.StatusPendingApprovals:
		if hiringTeamChange {
			approvalSteps := lo.Filter(record.Edges.ApprovalSteps, func(item *ent.HiringJobStep, index int) bool {
				return item.Status == hiringjobstep.StatusAccepted
			})
			if len(approvalSteps) > 0 {
				return nil, errors.New("model.hiring_jobs.validation.job_already_approving")
			}
		}
	case hiringjob.StatusOpened:
		if recTeamChange || hiringTeamChange {
			return nil, errors.New("model.hiring_jobs.validation.invalid_status_update_team")
		}
	case hiringjob.StatusClosed, hiringjob.StatusCancelled:
		return nil, errors.New("model.hiring_jobs.validation.invalid_status")
	}
	return nil, nil
}

// Path: repository/hiring_job.repository.go
