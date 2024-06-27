package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/predicate"
	"trec/ent/skilltype"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type SkillTypeService interface {
	// mutation
	CreateSkillType(ctx context.Context, input ent.NewSkillTypeInput, note string) (*ent.SkillTypeResponse, error)
	UpdateSkillType(ctx context.Context, skillTypeId uuid.UUID, input ent.UpdateSkillTypeInput, note string) (*ent.SkillTypeResponse, error)
	DeleteSkillType(ctx context.Context, skillTypeId uuid.UUID, note string) error
	// query
	GetSkillType(ctx context.Context, skillTypeId uuid.UUID) (*ent.SkillTypeResponse, error)
	GetSkillTypes(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillTypeFreeWord,
		filter *ent.SkillTypeFilter, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillTypeFreeWord,
		filter *ent.SkillTypeFilter, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeSelectionResponseGetAll, error)
}

type skillTypeSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewSkillTypeService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) SkillTypeService {
	return &skillTypeSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

func (svc *skillTypeSvcImpl) CreateSkillType(ctx context.Context, input ent.NewSkillTypeInput, note string) (*ent.SkillTypeResponse, error) {
	var result *ent.SkillType
	errString, err := svc.repoRegistry.SkillType().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.SkillType().CreateSkillType(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.SkillType().GetSkillType(ctx, result.ID)
	jsonString, err := svc.dtoRegistry.SkillType().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleSkillTypes, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.SkillTypeResponse{
		Data: result,
	}, nil
}

func (svc *skillTypeSvcImpl) UpdateSkillType(ctx context.Context, skillTypeId uuid.UUID, input ent.UpdateSkillTypeInput, note string) (*ent.SkillTypeResponse, error) {
	var result *ent.SkillType
	record, err := svc.repoRegistry.SkillType().GetSkillType(ctx, skillTypeId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	errString, err := svc.repoRegistry.SkillType().ValidName(ctx, skillTypeId, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.SkillType().UpdateSkillType(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.SkillType().GetSkillType(ctx, skillTypeId)
	jsonString, err := svc.dtoRegistry.SkillType().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, skillTypeId, audittrail.ModuleSkillTypes, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.SkillTypeResponse{
		Data: result,
	}, nil
}

func (svc *skillTypeSvcImpl) DeleteSkillType(ctx context.Context, skillTypeId uuid.UUID, note string) error {
	skillTypeRecord, err := svc.repoRegistry.SkillType().GetSkillType(ctx, skillTypeId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if len(skillTypeRecord.Edges.SkillEdges) != 0 {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, "model.skill_types.validation.cannot_delete", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.SkillType().DeleteSkillType(ctx, skillTypeRecord)
		if err != nil {
			return err
		}
		err = repoRegistry.Skill().DeleteBulkSkill(ctx, skillTypeId)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.SkillType().AuditTrailDelete(skillTypeRecord)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, skillTypeId, audittrail.ModuleSkillTypes, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *skillTypeSvcImpl) GetSkillType(ctx context.Context, skillTypeId uuid.UUID) (*ent.SkillTypeResponse, error) {
	skillTypeRecord, err := svc.repoRegistry.SkillType().GetSkillType(ctx, skillTypeId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.SkillTypeResponse{
		Data: skillTypeRecord,
	}, nil
}

func (svc skillTypeSvcImpl) GetSkillTypes(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillTypeFreeWord,
	filter *ent.SkillTypeFilter, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeResponseGetAll, error) {
	var edges []*ent.SkillTypeEdge
	skillTypes, count, page, perPage, err := svc.getAllSkillType(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(skillTypes, func(entity *ent.SkillType, index int) *ent.SkillTypeEdge {
		return &ent.SkillTypeEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.SkillTypeResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc skillTypeSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillTypeFreeWord,
	filter *ent.SkillTypeFilter, orderBy *ent.SkillTypeOrder) (*ent.SkillTypeSelectionResponseGetAll, error) {
	var edges []*ent.SkillTypeSelectionEdge
	skillTypes, count, page, perPage, err := svc.getAllSkillType(ctx, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(skillTypes, func(entity *ent.SkillType, index int) *ent.SkillTypeSelectionEdge {
		return &ent.SkillTypeSelectionEdge{
			Node: &ent.SkillTypeSelection{
				ID:   entity.ID.String(),
				Name: entity.Name,
				Skills: lo.Map(entity.Edges.SkillEdges, func(skill *ent.Skill, index int) *ent.SkillSelection {
					return &ent.SkillSelection{
						ID:   skill.ID.String(),
						Name: skill.Name,
					}
				}),
			},
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	return &ent.SkillTypeSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}, nil
}

func (svc skillTypeSvcImpl) getAllSkillType(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.SkillTypeFreeWord,
	filter *ent.SkillTypeFilter, orderBy *ent.SkillTypeOrder) ([]*ent.SkillType, int, int, int, error) {
	var page int
	var perPage int
	query := svc.repoRegistry.SkillType().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.SkillType().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(skilltype.FieldCreatedAt)
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
	skillTypes, err := svc.repoRegistry.SkillType().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return skillTypes, count, page, perPage, nil
}

// common function
func (svc *skillTypeSvcImpl) freeWord(query *ent.SkillTypeQuery, input *ent.SkillTypeFreeWord) {
	predicate := []predicate.SkillType{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, skilltype.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
	if len(predicate) > 0 {
		query.Where(skilltype.Or(predicate...))
	}
}

func (svc *skillTypeSvcImpl) filter(query *ent.SkillTypeQuery, input *ent.SkillTypeFilter) {
	if input != nil {
		if input.Name != nil {
			query.Where(skilltype.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
	}
}

// Path: service/skill_types.service.go
