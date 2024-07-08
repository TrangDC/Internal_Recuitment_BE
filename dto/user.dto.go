package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/user"
	"trec/models"
)

type UserDto interface {
	AuditTrailCreate(record *ent.User) (string, error)
	AuditTrailDelete(record *ent.User) (string, error)
	AuditTrailUpdate(oldRecord *ent.User, newRecord *ent.User) (string, error)
	AuditTrailUpdateTeam(oldRecord *ent.User, teamName string) string
}

type userDtoImpl struct {
}

func NewUserDto() UserDto {
	return &userDtoImpl{}
}

func (d userDtoImpl) AuditTrailCreate(record *ent.User) (string, error) {
	result := models.AuditTrailData{
		Module: UserI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d userDtoImpl) AuditTrailDelete(record *ent.User) (string, error) {
	result := models.AuditTrailData{
		Module: UserI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d userDtoImpl) AuditTrailUpdate(oldRecord *ent.User, newRecord *ent.User) (string, error) {
	result := models.AuditTrailData{
		Module: UserI18n,
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
			case "model.users.status":
				oldValueField = d.statusI18n(oldRecord.Status)
				newValueField = d.statusI18n(newRecord.Status)
			case "model.users.team":
				if oldRecord.Edges.MemberOfTeamEdges != nil {
					oldValueField = oldRecord.Edges.MemberOfTeamEdges.Name
				} else {
					oldValueField = ""
				}
				if newRecord.Edges.MemberOfTeamEdges != nil {
					newValueField = newRecord.Edges.MemberOfTeamEdges.Name
				} else {
					newValueField = ""
				}
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

func (d userDtoImpl) AuditTrailUpdateTeam(oldRecord *ent.User, teamName string) string {
	result := models.AuditTrailData{
		Module: UserI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	entity := []interface{}{}
	oldTeamName := ""
	if oldRecord.Edges.MemberOfTeamEdges != nil {
		oldTeamName = oldRecord.Edges.MemberOfTeamEdges.Name
	}
	entity = append(entity, models.AuditTrailUpdate{
		Field: "model.users.team",
		Value: models.ValueChange{
			OldValue: oldTeamName,
			NewValue: teamName,
		},
	})
	result.Update = append(result.Update, entity...)
	jsonObj, _ := json.Marshal(result)
	return string(jsonObj)
}

func (d userDtoImpl) recordAudit(record *ent.User) []interface{} {
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
		case "model.users.status":
			valueField = d.statusI18n(record.Status)
		case "model.users.team":
			valueField = record.Edges.MemberOfTeamEdges.Name
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d userDtoImpl) statusI18n(input user.Status) string {
	switch input {
	case user.StatusActive:
		return "model.users.status_enum.active"
	default:
		return "model.users.status_enum.inactive"
	}
}

func (d userDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.users.name"
	case "Description":
		return "model.users.work_email"
	case "Status":
		return "model.users.status"
	case "TeamID":
		return "model.users.team"
	}
	return ""
}
