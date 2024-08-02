package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/permission"
	"trec/ent/role"

	"github.com/google/uuid"
)

type RoleRepository interface {
	CreateRole(ctx context.Context, input ent.NewRoleInput) (*ent.Role, error)
	UpdateRole(ctx context.Context, record *ent.Role, input ent.UpdateRoleInput) (*ent.Role, error)
	DeleteRole(ctx context.Context, record *ent.Role) (*ent.Role, error)

	// query
	GetRole(ctx context.Context, id uuid.UUID) (*ent.Role, error)
	BuildQuery() *ent.RoleQuery
	BuildBaseQuery() *ent.RoleQuery
	BuildCount(ctx context.Context, query *ent.RoleQuery) (int, error)
	BuildList(ctx context.Context, query *ent.RoleQuery) ([]*ent.Role, error)

	// common function
	ValidName(ctx context.Context, teamId uuid.UUID, name string) (error, error)
}

type roleRepoImpl struct {
	client *ent.Client
}

func NewRoleRepository(client *ent.Client) RoleRepository {
	return &roleRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *roleRepoImpl) BuildCreate() *ent.RoleCreate {
	return rps.client.Role.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *roleRepoImpl) BuildUpdate() *ent.RoleUpdate {
	return rps.client.Role.Update().SetUpdatedAt(time.Now())
}

func (rps *roleRepoImpl) BuildDelete() *ent.RoleUpdate {
	return rps.client.Role.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *roleRepoImpl) BuildQuery() *ent.RoleQuery {
	return rps.client.Role.Query().Where(role.DeletedAtIsNil()).
		WithRolePermissionEdges(func(query *ent.EntityPermissionQuery) {
			query.
				WithPermissionEdges(func(query *ent.PermissionQuery) {
					query.Where(permission.DeletedAtIsNil())
				})
		}).
		WithUserEdges()
}
func (rps *roleRepoImpl) BuildBaseQuery() *ent.RoleQuery {
	return rps.client.Role.Query().Where(role.DeletedAtIsNil())
}

func (rps *roleRepoImpl) BuildGet(ctx context.Context, query *ent.RoleQuery) (*ent.Role, error) {
	return query.First(ctx)
}

func (rps *roleRepoImpl) BuildList(ctx context.Context, query *ent.RoleQuery) ([]*ent.Role, error) {
	return query.All(ctx)
}

func (rps *roleRepoImpl) BuildCount(ctx context.Context, query *ent.RoleQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *roleRepoImpl) BuildExist(ctx context.Context, query *ent.RoleQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *roleRepoImpl) BuildUpdateOne(ctx context.Context, record *ent.Role) *ent.RoleUpdateOne {
	return record.Update().SetUpdatedAt(time.Now())
}

func (rps *roleRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.RoleUpdateOne) (*ent.Role, error) {
	return update.Save(ctx)
}

// mutation
func (rps *roleRepoImpl) CreateRole(ctx context.Context, input ent.NewRoleInput) (*ent.Role, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(*input.Description)).Save(ctx)
}

func (rps *roleRepoImpl) UpdateRole(ctx context.Context, record *ent.Role, input ent.UpdateRoleInput) (*ent.Role, error) {
	return rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(*input.Name)).
		SetDescription(strings.TrimSpace(*input.Description)).Save(ctx)
}

func (rps *roleRepoImpl) DeleteRole(ctx context.Context, record *ent.Role) (*ent.Role, error) {
	update := rps.BuildUpdateOne(ctx, record).ClearUserRoles().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
	return update.Save(ctx)
}

// query
func (rps *roleRepoImpl) GetRole(ctx context.Context, id uuid.UUID) (*ent.Role, error) {
	query := rps.BuildQuery().Where(role.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *roleRepoImpl) ValidName(ctx context.Context, roleId uuid.UUID, name string) (error, error) {
	query := rps.BuildQuery().Where(role.NameEqualFold(strings.TrimSpace(name)))
	if roleId != uuid.Nil {
		query = query.Where(role.IDNEQ(roleId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.roles.validation.name_exist"), nil
	}
	return nil, nil
}
