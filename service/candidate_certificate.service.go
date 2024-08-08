package service

import (
	"context"
	"trec/ent"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateCertificateService interface {
	ProcessCandidateCertificateInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateCertificateInput, oldRecords []*ent.CandidateCertificate, repoRegistry repository.Repository) error
}

type candidateCertificateSvcImpl struct {
	logger *zap.Logger
}

func NewCandidateCertificateService(logger *zap.Logger) CandidateCertificateService {
	return &candidateCertificateSvcImpl{
		logger: logger,
	}
}

func (svc candidateCertificateSvcImpl) ProcessCandidateCertificateInput(ctx context.Context, candidateId uuid.UUID, input []*ent.CandidateCertificateInput,
	oldRecords []*ent.CandidateCertificate, repoRegistry repository.Repository) error {
	var newRecord []*ent.CandidateCertificateInput
	var updateRecords []*ent.CandidateCertificateInput
	var deleteAttachmentIds []uuid.UUID
	var currentIds []uuid.UUID
	newRecord = lo.Filter(input, func(entity *ent.CandidateCertificateInput, index int) bool {
		return entity.ID == ""
	})
	for _, v := range oldRecords {
		updateRecord, exist := lo.Find(input, func(entity *ent.CandidateCertificateInput) bool {
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
		err := repoRegistry.CandidateCertificate().BuildBulkDelete(ctx, currentIds)
		if err != nil {
			return err
		}
	}
	// Create new
	if len(newRecord) > 0 {
		err := repoRegistry.CandidateCertificate().BuildBulkCreate(ctx, newRecord, candidateId)
		if err != nil {
			return err
		}
	}
	// Update
	if len(updateRecords) > 0 {
		err := repoRegistry.CandidateCertificate().BuildBulkUpdate(ctx, updateRecords)
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
