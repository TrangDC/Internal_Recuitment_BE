package dto

import (
	"encoding/json"
	"fmt"
	"reflect"
	"trec/ent"
	"trec/models"
)

type RecTeamDto interface {
	AuditTrailCreate(record *ent.RecTeam) (string, error)
	AuditTrailDelete(record *ent.RecTeam) (string, error)
	AuditTrailUpdate(oldRecord *ent.RecTeam, newRecord *ent.RecTeam) (string, error)
}

type recTeamDtoImpl struct {
}

func NewRecTeamDto() RecTeamDto {
	return &recTeamDtoImpl{}
}

func (d recTeamDtoImpl) AuditTrailCreate(record *ent.RecTeam) (string, error) {
	result := models.AuditTrailData{
		Module: RecTeamI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d recTeamDtoImpl) AuditTrailDelete(record *ent.RecTeam) (string, error) {
	result := models.AuditTrailData{
		Module: RecTeamI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d recTeamDtoImpl) AuditTrailUpdate(oldRecord *ent.RecTeam, newRecord *ent.RecTeam) (string, error) {
	result := models.AuditTrailData{
		Module: RecTeamI18n,
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
			case "model.rec_teams.rec_team_leader":
				oldValueField = oldRecord.Edges.RecLeaderEdge.Name
				newValueField = newRecord.Edges.RecLeaderEdge.Name
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

func (d recTeamDtoImpl) recordAudit(record *ent.RecTeam) []interface{} {
	fmt.Println("====>", record.Edges.RecLeaderEdge)
	fmt.Println("====>", record.Edges.RecMemberEdges)
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
		case "model.rec_teams.rec_team_leader":
			valueField = record.Edges.RecLeaderEdge.Name
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
}

func (d recTeamDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.rec_teams.name"
	case "Description":
		return "model.rec_teams.description"
	case "LeaderID":
		return "model.rec_teams.rec_team_leader"
	}
	return ""
}
