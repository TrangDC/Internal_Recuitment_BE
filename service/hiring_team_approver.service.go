package service

import (
	"context"
	"trec/ent"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringTeamApproverService interface {
	HiringTeamApproverMutation(ctx context.Context, inputs []*ent.HiringTeamApproverInput, hiringTeamID uuid.UUID, currentApprover []*ent.HiringTeamApprover) error
}

type hiringTeamApproverServiceImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewHiringTeamApproverService(repoRegistry repository.Repository, logger *zap.Logger) HiringTeamApproverService {
	return &hiringTeamApproverServiceImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *hiringTeamApproverServiceImpl) HiringTeamApproverMutation(ctx context.Context, inputs []*ent.HiringTeamApproverInput, hiringTeamID uuid.UUID, currentApprovers []*ent.HiringTeamApprover) error {
	for _, input := range inputs {
		if input.ID == "" {
			// create new approver
			err := svc.repoRegistry.HiringTeamApprover().CreateHiringTeamApprover(ctx, input, hiringTeamID)
			if err != nil {
				return err
			}
			continue
		}
		// update existing approver
		err := svc.repoRegistry.HiringTeamApprover().UpdateHiringTeamApproverByID(ctx, input)
		if err != nil {
			return err
		}
	}
	for _, approver := range currentApprovers {
		if lo.NoneBy(inputs, func(input *ent.HiringTeamApproverInput) bool {
			return input.ID == approver.ID.String()
		}) {
			// delete approver
			err := svc.repoRegistry.HiringTeamApprover().DeleteHiringTeamApproverByID(ctx, approver.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
