package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatejobstep"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/middleware"
	"trec/models"

	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateJobRepository interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, failedReason []string) (*ent.CandidateJob, error)
	DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error
	UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobStatus, failedReason []string) (*ent.CandidateJob, error)
	UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error)
	DeleteRelationCandidateJob(ctx context.Context, recordId uuid.UUID) error
	// query
	GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error)
	BuildQuery() *ent.CandidateJobQuery
	BuildBaseQuery() *ent.CandidateJobQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error)
	BuildIDList(ctx context.Context, query *ent.CandidateJobQuery) ([]uuid.UUID, error)
	GetOneCandidateJob(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error)
	// common function
	ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) (error, error)
	ValidInput(ctx context.Context, input models.CandidateJobValidInput) ([]string, error, error)
	ValidStatus(oldStatus candidatejob.Status, newStatus ent.CandidateJobStatus) error
	GetDataForKeyword(ctx context.Context, record *ent.CandidateJob) (models.GroupModule, error)
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
	createTime := time.Now().UTC()
	return rps.client.CandidateJob.Create().SetUpdatedAt(createTime).SetCreatedAt(createTime)
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
			query.WithHiringTeamEdge().WithOwnerEdge()
		},
	).WithCreatedByEdge().WithCandidateJobStep(
		func(query *ent.CandidateJobStepQuery) {
			query.Order(ent.Asc(candidatejobstep.FieldCreatedAt))
		},
	).WithCandidateJobInterview(
		func(query *ent.CandidateInterviewQuery) {
			query.Where(candidateinterview.DeletedAtIsNil())
		},
	)
}

func (rps candidateJobRepoImpl) BuildBaseQuery() *ent.CandidateJobQuery {
	return rps.client.CandidateJob.Query().Where(candidatejob.DeletedAtIsNil())
}

func (rps candidateJobRepoImpl) GetOneCandidateJob(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error) {
	return query.First(ctx)
}

func (rps candidateJobRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error) {
	return query.First(ctx)
}

func (rps candidateJobRepoImpl) BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error) {
	return query.All(ctx)
}

func (rps candidateJobRepoImpl) BuildIDList(ctx context.Context, query *ent.CandidateJobQuery) ([]uuid.UUID, error) {
	return query.IDs(ctx)
}

func (rps candidateJobRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error) {
	return query.Count(ctx)
}

func (rps candidateJobRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateJobQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps candidateJobRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.CandidateJob) *ent.CandidateJobUpdateOne {
	return record.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateJobRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateJobUpdateOne) (*ent.CandidateJob, error) {
	return update.Save(ctx)
}

