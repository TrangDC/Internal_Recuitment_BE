package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"
)

type SkillDto interface {
	AuditTrailCreate(record *ent.Skill) (string, error)
	AuditTrailDelete(record *ent.Skill) (string, error)
	AuditTrailUpdate(oldRecord *ent.Skill, newRecord *ent.Skill) (string, error)
}

type skillDtoImpl struct {
}

func NewSkillDto() SkillDto {
	return &skillDtoImpl{}
}

func (d skillDtoImpl) AuditTrailCreate(record *ent.Skill) (string, error) {
	result := models.AuditTrailData{
		Module: SkillI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d skillDtoImpl) AuditTrailDelete(record *ent.Skill) (string, error) {
	result := models.AuditTrailData{
		Module: SkillI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d skillDtoImpl) AuditTrailUpdate(oldRecord *ent.Skill, newRecord *ent.Skill) (string, error) {
	result := models.AuditTrailData{
		Module: SkillI18n,
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
			case "model.skills.skill_type":
				if oldRecord.Edges.SkillTypeEdge != nil {
					oldValueField = oldRecord.Edges.SkillTypeEdge.Name
				} else {
					oldValueField = ""
				}
				if newRecord.Edges.SkillTypeEdge != nil {
					newValueField = newRecord.Edges.SkillTypeEdge.Name
				} else {
					newValueField = ""
				}
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

func (d skillDtoImpl) recordAudit(record *ent.Skill) []interface{} {
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
		case "model.skills.skill_type":
			if record.Edges.SkillTypeEdge != nil {
				valueField = record.Edges.SkillTypeEdge.Name
			} else {
				valueField = ""
			}
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d skillDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.skills.name"
	case "Description":
		return "model.skills.description"
	case "SkillTypeID":
		return "model.skills.skill_type"
	}
	return ""
}
