package repository

import (
	"context"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/candidatehistorycall"
	"trec/middleware"

	"github.com/google/uuid"
)

type CandidateHistoryCallRepository interface {
	CreateCandidateHistoryCall(ctx context.Context, input ent.NewCandidateHistoryCallInput) (*ent.CandidateHistoryCall, error)
	UpdateCandidateHistoryCall(ctx context.Context, record *ent.CandidateHistoryCall, input ent.UpdateCandidateHistoryCallInput) (*ent.CandidateHistoryCall, error)
	DeleteCandidateHistoryCall(ctx context.Context, record *ent.CandidateHistoryCall) (*ent.CandidateHistoryCall, error)

	// query
	GetCandidateHistoryCall(ctx context.Context, id uuid.UUID) (*ent.CandidateHistoryCall, error)
	BuildQuery() *ent.CandidateHistoryCallQuery
	BuildCount(ctx context.Context, query *ent.CandidateHistoryCallQuery) (int, error)
	BuildList(ctx context.Context, query *ent.CandidateHistoryCallQuery) ([]*ent.CandidateHistoryCall, error)
}

type candidateHistoryCallRepoImpl struct {
	client *ent.Client
}

func NewCandidateHistoryCallRepository(client *ent.Client) CandidateHistoryCallRepository {
	return &candidateHistoryCallRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *candidateHistoryCallRepoImpl) BuildCreate() *ent.CandidateHistoryCallCreate {
	return rps.client.CandidateHistoryCall.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *candidateHistoryCallRepoImpl) BuildUpdate() *ent.CandidateHistoryCallUpdate {
	return rps.client.CandidateHistoryCall.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateHistoryCallRepoImpl) BuildDelete() *ent.CandidateHistoryCallUpdate {
	return rps.client.CandidateHistoryCall.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *candidateHistoryCallRepoImpl) BuildQuery() *ent.CandidateHistoryCallQuery {
	return rps.client.CandidateHistoryCall.Query().Where(candidatehistorycall.DeletedAtIsNil()).WithCreatedByEdge().WithCandidateEdge()
}

func (rps *candidateHistoryCallRepoImpl) BuildGet(ctx context.Context, query *ent.CandidateHistoryCallQuery) (*ent.CandidateHistoryCall, error) {
	return query.First(ctx)
}

func (rps *candidateHistoryCallRepoImpl) BuildList(ctx context.Context, query *ent.CandidateHistoryCallQuery) ([]*ent.CandidateHistoryCall, error) {
	return query.All(ctx)
}

func (rps *candidateHistoryCallRepoImpl) BuildCount(ctx context.Context, query *ent.CandidateHistoryCallQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *candidateHistoryCallRepoImpl) BuildExist(ctx context.Context, query *ent.CandidateHistoryCallQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *candidateHistoryCallRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.CandidateHistoryCall) *ent.CandidateHistoryCallUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

func (rps *candidateHistoryCallRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.CandidateHistoryCallUpdateOne) (*ent.CandidateHistoryCall, error) {
	return update.Save(ctx)
}

// mutation
func (rps *candidateHistoryCallRepoImpl) CreateCandidateHistoryCall(ctx context.Context, input ent.NewCandidateHistoryCallInput) (*ent.CandidateHistoryCall, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	create := rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetCreatedByID(payload.UserID).
		SetDate(input.Date).
		SetCandidateID(uuid.MustParse(input.CandidateID)).
		SetContactTo(strings.TrimSpace(input.ContactTo)).
		SetType(candidatehistorycall.Type(input.Type)).
		SetDescription(strings.TrimSpace(input.Description))
	if !input.StartTime.IsZero() {
		create.SetStartTime(*input.StartTime)
	}
	if !input.EndTime.IsZero() {
		create.SetEndTime(*input.EndTime)
	}
	return create.Save(ctx)
}

func (rps *candidateHistoryCallRepoImpl) UpdateCandidateHistoryCall(ctx context.Context, record *ent.CandidateHistoryCall, input ent.UpdateCandidateHistoryCallInput) (*ent.CandidateHistoryCall, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetDate(input.Date).
		SetContactTo(strings.TrimSpace(input.ContactTo)).
		SetType(candidatehistorycall.Type(input.Type)).
		SetDescription(strings.TrimSpace(input.Description))
	if !input.StartTime.IsZero() {
		update.SetStartTime(*input.StartTime)
	} else {
		update.ClearStartTime()
	}
	if !input.EndTime.IsZero() {
		update.SetEndTime(*input.EndTime)
	} else {
		update.ClearEndTime()
	}
	return update.Save(ctx)
}

func (rps *candidateHistoryCallRepoImpl) DeleteCandidateHistoryCall(ctx context.Context, record *ent.CandidateHistoryCall) (*ent.CandidateHistoryCall, error) {
	update := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
	return update.Save(ctx)
}

// query
func (rps *candidateHistoryCallRepoImpl) GetCandidateHistoryCall(ctx context.Context, id uuid.UUID) (*ent.CandidateHistoryCall, error) {
	query := rps.BuildQuery().Where(candidatehistorycall.IDEQ(id))
	return rps.BuildGet(ctx, query)
}
