package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/internal/util"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type HiringJobService interface {
	// mutation
	CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput, note string) (*ent.HiringJobResponse, error)
	UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID, note string) (*ent.HiringJobResponse, error)
	UpdateHiringJobStatus(ctx context.Context, status ent.HiringJobStatus, id uuid.UUID, note string) (*ent.HiringJobResponse, error)
	DeleteHiringJob(ctx context.Context, id uuid.UUID, note string) error
	// query
	GetHiringJob(ctx context.Context, id uuid.UUID) (*ent.HiringJobResponse, error)
	GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
		filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobResponseGetAll, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
		filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobSelectionResponseGetAll, error)

	// resolved
	GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType
}
type hiringJobSvcImpl struct {
	repoRegistry repository.Repository
	dtoRegistry  dto.Dto
	logger       *zap.Logger
}

func NewHiringJobService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) HiringJobService {
	return &hiringJobSvcImpl{
		repoRegistry: repoRegistry,
		dtoRegistry:  dtoRegistry,
		logger:       logger,
	}
}

func (svc *hiringJobSvcImpl) CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput, note string) (*ent.HiringJobResponse, error) {
	var record *ent.HiringJob
	errString, err := svc.repoRegistry.HiringJob().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.HiringJob().ValidPriority(ctx, uuid.Nil, uuid.MustParse(input.TeamID), input.Priority)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if input.Amount == 0 && input.Status == ent.HiringJobStatusOpened {
		return nil, util.WrapGQLError(ctx, "model.hiring_jobs.validation.amount_neq_zero", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.HiringJob().CreateHiringJob(ctx, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().CreateAndUpdateEntitySkill(ctx, record.ID, input.EntitySkillRecords, nil, entityskill.EntityTypeHiringJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) DeleteHiringJob(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if len(record.Edges.CandidateJobEdges) > 0 {
		return util.WrapGQLError(ctx, "model.hiring_jobs.validation.candidate_job_exist", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.HiringJob().DeleteHiringJob(ctx, record)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().DeleteAllEntitySkill(ctx, record.ID)
		return err
	})
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return err
}

func (svc *hiringJobSvcImpl) UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID, note string) (*ent.HiringJobResponse, error) {
	var result *ent.HiringJob
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if input.Amount == 0 && record.Status == hiringjob.StatusOpened {
		return nil, util.WrapGQLError(ctx, "model.hiring_jobs.validation.amount_neq_zero", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err := svc.repoRegistry.HiringJob().ValidName(ctx, id, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.HiringJob().ValidPriority(ctx, id, uuid.MustParse(input.TeamID), input.Priority)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringJob().UpdateHiringJob(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().CreateAndUpdateEntitySkill(ctx, record.ID, input.EntitySkillRecords, record.Edges.HiringJobSkillEdges, entityskill.EntityTypeHiringJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) UpdateHiringJobStatus(ctx context.Context, status ent.HiringJobStatus, id uuid.UUID, note string) (*ent.HiringJobResponse, error) {
	var result *ent.HiringJob
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	candidateJobWithStatusOpen := lo.Filter(record.Edges.CandidateJobEdges, func(item *ent.CandidateJob, index int) bool {
		return ent.CandidateJobStatusAbleToClose.IsValid(ent.CandidateJobStatusAbleToClose(item.Status))
	})
	if len(candidateJobWithStatusOpen) > 0 && record.Status == hiringjob.StatusOpened && hiringjob.Status(status) == hiringjob.StatusClosed {
		return nil, util.WrapGQLError(ctx, "model.hiring_jobs.validation.candidate_job_open_exist", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, record, status)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) GetHiringJob(ctx context.Context, id uuid.UUID) (*ent.HiringJobResponse, error) {
	result, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobResponseGetAll, error) {
	var result *ent.HiringJobResponseGetAll
	var edges []*ent.HiringJobEdge
	var page int
	var perPage int
	var err error
	var count int
	var hiringJobs []*ent.HiringJob
	query := svc.repoRegistry.HiringJob().BuildQuery()
	hiringJobs, count, page, perPage, err = svc.getHiringJobs(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	edges = lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, index int) *ent.HiringJobEdge {
		return &ent.HiringJobEdge{
			Node: hiringJob,
			Cursor: ent.Cursor{
				Value: hiringJob.ID.String(),
			},
		}
	})
	result = &ent.HiringJobResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *hiringJobSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobSelectionResponseGetAll, error) {
	var result *ent.HiringJobSelectionResponseGetAll
	var edges []*ent.HiringJobSelectionEdge
	var page int
	var perPage int
	var err error
	var count int
	var hiringJobs []*ent.HiringJob
	query := svc.repoRegistry.HiringJob().BuildBaseQuery()
	hiringJobs, count, page, perPage, err = svc.getHiringJobs(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	edges = lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, index int) *ent.HiringJobSelectionEdge {
		return &ent.HiringJobSelectionEdge{
			Node: &ent.HiringJobSelection{
				ID:   hiringJob.ID.String(),
				Name: hiringJob.Name,
			},
			Cursor: ent.Cursor{
				Value: hiringJob.ID.String(),
			},
		}
	})
	result = &ent.HiringJobSelectionResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc hiringJobSvcImpl) GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType {
	return svc.dtoRegistry.EntitySkill().GroupSkillType(input)
}

func (svc *hiringJobSvcImpl) getHiringJobs(ctx context.Context, query *ent.HiringJobQuery, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) ([]*ent.HiringJob, int, int, int, error) {
	var page int
	var perPage int
	var err error
	var count int
	var hiringJobs []*ent.HiringJob
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	if ent.HiringJobOrderByAdditionalField.IsValid(ent.HiringJobOrderByAdditionalField(orderBy.Field.String())) {
		count, hiringJobs, err = svc.getListByAdditionalOrder(ctx, query, page, perPage, orderBy)
		if err != nil {
			return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	} else {
		count, hiringJobs, err = svc.getListByNormalOrder(ctx, query, page, perPage, orderBy)
		if err != nil {
			return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
		}
	}
	if err != nil {
		return nil, 0, 0, 0, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return hiringJobs, count, page, perPage, nil
}

func (svc hiringJobSvcImpl) getListByNormalOrder(ctx context.Context, query *ent.HiringJobQuery, page int, perPage int, orderBy ent.HiringJobOrderBy) (int, []*ent.HiringJob, error) {
	count, err := svc.repoRegistry.HiringJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return 0, nil, err
	}
	order := ent.Desc(strings.ToLower(orderBy.Field.String()))
	if orderBy.Direction == ent.OrderDirectionAsc {
		order = ent.Asc(strings.ToLower(orderBy.Field.String()))
	}
	query = query.Order(order)
	if page != 0 && perPage != 0 {
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return 0, nil, err
	}
	return count, hiringJobs, nil
}

func (svc hiringJobSvcImpl) getListByAdditionalOrder(ctx context.Context, query *ent.HiringJobQuery, page int, perPage int, orderBy ent.HiringJobOrderBy) (int, []*ent.HiringJob, error) {
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return 0, nil, err
	}
	count := len(hiringJobs)
	switch orderBy.Field {
	case ent.HiringJobOrderByFieldTotalCandidatesRecruited:
		sort.Slice(hiringJobs, func(i, j int) bool {
			if orderBy.Direction == ent.OrderDirectionAsc {
				return len(hiringJobs[i].Edges.CandidateJobEdges) < len(hiringJobs[j].Edges.CandidateJobEdges)
			} else {
				return len(hiringJobs[i].Edges.CandidateJobEdges) > len(hiringJobs[j].Edges.CandidateJobEdges)
			}
		})
	}
	// Split slice by page and perPage
	if page != 0 && perPage != 0 {
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(hiringJobs) {
			return count, nil, nil
		}
		if start <= len(hiringJobs) && end > len(hiringJobs) {
			return count, hiringJobs[start:], nil
		}
		hiringJobs = hiringJobs[start:end]
	}
	return count, hiringJobs, nil
}

// common function
func (svc *hiringJobSvcImpl) freeWord(hiringJobQuery *ent.HiringJobQuery, input *ent.HiringJobFreeWord) {
	if input != nil {
		if input.Name != nil {
			hiringJobQuery.Where(hiringjob.NameContainsFold(strings.TrimSpace(*input.Name)))
		}
	}
}

func (svc *hiringJobSvcImpl) filter(hiringJobQuery *ent.HiringJobQuery, input *ent.HiringJobFilter) {
	if input != nil {
		if input.Name != nil {
			hiringJobQuery.Where(hiringjob.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.TeamIds != nil {
			ids := lo.Map(input.TeamIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.TeamIDIn(ids...))
		}
		if input.Status != nil {
			hiringJobQuery.Where(hiringjob.StatusEQ(hiringjob.Status(*input.Status)))
		}
		if input.Priority != nil {
			hiringJobQuery.Where(hiringjob.PriorityEQ(*input.Priority))
		}
		if input.Location != nil {
			locations := lo.Map(input.Location, func(item *ent.LocationEnum, index int) hiringjob.Location {
				return hiringjob.Location(*item)
			})
			hiringJobQuery.Where(hiringjob.LocationIn(locations...))
		}
		if input.SkillIds != nil {
			ids := lo.Map(input.SkillIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.HasHiringJobSkillEdgesWith(
				entityskill.SkillIDIn(ids...),
			))
		}
		if input.CreatedByIds != nil {
			ids := lo.Map(input.CreatedByIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.CreatedByIn(ids...))
		}
	}
}
