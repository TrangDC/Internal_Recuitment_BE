package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidateeducate"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateEducateRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateEducateInput, candidateId uuid.UUID) ([]*ent.Attachment, error)
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateEducateInput) ([]*ent.Attachment, error)
	BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error
}

type candidateEducateRepoImpl struct {
	client *ent.Client
}

func NewCandidateEducateRepository(client *ent.Client) CandidateEducateRepository {
	return &candidateEducateRepoImpl{
		client: client,
	}
}

func (rps candidateEducateRepoImpl) BuildCreate() *ent.CandidateEducateCreate {
	return rps.client.CandidateEducate.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps candidateEducateRepoImpl) BuildUpdate() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateEducateRepoImpl) BuildDelete() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps candidateEducateRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateEducateInput, candidateId uuid.UUID) ([]*ent.Attachment, error) {
	var createBulk []*ent.CandidateEducateCreate
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
			SetSchoolName(strings.TrimSpace(v.SchoolName)).
			SetMajor(strings.TrimSpace(v.Major)).
			SetGpa(strings.TrimSpace(v.Gpa)).
			SetCandidateID(candidateId).
			SetLocation(strings.TrimSpace(v.Location)).
			SetOrderID(v.OrderID)
		if !v.StartDate.IsZero() {
			create.SetStartDate(*v.StartDate)
		}
		if !v.EndDate.IsZero() {
			create.SetEndDate(*v.EndDate)
		}
		createBulk = append(createBulk, create)
	}
	results, err := rps.client.CandidateEducate.CreateBulk(createBulk...).Save(ctx)
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
				RelationType: attachment.RelationTypeCandidateEducates,
				RelationID:   v.ID,
			})
		}
	}
	return attachments, nil
}

func (rps candidateEducateRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateEducateInput) ([]*ent.Attachment, error) {
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
		update := rps.client.CandidateEducate.UpdateOneID(uuid.MustParse(v.ID)).
			SetSchoolName(strings.TrimSpace(v.SchoolName)).
			SetMajor(strings.TrimSpace(v.Major)).
			SetGpa(strings.TrimSpace(v.Gpa)).
			SetLocation(strings.TrimSpace(v.Location)).
			SetDescription(strings.TrimSpace(v.Description)).
			SetOrderID(v.OrderID)
		if !v.StartDate.IsZero() {
			update.SetStartDate(*v.StartDate)
		} else {
			update.ClearStartDate()
		}
		if !v.EndDate.IsZero() {
			update.SetEndDate(*v.EndDate)
		} else {
			update.ClearEndDate()
		}
		_, err := update.Save(ctx)
		if err != nil {
			return attachments, err
		}
	}
	return attachments, nil
}

func (rps candidateEducateRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateEducate.Delete().Where(candidateeducate.IDNotIn(ids...)).Exec(ctx)
	return err
}
