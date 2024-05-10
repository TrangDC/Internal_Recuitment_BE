package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/team"
	"trec/ent/teammanager"
	"trec/ent/user"
	"trec/internal/util"

	"github.com/google/uuid"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, input ent.NewTeamInput, memberIds []uuid.UUID) (*ent.Team, error)
	UpdateTeam(ctx context.Context, model *ent.Team, input ent.UpdateTeamInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.Team, error)
	DeleteTeam(ctx context.Context, model *ent.Team, memberIds []uuid.UUID) (*ent.Team, error)

	// query
	GetTeam(ctx context.Context, id uuid.UUID) (*ent.Team, error)
	BuildQuery() *ent.TeamQuery
	BuildCount(ctx context.Context, query *ent.TeamQuery) (int, error)
	BuildList(ctx context.Context, query *ent.TeamQuery) ([]*ent.Team, error)

	// common function
	ValidName(ctx context.Context, teamId uuid.UUID, name string) error
	ValidUserInAnotherTeam(ctx context.Context, id uuid.UUID, memberIds []uuid.UUID) error
}

type teamRepoImpl struct {
	client *ent.Client
}

func NewTeamRepository(client *ent.Client) TeamRepository {
	return &teamRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *teamRepoImpl) BuildCreate() *ent.TeamCreate {
	return rps.client.Team.Create()
}

func (rps *teamRepoImpl) BuildUpdate() *ent.TeamUpdate {
	return rps.client.Team.Update().SetUpdatedAt(time.Now())
}

func (rps *teamRepoImpl) BuildDelete() *ent.TeamUpdate {
	return rps.client.Team.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *teamRepoImpl) BuildQuery() *ent.TeamQuery {
	return rps.client.Team.Query().Where(team.DeletedAtIsNil()).WithUserEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	)
}

func (rps *teamRepoImpl) BuildGet(ctx context.Context, query *ent.TeamQuery) (*ent.Team, error) {
	return query.First(ctx)
}

func (rps *teamRepoImpl) BuildList(ctx context.Context, query *ent.TeamQuery) ([]*ent.Team, error) {
	return query.All(ctx)
}

func (rps *teamRepoImpl) BuildCount(ctx context.Context, query *ent.TeamQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *teamRepoImpl) BuildExist(ctx context.Context, query *ent.TeamQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *teamRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.Team) *ent.TeamUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *teamRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.TeamUpdateOne) (*ent.Team, error) {
	return update.Save(ctx)
}

// mutation
func (rps *teamRepoImpl) CreateTeam(ctx context.Context, input ent.NewTeamInput, memberIds []uuid.UUID) (*ent.Team, error) {
	create := rps.BuildCreate().SetName(strings.TrimSpace(input.Name)).AddUserEdgeIDs(memberIds...).SetSlug(util.SlugGeneration(input.Name))
	return create.Save(ctx)
}

func (rps *teamRepoImpl) UpdateTeam(ctx context.Context, model *ent.Team, input ent.UpdateTeamInput, newMemberIds []uuid.UUID, removeMemberIds []uuid.UUID) (*ent.Team, error) {
	update := rps.BuildUpdateOne(ctx, model).SetName(strings.TrimSpace(input.Name)).SetSlug(util.SlugGeneration(input.Name)).
		AddUserEdgeIDs(newMemberIds...).RemoveUserEdgeIDs(removeMemberIds...)
	return rps.BuildSaveUpdateOne(ctx, update)
}

func (rps *teamRepoImpl) DeleteTeam(ctx context.Context, model *ent.Team, memberIds []uuid.UUID) (*ent.Team, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now()).RemoveUserEdgeIDs(memberIds...)
	return update.Save(ctx)
}

// query
func (rps *teamRepoImpl) GetTeam(ctx context.Context, id uuid.UUID) (*ent.Team, error) {
	query := rps.BuildQuery().Where(team.IDEQ(id)).WithUserEdges(
		func(query *ent.UserQuery) {
			query.Where(user.DeletedAtIsNil())
		},
	)
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *teamRepoImpl) ValidName(ctx context.Context, teamId uuid.UUID, name string) error {
	query := rps.BuildQuery().Where(team.NameEqualFold(strings.TrimSpace(name)))
	if teamId != uuid.Nil {
		query = query.Where(team.IDNEQ(teamId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("model.teams.validation.name_exist")
	}
	return nil
}

func (rps *teamRepoImpl) ValidUserInAnotherTeam(ctx context.Context, id uuid.UUID, memberIds []uuid.UUID) error {
	query := rps.BuildQuery().Where(team.HasUserEdgesWith(user.IDIn(memberIds...)), team.HasUserTeamsWith(teammanager.DeletedAtIsNil()))
	if id != uuid.Nil {
		query = query.Where(team.IDNEQ(id))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("model.teams.validation.user_in_another_team")
	}
	return nil
}
