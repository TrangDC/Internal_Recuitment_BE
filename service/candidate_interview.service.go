package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"trec/config"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/emailtemplate"
	"trec/ent/hiringjob"
	"trec/ent/hiringteam"
	"trec/ent/predicate"
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

type CandidateInterviewService interface {
	// mutation
	CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, note string) (*ent.CandidateInterviewResponse, error)
	DeleteCandidateInterview(ctx context.Context, id uuid.UUID, note string) error
	UpdateCandidateInterview(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewInput, note string) (*ent.CandidateInterviewResponse, error)
	UpdateCandidateInterviewSchedule(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterviewResponse, error)
	CreateCandidateInterview4Calendar(ctx context.Context, input ent.NewCandidateInterview4CalendarInput, note string) error
	UpdateCandidateInterviewStatus(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewStatusInput, note string) error

	// query
	GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterviewResponse, error)
	GetCandidateInterviews(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter ent.CandidateInterviewFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error)
	GetAllCandidateInterview4Calendar(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter *ent.CandidateInterviewCalendarFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error)

	// validate
	ValidateCandidateInterview(ctx context.Context, input ent.CandidateInterviewValidateInput) (*ent.CandidateInterviewResponseValidate, error)
}

type candidateInterviewSvcImpl struct {
	candidateInterviewerSvc CandidateInterviewerService
	emailSvc                EmailService
	outgoingEmailSvc        OutgoingEmailService
	serviceBusClient        servicebus.ServiceBus
	repoRegistry            repository.Repository
	dtoRegistry             dto.Dto
	logger                  *zap.Logger
}

func NewCandidateInterviewService(repoRegistry repository.Repository, serviceBusClient servicebus.ServiceBus, dtoRegistry dto.Dto, logger *zap.Logger, configs *config.Configurations) CandidateInterviewService {
	return &candidateInterviewSvcImpl{
		candidateInterviewerSvc: NewCandidateInterviewerService(repoRegistry, logger),
		emailSvc:                NewEmailService(repoRegistry, serviceBusClient, dtoRegistry, logger, configs),
		outgoingEmailSvc:        NewOutgoingEmailService(repoRegistry, logger),
		serviceBusClient:        serviceBusClient,
		repoRegistry:            repoRegistry,
		dtoRegistry:             dtoRegistry,
		logger:                  logger,
	}
}

