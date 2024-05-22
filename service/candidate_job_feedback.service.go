package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/attachment"
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
	CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error)
	UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error)
	// query
	GetCandidateJobFeedback(ctx context.Context, id uuid.UUID) (*ent.CandidateJobFeedbackResponse, error)
	GetCandidateJobFeedbacks(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFeedbackFreeWord, filter *ent.CandidateJobFeedbackFilter, orderBy *ent.CandidateJobFeedbackOrder) (*ent.CandidateJobFeedbackResponseGetAll, error)
	DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID) (error)
}
type candidateJobFeedbackSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	logger        *zap.Logger
}

func NewCandidateJobFeedbackService(repoRegistry repository.Repository, logger *zap.Logger) CandidateJobFeedbackService {
	return &candidateJobFeedbackSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		logger:        logger,
	}
}

func (svc *candidateJobFeedbackSvcImpl) CreateCandidateJobFeedback(ctx context.Context, input *ent.NewCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error) {
	var candidateJobFeedback *ent.CandidateJobFeedback
	err := svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, uuid.MustParse(input.CandidateJobID))
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
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().CreateCandidateJobFeedback(ctx, input)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJobFeedback.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
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
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) UpdateCandidateJobFeedback(ctx context.Context, id uuid.UUID, input *ent.UpdateCandidateJobFeedbackInput) (*ent.CandidateJobFeedbackResponse, error) {
	candidateJobFeedback, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		var err error
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().UpdateCandidateJobFeedback(ctx, candidateJobFeedback, input)
		if err != nil {
			return err
		}
		if input.Attachments != nil {
			_, err = svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJobFeedback.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
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
	return &ent.CandidateJobFeedbackResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobFeedbackSvcImpl) DeleteCandidateJobFeedback(ctx context.Context, id uuid.UUID) (error) {
	candidateJobFeedback, err := svc.repoRegistry.CandidateJobFeedback().GetCandidateJobFeedback(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJobFeedback().ValidJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJobFeedback, err = repoRegistry.CandidateJobFeedback().DeleteCandidateJobFeedback(ctx, candidateJobFeedback)
		return err
	})
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
