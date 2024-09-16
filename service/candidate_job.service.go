package service

import (
	"context"
	"encoding/json"
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
	"trec/ent/hiringteam"
	"trec/ent/predicate"
	"trec/ent/recteam"
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
	UpdateCandidateJob(ctx context.Context, input ent.UpdateCandidateJobInput, id uuid.UUID, note string) (*ent.CandidateJobResponse, error)

	// query
	GetCandidateJob(ctx context.Context, id uuid.UUID) (*ent.CandidateJobResponse, error)
	GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter *ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error)
	GetCandidateJobGroupByStatus(ctx context.Context, pagination *ent.PaginationInput,
		filter *ent.CandidateJobGroupByStatusFilter, freeWord *ent.CandidateJobGroupByStatusFreeWord, orderBy *ent.CandidateJobByOrder) (
		*ent.CandidateJobGroupByStatusResponse, error)
	GetCandidateJobGroupByInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateJobGroupByInterviewResponse, error)
	ValidProcessingCdJobExistByCdID(ctx context.Context, candidateID uuid.UUID) (bool, error)
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
	var inputValidate models.CandidateJobValidInput
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(uuid.MustParse(input.HiringJobID))).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.
				WithUserEdges(func(query *ent.UserQuery) { query.Where(user.DeletedAtIsNil()) }).
				WithHiringMemberEdges()
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.HiringTeamEdge) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	jsonValidate, _ := json.Marshal(input)
	json.Unmarshal(jsonValidate, &inputValidate)
	failedReason, errString, err := svc.repoRegistry.CandidateJob().ValidInput(ctx, inputValidate)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateJob, err = repoRegistry.CandidateJob().CreateCandidateJob(ctx, input, failedReason)
		if err != nil {
			return err
		}
		_, err := svc.attachmentSvc.CreateAttachment(ctx, input.Attachments, candidateJob.ID, attachment.RelationTypeCandidateJobs, repoRegistry)
		if err != nil {
			return err
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidatejob.StatusApplied, candidateJob.ID, repoRegistry)
		if err != nil {
			return err
		}
		if candidateJob.Status == candidatejob.StatusInterviewing || candidateJob.Status == candidatejob.StatusOffering {
			err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidatejob.StatusInterviewing, candidateJob.ID, repoRegistry)
			if err != nil {
				return err
			}
		}
		if candidateJob.Status == candidatejob.StatusOffering {
			err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidatejob.StatusOffering, candidateJob.ID, repoRegistry)
			if err != nil {
				return err
			}
		}
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
	err = svc.triggerEventSendEmail(ctx, result, nil)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateJobResponse{
		Data: result,
	}, nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJobStatus(ctx context.Context, input ent.UpdateCandidateJobStatus, id uuid.UUID, note string) error {
	var result *ent.CandidateJob
	var inputValidate models.CandidateJobValidInput
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx,
		svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(candidatejob.IDEQ(id)).WithHiringJobEdge())
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
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
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.HiringTeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if record.Edges.HiringJobEdge.Status == hiringjob.StatusClosed {
		return util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	jsonValidate, _ := json.Marshal(input)
	json.Unmarshal(jsonValidate, &inputValidate)
	inputValidate.CandidateId = record.CandidateID
	inputValidate.CandidateJobId = record.ID
	failedReason, errString, err := svc.repoRegistry.CandidateJob().ValidInput(ctx, inputValidate)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagValidateFail)
	}
	if errString != nil {
		return util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.CandidateJob().ValidStatus(record.Status, input.Status)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.CandidateJob().UpdateCandidateJobStatus(ctx, record, input, failedReason)
		if err != nil {
			return err
		}
		if record.Status == candidatejob.StatusApplied && result.Status == candidatejob.StatusOffering {
			err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, candidatejob.StatusInterviewing, result.ID, repoRegistry)
			if err != nil {
				return err
			}
		}
		err = svc.candidateJobStepSvc.CreateCandidateJobStep(ctx, result.Status, result.ID, repoRegistry)
		if err != nil {
			return err
		}
		// interviewing to offering/failed interview: cancel all invited/interviewing interviews
		if record.Status == candidatejob.StatusInterviewing && (result.Status == candidatejob.StatusOffering || result.Status == candidatejob.StatusFailedInterview) {
			return repoRegistry.CandidateInterview().UpdateBulkCandidateInterviewStatus(
				ctx,
				[]predicate.CandidateInterview{
					candidateinterview.CandidateJobID(id),
					candidateinterview.StatusIn(candidateinterview.StatusInvitedToInterview, candidateinterview.StatusInterviewing),
				},
				candidateinterview.StatusCancelled,
			)
		}
		return nil
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	err = svc.triggerEventSendEmail(ctx, record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	jsonString, err := svc.dtoRegistry.CandidateJob().AuditTrailUpdateStatus(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobSvcImpl) UpdateCandidateJob(ctx context.Context, input ent.UpdateCandidateJobInput, id uuid.UUID, note string) (*ent.CandidateJobResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(candidateJob.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
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
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.HiringTeamEdge) {
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
	if candidateJob.Status == candidatejob.StatusOffering {
		errString := svc.repoRegistry.CandidateJob().ValidOfferingInput(input.OnboardDate, input.OfferExpirationDate)
		if errString != nil {
			return nil, util.WrapGQLError(ctx, errString.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
		}
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateJob().UpdateCandidateJob(ctx, candidateJob, input)
		if err != nil {
			return err
		}
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
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeUpdate, note)
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
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
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
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(candidateJob.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
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
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionMutation(payload, hiringJob.Edges.HiringTeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if candidateJob.Edges.HiringJobEdge.Status == hiringjob.StatusClosed && candidateJob.Status != candidatejob.StatusApplied {
		return util.WrapGQLError(ctx, "model.candidate_job.validation.job_is_closed", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
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
		if err != nil {
			return err
		}
		err = repoRegistry.CandidateJob().DeleteRelationCandidateJob(ctx, candidateJob.ID)
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
	_, err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, candidateJob.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateJobSvcImpl) GetCandidateJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateJobFreeWord, filter *ent.CandidateJobFilter, orderBy *ent.CandidateJobOrder) (*ent.CandidateJobResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateJobResponseGetAll
	var edges []*ent.CandidateJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateJob().BuildQuery()
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
			hiringjob.DeletedAtIsNil(),
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
	hiringTeamIds := lo.Map(hiringJobs, func(hiringJob *ent.HiringJob, index int) uuid.UUID {
		return hiringJob.HiringTeamID
	})
	hiringTeams, err := svc.repoRegistry.HiringTeam().BuildList(ctx,
		svc.repoRegistry.HiringTeam().BuildBaseQuery().Where(hiringteam.IDIn(hiringTeamIds...)))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	svc.dtoRegistry.HiringJob().MappingEdge(hiringJobs, hiringTeams)
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
		FailedCv: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusFailedCv
		}),
		FailedInterview: lo.Filter(candidateJobs, func(candidateJob *ent.CandidateJob, index int) bool {
			return candidateJob.Status == candidatejob.StatusFailedInterview
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
		Hired:           svc.Pagination(sampleEdges.Hired, page, perPage),
		FailedCv:        svc.Pagination(sampleEdges.FailedCv, page, perPage),
		FailedInterview: svc.Pagination(sampleEdges.FailedInterview, page, perPage),
		OfferLost:       svc.Pagination(sampleEdges.OfferLost, page, perPage),
		Offering:        svc.Pagination(sampleEdges.Offering, page, perPage),
		ExStaff:         svc.Pagination(sampleEdges.ExStaff, page, perPage),
		Applied:         svc.Pagination(sampleEdges.Applied, page, perPage),
		Interviewing:    svc.Pagination(sampleEdges.Interviewing, page, perPage),
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
	var edge *ent.CandidateJobInterviewFeedback
	query := svc.repoRegistry.CandidateJob().BuildQuery().Where(candidatejob.IDEQ(id)).
		WithCandidateJobInterview(func(query *ent.CandidateInterviewQuery) {
			query.Where(candidateinterview.DeletedAtIsNil()).
				WithCreatedByEdge().WithInterviewerEdges().WithCandidateJobEdge().
				Order(ent.Desc(candidateinterview.FieldCreatedAt))
		}).
		WithCandidateJobFeedback(func(query *ent.CandidateJobFeedbackQuery) {
			query.Where(candidatejobfeedback.DeletedAtIsNil()).
				WithAttachmentEdges(func(query *ent.AttachmentQuery) {
					query.Where(attachment.DeletedAtIsNil(), attachment.RelationTypeEQ(attachment.RelationTypeCandidateJobFeedbacks))
				}).
				WithCreatedByEdge(func(query *ent.UserQuery) { query.WithHiringTeamEdges().WithMemberOfHiringTeamEdges() }).
				Order(ent.Desc(candidatejobfeedback.FieldCreatedAt))
		})
	svc.validPermissionGet(payload, query)
	candidateJob, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edge = &ent.CandidateJobInterviewFeedback{
		Interview: candidateJob.Edges.CandidateJobInterview,
		Feedback:  candidateJob.Edges.CandidateJobFeedback,
	}
	result := &ent.CandidateJobGroupByInterviewResponse{
		Data: edge,
	}
	return result, nil
}

func (svc *candidateJobSvcImpl) ValidProcessingCdJobExistByCdID(ctx context.Context, candidateID uuid.UUID) (bool, error) {
	processingCdJobExist, err := svc.repoRegistry.CandidateJob().BuildExist(
		ctx,
		svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(
			candidatejob.CandidateID(candidateID),
			candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing, candidatejob.StatusOffering),
		),
	)
	if err != nil {
		svc.logger.Error(err.Error())
		return false, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return processingCdJobExist, nil
}

// common function
func (svc *candidateJobSvcImpl) freeWord(candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobFreeWord) {
	var predicate []predicate.CandidateJob
	if input != nil {
		if input.Team != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.HasHiringTeamEdgeWith(
					hiringteam.NameContainsFold(strings.TrimSpace(*input.Team)),
				),
			))
		}
		if input.HiringJob != nil {
			predicate = append(predicate, candidatejob.HasHiringJobEdgeWith(
				hiringjob.NameContainsFold(strings.TrimSpace(*input.HiringJob)),
			))
		}
		if input.CandidateName != nil {
			predicate = append(predicate, candidatejob.HasCandidateEdgeWith(
				candidate.NameContainsFold(strings.TrimSpace(*input.CandidateName)),
			))
		}
		if input.CandidateEmail != nil {
			predicate = append(predicate, candidatejob.HasCandidateEdgeWith(
				candidate.EmailContainsFold(strings.TrimSpace(*input.CandidateEmail)),
			))
		}
	}
	if len(predicate) > 0 {
		candidateJobQuery.Where(candidatejob.Or(predicate...))
	}
}

func (svc *candidateJobSvcImpl) filter(ctx context.Context, candidateJobQuery *ent.CandidateJobQuery, input *ent.CandidateJobFilter) {
	if input == nil {
		return
	}
	if input.CandidateID != nil {
		candidateJobQuery.Where(candidatejob.CandidateIDEQ(uuid.MustParse(*input.CandidateID)))
	}
	if input.Status != nil {
		candidateJobQuery.Where(candidatejob.StatusEQ(candidatejob.Status(*input.Status)))
	}
	if input.HiringTeamIds != nil {
		hiringTeamIds := lo.Map(input.HiringTeamIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasHiringTeamEdgeWith(
				hiringteam.IDIn(hiringTeamIds...),
			),
		))
	}
	if input.RecInChargeIds != nil {
		recInChargeIds := lo.Map(input.RecInChargeIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.RecInChargeIDIn(recInChargeIds...))
	}
	if input.RecTeamIds != nil {
		recTeamIds := lo.Map(input.RecTeamIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasRecTeamEdgeWith(
				recteam.IDIn(recTeamIds...),
			),
		))
	}
	if input.HiringJobIds != nil {
		hiringJobIds := lo.Map(input.HiringJobIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.IDIn(hiringJobIds...),
		))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
	if input.FailedReasons != nil && len(input.FailedReasons) != 0 {
		var candidateJobIds []uuid.UUID
		queryString := "SELECT id FROM candidate_jobs WHERE "
		for i, reason := range input.FailedReasons {
			queryString += "failed_reason @> '[\"" + reason.String() + "\"]'::jsonb"
			if i != len(input.FailedReasons)-1 {
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
	if input.Levels != nil {
		levels := lo.Map(input.Levels, func(level ent.CandidateJobLevel, index int) candidatejob.Level {
			return candidatejob.Level(level)
		})
		candidateJobQuery.Where(candidatejob.LevelIn(levels...))
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
		if input.CandidateName != nil {
			predicate = append(predicate, candidatejob.HasCandidateEdgeWith(
				candidate.NameContainsFold(strings.TrimSpace(*input.CandidateName)),
			))
		}
		if input.CandidateEmail != nil {
			predicate = append(predicate, candidatejob.HasCandidateEdgeWith(
				candidate.EmailContainsFold(strings.TrimSpace(*input.CandidateEmail)),
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
	if input.Status != nil {
		candidateJobQuery.Where(candidatejob.StatusEQ(candidatejob.Status(*input.Status)))
	}
	if input.HiringJobIds != nil {
		hiringJobIds := lo.Map(input.HiringJobIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HiringJobIDIn(hiringJobIds...))
	}
	if input.HiringTeamIds != nil {
		hiringTeamIds := lo.Map(input.HiringTeamIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasHiringTeamEdgeWith(
				hiringteam.IDIn(hiringTeamIds...),
			),
		))
	}
	if input.RecInChargeIds != nil {
		recInChargeIds := lo.Map(input.RecInChargeIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.RecInChargeIDIn(recInChargeIds...))
	}
	if input.RecTeamIds != nil {
		recTeamIds := lo.Map(input.RecTeamIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasRecTeamEdgeWith(
				recteam.IDIn(recTeamIds...),
			),
		))
	}
	if input.Priorities != nil {
		priorities := lo.Map(input.Priorities, func(id int, index int) int {
			return id
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.PriorityIn(priorities...)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateJobQuery.Where(candidatejob.CreatedAtGTE(*input.FromDate), candidatejob.CreatedAtLTE(*input.ToDate))
	}
	if input.SkillIds != nil {
		skillIds := lo.Map(input.SkillIds, func(id string, index int) uuid.UUID {
			return uuid.MustParse(id)
		})
		candidateJobQuery.Where(candidatejob.HasHiringJobEdgeWith(
			hiringjob.HasHiringJobSkillEdgesWith(entityskill.SkillIDIn(skillIds...)),
		))
	}
	if input.Locations != nil {
		locations := lo.Map(input.Locations, func(location ent.LocationEnum, index int) hiringjob.Location {
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
	if input.Levels != nil {
		levels := lo.Map(input.Levels, func(level ent.CandidateJobLevel, index int) candidatejob.Level {
			return candidatejob.Level(level)
		})
		candidateJobQuery.Where(candidatejob.LevelIn(levels...))
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
	var (
		messages, outgoingEmails []models.MessageInput
		results                  []*ent.OutgoingEmail
	)
	emailTemplates, err := svc.repoRegistry.EmailTemplate().ValidAndGetEmailTemplates(ctx, oldRecord, newRecord)
	if err != nil {
		return err
	}
	if len(emailTemplates) == 0 {
		return nil
	}
	groupModule, err := svc.repoRegistry.CandidateJob().GetDataForKeyword(ctx, newRecord)
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
	results, err = svc.outgoingEmailSvc.CreateBulkOutgoingEmail(ctx, messages, newRecord.CandidateID)
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
	return svc.emailSvc.SentEmail(ctx, outgoingEmails)
}

// permission
func (svc candidateJobSvcImpl) validPermissionMutation(payload *middleware.Payload, teamRecord *ent.HiringTeam) bool {
	if payload.ForAll {
		return true
	}
	if payload.ForTeam {
		memberIds := lo.Map(teamRecord.Edges.HiringMemberEdges, func(item *ent.User, index int) uuid.UUID {
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
		query.Where(candidatejob.HasHiringJobEdgeWith(hiringjob.HasHiringTeamEdgeWith(
			hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))),
		)))
	}
}

// Path: service/candidate_job.service.go
