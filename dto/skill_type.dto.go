package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"
)

type SkillTypeDto interface {
	AuditTrailCreate(record *ent.SkillType) (string, error)
	AuditTrailDelete(record *ent.SkillType) (string, error)
	AuditTrailUpdate(oldRecord *ent.SkillType, newRecord *ent.SkillType) (string, error)
}

type skillTypeDtoImpl struct {
}

func NewSkillTypeDto() SkillTypeDto {
	return &skillTypeDtoImpl{}
}

func (d skillTypeDtoImpl) AuditTrailCreate(record *ent.SkillType) (string, error) {
	result := models.AuditTrailData{
		Module: SkillTypeI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d skillTypeDtoImpl) AuditTrailDelete(record *ent.SkillType) (string, error) {
	result := models.AuditTrailData{
		Module: SkillTypeI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d skillTypeDtoImpl) AuditTrailUpdate(oldRecord *ent.SkillType, newRecord *ent.SkillType) (string, error) {
	result := models.AuditTrailData{
		Module: SkillTypeI18n,
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

func (d skillTypeDtoImpl) recordAudit(record *ent.SkillType) []interface{} {
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

func (d skillTypeDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.skill_types.name"
	case "Description":
		return "model.skill_types.description"
	}
	return ""
}
