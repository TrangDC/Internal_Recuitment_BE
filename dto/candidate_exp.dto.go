package dto

import (
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateExpDto interface {
	ProcessAuditTrail(oldRecord []*ent.CandidateExp, newRecord []*ent.CandidateExp, recordAudit *models.AuditTrailData)
}

type candidateExpDtoImpl struct {
}

func NewCandidateExpDto() CandidateExpDto {
	return &candidateExpDtoImpl{}
}

// ProcessAuditTrails process audit trails for candidate exps.
//
// oldRecords contains the old records of candidate exps.
// newRecords contains the new records of candidate exps.
// newCdExp := record not in oldRecords but in newRecords
// delCdExp := record not in newRecords but in oldRecords
// updCdExp := record in both oldRecords and newRecords
func (d candidateExpDtoImpl) ProcessAuditTrail(oldRecords []*ent.CandidateExp, newRecords []*ent.CandidateExp, recordAudit *models.AuditTrailData) {
	cdExpAudit := models.AuditTrailData{
		Module:    CandidateExpI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return
	}
	oldRecordCps := lo.Map(oldRecords, func(entity *ent.CandidateExp, _ int) interface{} {
		return entity
	})
	newRecordCps := lo.Map(newRecords, func(entity *ent.CandidateExp, _ int) interface{} {
		return entity
	})
	credIds, updIds, delIds := FindCUDArray(oldRecordCps, newRecordCps)
	d.recordAuditCreated(credIds, newRecords, &cdExpAudit)
	d.recordAuditUpdated(updIds, oldRecords, newRecords, &cdExpAudit)
	d.recordAuditDeleted(delIds, oldRecords, &cdExpAudit)
	recordAudit.SubModule = append(recordAudit.SubModule, cdExpAudit)
}

func (d candidateExpDtoImpl) recordAuditCreated(creIds []uuid.UUID, newRecords []*ent.CandidateExp, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(newRecords, func(entity *ent.CandidateExp) bool {
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
			}
			entity = append(entity, models.AuditTrailCreateDelete{
				Field: fieldName,
				Value: valueField,
			})
		}
		recordAudit.Create = append(recordAudit.Create, entity)
	}
}

func (d candidateExpDtoImpl) recordAuditDeleted(creIds []uuid.UUID, oldRecords []*ent.CandidateExp, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(oldRecords, func(entity *ent.CandidateExp) bool {
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
			case "model.candidate_exps.is_current":
				valueField = d.isCurrentI18n(record.IsCurrent)
			}
			entity = append(entity, models.AuditTrailCreateDelete{
				Field: fieldName,
				Value: valueField,
			})
		}
		recordAudit.Delete = append(recordAudit.Delete, entity)
	}
}

func (d candidateExpDtoImpl) recordAuditUpdated(updIds []uuid.UUID, oldRecords []*ent.CandidateExp, newRecords []*ent.CandidateExp, recordAudit *models.AuditTrailData) {
	for _, v := range updIds {
		entity := []interface{}{}
		oldRecord, _ := lo.Find(oldRecords, func(entity *ent.CandidateExp) bool {
			return entity.ID == v
		})
		newRecord, _ := lo.Find(newRecords, func(entity *ent.CandidateExp) bool {
			return entity.ID == v
		})
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
				case "model.candidate_exps.is_current":
					oldValueField = d.isCurrentI18n(oldRecord.IsCurrent)
					newValueField = d.isCurrentI18n(newRecord.IsCurrent)
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
		recordAudit.Update = append(recordAudit.Update, entity)
	}
}

// common
func (d candidateExpDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Position":
		return "model.candidate_exps.position"
	case "Company":
		return "model.candidate_exps.company"
	case "Location":
		return "model.candidate_exps.location"
	case "Description":
		return "model.candidate_exps.description"
	case "StartDate":
		return "model.candidate_exps.start_date"
	case "EndDate":
		return "model.candidate_exps.end_date"
	case "IsCurrent":
		return "model.candidate_exps.is_current"
	}
	return ""
}

func (d *candidateExpDtoImpl) isCurrentI18n(input bool) string {
	switch input {
	case true:
		return "model.candidate_educates.is_current_enum.yes"
	default:
		return "model.candidate_educates.is_current_enum.no"
	}
}
