package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidateeducate"

	"github.com/google/uuid"
)

type CandidateEducateRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateEducateInput, candidateId uuid.UUID) error
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateEducateInput) error
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

func (rps candidateEducateRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateEducateInput, candidateId uuid.UUID) error {
	var createBulk []*ent.CandidateEducateCreate
	for _, v := range input {
		var attachmentEdge []*ent.Attachment
		for _, attachmentRecord := range v.Attachments {
			attachmentEdge = append(attachmentEdge, &ent.Attachment{
				DocumentID:   uuid.MustParse(attachmentRecord.DocumentID),
				DocumentName: strings.TrimSpace(attachmentRecord.DocumentName),
				RelationType: attachment.RelationTypeCandidateEducates,
			})
		}
		create := rps.BuildCreate().
			SetSchoolName(strings.TrimSpace(v.SchoolName)).
			SetMajor(strings.TrimSpace(v.Major)).
			SetGpa(strings.TrimSpace(v.Gpa)).
			SetCandidateID(candidateId).
			SetLocation(strings.TrimSpace(v.Location)).
			AddAttachmentEdges(attachmentEdge...)
		if !v.StartDate.IsZero() {
			create.SetStartDate(*v.StartDate)
		}
		if !v.EndDate.IsZero() {
			create.SetEndDate(*v.EndDate)
		}
		createBulk = append(createBulk, create)
	}
	_, err := rps.client.CandidateEducate.CreateBulk(createBulk...).Save(ctx)
	return err
}

func (rps candidateEducateRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateEducateInput) error {
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
		update := rps.client.CandidateEducate.UpdateOneID(uuid.MustParse(v.ID)).
			SetSchoolName(strings.TrimSpace(v.SchoolName)).
			SetMajor(strings.TrimSpace(v.Major)).
			SetGpa(strings.TrimSpace(v.Gpa)).
			SetLocation(strings.TrimSpace(v.Location)).
			SetDescription(strings.TrimSpace(v.Description)).
			AddAttachmentEdges(attachmentEdge...)
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
			return err
		}
	}
	return nil
}

func (rps candidateEducateRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateEducate.Delete().Where(candidateeducate.IDNotIn(ids...)).Exec(ctx)
	return err
}
