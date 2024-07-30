package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"
)

type JobPositionDto interface {
	AuditTrailCreate(record *ent.JobPosition) (string, error)
	AuditTrailDelete(record *ent.JobPosition) (string, error)
	AuditTrailUpdate(oldRecord *ent.JobPosition, newRecord *ent.JobPosition) (string, error)
}

type jobPositionDtoImpl struct {
}

func NewJobPositionDto() JobPositionDto {
	return &jobPositionDtoImpl{}
}

func (d jobPositionDtoImpl) AuditTrailCreate(record *ent.JobPosition) (string, error) {
	result := models.AuditTrailData{
		Module: JobPositionI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d jobPositionDtoImpl) AuditTrailDelete(record *ent.JobPosition) (string, error) {
	result := models.AuditTrailData{
		Module: JobPositionI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d jobPositionDtoImpl) AuditTrailUpdate(oldRecord *ent.JobPosition, newRecord *ent.JobPosition) (string, error) {
	result := models.AuditTrailData{
		Module: JobPositionI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	var entity []interface{}
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

func (d jobPositionDtoImpl) recordAudit(record *ent.JobPosition) []interface{} {
	var entity []interface{}
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
	return entity
}

func (d jobPositionDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.job_positions.name"
	case "Description":
		return "model.job_positions.description"
	}
	return ""
}
