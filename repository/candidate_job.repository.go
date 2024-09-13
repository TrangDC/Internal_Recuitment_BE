package repository

import (
	"context"
	"errors"
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
	"trec/ent/recteam"
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
	UpdateCandidateJob(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobInput) (*ent.CandidateJob, error)
	DeleteRelationCandidateJob(ctx context.Context, recordId uuid.UUID) error
	// query
	GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error)
	BuildQuery() *ent.CandidateJobQuery
	BuildBaseQuery() *ent.CandidateJobQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error)
	BuildIDList(ctx context.Context, query *ent.CandidateJobQuery) ([]uuid.UUID, error)
	BuildGetOne(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error)
	BuildExist(ctx context.Context, query *ent.CandidateJobQuery) (bool, error)
	GetOneCandidateJob(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error)
	// common function
	ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) (error, error)
	ValidInput(ctx context.Context, input models.CandidateJobValidInput) ([]string, error, error)
	ValidStatus(oldStatus candidatejob.Status, newStatus ent.CandidateJobStatus) error
	ValidOfferingInput(onboardDatePtr, offerExpDatePtr *time.Time) error
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
			query.WithHiringTeamEdge().WithOwnerEdge().WithRecTeamEdge()
		},
	).WithCreatedByEdge().WithCandidateJobStep(
		func(query *ent.CandidateJobStepQuery) {
			query.Order(ent.Asc(candidatejobstep.FieldCreatedAt))
		},
	).WithCandidateJobInterview(
		func(query *ent.CandidateInterviewQuery) {
			query.Where(candidateinterview.DeletedAtIsNil())
		},
	).WithRecInChargeEdge()
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

func (rps candidateJobRepoImpl) BuildGetOne(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error) {
	return query.First(ctx)
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
		SetCreatedBy(createdById).
		SetRecInChargeID(uuid.MustParse(input.RecInChargeID))
	if input.Status == ent.CandidateJobStatusOpenOffering {
		if input.OnboardDate != nil {
			create.SetOnboardDate(*input.OnboardDate)
		}
		if input.OfferExpirationDate != nil {
			create.SetOfferExpirationDate(*input.OfferExpirationDate)
		}
		create.SetLevel(candidatejob.Level(*input.Level))
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
	case ent.CandidateJobStatusFailedCv, ent.CandidateJobStatusFailedInterview, ent.CandidateJobStatusOfferLost:
		update.SetFailedReason(failedReason)
	case ent.CandidateJobStatusOffering:
		update.SetLevel(candidatejob.Level(input.Level.String()))
		if input.OnboardDate != nil {
			update.SetOnboardDate(*input.OnboardDate)
		}
		if input.OfferExpirationDate != nil {
			update.SetOfferExpirationDate(*input.OfferExpirationDate)
		}
	}
	return update.Save(ctx)
}

// fix it, it not remove attachment
func (rps candidateJobRepoImpl) UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error) {
	return rps.BuildUpdateOne(ctx, record).Save(ctx)
}

