package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/candidatejob"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/hiringjobstep"
	"trec/ent/hiringteam"
	"trec/ent/predicate"
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/ent/user"
	"trec/internal/util"
	"trec/middleware"
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
	GetHiringJobsGroupByStatus(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
		filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobGroupByStatusResponse, error)
	Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
		filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobSelectionResponseGetAll, error)

	// resolved
	GroupSkillType(input []*ent.EntitySkill) []*ent.EntitySkillType
}
type hiringJobSvcImpl struct {
	repoRegistry     repository.Repository
	dtoRegistry      dto.Dto
	logger           *zap.Logger
	hiringJobStepSvc HiringJobStepService
}

func NewHiringJobService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) HiringJobService {
	return &hiringJobSvcImpl{
		repoRegistry:     repoRegistry,
		dtoRegistry:      dtoRegistry,
		logger:           logger,
		hiringJobStepSvc: NewHiringJobStepService(repoRegistry, logger),
	}
}

func (svc *hiringJobSvcImpl) CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput, note string) (*ent.HiringJobResponse, error) {
	var record *ent.HiringJob
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	hiringTeam, err := svc.repoRegistry.HiringTeam().GetHiringTeam(ctx, uuid.MustParse(input.HiringTeamID))
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionCreate(payload, hiringTeam) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	errString, err := svc.repoRegistry.HiringJob().ValidName(ctx, uuid.Nil, input.Name, input.HiringTeamID)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.HiringJob().ValidPriority(ctx, uuid.Nil, uuid.MustParse(input.HiringTeamID), input.Priority)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.HiringJob().CreateHiringJob(ctx, input)
		if err != nil {
			return err
		}
		err = svc.repoRegistry.EntitySkill().CreateAndUpdateEntitySkill(ctx, record.ID, input.EntitySkillRecords, nil, entityskill.EntityTypeHiringJob)
		if err != nil {
			return err
		}
		return svc.hiringJobStepSvc.CreateBulkHiringJobSteps(ctx, repoRegistry, record)
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) DeleteHiringJob(ctx context.Context, id uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	currentUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	jobCreator, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(record.CreatedBy)))
	if !svc.validPermissionUpdateStatus(payload, currentUser, jobCreator, record.RecInChargeID, record.HiringTeamID, record.RecTeamID) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if record.Status == hiringjob.StatusPendingApprovals {
		approvalSteps := lo.Filter(record.Edges.ApprovalSteps, func(item *ent.HiringJobStep, index int) bool {
			return item.Status == hiringjobstep.StatusAccepted && item.UserID != record.CreatedBy
		})
		if len(approvalSteps) > 0 {
			return util.WrapGQLError(ctx, "model.hiring_jobs.validation.job_already_approving", http.StatusBadRequest, util.ErrorFlagValidateFail)
		}
	} else {
		return util.WrapGQLError(ctx, "model.hiring_jobs.validation.invalid_deleted_status", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.HiringJob().DeleteHiringJob(ctx, record)
		if err != nil {
			return err
		}
		err = repoRegistry.HiringJobStep().DeleteHiringJobStep(ctx, record.ID)
		if err != nil {
			return err
		}
		err = repoRegistry.HiringJob().DeleteRelationHiringJob(ctx, record.ID)
		return err
	})
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return err
}

