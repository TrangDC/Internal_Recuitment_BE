package dto

import (
	"reflect"
	"trec/ent"
	"trec/ent/hiringjobstep"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type HiringJobStepDto interface {
	ProcessAuditTrail(oldRecord []*ent.HiringJobStep, newRecord []*ent.HiringJobStep, recordAudit *models.AuditTrailData)
}

type hiringJobStepDtoImpl struct {
}

var hiringJobStepFieldI18n = map[string]models.I18nFormat{
	"Status": {
		AuditTrail: "model.hiring_job_steps.status",
		Email:      "Status",
	},
	"UserID": {
		AuditTrail: "model.hiring_job_steps.user",
		Email:      "Approver",
	},
}

func NewHiringJobStepDto() HiringJobStepDto {
	return &hiringJobStepDtoImpl{}
}

// ProcessAuditTrails process audit trails for hiring job steps.
//
// oldRecords contains the old records of hiring job steps.
// newRecords contains the new records of hiring job steps.
// newHiringJobStep := record not in oldRecords but in newRecords
// delHiringJobStep := record not in newRecords but in oldRecords
// updHiringJobStep := record in both oldRecords and newRecords
func (d hiringJobStepDtoImpl) ProcessAuditTrail(oldRecords []*ent.HiringJobStep, newRecords []*ent.HiringJobStep, recordAudit *models.AuditTrailData) {
	hiringJobStepAudit := models.AuditTrailData{
		Module:    HiringJobStepI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return
	}
	oldRecordCps := lo.Map(oldRecords, func(entity *ent.HiringJobStep, _ int) interface{} {
		return entity
	})
	newRecordCps := lo.Map(newRecords, func(entity *ent.HiringJobStep, _ int) interface{} {
		return entity
	})
	credIds, updIds, delIds := FindCUDArray(oldRecordCps, newRecordCps)
	d.recordAuditCreated(credIds, newRecords, &hiringJobStepAudit)
	d.recordAuditUpdated(updIds, oldRecords, newRecords, &hiringJobStepAudit)
	d.recordAuditDeleted(delIds, oldRecords, &hiringJobStepAudit)
	recordAudit.SubModule = append(recordAudit.SubModule, hiringJobStepAudit)
}

func (d hiringJobStepDtoImpl) recordAuditCreated(creIds []uuid.UUID, newRecords []*ent.HiringJobStep, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(newRecords, func(entity *ent.HiringJobStep) bool {
			return entity.ID == v
		})
		value := reflect.ValueOf(interface{}(record)).Elem()
		recordType := reflect.TypeOf(record).Elem()
		for i := 1; i < value.NumField(); i++ {
			field := recordType.Field(i)
			valueField := value.Field(i).Interface()
			fieldName := d.formatFieldI18n(field.Name)
			switch fieldName {
			case "":
				continue
			case "model.hiring_job_steps.status":
				valueField = d.statusI18n(record.Status)
			case "model.hiring_job_steps.user":
				valueField = record.Edges.ApprovalUser.Name
			}
			entity = append(entity, models.AuditTrailCreateDelete{
				Field: fieldName,
				Value: valueField,
			})
		}
		recordAudit.Create = append(recordAudit.Create, entity)
	}
}

func (d hiringJobStepDtoImpl) recordAuditDeleted(creIds []uuid.UUID, oldRecords []*ent.HiringJobStep, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(oldRecords, func(entity *ent.HiringJobStep) bool {
			return entity.ID == v
		})
		value := reflect.ValueOf(interface{}(record)).Elem()
		recordType := reflect.TypeOf(record).Elem()
		for i := 1; i < value.NumField(); i++ {
			field := recordType.Field(i)
			valueField := value.Field(i).Interface()
			fieldName := d.formatFieldI18n(field.Name)
			switch fieldName {
			case "":
				continue
			case "model.hiring_job_steps.status":
				valueField = d.statusI18n(record.Status)
			case "model.hiring_job_steps.user":
				valueField = record.Edges.ApprovalUser.Name
			}
			entity = append(entity, models.AuditTrailCreateDelete{
				Field: fieldName,
				Value: valueField,
			})
		}
		recordAudit.Delete = append(recordAudit.Delete, entity)
	}
}

func (d hiringJobStepDtoImpl) recordAuditUpdated(updIds []uuid.UUID, oldRecords []*ent.HiringJobStep, newRecords []*ent.HiringJobStep, recordAudit *models.AuditTrailData) {
	for _, v := range updIds {
		entity := []interface{}{}
		oldRecord, _ := lo.Find(oldRecords, func(entity *ent.HiringJobStep) bool {
			return entity.ID == v
		})
		newRecord, _ := lo.Find(newRecords, func(entity *ent.HiringJobStep) bool {
			return entity.ID == v
		})
		if newRecord.Status != hiringjobstep.StatusPending || oldRecord.UserID != newRecord.UserID {
			oldValue := reflect.ValueOf(interface{}(oldRecord)).Elem()
			newValue := reflect.ValueOf(interface{}(newRecord)).Elem()
			recordType := reflect.TypeOf(oldRecord).Elem()
			for i := 1; i < oldValue.NumField(); i++ {
				field := recordType.Field(i)
				oldValueField := oldValue.Field(i).Interface()
				newValueField := newValue.Field(i).Interface()
				fieldName := d.formatFieldI18n(field.Name)
				if field.PkgPath == "" && !reflect.DeepEqual(oldValueField, newValueField) {
					switch fieldName {
					case "":
						continue
					case "model.hiring_job_steps.status":
						oldValueField = d.statusI18n(oldRecord.Status)
						newValueField = d.statusI18n(newRecord.Status)
					case "model.hiring_job_steps.user":
						oldValueField = oldRecord.Edges.ApprovalUser.Name
						newValueField = newRecord.Edges.ApprovalUser.Name
					}
					entity = append(entity, models.AuditTrailUpdate{
						Field: fieldName,
						Value: models.ValueChange{
							OldValue: oldValueField,
							NewValue: newValueField,
						},
					})
				}
			}
		}
		if len(entity) > 0 {
			recordAudit.Update = append(recordAudit.Update, entity)
		}
	}
}

// common
func (d hiringJobStepDtoImpl) formatFieldI18n(input string) string {
	if v, ok := hiringJobStepFieldI18n[input]; ok {
		return v.AuditTrail
	}
	return ""
}

func (d hiringJobStepDtoImpl) statusI18n(input hiringjobstep.Status) string {
	switch input {
	case hiringjobstep.StatusAccepted:
		return "model.hiring_job_steps.status_enum.accepted"
	case hiringjobstep.StatusPending:
		return "model.hiring_job_steps.status_enum.pending"
	case hiringjobstep.StatusWaiting:
		return "model.hiring_job_steps.status_enum.waiting"
	case hiringjobstep.StatusRejected:
		return "model.hiring_job_steps.status_enum.rejected"
	}
	return ""
}
