package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidateinterview"
	"trec/ent/candidatejob"
	"trec/ent/hiringjob"
	"trec/ent/predicate"
	"trec/ent/team"
	"trec/ent/user"
	"trec/internal/util"
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
	repoRegistry            repository.Repository
	dtoRegistry             dto.Dto
	logger                  *zap.Logger
}

func NewCandidateInterviewService(repoRegistry repository.Repository, dtoRegistry dto.Dto, logger *zap.Logger) CandidateInterviewService {
	return &candidateInterviewSvcImpl{
		candidateInterviewerSvc: NewCandidateInterviewerService(repoRegistry, logger),
		repoRegistry:            repoRegistry,
		dtoRegistry:             dtoRegistry,
		logger:                  logger,
	}
}

func (svc *candidateInterviewSvcImpl) CreateCandidateInterview(ctx context.Context, input ent.NewCandidateInterviewInput, note string) (*ent.CandidateInterviewResponse, error) {
	var candidateInterview *ent.CandidateInterview
	var memberIds []uuid.UUID
	var inputValidate models.CandidateInterviewInputValidate
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
	var results []*ent.CandidateInterview
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
	var candidateInterview *ent.CandidateInterview
	var inputValidate models.CandidateInterviewInputValidate
	jsonString, _ := json.Marshal(input)
	json.Unmarshal(jsonString, &inputValidate)
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
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
	var candidateInterview *ent.CandidateInterview
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
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
	var candidateInterview *ent.CandidateInterview
	var inputValidate models.CandidateInterviewInputValidate
	var newMemberIds, removeMemberIds []uuid.UUID
	jsonString, _ := json.Marshal(input)
	json.Unmarshal(jsonString, &inputValidate)
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
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
	result, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	return &ent.CandidateInterviewResponse{
		Data: result,
	}, nil
}

func (svc *candidateInterviewSvcImpl) DeleteCandidateInterview(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.CandidateInterview().GetCandidateInterview(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
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
	var result *ent.CandidateInterviewResponseGetAll
	var edges []*ent.CandidateInterviewEdge
	var page int
	var perPage int
	candidateJob, err := svc.repoRegistry.CandidateJob().GetCandidateJob(ctx, uuid.MustParse(filter.CandidateJobID))
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, "model.candidate_interviews.validation.candidate_not_found", http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	query := svc.repoRegistry.CandidateInterview().BuildQuery().Where(
		candidateinterview.CandidateJobIDEQ(uuid.MustParse(filter.CandidateJobID)),
		candidateinterview.CandidateJobStatusEQ(candidateinterview.CandidateJobStatus(candidateJob.Status.String())),
	)
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
	if input.TeamId != nil {
		candidateInterviewQuery.Where(candidateinterview.HasCandidateJobEdgeWith(
			candidatejob.HasHiringJobEdgeWith(
				hiringjob.HasTeamEdgeWith(
					team.IDEQ(uuid.MustParse(*input.TeamId)),
				),
			),
		))
	}
	if input.HiringJobId != nil {
		candidateInterviewQuery.Where(candidateinterview.CandidateJobIDEQ(uuid.MustParse(*input.HiringJobId)))
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
