package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateInterviewService interface {
	// mutation
	CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput) (*ent.CandidateInterviewResponse, error)
	DeleteCandidateInterview(ctx context.Context, id uuid.UUID) error
	UpdateCandidateInterview(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewInput) (*ent.CandidateInterviewResponse, error)

	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterviewResponse, error)
	GetCandidateInterviews(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter ent.CandidateInterviewFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error)
}

type candidateInterviewSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	logger        *zap.Logger
}

func NewCandidateInterviewService(repoRegistry repository.Repository, logger *zap.Logger) CandidateInterviewService {
	return &candidateInterviewSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		logger:        logger,
	}
}

func (svc *candidateInterviewSvcImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput) (*ent.CandidateInterviewResponse, error) {
	var candidateInterview *ent.CandidateInterview
	var memberIds []uuid.UUID
	err := svc.repoRegistry.CandidateInterview().ValidateInput(ctx, uuid.MustParse(input.CandidateJobID), input.Title, uuid.Nil)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds = lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().CreateCandidateInterview(ctx, input, memberIds)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc candidateInterviewSvcImpl) UpdateCandidateInterview(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewInput) (*ent.CandidateInterviewResponse, error) {
	var candidateInterview *ent.CandidateInterview
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	newMemberIds, removeMemberIds := svc.updateMembers(record, memberIds)
	err = svc.repoRegistry.CandidateInterview().ValidateInput(ctx, uuid.Nil, input.Title, uuid.Nil)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().UpdateCandidateInterview(ctx, record, input, newMemberIds, removeMemberIds)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc *candidateInterviewSvcImpl) GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterviewResponse, error) {
	result, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc *candidateInterviewSvcImpl) DeleteCandidateInterview(ctx context.Context, id uuid.UUID) error {
	candidateInterview, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	memberIds := lo.Map(candidateInterview.Edges.InterviewerEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateInterview().DeleteCandidateInterview(ctx, candidateInterview, memberIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *candidateInterviewSvcImpl) GetCandidateInterviews(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter ent.CandidateInterviewFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error) {
	var result *ent.CandidateInterviewResponseGetAll
	var edges []*ent.CandidateInterviewEdge
	var page int
	var perPage int
	candidate, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, uuid.MustParse(filter.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_not_found", http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.CandidateInterview().BuildQuery().Where(
		candidateinterview.CandidateJobIDEQ(uuid.MustParse(filter.CandidateJobID)),
		candidateinterview.CandidateJobStatusEQ(candidateinterview.CandidateJobStatus(candidate.Status.String())),
	)
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateInterview().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidatejob.FieldCreatedAt)
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
	candidateInterviews, err := svc.repoRegistry.CandidateInterview().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateInterviews, func(candidateInterview *ent.CandidateInterview, index int) *ent.CandidateInterviewEdge {
		return &ent.CandidateInterviewEdge{
			Node: candidateInterview,
			Cursor: ent.Cursor{
				Value: candidateInterview.ID.String(),
			},
		}
	})
	result = &ent.CandidateInterviewResponseGetAll{
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
func (svc *candidateInterviewSvcImpl) freeWord(candidateInterviewQuery *ent.CandidateInterviewQuery, input *ent.CandidateInterviewFreeWord) {
	var predidacate []predicate.CandidateInterview
	if input != nil {
		if input.Title != nil {
			predidacate = append(predidacate, candidateinterview.TitleContainsFold(strings.TrimSpace(*input.Title)))
		}
		if input.Description != nil {
			predidacate = append(predidacate, candidateinterview.DescriptionContainsFold(strings.TrimSpace(*input.Description)))
		}
	}
	if len(predidacate) > 0 {
		candidateInterviewQuery.Where(candidateinterview.Or(predidacate...))
	}
}

func (svc *candidateInterviewSvcImpl) filter(candidateInterviewQuery *ent.CandidateInterviewQuery, input ent.CandidateInterviewFilter) {
	if input.InterviewDate != nil {
		candidateInterviewQuery.Where(candidateinterview.InterviewDateEQ(*input.InterviewDate))
	}
	if input.StartFrom != nil && input.EndAt != nil {
		candidateInterviewQuery.Where(candidateinterview.And(candidateinterview.StartFromGTE(*input.StartFrom), candidateinterview.EndAtLTE(*input.EndAt)))
	}
}

// common function
func (svc *candidateInterviewSvcImpl) updateMembers(record *ent.CandidateInterview, memberIds []uuid.UUID) ([]uuid.UUID, []uuid.UUID) {
	var newMemberIds []uuid.UUID
	var removeMemberIds []uuid.UUID
	currentMemberIds := lo.Map(record.Edges.InterviewerEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	newMemberIds = lo.Filter(memberIds, func(memberId uuid.UUID, index int) bool {
		return !lo.Contains(currentMemberIds, memberId)
	})
	removeMemberIds = lo.Filter(currentMemberIds, func(memberId uuid.UUID, index int) bool {
		return !lo.Contains(memberIds, memberId)
	})
	return newMemberIds, removeMemberIds
}
