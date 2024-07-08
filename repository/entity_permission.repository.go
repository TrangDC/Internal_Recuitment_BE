package repository

import (
	"context"
	"fmt"
	"time"
	"trec/ent"
	"trec/ent/entitypermission"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type EntityPermissionRepository interface {
	CreateAndUpdateEntityPermission(ctx context.Context, entityId uuid.UUID, input []*ent.NewEntityPermissionInput,
		entityPermissionRecord []*ent.EntityPermission, entityPermissionType entitypermission.EntityType) error
	DeleteAllEntityPermission(ctx context.Context, entityId uuid.UUID) error
	ValidActionPermission(ctx context.Context, input []*ent.NewEntityPermissionInput) (error, error)
}

type entityPermissionRepoImpl struct {
	client *ent.Client
}

type updateEntityPermissionRecord struct {
	RecordInput *ent.NewEntityPermissionInput
	Record      *ent.EntityPermission
}

func NewEntityPermissionRepository(client *ent.Client) EntityPermissionRepository {
	return &entityPermissionRepoImpl{
		client: client,
	}
}

func (rps entityPermissionRepoImpl) CreateBulkEntityPermission(ctx context.Context, input []*ent.NewEntityPermissionInput,
	entityId uuid.UUID, entityPermissionType entitypermission.EntityType) error {
	var recordCreate []*ent.EntityPermissionCreate
	for _, entity := range input {
		recordCreate = append(recordCreate,
			rps.client.EntityPermission.Create().
				SetEntityID(entityId).
				SetPermissionID(uuid.MustParse(entity.PermissionID)).
				SetEntityType(entityPermissionType).
				SetForOwner(entity.ForOwner).
				SetForAll(entity.ForAll).
				SetForTeam(entity.ForTeam).
				SetCreatedAt(time.Now().UTC()).
				SetUpdatedAt(time.Now().UTC()),
		)
	}
	_, err := rps.client.EntityPermission.CreateBulk(recordCreate...).Save(ctx)
	return err
}

func (rps entityPermissionRepoImpl) DeleteEntityPermission(ctx context.Context, entityPermissionIds []uuid.UUID) error {
	_, err := rps.client.EntityPermission.Delete().Where(entitypermission.IDIn(entityPermissionIds...)).Exec(ctx)
	return err
}

func (rps entityPermissionRepoImpl) UpdateEntityPermission(ctx context.Context, record *ent.EntityPermission, input *ent.NewEntityPermissionInput) error {
	_, err := rps.client.EntityPermission.UpdateOne(record).
		SetPermissionID(uuid.MustParse(input.PermissionID)).
		SetForOwner(input.ForOwner).
		SetForAll(input.ForAll).
		SetForTeam(input.ForTeam).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps entityPermissionRepoImpl) CreateAndUpdateEntityPermission(ctx context.Context, entityId uuid.UUID, input []*ent.NewEntityPermissionInput,
	entityPermissionRecord []*ent.EntityPermission, entityPermissionType entitypermission.EntityType) error {
	var newInput []*ent.NewEntityPermissionInput
	var deletedIds []uuid.UUID
	var updatedRecord []updateEntityPermissionRecord
	for _, entity := range input {
		if entity.ID == nil || *entity.ID == "" {
			newInput = append(newInput, entity)
		}
	}
	if len(entityPermissionRecord) > 0 {
		for _, entity := range entityPermissionRecord {
			inputRecord, exist := lo.Find(input, func(record *ent.NewEntityPermissionInput) bool {
				return record.ID != nil && *record.ID != "" && *record.ID == entity.ID.String()
			})
			if !exist {
				deletedIds = append(deletedIds, entity.ID)
			} else {
				if inputRecord.PermissionID != entity.PermissionID.String() || inputRecord.ForOwner != entity.ForOwner ||
					inputRecord.ForAll != entity.ForAll || inputRecord.ForTeam != entity.ForTeam {
					updatedRecord = append(updatedRecord, updateEntityPermissionRecord{
						RecordInput: inputRecord,
						Record:      entity,
					})
				}
			}
		}
	}
	if len(newInput) > 0 {
		err := rps.CreateBulkEntityPermission(ctx, newInput, entityId, entityPermissionType)
		if err != nil {
			return err
		}
	}
	if len(deletedIds) > 0 {
		err := rps.DeleteEntityPermission(ctx, deletedIds)
		if err != nil {
			return err
		}
	}
	if len(updatedRecord) > 0 {
		for _, record := range updatedRecord {
			err := rps.UpdateEntityPermission(ctx, record.Record, record.RecordInput)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (rps entityPermissionRepoImpl) DeleteAllEntityPermission(ctx context.Context, entityId uuid.UUID) error {
	_, err := rps.client.EntityPermission.Delete().Where(entitypermission.EntityID(entityId)).Exec(ctx)
	return err
}

// common
func (rps entityPermissionRepoImpl) ValidActionPermission(ctx context.Context, input []*ent.NewEntityPermissionInput) (error, error) {
	permissions, err := rps.client.Permission.Query().Where().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, record := range input {
		if (!record.ForAll && !record.ForOwner && !record.ForTeam) || (record.ForAll && record.ForOwner && record.ForTeam) {
			return fmt.Errorf("model.permissions.validation.invalid_permission"), nil
		}
		permissionRecord, exist := lo.Find(permissions, func(permission *ent.Permission) bool {
			return permission.ID.String() == record.PermissionID
		})
		if !exist {
			return fmt.Errorf("model.permissions.validation.permission_not_found"), nil
		}
		if record.ForAll && !permissionRecord.ForAll {
			return fmt.Errorf("model.permissions.validation.permission_not_allow_for_all"), nil
		}
		if record.ForOwner && !permissionRecord.ForOwner {
			return fmt.Errorf("model.permissions.validation.permission_not_allow_for_owner"), nil
		}
		if record.ForTeam && !permissionRecord.ForTeam {
			return fmt.Errorf("model.permissions.validation.permission_not_allow_for_team"), nil
		}
	}
	inputPermissionIds := lo.Map(input, func(record *ent.NewEntityPermissionInput, index int) string {
		return record.PermissionID
	})
	selectPermissions := lo.Filter(permissions, func(entity *ent.Permission, index int) bool {
		return lo.Contains(inputPermissionIds, entity.ID.String())
	})
	selectPermissionIds := lo.Map(selectPermissions, func(entity *ent.Permission, index int) uuid.UUID {
		return entity.ID
	})
	invalidPermissions := lo.Filter(selectPermissions, func(record *ent.Permission, index int) bool {
		return record.ParentID != uuid.Nil && !lo.Contains(selectPermissionIds, uuid.MustParse(record.ParentID.String()))
	})
	if len(invalidPermissions) > 0 {
		return fmt.Errorf("model.permissions.validation.missing_permission_parent"), nil
	}
	return nil, nil
}