func (svc *hiringJobSvcImpl) UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID, note string) (*ent.HiringJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.HiringJob
	if input.Amount == 0 {
		return nil, util.WrapGQLError(ctx, "model.hiring_jobs.validation.amount_neq_zero", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	currentUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	if !svc.validPermissionUpdate(payload, currentUser, record.Edges.OwnerEdge, record.HiringTeamID, record.RecTeamID) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	oldRecLeaderID := record.Edges.RecTeamEdge.LeaderID
	recTeamChange := record.RecTeamID != uuid.MustParse(input.RecTeamID)
	hiringTeamChange := record.HiringTeamID != uuid.MustParse(input.HiringTeamID)
	errString, err := svc.repoRegistry.HiringJob().ValidName(ctx, id, input.Name, input.HiringTeamID)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.HiringJob().ValidStatusWhenUpdate(ctx, record, input, recTeamChange, hiringTeamChange)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.HiringJob().ValidPriority(ctx, id, uuid.MustParse(input.HiringTeamID), input.Priority)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.RecTeam().ValidRecInCharge(ctx, uuid.MustParse(input.RecTeamID), uuid.MustParse(input.RecInChargeID))
	if err != nil {
		svc.logger.Error(err.Error())
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
		if err != nil {
			return err
		}
		result, _ = repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
		switch {
		case hiringTeamChange:
			err = repoRegistry.HiringJobStep().DeleteHiringJobStep(ctx, id)
			if err != nil {
				return err
			}
			return svc.hiringJobStepSvc.CreateBulkHiringJobSteps(ctx, repoRegistry, result)
		case recTeamChange:
			return svc.hiringJobStepSvc.UpdateHiringJobStepByRecLeader(ctx, repoRegistry, result, oldRecLeaderID)
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) UpdateHiringJobStatus(ctx context.Context, status ent.HiringJobStatus, id uuid.UUID, note string) (*ent.HiringJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.HiringJob
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if status == ent.HiringJobStatus(record.Status) {
		return &ent.HiringJobResponse{Data: record}, nil
	}
	currentUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	jobCreator, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(record.CreatedBy)))
	if !svc.validPermissionUpdateStatus(payload, currentUser, jobCreator, record.RecInChargeID, record.HiringTeamID, record.RecTeamID) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	err = svc.repoRegistry.HiringJob().ValidStatus(record.Status, status, record.Edges.CandidateJobEdges)
	if err != nil {
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, record, status)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ = svc.repoRegistry.HiringJob().GetHiringJob(ctx, record.ID)
	jsonString, err := svc.dtoRegistry.HiringJob().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error())
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) GetHiringJob(ctx context.Context, id uuid.UUID) (*ent.HiringJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.HiringJob().BuildQuery().Where(hiringjob.IDEQ(id))
	payloadUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	svc.validPermissionGet(payload, query, payloadUser)
	result, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.HiringJobResponse{
		Data: result,
	}, nil
}

func (svc *hiringJobSvcImpl) GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var (
		result     *ent.HiringJobResponseGetAll
		edges      []*ent.HiringJobEdge
		page       int
		perPage    int
		err        error
		count      int
		hiringJobs []*ent.HiringJob
	)
	query := svc.repoRegistry.HiringJob().BuildQuery()
	payloadUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	svc.validPermissionGet(payload, query, payloadUser)
	hiringJobs, count, page, perPage, err = svc.getHiringJobs(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error())
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

func (svc *hiringJobSvcImpl) GetHiringJobsGroupByStatus(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobGroupByStatusResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().
		WithCandidateJobEdges(func(query *ent.CandidateJobQuery) { query.Where(candidatejob.DeletedAtIsNil()) }).
		WithHiringTeamEdge(func(query *ent.HiringTeamQuery) { query.Where(hiringteam.DeletedAtIsNil()) }).
		WithHiringJobSkillEdges(func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).
				WithSkillEdge(func(query *ent.SkillQuery) {
					query.Where(skill.DeletedAtIsNil()).
						WithSkillTypeEdge(func(query *ent.SkillTypeQuery) { query.Where(skilltype.DeletedAtIsNil()) })
				})
		}).
		WithOwnerEdge().WithRecInChargeEdge().WithRecTeamEdge()
	payloadUser, _ := svc.repoRegistry.User().GetOneUser(ctx, svc.repoRegistry.User().BuildBaseQuery().Where(user.ID(payload.UserID)))
	svc.validPermissionGet(payload, query, payloadUser)
	hiringJobs, count, _, _, err := svc.getHiringJobs(ctx, query, nil, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	sampleEdge := &ent.HiringJobGroupByStatus{}
	for _, hiringJob := range hiringJobs {
		switch hiringJob.Status {
		case hiringjob.StatusPendingApprovals:
			sampleEdge.PendingApprovals = append(sampleEdge.PendingApprovals, hiringJob)
		case hiringjob.StatusOpened:
			sampleEdge.Opened = append(sampleEdge.Opened, hiringJob)
		case hiringjob.StatusClosed:
			sampleEdge.Closed = append(sampleEdge.Closed, hiringJob)
		case hiringjob.StatusCancelled:
			sampleEdge.Cancelled = append(sampleEdge.Cancelled, hiringJob)
		}
	}
	page := *pagination.Page
	perPage := *pagination.PerPage
	edge := &ent.HiringJobGroupByStatus{
		PendingApprovals: svc.pagination(sampleEdge.PendingApprovals, page, perPage),
		Opened:           svc.pagination(sampleEdge.Opened, page, perPage),
		Closed:           svc.pagination(sampleEdge.Closed, page, perPage),
		Cancelled:        svc.pagination(sampleEdge.Cancelled, page, perPage),
	}
	result := &ent.HiringJobGroupByStatusResponse{
		Data: edge,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc hiringJobSvcImpl) pagination(records []*ent.HiringJob, page int, perPage int) []*ent.HiringJob {
	if page != 0 && perPage != 0 {
		start := (page - 1) * perPage
		end := start + perPage
		if start > len(records) {
			return nil
		}
		if start <= len(records) && end > len(records) {
			return records[start:]
		}
		records = records[start:end]
	}
	return records
}

func (svc *hiringJobSvcImpl) Selections(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord,
	filter *ent.HiringJobFilter, orderBy ent.HiringJobOrderBy) (*ent.HiringJobSelectionResponseGetAll, error) {
	var (
		result     *ent.HiringJobSelectionResponseGetAll
		edges      []*ent.HiringJobSelectionEdge
		page       int
		perPage    int
		err        error
		count      int
		hiringJobs []*ent.HiringJob
	)
	query := svc.repoRegistry.HiringJob().BuildBaseQuery()
	hiringJobs, count, page, perPage, err = svc.getHiringJobs(ctx, query, pagination, freeWord, filter, orderBy)
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, err
	}
	edges = lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, _ int) *ent.HiringJobSelectionEdge {
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
	var (
		page       int
		perPage    int
		err        error
		count      int
		hiringJobs []*ent.HiringJob
	)
	svc.filter(ctx, query, filter)
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
		svc.logger.Error(err.Error())
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
		svc.logger.Error(err.Error())
		return 0, nil, err
	}
	return count, hiringJobs, nil
}

