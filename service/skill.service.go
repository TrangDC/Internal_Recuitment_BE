package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/predicate"
	"trec/ent/skill"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type SkillService interface {
	// mutation
	CreateSkill(ctx context.Context, input ent.NewSkillInput) (*ent.SkillResponse, error)
	DeleteSkill(ctx context.Context, id uuid.UUID) error
	UpdateSkill(ctx context.Context, id uuid.UUID, input ent.UpdateSkillInput) (*ent.SkillResponse, error)
	// query
	GetSkill(ctx context.Context, id uuid.UUID) (*ent.SkillResponse, error)
	GetSkills(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillFreeWord, filter *ent.SkillFilter, orderBy *ent.SkillOrder) (*ent.SkillResponseGetAll, error)
}

type skillSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewSkillService(repoRegistry repository.Repository, logger *zap.Logger) SkillService {
	return &skillSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

// mutation
func (svc *skillSvcImpl) CreateSkill(ctx context.Context, input ent.NewSkillInput) (*ent.SkillResponse, error) {
	var record *ent.Skill
	stringError, err := svc.repoRegistry.Skill().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.Skill().CreateSkill(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.SkillResponse{
		Data: record,
	}, nil
}

func (svc *skillSvcImpl) DeleteSkill(ctx context.Context, id uuid.UUID) error {
	record, err := svc.repoRegistry.Skill().GetSkill(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.Skill().DeleteSkill(ctx, record)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *skillSvcImpl) UpdateSkill(ctx context.Context, id uuid.UUID, input ent.UpdateSkillInput) (*ent.SkillResponse, error) {
	var result *ent.Skill
	record, err := svc.repoRegistry.Skill().GetSkill(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	stringError, err := svc.repoRegistry.Skill().ValidName(ctx, id, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.Skill().UpdateSkill(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.SkillResponse{
		Data: result,
	}, nil
}

// query
func (svc *skillSvcImpl) GetSkill(ctx context.Context, id uuid.UUID) (*ent.SkillResponse, error) {
	record, err := svc.repoRegistry.Skill().GetSkill(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.SkillResponse{
		Data: record,
	}, nil
}

func (svc *skillSvcImpl) GetSkills(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillFreeWord, filter *ent.SkillFilter, orderBy *ent.SkillOrder) (*ent.SkillResponseGetAll, error) {
	var result *ent.SkillResponseGetAll
	var edges []*ent.SkillEdge
	var page int
	var perPage int
	query := svc.repoRegistry.Skill().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.Skill().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(skill.FieldCreatedAt)
	if orderBy != nil {
		order = ent.Desc(strings.ToLower(orderBy.Field.String()))
		if orderBy.Direction == ent.OrderDirectionAsc {
			order = ent.Asc(strings.ToLower(orderBy.Field.String()))
		}
	}
	query = query.Order(order)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	skills, err := svc.repoRegistry.Skill().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(skills, func(skill *ent.Skill, index int) *ent.SkillEdge {
		return &ent.SkillEdge{
			Node: skill,
			Cursor: ent.Cursor{
				Value: skill.ID.String(),
			},
		}
	})
	result = &ent.SkillResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

// common function
func (svc *skillSvcImpl) freeWord(skillQuery *ent.SkillQuery, input *ent.SkillFreeWord) {
	predicate := []predicate.Skill{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, skill.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
	if len(predicate) > 0 {
		skillQuery.Where(skill.Or(predicate...))
	}
}

func (svc *skillSvcImpl) filter(skillQuery *ent.SkillQuery, input *ent.SkillFilter) {
	if input != nil {
		if input.Name != nil {
			skillQuery.Where(skill.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
	}
}
