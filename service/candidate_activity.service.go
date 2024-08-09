package service

import (
	"context"
	"net/http"
	"sort"
	"trec/ent"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatenote"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateActivityServiceService interface {
	GetAllCandidateActivities(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateActivityFilter, freeWord *ent.CandidateActivityFreeWord,
		orderBy *ent.CandidateActivityOrder) (*ent.CandidateActivityResponse, error)
}

type candidateActivitySvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewCandidateActivityServiceService(repoRegistry repository.Repository, logger *zap.Logger) CandidateActivityServiceService {
	return &candidateActivitySvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *candidateActivitySvcImpl) GetAllCandidateActivities(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateActivityFilter, freeWord *ent.CandidateActivityFreeWord,
	orderBy *ent.CandidateActivityOrder) (*ent.CandidateActivityResponse, error) {
	candidateId := uuid.MustParse(filter.CandidateID)
	var page int
	var perPage int
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	var referenceModels []models.CandidateActivityReference
	var candidateNoteResults []*ent.CandidateNote
	var candidateHistoryCallResults []*ent.CandidateHistoryCall
	var candidateInterviewResults []*ent.CandidateInterview
	// get candidate activities
	candidateInterviewQuery := svc.repoRegistry.CandidateInterview().BuildQuery().Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.CandidateID(candidateId)))
	candidateInterviews, err := svc.repoRegistry.CandidateInterview().BuildList(ctx, candidateInterviewQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateNoteQuery := svc.repoRegistry.CandidateNote().BuildQuery().Where(candidatenote.CandidateID(candidateId))
	candidateNotes, err := svc.repoRegistry.CandidateNote().BuildList(ctx, candidateNoteQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateHistoryCallQuery := svc.repoRegistry.CandidateHistoryCall().BuildQuery().Where(candidatehistorycall.CandidateID(candidateId))
	candidateHistoryCalls, err := svc.repoRegistry.CandidateHistoryCall().BuildList(ctx, candidateHistoryCallQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	count := len(candidateInterviews) + len(candidateNotes) + len(candidateHistoryCalls)
	// end
	referenceModels = append(referenceModels, lo.Map(candidateInterviews, func(entity *ent.CandidateInterview, index int) models.CandidateActivityReference {
		return models.CandidateActivityReference{
			Id:        entity.ID,
			CreatedAt: entity.CreatedAt,
		}
	})...)
	referenceModels = append(referenceModels, lo.Map(candidateNotes, func(entity *ent.CandidateNote, index int) models.CandidateActivityReference {
		return models.CandidateActivityReference{
			Id:        entity.ID,
			CreatedAt: entity.CreatedAt,
		}
	})...)
	referenceModels = append(referenceModels, lo.Map(candidateHistoryCalls, func(entity *ent.CandidateHistoryCall, index int) models.CandidateActivityReference {
		return models.CandidateActivityReference{
			Id:        entity.ID,
			CreatedAt: entity.CreatedAt,
		}
	})...)
	if orderBy.Direction == ent.OrderDirectionDesc {
		sort.Slice(referenceModels, func(i, j int) bool {
			return referenceModels[i].CreatedAt.After(referenceModels[j].CreatedAt)
		})
	} else {
		sort.Slice(referenceModels, func(i, j int) bool {
			return referenceModels[i].CreatedAt.Before(referenceModels[j].CreatedAt)
		})
	}
	if page != 0 && perPage != 0 {
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(referenceModels) {
			return nil, nil
		}
		if start <= len(referenceModels) && end > len(referenceModels) {
			referenceModels = referenceModels[start:]
		}
		referenceModels = referenceModels[start:end]
	}
	for _, referenceModel := range referenceModels {
		candidateNoteResults = append(candidateNoteResults, lo.Filter(candidateNotes, func(entity *ent.CandidateNote, index int) bool {
			return entity.ID == referenceModel.Id
		})...)
		candidateHistoryCallResults = append(candidateHistoryCallResults, lo.Filter(candidateHistoryCalls, func(entity *ent.CandidateHistoryCall, index int) bool {
			return entity.ID == referenceModel.Id
		})...)
		candidateInterviewResults = append(candidateInterviewResults, lo.Filter(candidateInterviews, func(entity *ent.CandidateInterview, index int) bool {
			return entity.ID == referenceModel.Id
		})...)
	}
	return &ent.CandidateActivityResponse{
		Data: &ent.CandidateActivity{
			CandidateNotes:        candidateNoteResults,
			CandidateHistoryCalls: candidateHistoryCallResults,
			CandidateInterviews:   candidateInterviewResults,
			Total:                 count,
		}}, nil
}