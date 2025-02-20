package service

import (
	"context"
	"net/http"
	"sync"
	"sync/atomic"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/hiringteamapprover"
	"trec/ent/recteam"
	"trec/ent/skill"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringJobStepService interface {
	CreateBulkHiringJobSteps(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob) error
	UpdateBulkHiringJobStepsStatus(ctx context.Context, input ent.UpdateHiringJobStepInput, note string) error
	UpdateHiringJobStepStatus(ctx context.Context, hiringJobID uuid.UUID, status ent.HiringJobStepStatusEnum, note string) error
	UpdateHiringJobStepByRecLeader(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob) error
}

type hiringJobStepImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
	dtoRegistry  dto.Dto
}

func NewHiringJobStepService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) HiringJobStepService {
	return &hiringJobStepImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
		dtoRegistry:  dtoRegistry,
	}
}

func (svc *hiringJobStepImpl) CreateBulkHiringJobSteps(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob) error {
	hiringTeamApprovers, err := repoRegistry.HiringTeamApprover().BuildList(
		ctx,
		repoRegistry.HiringTeamApprover().BuildQuery().
			Where(hiringteamapprover.HiringTeamID(hiringJob.HiringTeamID)).
			Order(ent.Asc(hiringteamapprover.FieldOrderID)),
	)
	if err != nil {
		return err
	}
	approvers := lo.Map(hiringTeamApprovers, func(approver *ent.HiringTeamApprover, _ int) *ent.User { return approver.Edges.UserEdge })
	recLeader, err := repoRegistry.User().GetOneUser(
		ctx,
		repoRegistry.User().BuildBaseQuery().
			Where(user.HasLeaderRecEdgeWith(recteam.DeletedAtIsNil(), recteam.ID(hiringJob.RecTeamID))),
	)
	if err != nil {
		return err
	}
	approvers = append(approvers, recLeader)
	creates := make([]*ent.HiringJobStepCreate, 0)
	loggedInUserID := ctx.Value(middleware.Payload{}).(*middleware.Payload).UserID
	orderID := 1
	status := hiringjobstep.StatusPending
	currentUserAccepted := false
	for i, approver := range approvers {
		create := repoRegistry.HiringJobStep().BuildCreate().
			SetHiringJobID(hiringJob.ID).SetUserID(approver.ID)
		if approver.ID == loggedInUserID && !currentUserAccepted {
			orderID = 1
			status = hiringjobstep.StatusAccepted
			creates = []*ent.HiringJobStepCreate{create.SetOrderID(orderID).SetStatus(status)}
			if i == len(approvers)-1 {
				_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusOpened)
				if err != nil {
					return err
				}
			}
			orderID++
			currentUserAccepted = true
			continue
		}
		if orderID == 1 {
			creates = append(creates, create.SetOrderID(orderID).SetStatus(status))
			orderID++
			continue
		}
		switch status {
		case hiringjobstep.StatusAccepted:
			status = hiringjobstep.StatusPending
		case hiringjobstep.StatusPending, hiringjobstep.StatusWaiting:
			status = hiringjobstep.StatusWaiting
		}
		creates = append(creates, create.SetOrderID(orderID).SetStatus(status))
		orderID++
	}
	_, err = repoRegistry.HiringJobStep().CreateBulkHiringJobSteps(ctx, creates)
	return err
}

func (svc *hiringJobStepImpl) UpdateBulkHiringJobStepsStatus(ctx context.Context, input ent.UpdateHiringJobStepInput, note string) error {
	switch input.Status {
	case ent.HiringJobStepStatusEnumPending:
		return nil
	case ent.HiringJobStepStatusEnumWaiting:
		return util.WrapGQLError(ctx, "model.hiring_job_steps.validation.invalid_status_update", http.StatusBadRequest, util.ErrorFlagValidateFail)
	default:
		wg := new(sync.WaitGroup)
		var hasErr int32
		for _, jobIDStr := range input.HiringJobIds {
			wg.Add(1)
			jobID := uuid.MustParse(jobIDStr)
			go func(jobID uuid.UUID) {
				defer wg.Done()
				err := svc.UpdateHiringJobStepStatus(ctx, jobID, input.Status, note)
				if err != nil {
					atomic.StoreInt32(&hasErr, 1)
				}
			}(jobID)
		}
		wg.Wait()
		if atomic.LoadInt32(&hasErr) == 1 {
			return util.WrapGQLError(
				ctx,
				"internal system error(s), pls see server logs for more details",
				http.StatusInternalServerError,
				util.ErrorFlagInternalError,
			)
		}
		return nil
	}
}

func (svc *hiringJobStepImpl) UpdateHiringJobStepStatus(ctx context.Context, hiringJobID uuid.UUID, status ent.HiringJobStepStatusEnum, note string) error {
	jobIDZapField := zap.String("job_id", hiringJobID.String())
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(
		ctx,
		svc.repoRegistry.HiringJob().BuildBaseQuery().
			Where(hiringjob.DeletedAtIsNil(), hiringjob.ID(hiringJobID)).
			WithApprovalSteps(func(query *ent.HiringJobStepQuery) {
				query.Order(ent.Asc(hiringjobstep.FieldOrderID))
			}).
			WithHiringJobSkillEdges(func(query *ent.EntitySkillQuery) {
				query.Where(entityskill.DeletedAtIsNil()).
					WithSkillEdge(func(sq *ent.SkillQuery) { sq.Where(skill.DeletedAtIsNil()) })
			}),
	)
	if err != nil {
		svc.logger.Error(err.Error(), jobIDZapField)
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	currentUserID := ctx.Value(middleware.Payload{}).(*middleware.Payload).UserID
	steps := hiringJob.Edges.ApprovalSteps
	currentStep, stepExists := lo.Find(steps, func(item *ent.HiringJobStep) bool {
		return item.UserID == currentUserID && item.Status == hiringjobstep.StatusPending
	})
	if !stepExists {
		return util.WrapGQLError(ctx, "model.hiring_job_steps.validation.invalid_user", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		updatedRec, err := repoRegistry.HiringJobStep().UpdateHiringJobStepStatus(ctx, currentStep, status)
		if err != nil {
			return err
		}
		switch updatedRec.Status {
		case hiringjobstep.StatusAccepted:
			if updatedRec.OrderID == len(steps) {
				_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusOpened)
				if err != nil {
					return err
				}
				break
			}
			_, err = repoRegistry.HiringJobStep().UpdateHiringJobStepStatus(ctx, steps[updatedRec.OrderID], ent.HiringJobStepStatusEnumPending)
			return err
		case hiringjobstep.StatusRejected:
			_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusCancelled)
			return err
		}
		return nil
	})
	if err != nil {
		svc.logger.Error(err.Error(), jobIDZapField)
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.HiringJob().GetHiringJob(ctx, hiringJobID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailUpdate(hiringJob, result)
	if err != nil {
		svc.logger.Error(err.Error(), jobIDZapField)
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, hiringJob.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), jobIDZapField)
	}
	return nil
}

func (svc *hiringJobStepImpl) UpdateHiringJobStepByRecLeader(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob) error {
	recLeaderStep := hiringJob.Edges.ApprovalSteps[len(hiringJob.Edges.ApprovalSteps)-1]
	_, err := repoRegistry.HiringJobStep().UpdateHiringJobStepByRecLeader(ctx, recLeaderStep, hiringJob.Edges.RecTeamEdge.LeaderID)
	return err
}
