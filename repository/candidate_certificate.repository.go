package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidatecertificate"

	"github.com/google/uuid"
)

type CandidateCertificateRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateCertificateInput, candidateId uuid.UUID) error
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateCertificateInput) error
	BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error
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
func (rps candidateCertificateRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateCertificateInput, candidateId uuid.UUID) error {
	var createBulk []*ent.CandidateCertificateCreate
	for _, v := range input {
		var attachmentEdge []*ent.Attachment
		for _, attachmentRecord := range v.Attachments {
			attachmentEdge = append(attachmentEdge, &ent.Attachment{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateCertificates,
			})
		}
		create := rps.BuildCreate().
			SetCandidateID(candidateId).
			SetName(strings.TrimSpace(v.Name)).
			SetScore(strings.TrimSpace(v.Score)).
			AddAttachmentEdges(attachmentEdge...)
		if !v.AchievedDate.IsZero() {
			create.SetAchievedDate(*v.AchievedDate)
		}
		createBulk = append(createBulk, create)
	}
	_, err := rps.client.CandidateCertificate.CreateBulk(createBulk...).Save(ctx)
	return err
}

func (rps candidateCertificateRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateCertificateInput) error {
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
		update := rps.client.CandidateCertificate.UpdateOneID(uuid.MustParse(v.ID)).
			SetName(strings.TrimSpace(v.Name)).
			SetScore(strings.TrimSpace(v.Score)).
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

func (rps candidateCertificateRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateCertificate.Delete().Where(candidatecertificate.IDNotIn(ids...)).Exec(ctx)
	return err
}
