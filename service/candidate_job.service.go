package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/team"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateJobService interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJobResponse, error)
	UpdateCandidateJobStatus(ctx context.Context, status ent.CandidateJobStatus, id uuid.UUID) (*ent.CandidateJobResponse, error)
	DeleteCandidateJob(ctx context.Context, id uuid.UUID) error
	UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID) (*ent.CandidateJobResponse, error)

	// query
	GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error)
	GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error)
}

type candidateJobSvcImpl struct {
	attachmentSvc AttachmentService
	repoRegistry  repository.Repository
	logger        *zap.Logger
}

func NewCandidateJobService(repoRegistry repository.Repository, logger *zap.Logger) CandidateJobService {
	return &candidateJobSvcImpl{
		attachmentSvc: NewAttachmentService(repoRegistry, logger),
		repoRegistry:  repoRegistry,
		logger:        logger,
	}
}

func (svc *candidateJobSvcImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput) (*ent.CandidateJobResponse, error) {
	var candidateJob *ent.CandidateJob
	err := svc.repoRegistry.CandidateJob().ValidStatus(ctx, uuid.Nil, &input.Status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJob, err = repoRegistry.CandidateJob().CreateCandidateJob(ctx, input)
		if err != nil {
			return err
		}
		_, err := svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, candidateJob.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobStatus(ctx context.Context, status ent.CandidateJobStatus, id uuid.UUID) (*ent.CandidateJobResponse, error) {
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.CandidateJob().ValidStatus(ctx, candidateJob.ID, &status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJob, err = repoRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, candidateJob, status)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID) (*ent.CandidateJobResponse, error) {
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err := svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error) {
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) DeleteCandidateJob(ctx context.Context, id uuid.UUID) error {
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.CandidateJob().DeleteCandidateJob(ctx, candidateJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	var result *ent.CandidateJobResponseGetAll
	var edges []*ent.CandidateJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.CandidateIDEQ(uuid.MustParse(filter.CandidateID)))
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJob().BuildCount(ctx, query)
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
	candidateJobs, err := svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(candidateJobs, func(candidateJob *ent.CandidateJob, index int) *ent.CandidateJobEdge {
		return &ent.CandidateJobEdge{
			Node: candidateJob,
			Cursor: ent.Cursor{
				Value: candidateJob.ID.String(),
			},
		}
	})
	result = &ent.CandidateJobResponseGetAll{
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
func (svc *candidateJobSvcImpl) freeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobFreeWord) {
	var predidacate []predicate.CandidateJob
	if input != nil {
		if input.Team != nil {
			predidacate = append(predidacate, candidatejob.HasHiringJobWith(
				hiringjob.HasTeamEdgeWith(
					team.NameEqualFold(strings.TrimSpace(*input.Team)),
				),
			))
		}
		if input.HiringJob != nil {
			predidacate = append(predidacate, candidatejob.HasHiringJobWith(
				hiringjob.NameEqualFold(strings.TrimSpace(*input.HiringJob)),
			))
		}
	}
	if len(predidacate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predidacate...))
	}
}

func (svc *candidateJobSvcImpl) filter(candidateJobQuery *ent.CandidateJobQuery, input ent.CandidateJobFilter) {
	if input.Status != nil {
		candidateJobQuery.Where(candidatejob.StatusEQ(candidatejob.Status(*input.Status)))
	}
	if input.TeamID != nil {
		candidateJobQuery.Where(candidatejob.HasHiringJobWith(
			hiringjob.HasTeamEdgeWith(
				team.IDEQ(uuid.MustParse(*input.TeamID)),
			),
		))
	}
	if input.HiringJobID != nil {
		candidateJobQuery.Where(candidatejob.HiringJobIDEQ(uuid.MustParse(*input.HiringJobID)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
}