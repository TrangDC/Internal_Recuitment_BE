package service

import (
	"context"
	"net/http"
	"sort"
	"trec/ent"
	"trec/ent/candidatehistorycall"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/candidatenote"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/outgoingemail"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateActivityService interface {
	GetAllCandidateActivities(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateActivityFilter, freeWord *ent.CandidateActivityFreeWord,
		orderBy ent.CandidateActivityOrder) (*ent.CandidateActivityResponse, error)
}

type candidateActivitySvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewCandidateActivityService(repoRegistry repository.Repository, logger *zap.Logger) CandidateActivityService {
	return &candidateActivitySvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *candidateActivitySvcImpl) GetAllCandidateActivities(ctx context.Context, pagination *ent.PaginationInput, filter ent.CandidateActivityFilter, freeWord *ent.CandidateActivityFreeWord,
	orderBy ent.CandidateActivityOrder) (*ent.CandidateActivityResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	candidateId := uuid.MustParse(filter.CandidateID)
	var page int
	var perPage int
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	referenceModels := []models.CandidateActivityReference{}
	candidateNoteResults := []*ent.CandidateNote{}
	candidateHistoryCallResults := []*ent.CandidateHistoryCall{}
	candidateInterviewResults := []*ent.CandidateInterview{}
	candidateJobFeedbackResults := []*ent.CandidateJobFeedback{}
	outgoingEmailResults := []*ent.OutgoingEmail{}
	// get candidate activities
	candidateInterviewQuery := svc.repoRegistry.CandidateInterview().BuildQuery().
		Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.CandidateID(candidateId))).
		WithInterviewerEdges().WithCreatedByEdge()
	candidateJobFeedbackQuery := svc.repoRegistry.CandidateJobFeedback().BuildQuery().
		Where(candidatejobfeedback.HasCandidateJobEdgeWith(candidatejob.CandidateID(candidateId))).
		WithAttachmentEdges().WithCreatedByEdge()
	candidateNoteQuery := svc.repoRegistry.CandidateNote().BuildQuery().Where(candidatenote.CandidateID(candidateId))
	candidateHistoryCallQuery := svc.repoRegistry.CandidateHistoryCall().BuildQuery().Where(candidatehistorycall.CandidateID(candidateId))
	outgoingEmailQuery := svc.repoRegistry.OutgoingEmail().BuildQuery().Where(outgoingemail.CandidateID(candidateId))
	// apply permission
	svc.validPermissionGet(payload, candidateInterviewQuery, candidateJobFeedbackQuery)
	// apply filter
	if filter.FromDate != nil && filter.ToDate != nil {
		candidateInterviewQuery.Where(candidateinterview.StartFromGTE(*filter.FromDate), candidateinterview.StartFromLTE(*filter.ToDate))
		candidateJobFeedbackQuery.Where(candidatejobfeedback.CreatedAtGTE(*filter.FromDate), candidatejobfeedback.CreatedAtLTE(*filter.ToDate))
		candidateNoteQuery.Where(candidatenote.CreatedAtGTE(*filter.FromDate), candidatenote.CreatedAtLTE(*filter.ToDate))
		candidateHistoryCallQuery.Where(candidatehistorycall.CreatedAtGTE(*filter.FromDate), candidatehistorycall.CreatedAtLTE(*filter.ToDate))
		outgoingEmailQuery.Where(outgoingemail.CreatedAtGTE(*filter.FromDate), outgoingemail.CreatedAtLTE(*filter.ToDate))
	}
	// apply free word
	if freeWord != nil && freeWord.FreeWord != nil {
		candidateInterviewQuery.Where(candidateinterview.TitleContainsFold(*freeWord.FreeWord))
		candidateJobFeedbackQuery.Where(candidatejobfeedback.FeedbackContainsFold(*freeWord.FreeWord))
		candidateNoteQuery.Where(candidatenote.NameContainsFold(*freeWord.FreeWord))
		candidateHistoryCallQuery.Where(candidatehistorycall.NameContainsFold(*freeWord.FreeWord))
		outgoingEmailQuery.Where(outgoingemail.Or(
			outgoingemail.SubjectContainsFold(*freeWord.FreeWord),
			outgoingemail.ContentContainsFold(*freeWord.FreeWord),
		))
	}
	// query
	candidateInterviews, err := svc.repoRegistry.CandidateInterview().BuildList(ctx, candidateInterviewQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateJobFeedbacks, err := svc.repoRegistry.CandidateJobFeedback().BuildList(ctx, candidateJobFeedbackQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateNotes, err := svc.repoRegistry.CandidateNote().BuildList(ctx, candidateNoteQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateHistoryCalls, err := svc.repoRegistry.CandidateHistoryCall().BuildList(ctx, candidateHistoryCallQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	outgoingEmails, err := svc.repoRegistry.OutgoingEmail().BuildList(ctx, outgoingEmailQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	count := len(candidateInterviews) + len(candidateJobFeedbacks) + len(candidateNotes) + len(candidateHistoryCalls) + len(outgoingEmails)
	// end
	referenceModels = append(referenceModels, lo.Map(candidateInterviews, func(entity *ent.CandidateInterview, index int) models.CandidateActivityReference {
		return models.CandidateActivityReference{
			Id:        entity.ID,
			CreatedAt: entity.CreatedAt,
		}
	})...)
	referenceModels = append(referenceModels, lo.Map(candidateJobFeedbacks, func(entity *ent.CandidateJobFeedback, index int) models.CandidateActivityReference {
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
	referenceModels = append(referenceModels, lo.Map(outgoingEmails, func(entity *ent.OutgoingEmail, index int) models.CandidateActivityReference {
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
		switch {
		case start > count:
			return nil, nil
		case start <= count && end > count:
			referenceModels = referenceModels[start:]
		default:
			referenceModels = referenceModels[start:end]
		}
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
		candidateJobFeedbackResults = append(candidateJobFeedbackResults, lo.Filter(candidateJobFeedbacks, func(entity *ent.CandidateJobFeedback, index int) bool {
			return entity.ID == referenceModel.Id
		})...)
		outgoingEmailResults = append(outgoingEmailResults, lo.Filter(outgoingEmails, func(entity *ent.OutgoingEmail, index int) bool {
			return entity.ID == referenceModel.Id
		})...)
	}
	return &ent.CandidateActivityResponse{
		Data: &ent.CandidateActivity{
			CandidateNotes:        candidateNoteResults,
			CandidateHistoryCalls: candidateHistoryCallResults,
			CandidateInterviews:   candidateInterviewResults,
			CandidateJobFeedbacks: candidateJobFeedbackResults,
			OutgoingEmails:        outgoingEmailResults,
			Total:                 count,
		}}, nil
}

// permission
func (svc candidateActivitySvcImpl) validPermissionGet(payload *middleware.Payload, candidateInterviewQuery *ent.CandidateInterviewQuery, candidateJobFeedbackQuery *ent.CandidateJobFeedbackQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForTeam {
		candidateInterviewQuery.Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HasHiringTeamEdgeWith(
			hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))),
		))))
		candidateJobFeedbackQuery.Where(candidatejobfeedback.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HasHiringTeamEdgeWith(
			hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))),
		))))
	}
	if payload.ForOwner {
		candidateInterviewQuery.Where(candidateinterview.HasInterviewerEdgesWith(
			user.IDEQ(payload.UserID), user.DeletedAtIsNil(),
		))
		candidateJobFeedbackQuery.Where(candidatejobfeedback.CreatedByEQ(payload.UserID))
	}
}

// Path: service/candidate_activity.service.go