func (svc hiringJobSvcImpl) getListByAdditionalOrder(ctx context.Context, query *ent.HiringJobQuery, page int, perPage int, orderBy ent.HiringJobOrderBy) (int, []*ent.HiringJob, error) {
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error())
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

func (svc *hiringJobSvcImpl) filter(ctx context.Context, hiringJobQuery *ent.HiringJobQuery, input *ent.HiringJobFilter) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	if input != nil {
		if input.Name != nil {
			hiringJobQuery.Where(hiringjob.NameEqualFold(strings.TrimSpace(*input.Name)))
		}
		if input.HiringTeamIds != nil {
			ids := lo.Map(input.HiringTeamIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.HiringTeamIDIn(ids...))
		}
		if input.RecTeamIds != nil {
			ids := lo.Map(input.RecTeamIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.RecTeamIDIn(ids...))
		}
		if input.HasRecInCharge != nil || input.RecInChargeIds != nil {
			var conditions []predicate.HiringJob
			if input.HasRecInCharge != nil {
				if *input.HasRecInCharge {
					conditions = append(conditions, hiringjob.RecInChargeIDIsNil())
				} else {
					conditions = append(conditions, hiringjob.RecInChargeIDNotNil())
				}
			}
			if input.RecInChargeIds != nil {
				ids := lo.Map(input.RecInChargeIds, func(item string, index int) uuid.UUID {
					return uuid.MustParse(item)
				})
				conditions = append(conditions, hiringjob.RecInChargeIDIn(ids...))
			}
			if len(conditions) > 0 {
				hiringJobQuery.Where(hiringjob.Or(conditions...))
			}
		}
		if input.Status != nil {
			hiringJobQuery.Where(hiringjob.StatusEQ(hiringjob.Status(*input.Status)))
		}
		if input.Priorities != nil {
			hiringJobQuery.Where(hiringjob.PriorityIn(input.Priorities...))
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
		if input.ForOwner != nil {
			if *input.ForOwner {
				hiringJobQuery.Where(hiringjob.HasHiringTeamEdgeWith(
					hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)),
				))
			} else {
				hiringJobQuery.Where(hiringjob.IDEQ(uuid.Nil))
			}
		}
		if input.ForHiringTeam != nil {
			if *input.ForHiringTeam {
				hiringJobQuery.Where(hiringjob.HasHiringTeamEdgeWith(
					hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))),
				))
			} else {
				hiringJobQuery.Where(hiringjob.IDEQ(uuid.Nil))
			}
		}
		if input.JobPositionIds != nil {
			ids := lo.Map(input.JobPositionIds, func(item string, index int) uuid.UUID {
				return uuid.MustParse(item)
			})
			hiringJobQuery.Where(hiringjob.JobPositionIDIn(ids...))
		}
		hiringJobStepPredicates := make([]predicate.HiringJobStep, 0)
		if input.ApproverID != nil {
			hiringJobStepPredicates = append(hiringJobStepPredicates, hiringjobstep.UserID(uuid.MustParse(*input.ApproverID)))
		}
		if input.ApproverStatus != nil {
			hiringJobStepPredicates = append(hiringJobStepPredicates, hiringjobstep.StatusEQ(hiringjobstep.Status(*input.ApproverStatus)))
		}
		if len(hiringJobStepPredicates) > 0 {
			hiringJobQuery.Where(hiringjob.HasApprovalStepsWith(hiringJobStepPredicates...))
		}
	}
}

