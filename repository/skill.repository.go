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

type SkillRepository interface {
	// mutation
	CreateSkill(ctx context.Context, input ent.NewSkillInput) (*ent.Skill, error)
	UpdateSkill(ctx context.Context, record *ent.Skill, input ent.UpdateSkillInput) (*ent.Skill, error)
	DeleteSkill(ctx context.Context, record *ent.Skill) error
	DeleteBulkSkill(ctx context.Context, skillTypeId uuid.UUID) error

	// query
	GetSkill(ctx context.Context, skillId uuid.UUID) (*ent.Skill, error)
	BuildQuery() *ent.SkillQuery
	BuildBaseQuery() *ent.SkillQuery
	BuildCount(ctx context.Context, query *ent.SkillQuery) (int, error)
	BuildList(ctx context.Context, query *ent.SkillQuery) ([]*ent.Skill, error)

	// common function
	ValidName(ctx context.Context, skillId uuid.UUID, name string) (error, error)
}

type skillRepoImpl struct {
	client *ent.Client
}

func NewSkillRepository(client *ent.Client) SkillRepository {
	return &skillRepoImpl{
		client: client,
	}
}

// Base function
func (rps skillRepoImpl) BuildCreate() *ent.SkillCreate {
	return rps.client.Skill.Create().SetUpdatedAt(time.Now().UTC()).SetCreatedAt(time.Now().UTC())
}

func (rps skillRepoImpl) BuildUpdate() *ent.SkillUpdate {
	return rps.client.Skill.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps skillRepoImpl) BuildDelete() *ent.SkillUpdate {
	return rps.client.Skill.Update().SetDeletedAt(time.Now().UTC())
}

func (rps skillRepoImpl) BuildQuery() *ent.SkillQuery {
	return rps.client.Skill.Query().Where(skill.DeletedAtIsNil(), skill.HasSkillTypeEdgeWith(
		skilltype.DeletedAtIsNil(),
	)).WithSkillTypeEdge()
}

func (rps skillRepoImpl) BuildBaseQuery() *ent.SkillQuery {
	return rps.client.Skill.Query().Where(skill.DeletedAtIsNil())
}

func (rps skillRepoImpl) GetOneSkill(ctx context.Context, query *ent.SkillQuery) (*ent.Skill, error) {
	return query.First(ctx)
}

func (rps skillRepoImpl) BuildGet(ctx context.Context, query *ent.SkillQuery) (*ent.Skill, error) {
	return query.First(ctx)
}

func (rps skillRepoImpl) BuildList(ctx context.Context, query *ent.SkillQuery) ([]*ent.Skill, error) {
	return query.All(ctx)
}

func (rps skillRepoImpl) BuildCount(ctx context.Context, query *ent.SkillQuery) (int, error) {
	return query.Count(ctx)
}

func (rps skillRepoImpl) BuildExist(ctx context.Context, query *ent.SkillQuery) (bool, error) {
	return query.Exist(ctx)
}

func (rps *skillRepoImpl) BuildUpdateOne(ctx context.Context, model *ent.Skill) *ent.SkillUpdateOne {
	return model.Update().SetUpdatedAt(time.Now().UTC())
}

func (rps *skillRepoImpl) BuildSaveUpdateOne(ctx context.Context, update *ent.SkillUpdateOne) (*ent.Skill, error) {
	return update.Save(ctx)
}

// mutation
func (rps *skillRepoImpl) CreateSkill(ctx context.Context, input ent.NewSkillInput) (*ent.Skill, error) {
	create := rps.BuildCreate().
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetSkillTypeID(uuid.MustParse(*input.SkillTypeID))
	return create.Save(ctx)
}

func (rps *skillRepoImpl) UpdateSkill(ctx context.Context, record *ent.Skill, input ent.UpdateSkillInput) (*ent.Skill, error) {
	update := rps.BuildUpdateOne(ctx, record).
		SetName(strings.TrimSpace(input.Name)).
		SetDescription(strings.TrimSpace(input.Description)).
		SetSkillTypeID(uuid.MustParse(*input.SkillTypeID))
	return update.Save(ctx)
}

func (rps *skillRepoImpl) DeleteSkill(ctx context.Context, record *ent.Skill) error {
	_, err := rps.BuildUpdateOne(ctx, record).SetDeletedAt(time.Now().UTC()).Save(ctx)
	return err
}

func (rps skillRepoImpl) DeleteBulkSkill(ctx context.Context, skillTypeId uuid.UUID) error {
	_, err := rps.BuildDelete().Where(skill.SkillTypeIDEQ(skillTypeId)).Save(ctx)
	return err
}

// query
func (rps *skillRepoImpl) GetSkill(ctx context.Context, skillId uuid.UUID) (*ent.Skill, error) {
	return rps.BuildQuery().Where(skill.IDEQ(skillId)).First(ctx)
}

// common function
func (rps *skillRepoImpl) ValidName(ctx context.Context, skillId uuid.UUID, name string) (error, error) {
	query := rps.BuildQuery().Where(skill.NameEqualFold(strings.TrimSpace(name)))
	if skillId != uuid.Nil {
		query = query.Where(skill.IDNEQ(skillId))
	}
	isExist, err := rps.BuildExist(ctx, query)
	if err != nil {
		return nil, err
	}
	if isExist {
		return fmt.Errorf("model.skills.validation.name_exist"), nil
	}
	return nil, nil
}
