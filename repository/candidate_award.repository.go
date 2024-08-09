package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidateaward"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateAwardRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateAwardInput, candidateId uuid.UUID) ([]*ent.Attachment, error)
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateAwardInput) ([]*ent.Attachment, error)
	BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error
}

type candidateAwardRepoImpl struct {
	client *ent.Client
}

func NewCandidateAwardRepository(client *ent.Client) CandidateAwardRepository {
	return &candidateAwardRepoImpl{
		client: client,
	}
}

func (rps candidateAwardRepoImpl) BuildCreate() *ent.CandidateAwardCreate {
	return rps.client.CandidateAward.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps candidateAwardRepoImpl) BuildUpdate() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateAwardRepoImpl) BuildDelete() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps candidateAwardRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateAwardInput, candidateId uuid.UUID) ([]*ent.Attachment, error) {
	var createBulk []*ent.CandidateAwardCreate
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
			SetOrderID(v.OrderID)
		if !v.AchievedDate.IsZero() {
			create.SetAchievedDate(*v.AchievedDate)
		}
		createBulk = append(createBulk, create)
	}
	results, err := rps.client.CandidateAward.CreateBulk(createBulk...).Save(ctx)
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
				RelationType: attachment.RelationTypeCandidateAwards,
				RelationID:   v.ID,
			})
		}
	}
	return attachments, nil
}

func (rps candidateAwardRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateAwardInput) ([]*ent.Attachment, error) {
	var attachments []*ent.Attachment
	for _, v := range input {
		for _, attachmentRecord := range v.Attachments {
			if attachmentRecord.ID != nil && *attachmentRecord.ID != "" {
				continue
			}
			attachments = append(attachments, &ent.Attachment{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateAwards,
				RelationID:   uuid.MustParse(v.ID),
			})
		}
		update := rps.client.CandidateAward.UpdateOneID(uuid.MustParse(v.ID)).
			SetName(strings.TrimSpace(v.Name)).
			SetOrderID(v.OrderID)
		if !v.AchievedDate.IsZero() {
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

func (rps candidateAwardRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateAward.Delete().Where(candidateaward.IDNotIn(ids...)).Exec(ctx)
	return err
}
