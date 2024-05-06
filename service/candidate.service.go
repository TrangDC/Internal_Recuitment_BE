package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateService interface {
	// mutation
	CreateCandidate(ctx context.Context, input *ent.NewCandidateInput) (*ent.CandidateResponse, error)
	UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID) (*ent.CandidateResponse, error)
	DeleteCandidate(ctx context.Context, id uuid.UUID) error
	SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool) error
	// query
	GetCandidate(ctx context.Context, id uuid.UUID) (*ent.CandidateResponse, error)
	GetCandidates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord, filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error)
}
type candidateSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewCandidateService(repoRegistry repository.Repository, logger *zap.Logger) CandidateService {
	return &candidateSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *candidateSvcImpl) CreateCandidate(ctx context.Context, input *ent.NewCandidateInput) (*ent.CandidateResponse, error) {
	var candidate *ent.Candidate
	err := svc.repoRegistry.Candidate().ValidEmail(ctx, uuid.Nil, input.Email)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidate, err = repoRegistry.Candidate().CreateCandidate(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.Candidate().GetCandidate(ctx, candidate.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID) (*ent.CandidateResponse, error) {
	candidate, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.Candidate().ValidEmail(ctx, id, input.Email)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidate, err = repoRegistry.Candidate().UpdateCandidate(ctx, candidate, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool) error {
	candidate, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return svc.repoRegistry.Candidate().SetBlackListCandidate(ctx, candidate, isBlackList)
}

func (svc *candidateSvcImpl) GetCandidate(ctx context.Context, id uuid.UUID) (*ent.CandidateResponse, error) {
	result, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) DeleteCandidate(ctx context.Context, id uuid.UUID) error {
	candidate, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.Candidate().DeleteCandidate(ctx, candidate)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *candidateSvcImpl) GetCandidates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord, filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error) {
	var result *ent.CandidateResponseGetAll
	var edges []*ent.CandidateEdge
	var page int
	var perPage int
	query := svc.repoRegistry.Candidate().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.Candidate().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(candidate.FieldCreatedAt)
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
	candidates, err := svc.repoRegistry.Candidate().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidates, func(candidate *ent.Candidate, index int) *ent.CandidateEdge {
		return &ent.CandidateEdge{
			Node: candidate,
			Cursor: ent.Cursor{
				Value: candidate.ID.String(),
			},
		}
	})
	result = &ent.CandidateResponseGetAll{
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
func (svc *candidateSvcImpl) freeWord(candidateQuery *ent.CandidateQuery, input *ent.CandidateFreeWord) {
	predicate := []predicate.Candidate{}
	if input != nil {
		if input.Name != nil {
			predicate = append(predicate, candidate.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
		if input.Email != nil {
			predicate = append(predicate, candidate.EmailContainsFold(strings.TrimSpace(*input.Email)))
		}
		if input.Phone != nil {
			predicate = append(predicate, candidate.PhoneContainsFold(strings.TrimSpace(*input.Phone)))
		}
	}
	if len(predicate) > 0 {
		candidateQuery.Where(candidate.Or(predicate...))
	}
}

func (svc *candidateSvcImpl) filter(candidateQuery *ent.CandidateQuery, input *ent.CandidateFilter) {
	if input != nil {
		if input.Name != nil {
			candidateQuery.Where(candidate.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.Email != nil {
			candidateQuery.Where(candidate.EmailEqualFold(strings.TrimSpace(*input.Email)))
		}
		if input.Phone != nil {
			candidateQuery.Where(candidate.PhoneEqualFold(strings.TrimSpace(*input.Phone)))
		}
		if input.DobFromDate != nil && input.DobToDate != nil {
			candidateQuery.Where(candidate.DobGTE(*input.DobFromDate), candidate.DobLTE(*input.DobToDate))
		}
		if input.IsBlackList != nil {
			candidateQuery.Where(candidate.IsBlacklist(*input.IsBlackList))
		}
		if input.FromDate != nil && input.ToDate != nil {
			candidateQuery.Where(candidate.CreatedAtGTE(*input.FromDate), candidate.CreatedAtLTE(*input.ToDate))
		}
	}
}
