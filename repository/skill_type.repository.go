package repository

import (
	"context"
	"fmt"
	"strings"
	"time"
	"trec/ent"
	"trec/ent/skill"
	"trec/ent/skilltype"

	"github.com/google/uuid"
)

type SkillTypeRepository interface {
	CreateSkillType(ctx context.Context, input ent.NewSkillTypeInput) (*ent.SkillType, error)
	UpdateSkillType(ctx context.Context, model *ent.SkillType, input ent.UpdateSkillTypeInput) (*ent.SkillType, error)
	DeleteSkillType(ctx context.Context, model *ent.SkillType) (*ent.SkillType, error)

	// query
	GetSkillType(ctx context.Context, id uuid.UUID) (*ent.SkillType, error)
	BuildQuery() *ent.SkillTypeQuery
	BuildCount(ctx context.Context, query *ent.SkillTypeQuery) (int, error)
	BuildList(ctx context.Context, query *ent.SkillTypeQuery) ([]*ent.SkillType, error)

	// common function
	ValidName(ctx context.Context, teamId uuid.UUID, name string) (error, error)
}

type skillTypeRepoImpl struct {
	client *ent.Client
}

func NewSkillTypeRepository(client *ent.Client) SkillTypeRepository {
	return &skillTypeRepoImpl{
		client: client,
	}
}

// Base functions
func (rps *skillTypeRepoImpl) BuildCreate() *ent.SkillTypeCreate {
	return rps.client.SkillType.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps *skillTypeRepoImpl) BuildUpdate() *ent.SkillTypeUpdate {
	return rps.client.SkillType.Update().SetUpdatedAt(time.Now())
}

func (rps *skillTypeRepoImpl) BuildDelete() *ent.SkillTypeUpdate {
	return rps.client.SkillType.Update().SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
}

func (rps *skillTypeRepoImpl) BuildQuery() *ent.SkillTypeQuery {
	return rps.client.SkillType.Query().Where(skilltype.DeletedAtIsNil()).WithSkillEdges(
		func(query *ent.SkillQuery) {
			query.Where(skill.DeletedAtIsNil()).Order(ent.Asc(skill.FieldCreatedAt))
		},
	)
}

func (rps *skillTypeRepoImpl) BuildGet(ctx context.Context, query *ent.SkillTypeQuery) (*ent.SkillType, error) {
	return query.First(ctx)
}

func (rps *skillTypeRepoImpl) BuildList(ctx context.Context, query *ent.SkillTypeQuery) ([]*ent.SkillType, error) {
	return query.All(ctx)
}

func (rps *skillTypeRepoImpl) BuildCount(ctx context.Context, query *ent.SkillTypeQuery) (int, error) {
	return query.Count(ctx)
}

func (rps *skillTypeRepoImpl) BuildExist(ctx context.Context, query *ent.SkillTypeQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *skillTypeRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.SkillType) *ent.SkillTypeUpdateOne {
	return model.Update().SetUpdatedAt(time.Now())
}

func (rps *skillTypeRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.SkillTypeUpdateOne) (*ent.SkillType, error) {
	return update.Save(ctx)
}

// mutation
func (rps *skillTypeRepoImpl) CreateSkillType(ctx context.Context, input ent.NewSkillTypeInput) (*ent.SkillType, error) {
	return rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).Save(ctx)
}

func (rps *skillTypeRepoImpl) UpdateSkillType(ctx context.Context, model *ent.SkillType, input ent.UpdateSkillTypeInput) (*ent.SkillType, error) {
	return rps.BuildUpdateOne(ctx, model).
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).Save(ctx)
}

func (rps *skillTypeRepoImpl) DeleteSkillType(ctx context.Context, model *ent.SkillType) (*ent.SkillType, error) {
	update := rps.BuildUpdateOne(ctx, model).SetDeletedAt(time.Now()).SetUpdatedAt(time.Now())
	return update.Save(ctx)
}

// query
func (rps *skillTypeRepoImpl) GetSkillType(ctx context.Context, id uuid.UUID) (*ent.SkillType, error) {
	query := rps.BuildQuery().Where(skilltype.IDEQ(id))
	return rps.BuildGet(ctx, query)
}

// common function
func (rps *skillTypeRepoImpl) ValidName(ctx context.Context, skillTypeId uuid.UUID, name string) (error, error) {
	query := rps.BuildQuery().Where(skilltype.NameEqualFold(strings.TrimSpace(name)))
	if skillTypeId != uuid.Nil {
		query = query.Where(skilltype.IDNEQ(skillTypeId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.skill_types.validation.name_exist"), nil
	}
	return nil, nil
}
