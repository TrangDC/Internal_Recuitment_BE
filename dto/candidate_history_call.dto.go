package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/candidatehistorycall"
	"trec/models"
)

type CandidateHistoryCallDto interface {
	AuditTrailCreate(record *ent.CandidateHistoryCall) (string, error)
	AuditTrailDelete(record *ent.CandidateHistoryCall) (string, error)
	AuditTrailUpdate(oldRecord *ent.CandidateHistoryCall, newRecord *ent.CandidateHistoryCall) (string, error)
}

type candidateHistoryCallDtoImpl struct {
}

func NewCandidateHistoryCallDto() CandidateHistoryCallDto {
	return &candidateHistoryCallDtoImpl{}
}

func (d candidateHistoryCallDtoImpl) AuditTrailCreate(record *ent.CandidateHistoryCall) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateHistoryCallI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d candidateHistoryCallDtoImpl) AuditTrailDelete(record *ent.CandidateHistoryCall) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateHistoryCallI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d candidateHistoryCallDtoImpl) AuditTrailUpdate(oldRecord *ent.CandidateHistoryCall, newRecord *ent.CandidateHistoryCall) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateHistoryCallI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	entity := []interface{}{}
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
			case "model.candidate_history_calls.type":
				oldValueField = d.typeI18n(oldRecord.Type)
				newValueField = d.typeI18n(newRecord.Type)
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
	result.Update = append(result.Update, entity...)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d candidateHistoryCallDtoImpl) recordAudit(record *ent.CandidateHistoryCall) []interface{} {
	entity := []interface{}{}
	value := reflect.ValueOf(interface{}(record)).Elem()
	recordType := reflect.TypeOf(record).Elem()
	for i := 1; i < value.NumField(); i++ {
		field := recordType.Field(i)
		valueField := value.Field(i).Interface()
		fieldName := d.formatFieldI18n(field.Name)
		switch fieldName {
		case "":
			continue
		case "model.candidate_history_calls.candidate":
			valueField = record.Edges.CandidateEdge.Name
		case "model.candidate_history_calls.created_by":
			valueField = record.Edges.CreatedByEdge.Name
		case "model.candidate_history_calls.type":
			valueField = d.typeI18n(record.Type)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d candidateHistoryCallDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidate_history_calls.name"
	case "CandidateID":
		return "model.candidate_history_calls.candidate"
	case "CreatedByID":
		return "model.candidate_history_calls.created_by"
	case "ContactTo":
		return "model.candidate_history_calls.contact_to"
	case "Description":
		return "model.candidate_history_calls.description"
	case "Type":
		return "model.candidate_history_calls.type"
	case "Date":
		return "model.candidate_history_calls.date"
	case "StartTime":
		return "model.candidate_history_calls.start_time"
	case "EndTime":
		return "model.candidate_history_calls.end_time"
	}
	return ""
}

func (d candidateHistoryCallDtoImpl) typeI18n(input candidatehistorycall.Type) string {
	switch input {
	case candidatehistorycall.TypeCandidate:
		return "model.candidate_history_calls.type_enum.candidate"
	case candidatehistorycall.TypeOthers:
		return "model.candidate_history_calls.type_enum.others"
	}
	return ""
}
