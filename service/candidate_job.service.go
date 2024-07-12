package service

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"trec/config"
	"trec/dto"
	"trec/ent"
	"trec/ent/attachment"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/candidatejobfeedback"
	"trec/ent/entityskill"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/skill"
	"trec/ent/skilltype"
	"trec/ent/team"
	"trec/ent/user"
	"trec/internal/servicebus"
	"trec/internal/util"
	"trec/middleware"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateJobService interface {
	// mutation
	CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, note string) (*ent.CandidateJobResponse, error)
	UpdateCandidateJobStatus(ctx context.Context, input ent.UpdateCandidateJobStatus, id uuid.UUID, note string) error
	DeleteCandidateJob(ctx context.Context, id uuid.UUID, note string) error
	UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID, note string) (*ent.CandidateJobResponse, error)

	// query
	GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error)
	GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error)
	GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput,
		filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (
		*ent.CandidateJobGroupByStatusResponse, error)
	GetCandidateJobGroupByInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateJobGroupByInterviewResponse, error)
}

type candidateJobSvcImpl struct {
	attachmentSvc       AttachmentService
	emailSvc            EmailService
	candidateJobStepSvc CandidateJobStepService
	outgoingEmailSvc    OutgoingEmailService
	serviceBusClient    servicebus.ServiceBus
	repoRegistry        repository.Repository
	dtoRegistry         dto.Dto
	logger              *zap.Logger
	configs             *config.Configurations
}

func NewCandidateJobService(repoRegistry repository.Repository, serviceBusClient servicebus.ServiceBus, dtoRegistry dto.Dto, logger *zap.Logger, configs *config.Configurations) CandidateJobService {
	return &candidateJobSvcImpl{
		attachmentSvc:       NewAttachmentService(repoRegistry, logger),
		candidateJobStepSvc: NewCandidateJobStepService(repoRegistry, logger),
		emailSvc:            NewEmailService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		outgoingEmailSvc:    NewOutgoingEmailService(repoRegistry, logger),
		serviceBusClient:    serviceBusClient,
		repoRegistry:        repoRegistry,
		dtoRegistry:         dtoRegistry,
		logger:              logger,
		configs:             configs,
	}
}