// mutation
func (rps candidateJobRepoImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, failedReason []string) (*ent.CandidateJob, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	createdById := payload.UserID
	create := rps.BuildCreate().
		SetHiringJobID(uuid.MustParse(input.HiringJobID)).
		SetUpdatedAt(time.Now().UTC()).
		SetCandidateID(uuid.MustParse(input.CandidateID)).
		SetStatus(candidatejob.Status(input.Status)).
		SetCreatedBy(createdById)
	if input.Status == ent.CandidateJobStatusOpenOffering {
		create.
			SetOnboardDate(*input.OnboardDate).
			SetOfferExpirationDate(*input.OfferExpirationDate).
			SetLevel(candidatejob.Level(*input.Level))
	}
	_, err := rps.client.Candidate.Update().Where(candidate.IDEQ(uuid.MustParse(input.CandidateID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = rps.client.HiringJob.Update().Where(hiringjob.IDEQ(uuid.MustParse(input.HiringJobID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return create.Save(ctx)
}

func (rps candidateJobRepoImpl) UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobStatus, failedReason []string) (*ent.CandidateJob, error) {
	update := rps.BuildUpdateOne(ctx, record).SetStatus(candidatejob.Status(input.Status.String()))
	switch input.Status {
	case ent.CandidateJobStatusKiv, ent.CandidateJobStatusOfferLost:
		update.SetFailedReason(failedReason)
	case ent.CandidateJobStatusOffering:
		update.
			SetOnboardDate(*input.OnboardDate).
			SetOfferExpirationDate(*input.OfferExpirationDate).
			SetLevel(candidatejob.Level(input.Level.String()))
	}
	return update.Save(ctx)
}

// fix it, it not remove attachment
func (rps candidateJobRepoImpl) UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error) {
	return rps.BuildUpdateOne(ctx, record).Save(ctx)
}

func (rps candidateJobRepoImpl) DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

func (rps candidateJobRepoImpl) DeleteRelationCandidateJob(ctx context.Context, recordId uuid.UUID) error {
	_, err := rps.client.CandidateJobStep.Delete().Where(candidatejobstep.CandidateJobID(recordId)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateInterview.Update().Where(candidateinterview.CandidateJobID(recordId)).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.Attachment.Update().Where(attachment.RelationID(recordId)).SetDeletedAt(time.Now().UTC()).Save(ctx)
	if err != nil {
		return err
	}
	_, err = rps.client.CandidateJobFeedback.Update().Where(candidatejobfeedback.CandidateJobID(recordId)).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

// query
func (rps candidateJobRepoImpl) GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error) {
	return rps.BuildQuery().Where(candidatejob.IDEQ(candidateId)).First(ctx)
}

func (rps candidateJobRepoImpl) GetDataForKeyword(ctx context.Context, record *ent.CandidateJob) (models.GroupModule, error) {
	var result models.GroupModule
	candidateQuery := rps.client.Candidate.Query().Where(candidate.DeletedAtIsNil(), candidate.IDEQ(record.CandidateID)).
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
	hiringJobQuery := rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil(), hiringjob.IDEQ(record.HiringJobID)).WithHiringJobSkillEdges(
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
	hiringJobRecord, err := hiringJobQuery.First(ctx)
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
		CandidateJob: record,
	}, nil
}

// common function
func (rps candidateJobRepoImpl) ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) (error, error) {
	candidateRecord, err := rps.client.Candidate.Query().Where(candidate.IDEQ(candidateId)).First(ctx)
	if err != nil {
		return nil, err
	}
	if candidateRecord.IsBlacklist {
		return fmt.Errorf("model.candidate_job.validation.candidate_is_blacklist"), nil
	}
	return nil, nil
}

func (rps candidateJobRepoImpl) ValidInput(ctx context.Context, input models.CandidateJobValidInput) ([]string, error, error) {
	var failedReason []string
	openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(s)
	})
	openStatus = append(openStatus, candidatejob.StatusHired)
	candidateRecord, err := rps.client.Candidate.Query().Where(candidate.IDEQ(input.CandidateId)).First(ctx)
	if err != nil {
		return failedReason, nil, err
	}
	if candidateRecord.IsBlacklist {
		return failedReason, fmt.Errorf("model.candidate_job.validation.candidate_is_blacklist"), nil
	}
	switch input.Status {
	case ent.CandidateJobStatusApplied, ent.CandidateJobStatusInterviewing, ent.CandidateJobStatusOffering:
		if input.Status == ent.CandidateJobStatusOffering {
			utc, _ := time.LoadLocation(carbon.UTC)
			currentTime := carbon.Now().StartOfDay().SetLocation(utc)
			if input.OnboardDate == nil {
				return failedReason, fmt.Errorf("model.candidate_job.validation.onboard_date_required"), nil
			}
			if input.OfferExpDate == nil {
				return failedReason, fmt.Errorf("model.candidate_job.validation.offer_exp_date_required"), nil
			}
			onboardDate := carbon.Parse(input.OnboardDate.String())
			offerExpDate := carbon.Parse(input.OfferExpDate.String())
			if currentTime.Gte(onboardDate) {
				return failedReason, fmt.Errorf("model.candidate_job.validation.invalid_onboard_date"), nil
			}
			if currentTime.Gte(offerExpDate) {
				return failedReason, fmt.Errorf("model.candidate_job.validation.invalid_offer_exp_date"), nil
			}
			if onboardDate.Lte(offerExpDate) {
				return failedReason, fmt.Errorf("model.candidate_job.validation.onboard_before_offer_exp"), nil
			}
		}
		query := rps.BuildQuery().Where(candidatejob.CandidateIDEQ(input.CandidateId))
		if input.CandidateJobId != uuid.Nil {
			query.Where(candidatejob.IDNEQ(input.CandidateJobId))
		}
		query = query.Where(candidatejob.StatusIn(openStatus...))
		isExist, _ := rps.BuildExist(ctx, query)
		if isExist {
			return failedReason, fmt.Errorf("model.candidate_job.validation.candidate_job_status_exist"), nil
		}
	case ent.CandidateJobStatusOfferLost, ent.CandidateJobStatusKiv:
		if len(input.FailedReason) == 0 {
			return failedReason, fmt.Errorf("model.candidate_job.validation.failed_reason_required"), nil
		}
		failedReason = lo.Map(input.FailedReason, func(s ent.CandidateJobFailedReason, index int) string {
			return s.String()
		})
	}
	return failedReason, nil, nil
}

func (rps candidateJobRepoImpl) ValidStatus(oldStatus candidatejob.Status, newStatus ent.CandidateJobStatus) error {
	isErrorStatus := false
	entOldStatus := ent.CandidateJobStatus(oldStatus)
	switch newStatus {
	case ent.CandidateJobStatusInterviewing:
		if entOldStatus != ent.CandidateJobStatusApplied {
			isErrorStatus = true
		}
	case ent.CandidateJobStatusOffering, ent.CandidateJobStatusKiv:
		switch entOldStatus {
		case ent.CandidateJobStatusApplied, ent.CandidateJobStatusInterviewing:
			isErrorStatus = false
		default:
			isErrorStatus = true
		}
	case ent.CandidateJobStatusHired, ent.CandidateJobStatusOfferLost:
		if entOldStatus != ent.CandidateJobStatusOffering {
			isErrorStatus = true
		}
	case ent.CandidateJobStatusExStaff:
		if entOldStatus != ent.CandidateJobStatusHired {
			isErrorStatus = true
		}
	}
	if isErrorStatus {
		return fmt.Errorf("model.candidate_jobs.validation.invalid_status")
	}
	return nil
}
