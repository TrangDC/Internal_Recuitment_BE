package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"
)

type RoleDto interface {
	AuditTrailCreate(record *ent.Role) (string, error)
	AuditTrailDelete(record *ent.Role) (string, error)
	AuditTrailUpdate(oldRecord *ent.Role, newRecord *ent.Role) (string, error)
}

type roleDtoImpl struct {
}

func NewRoleDto() RoleDto {
	return &roleDtoImpl{}
}

func (d roleDtoImpl) AuditTrailCreate(record *ent.Role) (string, error) {
	result := models.AuditTrailData{
		Module: RoleI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d roleDtoImpl) AuditTrailDelete(record *ent.Role) (string, error) {
	result := models.AuditTrailData{
		Module: RoleI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d roleDtoImpl) AuditTrailUpdate(oldRecord *ent.Role, newRecord *ent.Role) (string, error) {
	result := models.AuditTrailData{
		Module: RoleI18n,
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

func (d roleDtoImpl) recordAudit(record *ent.Role) []interface{} {
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
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d roleDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.roles.name"
	case "Description":
		return "model.roles.description"
	}
	return ""
}