func (svc *candidateJobSvcImpl) CreateCandidateJob(ctx context.Context, input *ent.NewCandidateJobInput, note string) (*ent.CandidateJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var candidateJob *ent.CandidateJob
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(uuid.MustParse(input.HiringJobID))).WithTeamEdge(
		func(query *ent.TeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.TeamEdge) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, uuid.MustParse(input.CandidateID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJob().ValidStatus(ctx, uuid.MustParse(input.CandidateID), uuid.Nil, &input.Status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJob, err = repoRegistry.CandidateJob().CreateCandidateJob(ctx, input)
		if err != nil {
			return err
		}
		_, err := svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		if err != nil {
			return err
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidateJob.Status, candidateJob.ID, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, candidateJob.ID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobStatus(ctx context.Context, input ent.UpdateCandidateJobStatus, id uuid.UUID, note string) error {
	var result *ent.CandidateJob
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx,
		svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(candidatejob.IDEQ(id)).WithHiringJobEdge())
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.HiringJobID)).WithTeamEdge(
		func(query *ent.TeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.TeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if record.Edges.HiringJobEdge.Status == hiringjob.StatusClosed {
		return util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, record.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err = svc.repoRegistry.CandidateJob().ValidStatus(ctx, record.CandidateID, id, &input.Status)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, record, input)
		if err != nil {
			return err
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, result.Status, result.ID, repoRegistry)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	err = svc.triggerEventSendEmail(ctx, record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailUpdateStatus(record.Status, candidatejob.Status(input.Status))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobAttachment(ctx context.Context, input ent.UpdateCandidateAttachment, id uuid.UUID, note string) (*ent.CandidateJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(candidateJob.HiringJobID)).WithTeamEdge(
		func(query *ent.TeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.TeamEdge) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusClosed {
		return nil, util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, candidateJob.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = svc.repoRegistry.Attachment().CreateAndUpdateAttachment(ctx, candidateJob.ID, input.Attachments, candidateJob.Edges.AttachmentEdges, attachment.RelationTypeCandidateJobs)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailUpdate(candidateJob, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.IDEQ(id))
	svc.validPermissionGet(payload, query)
	result, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) DeleteCandidateJob(ctx context.Context, id uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(candidateJob.HiringJobID)).WithTeamEdge(
		func(query *ent.TeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.TeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusClosed && candidateJob.Status != candidatejob.StatusApplied {
		return util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	// if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusOpened && !ent.CandidateJobStatusEnded.IsValid(ent.CandidateJobStatusEnded(candidateJob.Status)) {
	// 	return util.WrapGQLError(ctx, "model.candidate_job.validation.status_is_invalid_to_delete", http.StatusBadRequest, util.ErrorFlagValidateFail)
	// }
	errString, err := svc.repoRegistry.CandidateJob().ValidUpsetByCandidateIsBlacklist(ctx, candidateJob.CandidateID)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.CandidateJob().DeleteCandidateJob(ctx, candidateJob)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailCreate(candidateJob)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, candidateJob.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateJobResponseGetAll
	var edges []*ent.CandidateJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.CandidateIDEQ(uuid.MustParse(filter.CandidateID)))
	svc.validPermissionGet(payload, query)
	svc.filter(ctx, query, filter)
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

func (svc candidateJobSvcImpl) GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput,
	filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (
	*ent.CandidateJobGroupByStatusResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateJobGroupByStatusResponse
	var edges *ent.CandidateJobGroupByStatus
	var page int
	var perPage int
	var err error
	var candidateJobs []*ent.CandidateJob
	query := svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(
		candidatejob.HasCandidateEdgeWith(
			candidate.DeletedAtIsNil(), candidate.IsBlacklist(false),
		), candidatejob.HasHiringJobEdgeWith(
			hiringjob.DeletedAtIsNil(), hiringjob.StatusNEQ(hiringjob.StatusClosed),
		))
	svc.validPermissionGet(payload, query)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
	}
	svc.customFilter(query, filter)
	svc.customFreeWord(query, freeWord)
	count, err := svc.repoRegistry.CandidateJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if orderBy != nil {
		if !ent.CandidateJobOrderByAdditionalField.IsValid(ent.CandidateJobOrderByAdditionalField(orderBy.Field.String())) {
			order := ent.Desc(strings.ToLower(orderBy.Field.String()))
			if orderBy.Direction == ent.OrderDirectionAsc {
				order = ent.Asc(strings.ToLower(orderBy.Field.String()))
			}
			query = query.Order(order)
		}
	}
	candidateJobs, err = svc.repoRegistry.CandidateJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}

	candidateJobIds := lo.Map(candidateJobs, func(candidateJob *ent.CandidateJob, index int) uuid.UUID {
		return candidateJob.ID
	})
	candidateIds := lo.Map(candidateJobs, func(candidateJob *ent.CandidateJob, index int) uuid.UUID {
		return candidateJob.CandidateID
	})
	hiringJobIds := lo.Map(candidateJobs, func(candidateJob *ent.CandidateJob, index int) uuid.UUID {
		return candidateJob.HiringJobID
	})

	interviews, err := svc.repoRegistry.CandidateInterview().BuildList(ctx,
		svc.repoRegistry.CandidateInterview().BuildBaseQuery().Where(candidateinterview.CandidateJobIDIn(candidateJobIds...)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidates, err := svc.repoRegistry.Candidate().BuildList(ctx,
		svc.repoRegistry.Candidate().BuildBaseQuery().Where(candidate.IDIn(candidateIds...)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx,
		svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDIn(hiringJobIds...)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	teamIds := lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, index int) uuid.UUID {
		return hiringJob.TeamID
	})
	teams, err := svc.repoRegistry.Team().BuildList(ctx,
		svc.repoRegistry.Team().BuildBaseQuery().Where(team.IDIn(teamIds...)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	svc.dtoRegistry.HiringJob().MappingEdge(hiringJobs, teams)
	svc.dtoRegistry.CandidateJob().MappingEdge(candidateJobs, candidates, interviews, hiringJobs)
	if orderBy != nil {
		if ent.CandidateJobOrderByAdditionalField.IsValid(ent.CandidateJobOrderByAdditionalField(orderBy.Field.String())) {
			switch orderBy.Field {
			case ent.CandidateJobOrderByFieldPriority:
				sort.Slice(candidateJobs, func(i, j int) bool {
					if orderBy.Direction == ent.OrderDirectionAsc {
						return candidateJobs[i].Edges.HiringJobEdge.Priority < candidateJobs[j].Edges.HiringJobEdge.Priority
					} else {
						return candidateJobs[i].Edges.HiringJobEdge.Priority > candidateJobs[j].Edges.HiringJobEdge.Priority
					}
				})
			case ent.CandidateJobOrderByFieldCreatedAt:
				sort.Slice(candidateJobs, func(i, j int) bool {
					if orderBy.Direction == ent.OrderDirectionAsc {
						return candidateJobs[i].CreatedAt.Before(candidateJobs[j].CreatedAt)
					} else {
						return candidateJobs[i].CreatedAt.After(candidateJobs[j].CreatedAt)
					}
				})
			}
		}
	}
	sampleEdges := &ent.CandidateJobGroupByStatus{
		Hired: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusHired
		}),
		Kiv: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusKiv
		}),
		OfferLost: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusOfferLost
		}),
		Offering: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusOffering
		}),
		ExStaff: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusExStaff
		}),
		Applied: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusApplied
		}),
		Interviewing: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusInterviewing
		}),
	}
	edges = &ent.CandidateJobGroupByStatus{
		Hired:        svc.Pagination(sampleEdges.Hired, page, perPage),
		Kiv:          svc.Pagination(sampleEdges.Kiv, page, perPage),
		OfferLost:    svc.Pagination(sampleEdges.OfferLost, page, perPage),
		Offering:     svc.Pagination(sampleEdges.Offering, page, perPage),
		ExStaff:      svc.Pagination(sampleEdges.ExStaff, page, perPage),
		Applied:      svc.Pagination(sampleEdges.Applied, page, perPage),
		Interviewing: svc.Pagination(sampleEdges.Interviewing, page, perPage),
	}
	result = &ent.CandidateJobGroupByStatusResponse{
		Data: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobGroupByInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateJobGroupByInterviewResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var edges *ent.CandidateJobGroupByInterview
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.IDEQ(id)).WithCandidateJobInterview(
		func(query *ent.CandidateInterviewQuery) {
			query.Where(candidateinterview.DeletedAtIsNil()).WithCreatedByEdge().WithInterviewerEdges().WithCandidateJobEdge().
				Order(ent.Desc(candidateinterview.FieldCreatedAt))
		},
	).WithCandidateJobFeedback(
		func(query *ent.CandidateJobFeedbackQuery) {
			query.Where(candidatejobfeedback.DeletedAtIsNil()).WithAttachmentEdges(
				func(query *ent.AttachmentQuery) {
					query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobFeedbacks))
				},
			).WithCreatedByEdge().Order(ent.Desc(candidatejobfeedback.FieldCreatedAt))
		},
	)
	svc.validPermissionGet(payload, query)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = &ent.CandidateJobGroupByInterview{
		Applied: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusApplied
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusApplied
			}),
		},
		Interviewing: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusInterviewing
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusInterviewing
			}),
		},
		Offering: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusOffering
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusOffering
			}),
		},
		Hired: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusHired
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusHired
			}),
		},
		OfferLost: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusOfferLost
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusOfferLost
			}),
		},
		Kiv: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusKiv
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusKiv
			}),
		},
		ExStaff: &ent.CandidateJobGroupInterviewFeedback{
			Interview: lo.Filter(candidateJob.Edges.CandidateJobInterview, func(candidateInterview *ent.CandidateInterview, index int) bool {
				return candidateInterview.CandidateJobStatus == candidateinterview.CandidateJobStatusExStaff
			}),
			Feedback: lo.Filter(candidateJob.Edges.CandidateJobFeedback, func(candidateJobFeedback *ent.CandidateJobFeedback, index int) bool {
				return candidateJobFeedback.CandidateJobStatus == candidatejobfeedback.CandidateJobStatusExStaff
			}),
		},
	}
	result := &ent.CandidateJobGroupByInterviewResponse{
		Data: edges,
	}
	return result, nil
}

