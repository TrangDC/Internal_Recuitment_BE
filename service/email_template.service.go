package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/emailtemplate"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type EmailTemplateService interface {
	// mutation
	CreateEmailTemplate(ctx context.Context, input ent.NewEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error)
	UpdateEmailTemplate(ctx context.Context, emailTpId uuid.UUID, input ent.UpdateEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error)
	UpdateEmailTemplateStatus(ctx context.Context, emailTpId uuid.UUID, input ent.UpdateEmailTemplateStatusInput, note string) error
	DeleteEmailTemplate(ctx context.Context, emailTpId uuid.UUID, note string) error
	// query
	GetEmailTemplate(ctx context.Context, emailTpId uuid.UUID) (*ent.EmailTemplateResponse, error)
	GetEmailTemplates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.EmailTemplateFreeWord,
		filter *ent.EmailTemplateFilter, orderBy ent.EmailTemplateOrder) (*ent.EmailTemplateResponseGetAll, error)
	GetAllEmailTemplateKeyword(filter ent.EmailTemplateKeywordFilter) (*ent.GetEmailTemplateKeywordResponse, error)
}

type emailtemplateSvcImpl struct {
	userSvcImpl  UserService
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewEmailTemplateService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) EmailTemplateService {
	return &emailtemplateSvcImpl{
		userSvcImpl:  NewUserService(repoRegistry, dtoRegistry, logger),
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

func (svc *emailtemplateSvcImpl) CreateEmailTemplate(ctx context.Context, input ent.NewEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error) {
	var result *ent.EmailTemplate
	var roleIds []uuid.UUID
	var err error
	var record *ent.EmailTemplate
	// err = svc.repoRegistry.EmailTemplate().ValidKeywordInput(input.Subject, input.Content, input.Event)
	// if err != nil {
	// 	return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	// }
	event, err := svc.repoRegistry.EmailEvent().GetEmailEvent(ctx, uuid.MustParse(input.EventID))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.EmailTemplate().ValidSentToAction(event, input.SendTo)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	sendTo := lo.Map(input.SendTo, func(s ent.EmailTemplateSendTo, index int) string {
		return s.String()
	})
	if len(sendTo) == 0 {
		return nil, util.WrapGQLError(ctx, "model.email_template.validation.send_to_required", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	roleIds = lo.Map(input.RoleIds, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	if len(roleIds) == 0 && lo.Contains(sendTo, string(ent.EmailTemplateSendToRole)) {
		return nil, fmt.Errorf("model.email_template.validation.select_role_required")
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.EmailTemplate().CreateEmailTemplate(ctx, input, roleIds, sendTo)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.EmailTemplate().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleEmailTemplates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.EmailTemplateResponse{
		Data: result,
	}, nil
}

func (svc *emailtemplateSvcImpl) UpdateEmailTemplate(ctx context.Context, emailTpId uuid.UUID, input ent.UpdateEmailTemplateInput, note string) (*ent.EmailTemplateResponse, error) {
	var roleIds []uuid.UUID
	var result *ent.EmailTemplate
	// err := svc.repoRegistry.EmailTemplate().ValidKeywordInput(input.Subject, input.Content, input.Event)
	// if err != nil {
	// 	return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	// }
	record, err := svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, emailTpId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if len(input.RoleIds) != 0 {
		roleIds = lo.Map(input.RoleIds, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
	}
	sendTo := lo.Map(input.SendTo, func(s ent.EmailTemplateSendTo, index int) string {
		return s.String()
	})
	if len(sendTo) == 0 {
		return nil, util.WrapGQLError(ctx, "model.email_template.validation.send_to_required", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	newRoleIds, removeRoleIds := svc.updateRoleIds(record, roleIds)
	if (len(roleIds)-len(removeRoleIds)+len(newRoleIds) == 0) && lo.Contains(sendTo, string(ent.EmailTemplateSendToRole)) {
		return nil, fmt.Errorf("model.email_template.validation.select_role_required")
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.EmailTemplate().UpdateEmailTemplate(ctx, record, input, newRoleIds, removeRoleIds, sendTo)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, emailTpId)
	jsonString, err := svc.dtoRegistry.EmailTemplate().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, emailTpId, audittrail.ModuleEmailTemplates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.EmailTemplateResponse{
		Data: result,
	}, nil
}

func (svc emailtemplateSvcImpl) UpdateEmailTemplateStatus(ctx context.Context, emailTpId uuid.UUID, input ent.UpdateEmailTemplateStatusInput, note string) error {
	emailTp, err := svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, emailTpId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.EmailTemplate().UpdateEmailTemplateStatus(ctx, emailTp, input)
		return err
	})
	result, _ := svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, emailTpId)
	jsonString, err := svc.dtoRegistry.EmailTemplate().AuditTrailUpdate(emailTp, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, emailTpId, audittrail.ModuleEmailTemplates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *emailtemplateSvcImpl) DeleteEmailTemplate(ctx context.Context, emailTpId uuid.UUID, note string) error {
	emailTp, err := svc.repoRegistry.EmailTemplate().GetEmailTemplate(ctx, emailTpId)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	roleIds := lo.Map(emailTp.Edges.RoleEdges, func(item *ent.Role, index int) uuid.UUID {
		return item.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.EmailTemplate().DeleteEmailTemplate(ctx, emailTp, roleIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.EmailTemplate().AuditTrailDelete(emailTp)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, emailTpId, audittrail.ModuleEmailTemplates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *emailtemplateSvcImpl) GetEmailTemplate(ctx context.Context, emailTpId uuid.UUID) (*ent.EmailTemplateResponse, error) {
	query := svc.repoRegistry.EmailTemplate().BuildQuery()
	emailTp, err := svc.repoRegistry.EmailTemplate().BuildGetOne(ctx, query.Where(emailtemplate.IDEQ(emailTpId)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.EmailTemplateResponse{
		Data: emailTp,
	}, nil
}

func (svc *emailtemplateSvcImpl) GetEmailTemplates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.EmailTemplateFreeWord, filter *ent.EmailTemplateFilter, orderBy ent.EmailTemplateOrder) (*ent.EmailTemplateResponseGetAll, error) {
	var result *ent.EmailTemplateResponseGetAll
	var edges []*ent.EmailTemplateEdge
	var page int
	var perPage int
	var emailTps []*ent.EmailTemplate
	var count int
	var err error
	query := svc.repoRegistry.EmailTemplate().BuildQuery()
	svc.freeWord(query, freeWord)
	svc.filter(ctx, query, filter)
	count, err = svc.repoRegistry.EmailTemplate().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(strings.ToLower(orderBy.Field.String()))
	if orderBy.Direction == ent.OrderDirectionAsc {
		order = ent.Asc(strings.ToLower(orderBy.Field.String()))
	}
	query = query.Order(order)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	emailTps, err = svc.repoRegistry.EmailTemplate().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(emailTps, func(entity *ent.EmailTemplate, index int) *ent.EmailTemplateEdge {
		return &ent.EmailTemplateEdge{
			Node: entity,
			Cursor: ent.Cursor{
				Value: entity.ID.String(),
			},
		}
	})
	result = &ent.EmailTemplateResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *emailtemplateSvcImpl) GetAllEmailTemplateKeyword(filter ent.EmailTemplateKeywordFilter) (*ent.GetEmailTemplateKeywordResponse, error) {
	result := &ent.EmailTemplateKeyword{
		General:      models.GeneralEmailTpKeywordJson,
		HiringTeam:   models.HiringTeamEmailTpKeywordJson,
		HiringJob:    models.HiringJobEmailTpKeywordJson,
		Candidate:    models.CandidateEmailTpKeywordJson,
		CandidateJob: models.CandidateAppEmailTpKeywordJson,
		Link:         models.LinkEmailTpKeywordCandidateJson,
	}
	if ent.EmailTemplateInterviewEventEnum.IsValid(ent.EmailTemplateInterviewEventEnum(filter.Event)) {
		result.Interview = models.InterviewEmailTpKeywordJson
		result.Link = append(result.Link, models.LinkEmailTpKeywordInterviewJson...)
	}
	return &ent.GetEmailTemplateKeywordResponse{
		Data: result,
	}, nil
}

// common function
func (svc *emailtemplateSvcImpl) freeWord(query *ent.EmailTemplateQuery, input *ent.EmailTemplateFreeWord) {
	if input != nil {
		if input.Subject != nil {
			query.Where(emailtemplate.SubjectContainsFold(*input.Subject))
		}
	}
}

func (svc *emailtemplateSvcImpl) filter(ctx context.Context, query *ent.EmailTemplateQuery, input *ent.EmailTemplateFilter) {
	if input != nil {
		if input.EventIds != nil {
			eventIDs := lo.Map(input.EventIds, func(idStr *string, _ int) uuid.UUID {
				return uuid.MustParse(*idStr)
			})
			query.Where(emailtemplate.EventIDIn(eventIDs...))
		}
		if input.Status != nil {
			query.Where(emailtemplate.StatusEQ(emailtemplate.Status(*input.Status)))
		}
		if input.SendTo != nil && len(input.SendTo) != 0 {
			emailTpIds := []uuid.UUID{}
			queryString := "SELECT id FROM email_templates WHERE "
			for i, reason := range input.SendTo {
				queryString += "send_to @> '[\"" + reason.String() + "\"]'::jsonb"
				if i != len(input.SendTo)-1 {
					queryString += " AND "
				}
			}
			queryString += ";"
			rows, _ := query.QueryContext(ctx, queryString)
			if rows != nil {
				for rows.Next() {
					var id uuid.UUID
					rows.Scan(&id)
					emailTpIds = append(emailTpIds, id)
				}
				query.Where(emailtemplate.IDIn(emailTpIds...))
			} else {
				query.Where(emailtemplate.IDEQ(uuid.Nil))
			}
		}
	}
}

func (svc *emailtemplateSvcImpl) updateRoleIds(record *ent.EmailTemplate, roleIds []uuid.UUID) ([]uuid.UUID, []uuid.UUID) {
	var newRoleIds []uuid.UUID
	var removeRoleIds []uuid.UUID
	currentRoleIds := lo.Map(record.Edges.RoleEdges, func(item *ent.Role, index int) uuid.UUID {
		return item.ID
	})
	newRoleIds = lo.Filter(roleIds, func(roleId uuid.UUID, index int) bool {
		return !lo.Contains(currentRoleIds, roleId)
	})
	removeRoleIds = lo.Filter(currentRoleIds, func(roleId uuid.UUID, index int) bool {
		return !lo.Contains(roleIds, roleId)
	})
	return newRoleIds, removeRoleIds
}

// Path: service/emailTp.service.go
