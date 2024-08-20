package service

import (
	"context"
	"trec/ent"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateExpService interface {
	ProcessCandidateExpInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateExpInput, repoRegistry repository.Repository) error
}

type candidateExpSvcImpl struct {
	logger *zap.Logger
}

func NewCandidateExpService(logger *zap.Logger) CandidateExpService {
	return &candidateExpSvcImpl{
		logger: logger,
	}
}

func (svc candidateExpSvcImpl) ProcessCandidateExpInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateExpInput, repoRegistry repository.Repository) error {
	var newRecord []*ent.CandidateExpInput
	var updateRecord []*ent.CandidateExpInput
	for _, v := range input {
		if v.ID == "" {
			newRecord = append(newRecord, v)
		} else {
			updateRecord = append(updateRecord, v)
		}
	}
	currentIds := lo.Map(updateRecord, func(v *ent.CandidateExpInput, index int) uuid.UUID {
		return uuid.MustParse(v.ID)
	})
	// Delete
	err := repoRegistry.CandidateExp().BuildBulkDelete(ctx, currentIds, candidateId)
	if err != nil {
		return err
	}
	// Create new
	if len(newRecord) > 0 {
		err := repoRegistry.CandidateExp().BuildBulkCreate(ctx, newRecord, candidateId)
		if err != nil {
			return err
		}
	}
	// Update
	if len(updateRecord) > 0 {
		err := repoRegistry.CandidateExp().BuildBulkUpdate(ctx, updateRecord)
		if err != nil {
			return err
		}
	}
	return nil
}
