package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/user"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UserDto interface {
	AuditTrailCreate(record *ent.User) (string, error)
	AuditTrailDelete(record *ent.User) (string, error)
	AuditTrailUpdate(oldRecord *ent.User, newRecord *ent.User) (string, error)
	AuditTrailUpdateHiringTeam(oldRecord *ent.User, hiringTeamName string) string
	AuditTrailUpdateRecTeam(oldRecord *ent.User, recTeamName string) string
	NewUserEntityPermissionInput(rolePermissions []*ent.EntityPermission) []*ent.NewEntityPermissionInput
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
			case "model.users.hiring_team":
				oldValueField = ""
				if oldRecord.Edges.MemberOfHiringTeamEdges != nil {
					oldValueField = oldRecord.Edges.MemberOfHiringTeamEdges.Name
				}
				newValueField = ""
				if newRecord.Edges.MemberOfHiringTeamEdges != nil {
					newValueField = newRecord.Edges.MemberOfHiringTeamEdges.Name
				}
			case "model.users.rec_team":
				oldValueField = ""
				if oldRecord.Edges.RecTeams != nil {
					oldValueField = oldRecord.Edges.RecTeams.Name
				}
				newValueField = ""
				if newRecord.Edges.RecTeams != nil {
					newValueField = newRecord.Edges.RecTeams.Name
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

func (d userDtoImpl) AuditTrailUpdateHiringTeam(oldRecord *ent.User, hiringTeamName string) string {
	result := models.AuditTrailData{
		Module: UserI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	entity := []interface{}{}
	oldTeamName := ""
	if oldRecord.Edges.MemberOfHiringTeamEdges != nil {
		oldTeamName = oldRecord.Edges.MemberOfHiringTeamEdges.Name
	}
	entity = append(entity, models.AuditTrailUpdate{
		Field: "model.users.hiring_team",
		Value: models.ValueChange{
			OldValue: oldTeamName,
			NewValue: hiringTeamName,
		},
	})
	result.Update = append(result.Update, entity...)
	jsonObj, _ := json.Marshal(result)
	return string(jsonObj)
}

func (d userDtoImpl) NewUserEntityPermissionInput(rolePermissions []*ent.EntityPermission) []*ent.NewEntityPermissionInput {
	inputByPermissionID := make(map[uuid.UUID]*ent.NewEntityPermissionInput)
	lo.ForEach(rolePermissions, func(rolePermission *ent.EntityPermission, _ int) {
		permissionID := rolePermission.Edges.PermissionEdges.ID
		if _, exist := inputByPermissionID[permissionID]; !exist {
			inputByPermissionID[permissionID] = &ent.NewEntityPermissionInput{
				ForOwner:     rolePermission.ForOwner,
				ForTeam:      rolePermission.ForTeam,
				ForAll:       rolePermission.ForAll,
				PermissionID: permissionID.String(),
			}
			return
		}
		if rolePermission.ForAll {
			inputByPermissionID[permissionID].ForAll = true
			inputByPermissionID[permissionID].ForTeam = false
			inputByPermissionID[permissionID].ForOwner = false
			return
		}
		if rolePermission.ForTeam && inputByPermissionID[permissionID].ForOwner {
			inputByPermissionID[permissionID].ForTeam = true
			inputByPermissionID[permissionID].ForOwner = false
			return
		}
	})
	return lo.Values(inputByPermissionID)
}

func (d userDtoImpl) AuditTrailUpdateRecTeam(oldRecord *ent.User, recTeamName string) string {
	result := models.AuditTrailData{
		Module: UserI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	entity := []interface{}{}
	oldTeamName := ""
	if oldRecord.Edges.RecTeams != nil {
		oldTeamName = oldRecord.Edges.RecTeams.Name
	}
	entity = append(entity, models.AuditTrailUpdate{
		Field: "model.users.rec_team",
		Value: models.ValueChange{
			OldValue: oldTeamName,
			NewValue: recTeamName,
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
		case "model.users.hiring_team":
			valueField = ""
			if record.Edges.MemberOfHiringTeamEdges != nil {
				valueField = record.Edges.MemberOfHiringTeamEdges.Name
			}
		case "model.users.rec_team":
			valueField = ""
			if record.Edges.RecTeams != nil {
				valueField = record.Edges.RecTeams.Name
			}
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
	case "HiringTeamID":
		return "model.users.hiring_team"
	case "RecTeamID":
		return "model.users.rec_team"
	}
	return ""
}
