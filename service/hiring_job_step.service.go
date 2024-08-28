package service

import (
	"context"
	"net/http"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/hiringteamapprover"
	"trec/ent/recteam"
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
	UpdateBulkHiringJobStepsStatus(ctx context.Context, input ent.UpdateHiringJobStepInput) error
	UpdateHiringJobStepByRecLeader(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob, newRecLeaderID uuid.UUID) error
}

type hiringJobStepImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewHiringJobStepService(repoRegistry repository.Repository, logger *zap.Logger) HiringJobStepService {
	return &hiringJobStepImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
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

func (svc *hiringJobStepImpl) UpdateBulkHiringJobStepsStatus(ctx context.Context, input ent.UpdateHiringJobStepInput) error {
	switch input.Status {
	case ent.HiringJobStepStatusEnumPending:
		return nil
	case ent.HiringJobStepStatusEnumWaiting:
		return util.WrapGQLError(ctx, "model.hiring_job_steps.validation.invalid_status_update", http.StatusBadRequest, util.ErrorFlagValidateFail)
	default:
		err := svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
			for _, jobIDStr := range input.HiringJobIds {
				err := svc.updateHiringJobStepStatus(ctx, repoRegistry, uuid.MustParse(jobIDStr), input.Status)
				if err != nil {
					svc.logger.Error(err.Error(), zap.String("job_id", jobIDStr))
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func (svc *hiringJobStepImpl) updateHiringJobStepStatus(ctx context.Context, repoRegistry repository.Repository, hiringJobID uuid.UUID, status ent.HiringJobStepStatusEnum) error {
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(
		ctx,
		svc.repoRegistry.HiringJob().BuildBaseQuery().
			Where(hiringjob.DeletedAtIsNil(), hiringjob.ID(hiringJobID)).
			WithApprovalSteps(func(query *ent.HiringJobStepQuery) {
				query.Order(ent.Asc(hiringjobstep.FieldOrderID))
			}),
	)
	if err != nil {
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
	updatedRec, err := repoRegistry.HiringJobStep().UpdateHiringJobStepStatus(ctx, currentStep, status)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	switch updatedRec.Status {
	case hiringjobstep.StatusAccepted:
		if updatedRec.OrderID == len(steps) {
			_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusOpened)
			if err != nil {
				return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
			}
		}
		_, err = repoRegistry.HiringJobStep().UpdateHiringJobStepStatus(ctx, steps[updatedRec.OrderID+1], ent.HiringJobStepStatusEnumPending)
		if err != nil {
			return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	case hiringjobstep.StatusRejected:
		_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusCancelled)
		if err != nil {
			return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	return nil
}

func (svc *hiringJobStepImpl) UpdateHiringJobStepByRecLeader(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob, oldRecLeaderID uuid.UUID) error {
	recLeaderStep, found := lo.Find(hiringJob.Edges.ApprovalSteps, func(step *ent.HiringJobStep) bool {
		return step.UserID == oldRecLeaderID && step.OrderID != 1 && step.Status != hiringjobstep.StatusAccepted
	})
	if !found || recLeaderStep == nil {
		return util.WrapGQLError(ctx, "model.hiring_job_steps.validation.invalid_rec_leader", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	_, err := repoRegistry.HiringJobStep().BuildUpdateOne(ctx, recLeaderStep).SetUserID(hiringJob.Edges.RecTeamEdge.LeaderID).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