func (rps candidateJobRepoImpl) UpdateCandidateJob(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobInput) (*ent.CandidateJob, error) {
	update := rps.BuildUpdateOne(ctx, record).SetRecInChargeID(uuid.MustParse(input.RecInChargeID))
	if input.OnboardDate != nil {
		update.SetOnboardDate(*input.OnboardDate)
	}
	if input.OfferExpirationDate != nil {
		update.SetOfferExpirationDate(*input.OfferExpirationDate)
	}
	return update.Save(ctx)
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
		WithReferenceUserEdge().
		WithCandidateSkillEdges(func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).
				WithSkillEdge(func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).
						WithSkillTypeEdge(func(stq *ent.SkillTypeQuery) { stq.Where(skilltype.DeletedAtIsNil()) })
				})
		})
	hiringJobQuery := rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil(), hiringjob.IDEQ(record.HiringJobID)).
		WithHiringJobSkillEdges(func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).
				WithSkillEdge(func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(func(stq *ent.SkillTypeQuery) { stq.Where(skilltype.DeletedAtIsNil()) })
				})
		}).
		WithOwnerEdge().WithRecInChargeEdge()
	candidateRecord, err := candidateQuery.First(ctx)
	if err != nil {
		return result, err
	}
	hiringJobRecord, err := hiringJobQuery.First(ctx)
	if err != nil {
		return result, err
	}
	hiringTeamRecord, err := rps.client.HiringTeam.Query().Where(hiringteam.DeletedAtIsNil(), hiringteam.IDEQ(hiringJobRecord.HiringTeamID)).
		WithUserEdges().WithApproversUsers().WithHiringMemberEdges().
		First(ctx)
	if err != nil {
		return result, nil
	}
	recTeamRecord, err := rps.client.RecTeam.Query().Where(recteam.DeletedAtIsNil(), recteam.IDEQ(hiringJobRecord.RecTeamID)).
		WithRecLeaderEdge().WithRecMemberEdges().
		First(ctx)
	if err != nil {
		return result, nil
	}
	return models.GroupModule{
		Candidate:    candidateRecord,
		HiringJob:    hiringJobRecord,
		HiringTeam:   hiringTeamRecord,
		CandidateJob: record,
		RecTeam:      recTeamRecord,
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
	candidateRecord, err := rps.client.Candidate.Query().
		Where(candidate.IDEQ(input.CandidateId)).
		WithCandidateJobEdges(func(query *ent.CandidateJobQuery) {
			query.Where(candidatejob.DeletedAtIsNil())
			if input.CandidateJobId != uuid.Nil {
				query.Where(candidatejob.IDNEQ(input.CandidateJobId))
			}
		}).
		First(ctx)
	if err != nil {
		return failedReason, nil, err
	}
	if candidateRecord.IsBlacklist {
		return failedReason, errors.New("model.candidate_job.validation.candidate_is_blacklist"), nil
	}
	for _, cdJob := range candidateRecord.Edges.CandidateJobEdges {
		if cdJob.Status == candidatejob.StatusHired {
			return failedReason, fmt.Errorf("model.candidate_job.validation.candidate_is_hired"), nil
		}
		cdJobIsProcessing := cdJob.Status == candidatejob.StatusApplied || cdJob.Status == candidatejob.StatusInterviewing || cdJob.Status == candidatejob.StatusOffering
		if cdJob.HiringJobID == input.HiringJobId && cdJobIsProcessing {
			return failedReason, errors.New("model.candidate_job.validation.same_hiring_job"), nil
		}
	}
	switch input.Status {
	case ent.CandidateJobStatusOffering:
		return failedReason, rps.ValidOfferingInput(input.OnboardDate, input.OfferExpDate), nil
	case ent.CandidateJobStatusOfferLost, ent.CandidateJobStatusFailedCv, ent.CandidateJobStatusFailedInterview:
		if len(input.FailedReason) == 0 {
			return failedReason, errors.New("model.candidate_job.validation.failed_reason_required"), nil
		}
		failedReason = lo.Map(input.FailedReason, func(s ent.CandidateJobFailedReason, _ int) string {
			return s.String()
		})
	}
	return failedReason, nil, nil
}

func (rps candidateJobRepoImpl) ValidStatus(oldStatus candidatejob.Status, newStatus ent.CandidateJobStatus) error {
	isErrorStatus := false
	entOldStatus := ent.CandidateJobStatus(oldStatus)
	switch newStatus {
	case ent.CandidateJobStatusInterviewing, ent.CandidateJobStatusFailedCv:
		if entOldStatus != ent.CandidateJobStatusApplied {
			isErrorStatus = true
		}
	case ent.CandidateJobStatusOffering:
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
	case ent.CandidateJobStatusFailedInterview:
		if entOldStatus != ent.CandidateJobStatusInterviewing {
			isErrorStatus = true
		}
	}
	if isErrorStatus {
		return fmt.Errorf("model.candidate_jobs.validation.invalid_status")
	}
	return nil
}

func (rps *candidateJobRepoImpl) ValidOfferingInput(onboardDatePtr, offerExpDatePtr *time.Time) error {
	utc, _ := time.LoadLocation(carbon.UTC)
	currentTime := carbon.Now().StartOfDay().SetLocation(utc)
	var onboardDate, offerExpDate carbon.Carbon
	if onboardDatePtr != nil {
		onboardDate = carbon.Parse(onboardDatePtr.String())
		if currentTime.Gte(onboardDate) {
			return errors.New("model.candidate_job.validation.invalid_onboard_date")
		}
	}
	if offerExpDatePtr != nil {
		offerExpDate = carbon.Parse(offerExpDatePtr.String())
		if currentTime.Gte(offerExpDate) {
			return errors.New("model.candidate_job.validation.invalid_offer_exp_date")
		}
		if !onboardDate.IsZero() && onboardDate.Lte(offerExpDate) {
			return errors.New("model.candidate_job.validation.onboard_before_offer_exp")
		}
	}
	return nil
}
