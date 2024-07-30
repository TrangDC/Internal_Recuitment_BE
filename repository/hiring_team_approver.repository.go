package repository

import (
	"context"
	"time"
	"trec/ent"

	"github.com/google/uuid"
)

type HiringTeamApproverRepository interface {
	CreateHiringTeamApprover(ctx context.Context, input *ent.HiringTeamApproverInput, hiringTeamID uuid.UUID) error
	UpdateHiringTeamApproverByID(ctx context.Context, input *ent.HiringTeamApproverInput) error
	DeleteHiringTeamApproverByID(ctx context.Context, id uuid.UUID) error
}

type hiringTeamApproverRepoImpl struct {
	client *ent.Client
}

func NewHiringTeamApproverRepository(client *ent.Client) HiringTeamApproverRepository {
	return &hiringTeamApproverRepoImpl{
		client: client,
	}
}

func (rps *hiringTeamApproverRepoImpl) CreateHiringTeamApprover(ctx context.Context, input *ent.HiringTeamApproverInput, hiringTeamID uuid.UUID) error {
	_, err := rps.client.HiringTeamApprover.Create().
		SetUserID(uuid.MustParse(input.UserID)).SetHiringTeamID(hiringTeamID).
		SetOrderID(input.OrderID).
		Save(ctx)
	return err
}

func (rps *hiringTeamApproverRepoImpl) UpdateHiringTeamApproverByID(ctx context.Context, input *ent.HiringTeamApproverInput) error {
	_, err := rps.client.HiringTeamApprover.UpdateOneID(uuid.MustParse(input.ID)).
		SetOrderID(input.OrderID).
		SetUpdatedAt(time.Now().UTC()).Save(ctx)
	return err
}

func (rps *hiringTeamApproverRepoImpl) DeleteHiringTeamApproverByID(ctx context.Context, id uuid.UUID) error {
	return rps.client.HiringTeamApprover.DeleteOneID(id).Exec(ctx)
}
