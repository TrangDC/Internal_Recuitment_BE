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
	"trec/ent/hiringjob"
	"trec/internal/util"
	"trec/models"
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
	GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord, filter *ent.HiringJobFilter, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error)
}
type hiringJobSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewHiringJobService(repoRegistry repository.Repository, logger *zap.Logger) HiringJobService {
	return &hiringJobSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc *hiringJobSvcImpl) CreateHiringJob(ctx context.Context, input *ent.NewHiringJobInput, note string) (*ent.HiringJobResponse, error) {
	var record *ent.HiringJob
	err := svc.repoRegistry.HiringJob().ValidName(ctx, uuid.Nil, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		record, err = repoRegistry.HiringJob().CreateHiringJob(ctx, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error())
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	svc.getAdditionalInfo(ctx, record)
	recordChanges := svc.recordCreateDelete(record, audittrail.ActionTypeCreate)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, string(recordChangesJson), audittrail.ActionTypeCreate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return &ent.HiringJobResponse{
		Data: record,
	}, nil
}

func (svc *hiringJobSvcImpl) DeleteHiringJob(ctx context.Context, id uuid.UUID, note string) error {
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		err = repoRegistry.HiringJob().DeleteHiringJob(ctx, record)
		return err
	})
	recordChanges := svc.recordCreateDelete(record, audittrail.ActionTypeDelete)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, record.ID, audittrail.ModuleHiringJobs, string(recordChangesJson), audittrail.ActionTypeDelete, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	return nil
}

func (svc *hiringJobSvcImpl) UpdateHiringJob(ctx context.Context, input *ent.UpdateHiringJobInput, id uuid.UUID, note string) (*ent.HiringJobResponse, error) {
	var result *ent.HiringJob
	record, err := svc.repoRegistry.HiringJob().GetHiringJob(ctx, id)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagNotFound)
	}
	err = svc.repoRegistry.HiringJob().ValidName(ctx, id, input.Name)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusBadRequest, util.ErrorFlagValidateFail)
	}
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringJob().UpdateHiringJob(ctx, record, input)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	svc.getAdditionalInfo(ctx, result)
	recordChanges := svc.recordUpdate(record, result)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleHiringJobs, string(recordChangesJson), audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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
	err = svc.repoRegistry.DoInTx(ctx, func(ctx context.Context, repoRegistry repository.Repository) error {
		result, err = repoRegistry.HiringJob().UpdateHiringJobStatus(ctx, record, status)
		return err
	})
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	recordChanges := svc.recordUpdate(record, result)
	recordChangesJson, _ := json.Marshal([]interface{}{recordChanges})
	err = svc.repoRegistry.AuditTrail().AuditTrailMutation(ctx, result.ID, audittrail.ModuleHiringJobs, string(recordChangesJson), audittrail.ActionTypeUpdate, note)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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

func (svc *hiringJobSvcImpl) GetHiringJobs(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.HiringJobFreeWord, filter *ent.HiringJobFilter, orderBy *ent.HiringJobOrder) (*ent.HiringJobResponseGetAll, error) {
	var result *ent.HiringJobResponseGetAll
	var edges []*ent.HiringJobEdge
	var page int
	var perPage int
	query := svc.repoRegistry.HiringJob().BuildQuery()
	svc.filter(query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.HiringJob().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(hiringjob.FieldCreatedAt)
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
	hiringJobs, err := svc.repoRegistry.HiringJob().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
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
	}
}

func (svc *hiringJobSvcImpl) recordCreateDelete(record *ent.HiringJob, auditTrailType audittrail.ActionType) models.AuditTrailData {
	auditTrail := models.AuditTrailData{
		Module: "models.hiring_jobs.model_name",
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
		fieldName := dto.FormatHiringJobField(field.Name)
		switch fieldName {
		case "":
			continue
		case "model.hiring_jobs.location":
			valueField = dto.LocationI18n(record.Location)
		case "model.hiring_jobs.status":
			valueField = dto.StatusI18n(record.Status)
		case "model.hiring_jobs.salary_type":
			valueField = dto.SalaryTypeI18n(record.SalaryType)
		case "model.hiring_jobs.currency":
			valueField = dto.CurrencyI18n(record.Currency)
		case "model.hiring_jobs.team":
			valueField = record.Edges.TeamEdge.Name
		case "model.hiring_jobs.created_by":
			valueField = record.Edges.OwnerEdge.Name
		}
		result = append(result, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	if auditTrailType == audittrail.ActionTypeCreate {
		auditTrail.Create = append(auditTrail.Create, result...)
	} else {
		auditTrail.Delete = append(auditTrail.Delete, result...)
	}
	return auditTrail
}

func (svc *hiringJobSvcImpl) recordUpdate(oldRecord *ent.HiringJob, newRecord *ent.HiringJob) models.AuditTrailData {
	auditTrail := models.AuditTrailData{
		Module: "models.hiring_jobs.model_name",
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
		fieldName := dto.FormatHiringJobField(field.Name)
		if field.PkgPath == "" && !reflect.DeepEqual(oldValueField, newValueField) {
			switch fieldName {
			case "":
				continue
			case "model.hiring_jobs.location":
				oldValueField = dto.LocationI18n(oldRecord.Location)
				newValueField = dto.LocationI18n(newRecord.Location)
			case "model.hiring_jobs.status":
				oldValueField = dto.StatusI18n(oldRecord.Status)
				newValueField = dto.StatusI18n(newRecord.Status)
			case "model.hiring_jobs.salary_type":
				oldValueField = dto.SalaryTypeI18n(oldRecord.SalaryType)
				newValueField = dto.SalaryTypeI18n(newRecord.SalaryType)
			case "model.hiring_jobs.currency":
				oldValueField = dto.CurrencyI18n(oldRecord.Currency)
				newValueField = dto.CurrencyI18n(newRecord.Currency)
			case "model.hiring_jobs.team":
				oldValueField = oldRecord.Edges.TeamEdge.Name
				newValueField = newRecord.Edges.TeamEdge.Name
			case "model.hiring_jobs.created_by":
				oldValueField = oldRecord.Edges.OwnerEdge.Name
				newValueField = newRecord.Edges.OwnerEdge.Name
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

func (svc *hiringJobSvcImpl) getAdditionalInfo(ctx context.Context, record *ent.HiringJob) {
	teamRecord, _ := svc.repoRegistry.Team().GetTeam(ctx, record.TeamID)
	userRecord, _ := svc.repoRegistry.User().GetUser(ctx, record.CreatedBy)
	record.Edges.TeamEdge = teamRecord
	record.Edges.OwnerEdge = userRecord
}
