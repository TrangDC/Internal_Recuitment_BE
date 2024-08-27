package service

import (
	"context"
	"trec/ent"
	"trec/ent/hiringjobstep"
	"trec/ent/hiringteamapprover"
	"trec/ent/recteam"
	"trec/ent/user"
	"trec/middleware"
	"trec/repository"

	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringJobStepService interface {
	CreateBulkHiringJobSteps(ctx context.Context, repoRegistry repository.Repository, hiringJob *ent.HiringJob) error
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
	for i, approver := range approvers {
		create := repoRegistry.HiringJobStep().BuildCreate().
			SetHiringJobID(hiringJob.ID).SetUserID(approver.ID)
		if approver.ID == loggedInUserID {
			orderID = 1
			status = hiringjobstep.StatusAccepted
			creates = []*ent.HiringJobStepCreate{create.SetOrderID(orderID).SetStatus(status)}
			if i == len(approvers)-1 {
				_, err := repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, hiringJob, ent.HiringJobStatusOpened)
				if err != nil {
					return err
				}
				break
			}
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