// common function
func (svc *candidateJobSvcImpl) freeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobFreeWord) {
	var predicate []predicate.CandidateJob
	if input != nil {
		if input.Team != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.HasTeamEdgeWith(
					team.NameEqualFold(strings.TrimSpace(*input.Team)),
				),
			))
		}
		if input.HiringJob != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.NameEqualFold(strings.TrimSpace(*input.HiringJob)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predicate...))
	}
}

func (svc *candidateJobSvcImpl) filter(ctx context.Context, candidateJobQuery *ent.CandidateJobQuery, input ent.CandidateJobFilter) {
	if input.Status != nil {
		candidateJobQuery.Where(candidatejob.StatusEQ(candidatejob.Status(*input.Status)))
	}
	if input.TeamID != nil {
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
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
	if input.FailedReason != nil && len(input.FailedReason) != 0 {
		candidateJobIds := []uuid.UUID{}
		queryString := "SELECT id FROM candidate_jobs WHERE "
		for i, reason := range input.FailedReason {
			queryString += "failed_reason @> '[\"" + reason.String() + "\"]'::jsonb"
			if i != len(input.FailedReason)-1 {
				queryString += " AND "
			}
		}
		queryString += ";"
		rows, _ := candidateJobQuery.QueryContext(ctx, queryString)
		if rows != nil {
			for rows.Next() {
				var id uuid.UUID
				rows.Scan(&id)
				candidateJobIds = append(candidateJobIds, id)
			}
			candidateJobQuery.Where(candidatejob.IDIn(candidateJobIds...))
		} else {
			candidateJobQuery.Where(candidatejob.IDEQ(uuid.Nil))
		}
	}
}

func (svc *candidateJobSvcImpl) customFreeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobGroupByStatusFreeWord) {
	var predicate []predicate.CandidateJob
	if input != nil {
		if input.JobTitle != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.NameContainsFold(strings.TrimSpace(*input.JobTitle)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predicate...))
	}
}

