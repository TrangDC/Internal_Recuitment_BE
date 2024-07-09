package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/samber/lo"
)

type TeamDto interface {
	AuditTrailCreate(record *ent.Team) (string, error)
	AuditTrailDelete(record *ent.Team) (string, error)
	AuditTrailUpdate(oldRecord *ent.Team, newRecord *ent.Team) (string, error)
}

type teamDtoImpl struct {
}

func NewTeamDto() TeamDto {
	return &teamDtoImpl{}
}

func (d teamDtoImpl) AuditTrailCreate(record *ent.Team) (string, error) {
	result := models.AuditTrailData{
		Module: TeamI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d teamDtoImpl) AuditTrailDelete(record *ent.Team) (string, error) {
	result := models.AuditTrailData{
		Module: TeamI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d teamDtoImpl) AuditTrailUpdate(oldRecord *ent.Team, newRecord *ent.Team) (string, error) {
	result := models.AuditTrailData{
		Module: TeamI18n,
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
	entity = d.teamMemberAuditTrailUpdate(oldRecord, newRecord, entity)
	result.Update = append(result.Update, entity...)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d teamDtoImpl) recordAudit(record *ent.Team) []interface{} {
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
	entity = d.teamMemberAuditTrail(record, entity)
	return entity
}

func (d teamDtoImpl) teamMemberAuditTrail(record *ent.Team, atInterface []interface{}) []interface{} {
	if len(record.Edges.UserEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.UserEdges, func(document *ent.User, index int) string {
		return document.Name
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.teams.members",
		Value: attachmentNames,
	})
	return atInterface
}

func (d teamDtoImpl) teamMemberAuditTrailUpdate(oldRecord *ent.Team, newRecord *ent.Team, atInterface []interface{}) []interface{} {
	if len(oldRecord.Edges.UserEdges) == 0 && len(newRecord.Edges.UserEdges) == 0 {
		return atInterface
	}
	oldUserNames := lo.Map(oldRecord.Edges.UserEdges, func(document *ent.User, index int) string {
		return document.Name
	})
	newUserNames := lo.Map(newRecord.Edges.UserEdges, func(document *ent.User, index int) string {
		return document.Name
	})
	if !CompareArray(oldUserNames, newUserNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.teams.members",
			Value: models.ValueChange{
				OldValue: oldUserNames,
				NewValue: newUserNames,
			},
		})
	}
	return atInterface
}

func (d teamDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.teams.name"
	}
	return ""
}