func (svc *candidateInterviewSvcImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, note string) (*ent.CandidateInterviewResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var candidateInterview *ent.CandidateInterview
	var memberIds []uuid.UUID
	var inputValidate models.CandidateInterviewInputValidate
	query := svc.repoRegistry.CandidateJob().BuildBaseQuery().Where(candidatejob.IDEQ(uuid.MustParse(input.CandidateJobID)))
	candidateJob, err := svc.repoRegistry.CandidateJob().GetOneCandidateJob(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(candidateJob.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionCreate(payload, hiringJob.Edges.HiringTeamEdge) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	jsonString, _ := json.Marshal(input)
	json.Unmarshal(jsonString, &inputValidate)
	candidateJobStatus, stringError, err := svc.repoRegistry.CandidateInterview().ValidateInput(ctx, uuid.Nil, inputValidate)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds = lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().CreateCandidateInterview(ctx, input, memberIds, candidateJobStatus)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	err = svc.triggerEventSendEmail(ctx, result, candidateJob, emailtemplate.EventCreatedInterview)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	if result.Status == candidateinterview.StatusInvitedToInterview {
		err = svc.scheduleUpdateStatus(ctx, result)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
	}
	atJsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailCreate(result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, atJsonString, audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc candidateInterviewSvcImpl) CreateCandidateInterview4Calendar(ctx context.Context, input ent.NewCandidateInterview4CalendarInput, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var results []*ent.CandidateInterview
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(uuid.MustParse(input.JobID))).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionCreate(payload, hiringJob.Edges.HiringTeamEdge) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	candidateJobs, stringError, err := svc.repoRegistry.CandidateInterview().ValidateCreateBulkInput(ctx, input)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	memberIds := lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		results, err = repoRegistry.CandidateInterview().CreateBulkCandidateInterview(ctx, candidateJobs, memberIds, input)
		return err
	})
	for _, candidateInterview := range results {
		candidateJobEdge, _ := lo.Find(candidateJobs, func(candidateJob *ent.CandidateJob) bool {
			return candidateJob.ID == candidateInterview.CandidateJobID
		})
		err = svc.triggerEventSendEmail(ctx, candidateInterview, candidateJobEdge, emailtemplate.EventUpdatingInterview)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
	}
	candidateInterviewIds := lo.Map(results, func(candidateInterview *ent.CandidateInterview, index int) uuid.UUID {
		return candidateInterview.ID
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	candidateInterviews, err := svc.repoRegistry.CandidateInterview().BuildList(ctx,
		svc.repoRegistry.CandidateInterview().BuildQuery().Where(candidateinterview.IDIn(candidateInterviewIds...)))
	var createBulkAuditTrail []models.CandidateInterviewAuditTrail
	for _, entity := range candidateInterviews {
		atJsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailCreate(entity)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
		createBulkAuditTrail = append(createBulkAuditTrail, models.CandidateInterviewAuditTrail{
			RecordId:   entity.Edges.CandidateJobEdge.CandidateID,
			JsonString: atJsonString,
		})
	}
	err = svc.repoRegistry.AuditTrail().CreateBulkCandidateInterviewAt(ctx, createBulkAuditTrail, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc candidateInterviewSvcImpl) UpdateCandidateInterview(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewInput, note string) (*ent.CandidateInterviewResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var candidateInterview *ent.CandidateInterview
	var inputValidate models.CandidateInterviewInputValidate
	jsonString, _ := json.Marshal(input)
	json.Unmarshal(jsonString, &inputValidate)
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.Edges.CandidateJobEdge.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, hiringJob.Edges.HiringTeamEdge, record) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if time.Now().UTC().After(record.EndAt) {
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_interview_ended", http.StatusBadRequest, util.ErrorFlagCanNotUpdate)
	}
	if record.CandidateJobStatus.String() != record.Edges.CandidateJobEdge.Status.String() {
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_job_status_changed", http.StatusBadRequest, util.ErrorFlagCanNotUpdate)
	}
	memberIds := lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
		return uuid.MustParse(member)
	})
	newMemberIds, removeMemberIds := svc.updateMembers(record, memberIds)
	_, stringError, err := svc.repoRegistry.CandidateInterview().ValidateInput(ctx, id, inputValidate)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().UpdateCandidateInterview(ctx, record, input, newMemberIds, removeMemberIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	err = svc.triggerEventSendEmail(ctx, candidateInterview, record.Edges.CandidateJobEdge, emailtemplate.EventUpdatingInterview)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	result, _ := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	atJsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, atJsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc candidateInterviewSvcImpl) UpdateCandidateInterviewStatus(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewStatusInput, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var candidateInterview *ent.CandidateInterview
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.Edges.CandidateJobEdge.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, hiringJob.Edges.HiringTeamEdge, record) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if candidatejob.Status(record.CandidateJobStatus) != record.Edges.CandidateJobEdge.Status {
		return util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_job_status_changed", http.StatusBadRequest, util.ErrorFlagCanNotUpdate)
	}
	err = svc.repoRegistry.CandidateInterview().ValidCandidateInterviewStatus(record, input.Status)
	if err != nil {
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().UpdateCandidateInterviewStatus(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if record.Status != candidateinterview.StatusCancelled && input.Status == ent.CandidateInterviewStatusEditableCancelled {
		err = svc.triggerEventSendEmail(ctx, candidateInterview, record.Edges.CandidateJobEdge, emailtemplate.EventCancelInterview)
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
		}
	}
	result, _ := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	atJsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, atJsonString, audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc candidateInterviewSvcImpl) UpdateCandidateInterviewSchedule(ctx context.Context, id uuid.UUID, input ent.UpdateCandidateInterviewScheduleInput) (*ent.CandidateInterviewResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var candidateInterview *ent.CandidateInterview
	var inputValidate models.CandidateInterviewInputValidate
	var newMemberIds, removeMemberIds []uuid.UUID
	jsonString, _ := json.Marshal(input)
	json.Unmarshal(jsonString, &inputValidate)
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.Edges.CandidateJobEdge.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, hiringJob.Edges.HiringTeamEdge, record) {
		return nil, util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if time.Now().UTC().After(record.EndAt) {
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_interview_ended", http.StatusBadRequest, util.ErrorFlagCanNotUpdate)
	}
	if record.CandidateJobStatus.String() != record.Edges.CandidateJobEdge.Status.String() {
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_job_status_changed", http.StatusBadRequest, util.ErrorFlagCanNotUpdate)
	}
	inputValidate.CandidateJobId = record.CandidateJobID
	inputValidate.Title = record.Title
	_, stringError, err := svc.repoRegistry.CandidateInterview().ValidateInput(ctx, id, inputValidate)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	if stringError != nil {
		svc.logger.Error(stringError.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, stringError.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	if input.Interviewer != nil || len(input.Interviewer) == 0 {
		memberIds := lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
		newMemberIds, removeMemberIds = svc.updateMembers(record, memberIds)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		candidateInterview, err = repoRegistry.CandidateInterview().UpdateCandidateInterviewSchedule(ctx, record, input, newMemberIds, removeMemberIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	result, _ := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, candidateInterview.ID)
	atJsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailUpdate(record, result)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, atJsonString, audittrail.ActionTypeUpdate, "")
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc *candidateInterviewSvcImpl) GetCandidateInterview(ctx context.Context, id uuid.UUID) (*ent.CandidateInterviewResponse, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	query := svc.repoRegistry.CandidateInterview().BuildQuery().Where(candidateinterview.IDEQ(id))
	svc.validPermissionGet(payload, query)
	result, err := svc.repoRegistry.CandidateInterview().BuildGetOne(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc *candidateInterviewSvcImpl) DeleteCandidateInterview(ctx context.Context, id uuid.UUID, note string) error {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	hiringJobQuery := svc.repoRegistry.HiringJob().BuildBaseQuery().Where(hiringjob.IDEQ(record.Edges.CandidateJobEdge.HiringJobID)).WithHiringTeamEdge(
		func(query *ent.HiringTeamQuery) {
			query.WithUserEdges(
				func(query *ent.UserQuery) {
					query.Where(user.DeletedAtIsNil())
				},
			)
		},
	)
	hiringJob, err := svc.repoRegistry.HiringJob().BuildGetOne(ctx, hiringJobQuery)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusNotFound, util.ErrorFlagNotFound)
	}
	if !svc.validPermissionUpdate(payload, hiringJob.Edges.HiringTeamEdge, record) {
		return util.WrapGQLError(ctx, "Permission Denied", http.StatusForbidden, util.ErrorFlagPermissionDenied)
	}
	if record.CandidateJobStatus.String() != record.Edges.CandidateJobEdge.Status.String() {
		return util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_job_status_changed", http.StatusBadRequest, util.ErrorFlagCanNotDelete)
	}
	memberIds := lo.Map(record.Edges.InterviewerEdges, func(user *ent.User, index int) uuid.UUID {
		return user.ID
	})
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		_, err = repoRegistry.CandidateInterview().DeleteCandidateInterview(ctx, record, memberIds)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	jsonString, err := svc.dtoRegistry.CandidateInterview().AuditTrailDelete(record)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil
	}
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.Edges.CandidateJobEdge.CandidateID, audittrail.ModuleCandidates, jsonString, audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

func (svc *candidateInterviewSvcImpl) GetCandidateInterviews(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter ent.CandidateInterviewFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateInterviewResponseGetAll
	var edges []*ent.CandidateInterviewEdge
	var page int
	var perPage int
	query := svc.repoRegistry.CandidateInterview().BuildQuery()
	if filter.CandidateJobID != nil {
		candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, uuid.MustParse(*filter.CandidateJobID))
		if err != nil {
			svc.logger.Error(err.Error(), zap.Error(err))
			return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_not_found", http.StatusNotFound, util.ErrorFlagNotFound)
		}
		query.Where(
			candidateinterview.CandidateJobIDEQ(uuid.MustParse(*filter.CandidateJobID)),
			candidateinterview.CandidateJobStatusEQ(candidateinterview.CandidateJobStatus(candidateJob.Status.String())),
		)
	}
	svc.validPermissionGet(payload, query)
	var newFilter models.CandidateInterviewFilter
	jsonString, _ := json.Marshal(filter)
	json.Unmarshal(jsonString, &newFilter)
	svc.filter(query, newFilter)
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

func (svc *candidateInterviewSvcImpl) GetAllCandidateInterview4Calendar(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateInterviewFreeWord, filter *ent.CandidateInterviewCalendarFilter, orderBy *ent.CandidateInterviewOrder) (*ent.CandidateInterviewResponseGetAll, error) {
	payload := ctx.Value(middleware.Payload{}).(*middleware.Payload)
	var result *ent.CandidateInterviewResponseGetAll
	var edges []*ent.CandidateInterviewEdge
	var page int
	var perPage int
	candidateJobStatusEnded := lo.Map(ent.AllCandidateJobStatusEnded, func(status ent.CandidateJobStatusEnded, index int) candidatejob.Status {
		return candidatejob.Status(status)
	})
	query := svc.repoRegistry.CandidateInterview().BuildQuery().Where(
		candidateinterview.HasCandidateJobEdgeWith(
			candidatejob.DeletedAtIsNil(), candidatejob.StatusNotIn(candidateJobStatusEnded...),
			candidatejob.HasCandidateEdgeWith(
				candidate.DeletedAtIsNil(), candidate.IsBlacklist(false),
			),
		),
	)
	svc.validPermissionGet(payload, query)
	var newFilter models.CandidateInterviewFilter
	jsonString, _ := json.Marshal(filter)
	json.Unmarshal(jsonString, &newFilter)
	svc.filter(query, newFilter)
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

// third func
func (svc candidateInterviewSvcImpl) triggerEventSendEmail(ctx context.Context, interviewRecord *ent.CandidateInterview,
	candidateJob *ent.CandidateJob, event emailtemplate.Event) error {
	var messages []models.MessageInput
	var outgoingEmails []models.MessageInput
	var results []*ent.OutgoingEmail
	emailTemplates, err := svc.repoRegistry.EmailTemplate().GetEmailTpInterviewEvent(ctx, event)
	if err != nil {
		return err
	}
	if len(emailTemplates) == 0 {
		return nil
	}
	groupModule, err := svc.repoRegistry.CandidateInterview().GetDataForKeyword(ctx, interviewRecord, candidateJob)
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
	results, err = svc.outgoingEmailSvc.CreateBulkOutgoingEmail(ctx, messages, candidateJob.CandidateID)
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

func (svc candidateInterviewSvcImpl) scheduleUpdateStatus(ctx context.Context, record *ent.CandidateInterview) error {
	err := svc.serviceBusClient.SendInterviewScheduleMessage(ctx, models.MessageOutput{
		ID:        record.ID.String(),
		IsSuccess: true,
		QueueName: servicebus.InterviewScheduleQueue,
	}, record.StartFrom)
	return err
}

// validate
func (svc *candidateInterviewSvcImpl) ValidateCandidateInterview(ctx context.Context, input ent.CandidateInterviewValidateInput) (*ent.CandidateInterviewResponseValidate, error) {
	// if  candidate_id is not nil => hiring_job_id
	// validate status
	// validate title
	// validate schedule
	// validate interviewer
	return nil, nil
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

func (svc *candidateInterviewSvcImpl) filter(candidateInterviewQuery *ent.CandidateInterviewQuery, input models.CandidateInterviewFilter) {
	if input.InterviewDate != nil {
		candidateInterviewQuery.Where(candidateinterview.InterviewDateEQ(*input.InterviewDate))
	}
	if input.StartFrom != nil && input.EndAt != nil {
		candidateInterviewQuery.Where(candidateinterview.And(candidateinterview.StartFromGTE(*input.StartFrom), candidateinterview.EndAtLTE(*input.EndAt)))
	}
	if input.FromDate != nil && input.ToDate != nil {
		candidateInterviewQuery.Where(candidateinterview.And(candidateinterview.CreatedAtGTE(*input.FromDate), candidateinterview.CreatedAtLTE(*input.ToDate)))
	}
	if input.InterviewDateFrom != nil && input.InterviewDateTo != nil {
		candidateInterviewQuery.Where(candidateinterview.And(candidateinterview.InterviewDateGTE(*input.InterviewDateFrom), candidateinterview.InterviewDateLTE(*input.InterviewDateTo)))
	}
	if input.Interviewer != nil {
		userIds := lo.Map(input.Interviewer, func(member string, index int) uuid.UUID {
			return uuid.MustParse(member)
		})
		candidateInterviewQuery.Where(candidateinterview.HasInterviewerEdgesWith(user.IDIn(userIds...)))
	}
	if input.HiringTeamId != nil {
		candidateInterviewQuery.Where(candidateinterview.HasCandidateJobEdgeWith(
			candidatejob.HasHiringJobEdgeWith(
				hiringjob.HasHiringTeamEdgeWith(
					hiringteam.IDEQ(uuid.MustParse(*input.HiringTeamId)),
				),
			),
		))
	}
	if input.HiringJobId != nil {
		candidateInterviewQuery.Where(candidateinterview.CandidateJobIDEQ(uuid.MustParse(*input.HiringJobId)))
	}
	if input.CandidateId != nil {
		candidateInterviewQuery.Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.CandidateIDEQ(uuid.MustParse(*input.CandidateId))))
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

// permission
func (svc candidateInterviewSvcImpl) validPermissionCreate(payload *middleware.Payload, teamRecord *ent.HiringTeam) bool {
	if payload.ForAll {
		return true
	}
	memberIds := lo.Map(teamRecord.Edges.HiringMemberEdges, func(item *ent.User, index int) uuid.UUID {
		return item.ID
	})
	managerIds := lo.Map(teamRecord.Edges.UserEdges, func(item *ent.User, index int) uuid.UUID {
		return item.ID
	})
	if payload.ForTeam {
		if lo.Contains(memberIds, payload.UserID) || lo.Contains(managerIds, payload.UserID) {
			return true
		}
	}
	return false
}

func (svc candidateInterviewSvcImpl) validPermissionUpdate(payload *middleware.Payload, teamRecord *ent.HiringTeam, record *ent.CandidateInterview) bool {
	interviewerIds := lo.Map(record.Edges.InterviewerEdges, func(item *ent.User, index int) uuid.UUID {
		return item.ID
	})
	if payload.ForOwner {
		return lo.Contains(interviewerIds, payload.UserID)
	}
	return svc.validPermissionCreate(payload, teamRecord)
}

func (svc candidateInterviewSvcImpl) validPermissionGet(payload *middleware.Payload, query *ent.CandidateInterviewQuery) {
	if payload.ForAll {
		return
	}
	if payload.ForTeam {
		query.Where(candidateinterview.HasCandidateJobEdgeWith(candidatejob.HasHiringJobEdgeWith(hiringjob.HasHiringTeamEdgeWith(
			hiringteam.Or(hiringteam.HasUserEdgesWith(user.IDEQ(payload.UserID)), hiringteam.HasHiringMemberEdgesWith(user.IDEQ(payload.UserID))),
		))))
	}
	if payload.ForOwner {
		query.Where(candidateinterview.HasInterviewerEdgesWith(
			user.IDEQ(payload.UserID), user.DeletedAtIsNil(),
		))
	}
}

// Path: service/candidate_interview.service.go