func (svc *candidateJobSvcImpl) customFilter(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobGroupByStatusFilter) {
	if input == nil {
		return
	}
	if input.HiringJobID != nil {
		hiringJobIds := lo.Map(input.HiringJobID, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HiringJobIDIn(hiringJobIds...))
	}
	if input.TeamID != nil {
		teamIds := lo.Map(input.TeamID, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasTeamEdgeWith(
				team.IDIn(teamIds...),
			),
		))
	}
	if input.Priority != nil {
		priority := lo.Map(input.Priority, func(id int, index int) int {
			return id
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.PriorityIn(priority...)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
	if input.SkillID != nil {
		skillIds := lo.Map(input.SkillID, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasHiringJobSkillEdgesWith(entityskill.SkillIDIn(skillIds...)),
		))
	}
	if input.Location != nil {
		locations := lo.Map(input.Location, func(location ent.LocationEnum, index int) hiringjob.Location {
			return hiringjob.Location(location)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.LocationIn(locations...),
		))
	}
	if input.CreatedByIds != nil {
		createdByIds := lo.Map(input.CreatedByIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.CreatedByIn(createdByIds...))
	}
}

func (svc candidateJobSvcImpl) Pagination(records []*ent.CandidateJob, page int, perPage int) []*ent.CandidateJob {
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

// third func
func (svc candidateJobSvcImpl) triggerEventSendEmail(ctx context.Context, oldRecord, newRecord *ent.CandidateJob) error {
	var messages []models.MessageInput
	var outgoingEmails []models.MessageInput
	var results []*ent.OutgoingEmail
	emailTemplates, err := svc.repoRegistry.EmailTemplate().ValidAndGetEmailTemplates(ctx, oldRecord, newRecord)
	if err != nil {
		return err
	}
	if len(emailTemplates) == 0 {
		return nil
	}
	groupModule, err := svc.getDataForKeyword(ctx, newRecord)
	if err != nil {
		return err
	}
	users, err := svc.repoRegistry.User().BuildList(ctx, svc.repoRegistry.User().BuildBaseQuery())
	if err != nil {
		return err
	}
	for _, entity := range emailTemplates {
		messages = append(messages, svc.emailSvc.GenerateEmail(ctx, users, entity, groupModule)...)
	}
	results, err = svc.outgoingEmailSvc.CreateBulkOutgoingEmail(ctx, messages)
	if err != nil {
		return err
	}
	outgoingEmails = lo.Map(results, func(entity *ent.OutgoingEmail, index int) models.MessageInput {
		return models.MessageInput{
			ID:        entity.ID,
			To:        entity.To,
			Cc:        entity.Cc,
			Bcc:       entity.Bcc,
			Subject:   entity.Subject,
			Content:   entity.Content,
			Signature: entity.Signature,
		}
	})
	err = svc.emailSvc.SentEmail(ctx, outgoingEmails)
	if err != nil {
		return err
	}
	return err
}

// permission
func (svc candidateJobSvcImpl) validPermissionMutation(payload *middleware.Payload, teamRecord *ent.Team) bool {
	if payload.ForAll {
		return true
	}
	if payload.ForTeam {
		memberIds := lo.Map(teamRecord.Edges.MemberEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		managerIds := lo.Map(teamRecord.Edges.UserEdges, func(item *ent.User, index int) uuid.UUID {
			return item.ID
		})
		if lo.Contains(memberIds, payload.UserID) || lo.Contains(managerIds, payload.UserID) {
			return true
		}
	}
	return false
}

func (svc candidateJobSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.CandidateJobQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForTeam {
		query.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.HasTeamEdgeWith(
			team.Or(team.HasUserEdgesWith(user.IDEQ(payload.UserID)), team.HasMemberEdgesWith(user.IDEQ(payload.UserID))),
		)))
	}
}

func (svc candidateJobSvcImpl) getDataForKeyword(ctx context.Context, record *ent.CandidateJob) (models.GroupModule, error) {
	var result models.GroupModule
	candidateQuery := svc.repoRegistry.Candidate().BuildBaseQuery().Where(candidate.IDEQ(record.CandidateID)).
		WithReferenceUserEdge().WithCandidateSkillEdges(
		func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).WithSkillEdge(
				func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(
						func(stq *ent.SkillTypeQuery) {
							stq.Where(skilltype.DeletedAtIsNil())
						},
					)
				},
			)
		},
	)
	hiringjobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.HiringJobID)).WithHiringJobSkillEdges(
		func(query *ent.EntitySkillQuery) {
			query.Where(entityskill.DeletedAtIsNil()).Order(ent.Asc(entityskill.FieldOrderID)).WithSkillEdge(
				func(sq *ent.SkillQuery) {
					sq.Where(skill.DeletedAtIsNil()).WithSkillTypeEdge(
						func(stq *ent.SkillTypeQuery) {
							stq.Where(skilltype.DeletedAtIsNil())
						},
					)
				},
			)
		},
	).WithOwnerEdge()
	candidateRecord, err := svc.repoRegistry.Candidate().BuildGet(ctx, candidateQuery)
	if err != nil {
		return result, err
	}
	hiringJobRecord, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringjobQuery)
	if err != nil {
		return result, err
	}
	teamRecord, err := svc.repoRegistry.Team().BuildGetOne(ctx, svc.repoRegistry.Team().BuildBaseQuery().Where(team.IDEQ(hiringJobRecord.TeamID)).WithUserEdges())
	if err != nil {
		return result, nil
	}
	return models.GroupModule{
		Candidate:    candidateRecord,
		HiringJob:    hiringJobRecord,
		Team:         teamRecord,
		CandidateJob: record,
	}, nil
}

// Path: service/candidate_job.service.go
