package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"
)

type CandidateDto interface {
	AuditTrailCreate(record *ent.Candidate) (string, error)
	AuditTrailDelete(record *ent.Candidate) (string, error)
	AuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate) (string, error)
}

type candidateDtoImpl struct {
}

func NewCandidateDto() CandidateDto {
	return &candidateDtoImpl{}
}

func (d *candidateDtoImpl) AuditTrailCreate(record *ent.Candidate) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateDtoImpl) AuditTrailDelete(record *ent.Candidate) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateDtoImpl) AuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate) (string, error) {
	auditTrail := models.AuditTrailData{
		Module: CandidateI18n,
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
		fieldName := d.formatFieldI18n(field.Name)
		if field.PkgPath == "" && !reflect.DeepEqual(oldValueField, newValueField) {
			switch fieldName {
			case "":
				continue
			case "model.candidates.is_blacklist":
				oldValueField = d.isBlacklistI18n(oldRecord.IsBlacklist)
				newValueField = d.isBlacklistI18n(newRecord.IsBlacklist)
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
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}


func (d *candidateDtoImpl) recordAudit(record *ent.Candidate) ([]interface{}) {
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
		case "model.candidates.is_blacklist":
			valueField = d.isBlacklistI18n(record.IsBlacklist)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d *candidateDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidates.name"
	case "Email":
		return "model.candidates.email"
	case "Phone":
		return "model.candidates.phone"
	case "Dob":
		return "model.candidates.dob"
	case "IsBlacklist":
		return "model.candidates.is_blacklist"
	}
	return ""
}

func (d *candidateDtoImpl) isBlacklistI18n(input bool) string {
	switch input {
	case true:
		return "model.candidates.is_blacklist_enum.yes"
	default:
		return "model.candidates.is_blacklist_enum.no"
	}
}
