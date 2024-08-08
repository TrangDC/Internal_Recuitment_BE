package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidateaward"

	"github.com/google/uuid"
)

type CandidateAwardRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateAwardInput, candidateId uuid.UUID) error
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateAwardInput) error
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

func (rps candidateAwardRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateAwardInput, candidateId uuid.UUID) error {
	var createBulk []*ent.CandidateAwardCreate
	for _, v := range input {
		var attachmentEdge []*ent.Attachment
		for _, attachmentRecord := range v.Attachments {
			attachmentEdge = append(attachmentEdge, &ent.Attachment{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateAwards,
			})
		}
		create := rps.BuildCreate().
			SetCandidateID(candidateId).
			SetName(strings.TrimSpace(v.Name)).
			AddAttachmentEdges(attachmentEdge...)
		if !v.AchievedDate.IsZero() {
			create.SetAchievedDate(*v.AchievedDate)
		}
		createBulk = append(createBulk, create)
	}
	_, err := rps.client.CandidateAward.CreateBulk(createBulk...).Save(ctx)
	return err
}

func (rps candidateAwardRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateAwardInput) error {
	for _, v := range input {
		var attachmentEdge []*ent.Attachment
		for _, attachmentRecord := range v.Attachments {
			if attachmentRecord.ID != nil && *attachmentRecord.ID != "" {
				continue
			}
			attachmentEdge = append(attachmentEdge, &ent.Attachment{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateAwards,
			})
		}
		update := rps.client.CandidateAward.UpdateOneID(uuid.MustParse(v.ID)).
			SetName(strings.TrimSpace(v.Name)).
			AddAttachmentEdges(attachmentEdge...)
		if !v.AchievedDate.IsZero() {
			update.SetAchievedDate(*v.AchievedDate)
		} else {
			update.ClearAchievedDate()
		}
		_, err := update.Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rps candidateAwardRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateAward.Delete().Where(candidateaward.IDNotIn(ids...)).Exec(ctx)
	return err
}
