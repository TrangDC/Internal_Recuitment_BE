package repository

import (
	"context"
	"trec/ent"
	"trec/ent/permissiongroup"
)

type PermissionGroupRepository interface {
	GetAllPermissionGroups(ctx context.Context) ([]*ent.PermissionGroup, error)
}

type groupPermissionRepoImpl struct {
	client *ent.Client
}

func NewPermissionGroupRepository(client *ent.Client) PermissionGroupRepository {
	return &groupPermissionRepoImpl{
		client: client,
	}
}

func (rps groupPermissionRepoImpl) GetAllPermissionGroups(ctx context.Context) ([]*ent.PermissionGroup, error) {
	return rps.client.PermissionGroup.Query().WithPermissionEdges(
		func(q *ent.PermissionQuery) {
			q.Order(ent.Asc(permissiongroup.FieldOrderID))
		},
	).Order(ent.Asc(permissiongroup.FieldOrderID)).All(ctx)
}
