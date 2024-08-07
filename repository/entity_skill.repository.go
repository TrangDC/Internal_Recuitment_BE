package repository

import (
	"context"
	"time"
	"trec/ent"
	"trec/ent/entityskill"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type EntitySkillRepository interface {
	CreateAndUpdateEntitySkill(ctx context.Context, entityId uuid.UUID, input []*ent.EntitySkillRecordInput,
		entitySkillRecord []*ent.EntitySkill, entitySkillType entityskill.EntityType) error
	DeleteAllEntitySkill(ctx context.Context, entityId uuid.UUID) error
	DeleteBulkEntitySkillBySkillID(ctx context.Context, skillID uuid.UUID) error
}

type entitySkillRepoImpl struct {
	client *ent.Client
}

type updateEntitySkillRecord struct {
	RecordInput *ent.EntitySkillRecordInput
	Record      *ent.EntitySkill
}

func NewEntitySkillRepository(client *ent.Client) EntitySkillRepository {
	return &entitySkillRepoImpl{
		client: client,
	}
}

func (rps entitySkillRepoImpl) CreateBulkEntitySkill(ctx context.Context, input []*ent.EntitySkillRecordInput,
	entityId uuid.UUID, entitySkillType entityskill.EntityType) error {
	var recordCreate []*ent.EntitySkillCreate
	for _, entity := range input {
		recordCreate = append(recordCreate,
			rps.client.EntitySkill.Create().
				SetEntityID(entityId).
				SetEntityType(entitySkillType).
				SetSkillID(uuid.MustParse(entity.SkillID)).
				SetOrderID(entity.OrderID).
				SetCreatedAt(time.Now().UTC()).
				SetUpdatedAt(time.Now().UTC()),
		)
	}
	_, err := rps.client.EntitySkill.CreateBulk(recordCreate...).Save(ctx)
	return err
}

func (rps entitySkillRepoImpl) DeleteEntitySkill(ctx context.Context, entitySkillIds []uuid.UUID) error {
	_, err := rps.client.EntitySkill.Update().Where(entityskill.IDIn(entitySkillIds...)).
		SetDeletedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps entitySkillRepoImpl) UpdateEntitySkill(ctx context.Context, record *ent.EntitySkill, input *ent.EntitySkillRecordInput) error {
	_, err := rps.client.EntitySkill.UpdateOne(record).
		SetUpdatedAt(time.Now().UTC()).
		SetOrderID(input.OrderID).
		Save(ctx)
	return err
}

func (rps entitySkillRepoImpl) CreateAndUpdateEntitySkill(ctx context.Context, entityId uuid.UUID, input []*ent.EntitySkillRecordInput,
	entitySkillRecord []*ent.EntitySkill, entitySkillType entityskill.EntityType) error {
	var newInput []*ent.EntitySkillRecordInput
	var deletedIds []uuid.UUID
	var updatedRecord []updateEntitySkillRecord
	for _, entity := range input {
		if entity.ID == nil || *entity.ID == "" {
			newInput = append(newInput, entity)
		}
	}
	if len(entitySkillRecord) > 0 {
		for _, entity := range entitySkillRecord {
			inputRecord, exist := lo.Find(input, func(record *ent.EntitySkillRecordInput) bool {
				return record.ID != nil && *record.ID != "" && *record.ID == entity.ID.String()
			})
			if !exist {
				deletedIds = append(deletedIds, entity.ID)
			} else {
				if inputRecord.OrderID != entity.OrderID {
					updatedRecord = append(updatedRecord, updateEntitySkillRecord{
						RecordInput: inputRecord,
						Record:      entity,
					})
				}
			}
		}
	}
	if len(newInput) > 0 {
		err := rps.CreateBulkEntitySkill(ctx, newInput, entityId, entitySkillType)
		if err != nil {
			return err
		}
	}
	if len(deletedIds) > 0 {
		err := rps.DeleteEntitySkill(ctx, deletedIds)
		if err != nil {
			return err
		}
	}
	if len(updatedRecord) > 0 {
		for _, record := range updatedRecord {
			err := rps.UpdateEntitySkill(ctx, record.Record, record.RecordInput)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (rps entitySkillRepoImpl) DeleteAllEntitySkill(ctx context.Context, entityId uuid.UUID) error {
	_, err := rps.client.EntitySkill.Update().Where(entityskill.EntityID(entityId)).
		SetDeletedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (rps entitySkillRepoImpl) DeleteBulkEntitySkillBySkillID(ctx context.Context, skillID uuid.UUID) error {
	_, err := rps.client.EntitySkill.Delete().Where(entityskill.SkillID(skillID)).Exec(ctx)
	return err
}
