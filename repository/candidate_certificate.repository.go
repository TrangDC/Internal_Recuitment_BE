package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidatecertificate"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateCertificateRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateCertificateInput, candidateId uuid.UUID) ([]*ent.Attachment, error)
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateCertificateInput) ([]*ent.Attachment, error)
	BuildBulkDelete(ctx context.Context, ids []uuid.UUID, candidateId uuid.UUID) error
}

type candidateCertificateRepoImpl struct {
	client *ent.Client
}

func NewCandidateCertificateRepository(client *ent.Client) CandidateCertificateRepository {
	return &candidateCertificateRepoImpl{
		client: client,
	}
}

func (rps candidateCertificateRepoImpl) BuildCreate() *ent.CandidateCertificateCreate {
	return rps.client.CandidateCertificate.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps candidateCertificateRepoImpl) BuildUpdate() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateCertificateRepoImpl) BuildDelete() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}
func (rps candidateCertificateRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateCertificateInput, candidateId uuid.UUID) ([]*ent.Attachment, error) {
	var createBulk []*ent.CandidateCertificateCreate
	var attachmentEdges []models.CreateBulkAttachmentInput
	var attachments []*ent.Attachment
	for _, v := range input {
		id := uuid.New()
		for _, attachmentRecord := range v.Attachments {
			attachmentEdges = append(attachmentEdges, models.CreateBulkAttachmentInput{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				OrderID:      v.OrderID,
			})
		}
		create := rps.BuildCreate().
			SetID(id).
			SetCandidateID(candidateId).
			SetName(strings.TrimSpace(v.Name)).
			SetScore(strings.TrimSpace(v.Score)).
			SetOrderID(v.OrderID)
		if v.AchievedDate != nil && !v.AchievedDate.IsZero() {
			create.SetAchievedDate(*v.AchievedDate)
		}
		createBulk = append(createBulk, create)
	}
	results, err := rps.client.CandidateCertificate.CreateBulk(createBulk...).Save(ctx)
	if err != nil {
		return attachments, err
	}
	for _, v := range results {
		attachmentEdge := lo.Filter(attachmentEdges, func(e models.CreateBulkAttachmentInput, index int) bool {
			return e.OrderID == v.OrderID
		})
		for _, attachmentRecord := range attachmentEdge {
			attachments = append(attachments, &ent.Attachment{
				DocumentID:   attachmentRecord.DocumentID,
				DocumentName: attachmentRecord.DocumentName,
				RelationType: attachment.RelationTypeCandidateCertificates,
				RelationID:   v.ID,
			})
		}
	}
	return attachments, nil
}

func (rps candidateCertificateRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateCertificateInput) ([]*ent.Attachment, error) {
	var attachments []*ent.Attachment
	for _, v := range input {
		for _, attachmentRecord := range v.Attachments {
			if attachmentRecord.ID != nil && *attachmentRecord.ID != "" {
				continue
			}
			attachments = append(attachments, &ent.Attachment{
				ID:           uuid.New(),
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateAwards,
				RelationID:   uuid.MustParse(v.ID),
			})
		}
		update := rps.client.CandidateCertificate.UpdateOneID(uuid.MustParse(v.ID)).
			SetName(strings.TrimSpace(v.Name)).
			SetScore(strings.TrimSpace(v.Score)).
			SetOrderID(v.OrderID)
		if v.AchievedDate != nil && !v.AchievedDate.IsZero() {
			update.SetAchievedDate(*v.AchievedDate)
		} else {
			update.ClearAchievedDate()
		}
		_, err := update.Save(ctx)
		if err != nil {
			return attachments, err
		}
	}
	return attachments, nil
}

func (rps candidateCertificateRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID, candidateId uuid.UUID) error {
	_, err := rps.client.CandidateCertificate.Delete().Where(candidatecertificate.IDNotIn(ids...), candidatecertificate.CandidateIDEQ(candidateId)).Exec(ctx)
	return err
}
