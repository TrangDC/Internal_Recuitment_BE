package service

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"trec/dto"
	"trec/ent"
	"trec/ent/audittrail"
	"trec/ent/candidate"
	"trec/ent/candidatejob"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type CandidateService interface {
	// mutation
	CreateCandidate(ctx context.Context, input *ent.NewCandidateInput, note string) (*ent.CandidateResponse, error)
	UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID, note string) (*ent.CandidateResponse, error)
	DeleteCandidate(ctx context.Context, id uuid.UUID, note string) error
	SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool, note string) error
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

func (svc *candidateSvcImpl) CreateCandidate(ctx context.Context, input *ent.NewCandidateInput, note string) (*ent.CandidateResponse, error) {
	var record *ent.Candidate
	err := svc.repoRegistry.Candidate().ValidEmail(ctx, uuid.Nil, input.Email)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.Candidate().CreateCandidate(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	recordChanges := svc.recordCreateDelete(record, audittrail.ActionTypeCreate)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleCandidates, string(recordChangesJson), audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateResponse{
		Data: record,
	}, nil
}

func (svc *candidateSvcImpl) DeleteCandidate(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	if len(record.Edges.CandidateJobEdges) > 0 {
		return util.WrapGQLError(ctx, "model.candidates.validation.candidate_job_exist", http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.Candidate().DeleteCandidate(ctx, record)
		return err
	})
	recordChanges := svc.recordCreateDelete(record, audittrail.ActionTypeDelete)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleCandidates, string(recordChangesJson), audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *candidateSvcImpl) UpdateCandidate(ctx context.Context, input *ent.UpdateCandidateInput, id uuid.UUID, note string) (*ent.CandidateResponse, error) {
	var result *ent.Candidate
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
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
		result, err = repoRegistry.Candidate().UpdateCandidate(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	recordChanges := svc.recordUpdate(record, result)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleCandidates, string(recordChangesJson), audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.CandidateResponse{
		Data: result,
	}, nil
}

func (svc *candidateSvcImpl) SetBlackListCandidate(ctx context.Context, id uuid.UUID, isBlackList bool, note string) error {
	var result *ent.Candidate
	record, err := svc.repoRegistry.Candidate().GetCandidate(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.Candidate().SetBlackListCandidate(ctx, record, isBlackList)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	recordChanges := svc.recordUpdate(record, result)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleHiringJobs, string(recordChangesJson), audittrail.ActionTypeUpdate, note)
	return nil
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

func (svc *candidateSvcImpl) GetCandidates(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.CandidateFreeWord, filter *ent.CandidateFilter, orderBy *ent.CandidateOrder) (*ent.CandidateResponseGetAll, error) {
	var result *ent.CandidateResponseGetAll
	var edges []*ent.CandidateEdge
	var page int
	var perPage int
	query := svc.repoRegistry.Candidate().BuildQuery()
	svc.filter(ctx, query, filter)
	svc.freeWord(query, freeWord)
	if filter != nil && filter.JobID != nil {
		if filter.IsAbleToInterview != nil && *filter.IsAbleToInterview {
			query = query.Where(candidate.HasCandidateJobEdgesWith(
				candidatejob.HiringJobIDEQ(uuid.MustParse(*filter.JobID)),
				candidatejob.StatusIn(candidatejob.StatusApplied, candidatejob.StatusInterviewing),
			))
		} else {
			query = query.Where(candidate.HasCandidateJobEdgesWith(candidatejob.HiringJobIDEQ(uuid.MustParse(*filter.JobID))))
		}
	}
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

func (svc *candidateSvcImpl) filter(ctx context.Context, candidateQuery *ent.CandidateQuery, input *ent.CandidateFilter) {
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
		if input.Status != nil {
			candidateQuery.Where(candidate.HasCandidateJobEdgesWith(
				candidatejob.StatusEQ(candidatejob.Status(*input.Status)),
			))
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
			rows, _ := candidateQuery.QueryContext(ctx, queryString)
			if rows != nil {
				for rows.Next() {
					var id uuid.UUID
					rows.Scan(&id)
					candidateJobIds = append(candidateJobIds, id)
				}
				candidateQuery.Where(candidate.HasCandidateJobEdgesWith(
					candidatejob.IDIn(candidateJobIds...),
				))
			} else {
				candidateQuery.Where(candidate.IDIn(uuid.Nil))
			}
		}
	}
}

func (svc *candidateSvcImpl) recordCreateDelete(record *ent.Candidate, auditTrailType audittrail.ActionType) models.AuditTrailData {
	auditTrail := models.AuditTrailData{
		Module: "model.candidates.model_name",
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	result := []interface{}{}
	value := reflect.ValueOf(interface{}(record)).Elem()
	recordType := reflect.TypeOf(record).Elem()
	for i := 1; i < value.NumField(); i++ {
		field := recordType.Field(i)
		valueField := value.Field(i).Interface()
		fieldName := dto.FormatCandidateField(field.Name)
		switch fieldName {
		case "":
			continue
		case "model.candidates.is_blacklist":
			valueField = dto.IsBlacklistI18n(record.IsBlacklist)
		}
		result = append(result, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	if auditTrailType == audittrail.ActionTypeCreate {
		auditTrail.Create = append(auditTrail.Create, result...)
	}
	if auditTrailType == audittrail.ActionTypeDelete {
		auditTrail.Delete = append(auditTrail.Delete, result...)
	}
	return auditTrail
}

func (svc *candidateSvcImpl) recordUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate) models.AuditTrailData {
	auditTrail := models.AuditTrailData{
		Module: "model.candidates.model_name",
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	result := []interface{}{}
	oldValue := reflect.ValueOf(interface{}(oldRecord)).Elem()
	newValue := reflect.ValueOf(interface{}(newRecord)).Elem()
	recordType := reflect.TypeOf(oldRecord).Elem()
	for i := 1; i < oldValue.NumField(); i++ {
		field := recordType.Field(i)
		oldValueField := oldValue.Field(i).Interface()
		newValueField := newValue.Field(i).Interface()
		fieldName := dto.FormatCandidateField(field.Name)
		if field.PkgPath == "" && !reflect.DeepEqual(oldValueField, newValueField) {
			switch fieldName {
			case "":
				continue
			case "model.candidates.is_blacklist":
				oldValueField = dto.IsBlacklistI18n(oldRecord.IsBlacklist)
				newValueField = dto.IsBlacklistI18n(newRecord.IsBlacklist)
			}
			result = append(result, models.AuditTrailUpdate{
				Field: fieldName,
				Value: models.ValueChange{
					OldValue: oldValueField,
					NewValue: newValueField,
				},
			})
		}
	}
	auditTrail.Update = append(auditTrail.Update, result...)
	return auditTrail
}
