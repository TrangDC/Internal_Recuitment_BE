package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/recteam"
	"trec/ent/user"

	"github.com/google/uuid"
)

type RecTeamRepository interface {
	//mutation
	CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput) (*ent.RecTeam, error)

	// query
	GetRecTeam(ctx context.Context, id uuid.UUID) (*ent.RecTeam, error)
	BuildQuery() *ent.RecTeamQuery
	BuildBaseQuery() *ent.RecTeamQuery
	BuildCount(ctx context.Context, query *ent.RecTeamQuery) (int, error)
	BuildList(ctx context.Context, query *ent.RecTeamQuery) ([]*ent.RecTeam, error)
	BuildGetOne(ctx context.Context, query *ent.RecTeamQuery) (*ent.RecTeam, error)
	DeleteRelationRecTeam(ctx context.Context, recTeamId uuid.UUID) error

	// common function
	ValidInput(ctx context.Context, recTeamId uuid.UUID, name string, leaderID uuid.UUID) (error, error)
}

type recTeamRepoImpl struct {
	client *ent.Client
}

func NewRecTeamRepository(client *ent.Client) RecTeamRepository {
	return &recTeamRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *recTeamRepoImpl) BuildCreate() *ent.RecTeamCreate {
	return rps.client.RecTeam.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *recTeamRepoImpl) BuildUpdate() *ent.RecTeamUpdate {
	return rps.client.RecTeam.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *recTeamRepoImpl) BuildDelete() *ent.RecTeamUpdate {
	return rps.client.RecTeam.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps *recTeamRepoImpl) BuildQuery() *ent.RecTeamQuery {
	return rps.client.RecTeam.Query().Where(recteam.DeletedAtIsNil()).
		WithRecLeaderEdge(
			func(query *ent.UserQuery) {
				query.Where(user.DeletedAtIsNil())
			},
		).WithRecMemberEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	)
}

func (rps *recTeamRepoImpl) BuildBaseQuery() *ent.RecTeamQuery {
	return rps.client.RecTeam.Query().Where(recteam.DeletedAtIsNil())
}

func (rps *recTeamRepoImpl) BuildCount(ctx context.Context, query *ent.RecTeamQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *recTeamRepoImpl) BuildList(ctx context.Context, query *ent.RecTeamQuery) ([]*ent.RecTeam, error) {
	return query.All(ctx)
}

func (rps *recTeamRepoImpl) BuildGetOne(ctx context.Context, query *ent.RecTeamQuery) (*ent.RecTeam, error) {
	return query.First(ctx)
}

func (rps *recTeamRepoImpl) DeleteRelationRecTeam(ctx context.Context, recTeamId uuid.UUID) error {
	return rps.client.RecTeam.DeleteOneID(recTeamId).Exec(ctx)
}

func (rps *recTeamRepoImpl) BuildExist(ctx context.Context, query *ent.RecTeamQuery) (bool, error) {
	return query.Exist(ctx)
}

// mutation
func (rps *recTeamRepoImpl) CreateRecTeam(ctx context.Context, input ent.NewRecTeamInput) (*ent.RecTeam, error) {
	create := rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetRecLeaderEdgeID(uuid.MustParse(input.LeaderID))
	return create.Save(ctx)
}

// query
func (rps *recTeamRepoImpl) GetRecTeam(ctx context.Context, id uuid.UUID) (*ent.RecTeam, error) {
	return rps.BuildQuery().Where(recteam.IDEQ(id)).First(ctx)
}

// common function
func (rps *recTeamRepoImpl) ValidInput(ctx context.Context, recTeamID uuid.UUID, name string, leaderID uuid.UUID) (error, error) {
	query := rps.BuildQuery().Where(recteam.NameEqualFold(strings.TrimSpace(name)))
	if recTeamID != uuid.Nil {
			query = query.Where(recteam.IDNEQ(recTeamID))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
			return nil, err
	}
	if isExist {
			return fmt.Errorf("model.rec_teams.validation.name_exist"), nil
	}

	query = rps.BuildQuery().Where(recteam.LeaderID(leaderID))
	isExist, err = rps.BuildExist(ctx, query)
	if err != nil {
			return nil, err
	}
	if isExist {
			return fmt.Errorf("model.rec_teams.validation.is_leader_in_another_rec_team"), nil
	}

	return nil, nil
}


