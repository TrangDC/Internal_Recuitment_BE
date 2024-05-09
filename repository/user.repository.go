package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/user"

	"github.com/google/uuid"
)

type UserRepository interface {
	// query
	GetUser(ctx context.Context, id uuid.UUID) (*ent.User, error)
	BuildQuery() *ent.UserQuery
	BuildCount(ctx context.Context, query *ent.UserQuery) (int, error)
	BuildList(ctx context.Context, query *ent.UserQuery) ([]*ent.User, error)
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
	return rps.client.User.Create()
}

func (rps *userRepoImpl) BuildUpdate() *ent.UserUpdate {
	return rps.client.User.Update().SetUpdatedAt(time.Now())
}

func (rps *userRepoImpl) BuildDelete() *ent.UserUpdate {
	return rps.client.User.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *userRepoImpl) BuildQuery() *ent.UserQuery {
	return rps.client.User.Query().Where(user.DeletedAtIsNil())
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
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *userRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.UserUpdateOne) (*ent.User, error) {
	return update.Save(ctx)
}

func (rps *userRepoImpl) GetUser(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	query := rps.BuildQuery().Where(user.IDEQ(id))
	return rps.BuildGet(ctx, query)
}
