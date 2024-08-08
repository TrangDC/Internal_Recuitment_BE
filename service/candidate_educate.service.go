package service

import (
	"context"
	"trec/ent"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateEducateService interface {
	ProcessCandidateEducateInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateEducateInput, oldRecords []*ent.CandidateEducate, repoRegistry repository.Repository) error
}

type candidateEducateSvcImpl struct {
	logger *zap.Logger
}

func NewCandidateEducateService(logger *zap.Logger) CandidateEducateService {
	return &candidateEducateSvcImpl{
		logger: logger,
	}
}

func (svc candidateEducateSvcImpl) ProcessCandidateEducateInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateEducateInput,
	oldRecords []*ent.CandidateEducate, repoRegistry repository.Repository) error {
	var newRecord []*ent.CandidateEducateInput
	var updateRecords []*ent.CandidateEducateInput
	var deleteAttachmentIds []uuid.UUID
	var currentIds []uuid.UUID
	newRecord = lo.Filter(input, func(entity *ent.CandidateEducateInput, index int) bool {
		return entity.ID == ""
	})
	for _, v := range oldRecords {
		updateRecord, exist := lo.Find(input, func(entity *ent.CandidateEducateInput) bool {
			return entity.ID == v.ID.String()
		})
		if !exist {
			deleteAttachmentIds = append(deleteAttachmentIds, lo.Map(v.Edges.AttachmentEdges, func(entity *ent.Attachment, index int) uuid.UUID {
				return entity.ID
			})...)
		} else {
			updateRecords = append(updateRecords, updateRecord)
			currentIds = append(currentIds, v.ID)
			for _, oldRecordAttachment := range v.Edges.AttachmentEdges {
				_, exist := lo.Find(updateRecord.Attachments, func(entity *ent.NewAttachmentInput) bool {
					return *entity.ID == oldRecordAttachment.ID.String()
				})
				if !exist {
					deleteAttachmentIds = append(deleteAttachmentIds, oldRecordAttachment.ID)
				}
			}
		}
	}
	// Delete
	if len(currentIds) > 0 {
		err := repoRegistry.CandidateEducate().BuildBulkDelete(ctx, currentIds)
		if err != nil {
			return err
		}
	}
	// Create new
	if len(newRecord) > 0 {
		err := repoRegistry.CandidateEducate().BuildBulkCreate(ctx, newRecord, candidateId)
		if err != nil {
			return err
		}
	}
	// Update
	if len(updateRecords) > 0 {
		err := repoRegistry.CandidateEducate().BuildBulkUpdate(ctx, updateRecords)
		if err != nil {
			return err
		}
	}
	if len(deleteAttachmentIds) > 0 {
		err := repoRegistry.Attachment().RemoveBulkAttachment(ctx, deleteAttachmentIds)
		if err != nil {
			return err
		}
	}
	return nil
}