// permission
func (svc hiringJobSvcImpl) validPermissionUpdate(payload *middleware.Payload, currentUser, jobCreator *ent.User, jobHiringTeamID, jobRecTeamID uuid.UUID) bool {
	if payload.ForAll {
		return true
	}
	if payload.ForTeam {
		return currentUser.HiringTeamID == jobHiringTeamID || currentUser.HiringTeamID == jobRecTeamID
	}
	return currentUser.ID == jobCreator.ID
}

func (svc hiringJobSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.HiringJobQuery, payloadUser *ent.User) {
	if payload.ForAll {
		return
	}
	if payload.ForTeam {
		teamPredidates := make([]predicate.HiringJob, 0)
		if payloadUser.HiringTeamID != uuid.Nil {
			teamPredidates = append(teamPredidates, hiringjob.HiringTeamID(payloadUser.HiringTeamID))
		}
		if payloadUser.RecTeamID != uuid.Nil {
			teamPredidates = append(teamPredidates, hiringjob.RecTeamID(payloadUser.RecTeamID))
		}
		query.Where(hiringjob.Or(teamPredidates...))
		return
	}
	query.Where(hiringjob.Or(
		hiringjob.CreatedBy(payload.UserID),
		hiringjob.RecInChargeID(payload.UserID),
	))
}

func (svc hiringJobSvcImpl) validPermissionCreate(payload *middleware.Payload, hiringTeam *ent.HiringTeam) bool {
	if payload.ForAll {
		return true
	}
	return lo.ContainsBy(hiringTeam.Edges.HiringMemberEdges, func(item *ent.User) bool { return item.ID == payload.UserID })
}

func (svc hiringJobSvcImpl) validPermissionUpdateStatus(payload *middleware.Payload, currentUser, jobCreator *ent.User, recInChargeID, jobHiringTeamID, jobRecTeamID uuid.UUID) bool {
	if payload.ForAll {
		return true
	}
	if payload.ForTeam {
		return currentUser.HiringTeamID == jobHiringTeamID || currentUser.RecTeamID == jobRecTeamID
	}
	return currentUser.ID == jobCreator.ID || currentUser.ID == recInChargeID
}

// Path: service/hiring_job.service.go
