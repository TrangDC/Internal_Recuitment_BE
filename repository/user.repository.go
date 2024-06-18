package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/team"
	"trec/ent/user"

	"github.com/google/uuid"
)

type UserRepository interface {
	// mutation
	CreateUser(ctx context.Context, input *ent.NewUserInput) (*ent.User, error)
	UpdateUser(ctx context.Context, record *ent.User, input *ent.UpdateUserInput) (*ent.User, error)
	DeleteUser(ctx context.Context, record *ent.User) error
	UpdateUserStatus(ctx context.Context, record *ent.User, status user.Status) (*ent.User, error)
	// query
	GetUser(ctx context.Context, userId uuid.UUID) (*ent.User, error)
	BuildQuery() *ent.UserQuery
	BuildCount(ctx context.Context, query *ent.UserQuery) (int, error)
	BuildList(ctx context.Context, query *ent.UserQuery) ([]*ent.User, error)
	// common function
	ValidWorkEmail(ctx context.Context, userId uuid.UUID, workEmail string) error
}

type userRepoImpl struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *userRepoImpl) BuildCreate() *ent.UserCreate {
	return rps.client.User.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *userRepoImpl) BuildUpdate() *ent.UserUpdate {
	return rps.client.User.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *userRepoImpl) BuildDelete() *ent.UserUpdate {
	return rps.client.User.Update().SetDeletedAt(time.Now().UTC()).SetUpdatedAt(time.Now().UTC())
}

func (rps *userRepoImpl) BuildQuery() *ent.UserQuery {
	return rps.client.User.Query().Where(user.DeletedAtIsNil()).WithTeamEdges(
		func(query *ent.TeamQuery) {
			query.Where(team.DeletedAtIsNil())
		},
	)
}

func (rps *userRepoImpl) BuildParanoidQuery() *ent.UserQuery {
	return rps.client.User.Query()
}

func (rps *userRepoImpl) BuildGet(ctx context.Context, query *ent.UserQuery) (*ent.User, error) {
	return query.First(ctx)
}

func (rps *userRepoImpl) BuildList(ctx context.Context, query *ent.UserQuery) ([]*ent.User, error) {
	return query.All(ctx)
}

func (rps *userRepoImpl) BuildCount(ctx context.Context, query *ent.UserQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *userRepoImpl) BuildExist(ctx context.Context, query *ent.UserQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *userRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.User) *ent.UserUpdateOne {
	return model.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *userRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.UserUpdateOne) (*ent.User, error) {
	return update.Save(ctx)
}

// mutation
func (rps *userRepoImpl) CreateUser(ctx context.Context, input *ent.NewUserInput) (*ent.User, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetWorkEmail(strings.TrimSpace(input.WorkEmail)).
		Save(ctx)
}

func (rps *userRepoImpl) UpdateUser(ctx context.Context, record *ent.User, input *ent.UpdateUserInput) (*ent.User, error) {
	return rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetWorkEmail(strings.TrimSpace(input.WorkEmail)).
		SetStatus(user.Status(input.Status)).
		Save(ctx)
}

func (rps *userRepoImpl) UpdateUserStatus(ctx context.Context, record *ent.User, status user.Status) (*ent.User, error) {
	return rps.BuildUpdateOne(ctx, record).SetStatus(status).Save(ctx)
}

func (rps *userRepoImpl) DeleteUser(ctx context.Context, record *ent.User) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

// query
func (rps *userRepoImpl) GetUser(ctx context.Context, userId uuid.UUID) (*ent.User, error) {
	query := rps.BuildQuery().Where(user.IDEQ(userId))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *userRepoImpl) ValidWorkEmail(ctx context.Context, userId uuid.UUID, workEmail string) error {
	query := rps.BuildParanoidQuery().Where(user.WorkEmailEqualFold(strings.TrimSpace(strings.ToLower(workEmail))))
	if userId != uuid.Nil {
		query = query.Where(user.IDNEQ(userId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("model.users.validation.work_email_exist")
	}
	return nil
}
