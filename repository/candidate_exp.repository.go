package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidateexp"

	"github.com/google/uuid"
)

type CandidateExpRepository interface {
	BuildBulkCreate(ctx context.Context, input []*ent.CandidateExpInput, candidateId uuid.UUID) error
	BuildBulkUpdate(ctx context.Context, input []*ent.CandidateExpInput) error
	BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error
}

type candidateExpRepoImpl struct {
	client *ent.Client
}

func NewCandidateExpRepository(client *ent.Client) CandidateExpRepository {
	return &candidateExpRepoImpl{
		client: client,
	}
}

func (rps candidateExpRepoImpl) BuildCreate() *ent.CandidateExpCreate {
	return rps.client.CandidateExp.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps candidateExpRepoImpl) BuildUpdate() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps candidateExpRepoImpl) BuildDelete() *ent.CandidateUpdate {
	return rps.client.Candidate.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps candidateExpRepoImpl) BuildBulkCreate(ctx context.Context, input []*ent.CandidateExpInput, candidateId uuid.UUID) error {
	var createBulk []*ent.CandidateExpCreate
	for _, v := range input {
		create := rps.BuildCreate().
			SetCandidateID(candidateId).
			SetPosition(strings.TrimSpace(v.Position)).
			SetCompany(strings.TrimSpace(v.Company)).
			SetLocation(strings.TrimSpace(v.Location)).
			SetDescription(strings.TrimSpace(v.Description)).
			SetOrderID(v.OrderID)
		if !v.StartDate.IsZero() {
			create.SetStartDate(*v.StartDate)
		}
		if !v.EndDate.IsZero() {
			create.SetEndDate(*v.EndDate)
		}
		createBulk = append(createBulk, create)
	}
	_, err := rps.client.CandidateExp.CreateBulk(createBulk...).Save(ctx)
	return err
}

func (rps candidateExpRepoImpl) BuildBulkUpdate(ctx context.Context, input []*ent.CandidateExpInput) error {
	for _, v := range input {
		update := rps.client.CandidateExp.UpdateOneID(uuid.MustParse(v.ID)).
			SetPosition(strings.TrimSpace(v.Position)).
			SetCompany(strings.TrimSpace(v.Company)).
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
			return err
		}
	}
	return nil
}

func (rps candidateExpRepoImpl) BuildBulkDelete(ctx context.Context, ids []uuid.UUID) error {
	_, err := rps.client.CandidateExp.Delete().Where(candidateexp.IDNotIn(ids...)).Exec(ctx)
	return err
}
