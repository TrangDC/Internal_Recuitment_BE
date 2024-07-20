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
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/ent/team"
	"trec/middleware"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateJobRepository interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJob, error)
	DeleteCandidateJob(ctx context.Context, record *ent.CandidateJob) error
	UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobStatus) (*ent.CandidateJob, error)
	UpsetCandidateAttachment(ctx context.Context, record *ent.CandidateJob) (*ent.CandidateJob, error)
	DeleteRelationCandidateJob(ctx context.Context, recordId uuid.UUID) error
	// query
	GetCandidateJob(ctx context.Context, candidateId uuid.UUID) (*ent.CandidateJob, error)
	BuildQuery() *ent.CandidateJobQuery
	BuildCount(ctx context.Context, query *ent.CandidateJobQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateJobQuery) ([]*ent.CandidateJob, error)
	GetOneCandidateJob(ctx context.Context, query *ent.CandidateJobQuery) (*ent.CandidateJob, error)
	// common function
	ValidStatus(ctx context.Context, candidateId uuid.UUID, candidateJobId uuid.UUID, status ent.CandidateJobStatus, onboardDate *time.Time, offerExpDate *time.Time) (error, error)
	ValidUpsetByCandidateIsBlacklist(ctx context.Context, candidateId uuid.UUID) (error, error)
	BuildBaseQuery() *ent.CandidateJobQuery
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
			query.WithTeamEdge().WithOwnerEdge()
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
func (rps candidateJobRepoImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJob, error) {
	_, err := rps.client.Candidate.Update().Where(candidate.IDEQ(uuid.MustParse(input.CandidateID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = rps.client.HiringJob.Update().Where(hiringjob.IDEQ(uuid.MustParse(input.HiringJobID))).SetUpdatedAt(time.Now().UTC()).SetLastApplyDate(time.Now().UTC()).Save(ctx)
	if err != nil {
		return nil, err
	}
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	createdById := payload.UserID
	create := rps.BuildCreate().
		SetHiringJobID(uuid.MustParse(input.HiringJobID)).
		SetUpdatedAt(time.Now().UTC()).
		SetCandidateID(uuid.MustParse(input.CandidateID)).
		SetStatus(candidatejob.Status(input.Status)).
		SetCreatedBy(createdById)
	if input.Status == ent.CandidateJobStatusOffering {
		create.
			SetOnboardDate(*input.OnboardDate).
			SetOfferExpirationDate(*input.OfferExpirationDate)
	}
	return create.Save(ctx)
}

func (rps candidateJobRepoImpl) UpdateCandidateJobStatus(ctx context.Context, record *ent.CandidateJob, input ent.UpdateCandidateJobStatus) (*ent.CandidateJob, error) {
	update := rps.BuildUpdateOne(ctx, record).SetStatus(candidatejob.Status(input.Status.String()))
	if input.Status == ent.CandidateJobStatusOffering {
		update.
			SetOnboardDate(*input.OnboardDate).
			SetOfferExpirationDate(*input.OfferExpirationDate)
	}
	if ent.CandidateJobStatusFailed.IsValid(ent.CandidateJobStatusFailed(input.Status)) {
		if input.FailedReason == nil && len(input.FailedReason) == 0 {
			return nil, fmt.Errorf("model.candidate_job.validation.failed_reason_required")
		}
		reason := lo.Map(input.FailedReason, func(s ent.CandidateJobFailedReason, index int) string {
			return s.String()
		})
		update.SetFailedReason(reason)
	} else {
		update.ClearFailedReason()
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
	hiringjobQuery := rps.client.HiringJob.Query().Where(hiringjob.DeletedAtIsNil(), hiringjob.IDEQ(record.HiringJobID)).WithHiringJobSkillEdges(
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
	teamRecord, err := rps.client.Team.Query().Where(team.DeletedAtIsNil(), team.IDEQ(hiringJobRecord.TeamID)).WithUserEdges().First(ctx)
	if err != nil {
		return result, nil
	}
	return models.GroupModule{
		Candidate:    candidateRecord,
		HiringJob:    hiringJobRecord,
		Team:         teamRecord,
		CandidateJob: record,
	}, nil
}

// common function
func (rps candidateJobRepoImpl) ValidStatus(ctx context.Context, candidateId uuid.UUID, candidateJobId uuid.UUID, status ent.CandidateJobStatus, onboardDate *time.Time, offerExpDate *time.Time) (error, error) {
	openStatus := lo.Map(ent.AllCandidateJobStatusOpen, func(s ent.CandidateJobStatusOpen, index int) candidatejob.Status {
		return candidatejob.Status(s)
	})
	openStatus = append(openStatus, candidatejob.StatusHired)
	if !lo.Contains(openStatus, candidatejob.Status(status)) {
		return nil, nil
	}
	if status == ent.CandidateJobStatusOffering {
		currentTime := time.Now().UTC()
		currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
		if onboardDate == nil {
			return fmt.Errorf("model.candidate_job.validation.onboard_date_required"), nil
		}
		if offerExpDate == nil {
			return fmt.Errorf("model.candidate_job.validation.offer_exp_date_required"), nil
		}
		if onboardDate.Before(currentTime) {
			return fmt.Errorf("model.candidate_job.validation.invalid_onboard_date"), nil
		}
		if offerExpDate.Before(currentTime) {
			return fmt.Errorf("model.candidate_job.validation.invalid_offer_exp_date"), nil
		}
		if onboardDate.Compare(*offerExpDate) <= 0 {
			return fmt.Errorf("model.candidate_job.validation.onboard_before_offer_exp"), nil
		}
	}
	query := rps.BuildQuery().Where(candidatejob.CandidateIDEQ(candidateId))
	if candidateJobId != uuid.Nil {
		query.Where(candidatejob.IDNEQ(candidateJobId))
	}
	query = query.Where(candidatejob.StatusIn(openStatus...))
	isExist, _ := rps.BuildExist(ctx, query)
	if isExist {
		return fmt.Errorf("model.candidate_job.validation.candidate_job_status_exist"), nil
	}
	return nil, nil
}

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
