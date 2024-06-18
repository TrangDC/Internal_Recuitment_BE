package service

import (
	"context"
	"net/http"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidatejobfeedback"
	"trec/ent/predicate"
	"trec/ent/user"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateJobFeedbackService interface {
	// mutation
	CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error)
	UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error)
	DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID, note string) error
	// query
	GetCandidateJobFeedback(ctx context.Context, id uuid.UUID) (*ent.CandidateJobFeedbackResponse, error)
	GetCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFeedbackFreeWord, filter *ent.CandidateJobFeedbackFilter, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error)
}
type candidateJobFeedbackSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	dtoRegistry   dto.Dto
	logger        *zap.Logger
}

func NewCandidateJobFeedbackService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateJobFeedbackService {
	return &candidateJobFeedbackSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		dtoRegistry:   dtoRegistry,
		logger:        logger,
	}
}

func (svc *candidateJobFeedbackSvcImpl) CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error) {
	var candidateJobFeedback *ent.CandidateJobFeedback
	status, err := svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, uuid.MustParse(input.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, uuid.MustParse(input.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().CreateCandidateJobFeedback(ctx, input, status)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJobFeedback.ID, attachment.RelationTypeCandidateJobFeedbacks, repoRegistry)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, candidateJobFeedback.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput, note string) (*ent.CandidateJobFeedbackResponse, error) {
	record, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	userId := ctx.Value("user_id").(uuid.UUID)
	if record.CreatedBy != userId {
		return nil, util.WrapGQLError(ctx, "model.candidate_job_feedbacks.validation.editor_not_is_owner", http.StatusBadGateway, util.ErrorFlagValidateFail)
	}
	_, err = svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, record.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		result, err := repoRegistry.CandidateJobFeedback().UpdateCandidateJobFeedback(ctx, record, input)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, result.ID, attachment.RelationTypeCandidateJobFeedbacks, repoRegistry)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID, note string) error {
	candidateJobFeedback, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	_, err = svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, candidateJobFeedback.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, candidateJobFeedback.CandidateJobID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().DeleteCandidateJobFeedback(ctx, candidateJobFeedback)
		if err != nil {
			return err
		}
		err = svc.attachmentSvc.RemoveAttachment(ctx, candidateJobFeedback.ID, repoRegistry)
		return err
	})
	jsonString, err := svc.dtoRegistry.CandidateJobFeedback().AuditTrailDelete(candidateJobFeedback)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, candidateJobFeedback.ID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return err
}

func (svc *candidateJobFeedbackSvcImpl) GetCandidateJobFeedback(ctx context.Context, id uuid.UUID) (*ent.CandidateJobFeedbackResponse, error) {
	result, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) GetCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFeedbackFreeWord, filter *ent.CandidateJobFeedbackFilter, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error) {
	var result *ent.CandidateJobFeedbackResponseGetAll
	var edges []*ent.CandidateJobFeedbackEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJobFeedback().BuildQuery().Where(candidatejobfeedback.CandidateJobID(uuid.MustParse(filter.CandidateJobID)))
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJobFeedback().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidatejobfeedback.FieldCreatedAt)
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
	candidateJobFeedbacks, err := svc.repoRegistry.CandidateJobFeedback().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateJobFeedbacks, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) *ent.CandidateJobFeedbackEdge {
		return &ent.CandidateJobFeedbackEdge{
			Node: candidateJobFeedback,
			Cursor: ent.Cursor{
				Value: candidateJobFeedback.ID.String(),
			},
		}
	})
	result = &ent.CandidateJobFeedbackResponseGetAll{
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
func (svc *candidateJobFeedbackSvcImpl) freeWord(candidateJobFeedbackQuery *ent.CandidateJobFeedbackQuery, input *ent.CandidateJobFeedbackFreeWord) {
	predicate := []predicate.CandidateJobFeedback{}
	if input != nil {
		if input.Feedback != nil {
			predicate = append(predicate, candidatejobfeedback.FeedbackContainsFold(strings.TrimSpace(*input.Feedback)))
		}
		if input.UserName != nil {
			predicate = append(predicate, candidatejobfeedback.HasCreatedByEdgeWith(
				user.NameContainsFold(strings.TrimSpace(*input.UserName)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobFeedbackQuery.Where(candidatejobfeedback.Or(predicate...))
	}
}

func (svc *candidateJobFeedbackSvcImpl) filter(candidateJobFeedbackQuery *ent.CandidateJobFeedbackQuery, input *ent.CandidateJobFeedbackFilter) {
	if input.CreatedBy != nil {
		candidateJobFeedbackQuery.Where(candidatejobfeedback.CreatedByEQ(uuid.MustParse(*input.CreatedBy)))
	}
}
