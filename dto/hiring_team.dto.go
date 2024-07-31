package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/samber/lo"
)

type HiringTeamDto interface {
	AuditTrailCreate(record *ent.HiringTeam) (string, error)
	AuditTrailDelete(record *ent.HiringTeam) (string, error)
	AuditTrailUpdate(oldRecord *ent.HiringTeam, newRecord *ent.HiringTeam) (string, error)
}

type hiringTeamDtoImpl struct {
}

func NewHiringTeamDto() HiringTeamDto {
	return &hiringTeamDtoImpl{}
}

func (d hiringTeamDtoImpl) AuditTrailCreate(record *ent.HiringTeam) (string, error) {
	result := models.AuditTrailData{
		Module: HiringTeamI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringTeamDtoImpl) AuditTrailDelete(record *ent.HiringTeam) (string, error) {
	result := models.AuditTrailData{
		Module: HiringTeamI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringTeamDtoImpl) AuditTrailUpdate(oldRecord *ent.HiringTeam, newRecord *ent.HiringTeam) (string, error) {
	result := models.AuditTrailData{
		Module: HiringTeamI18n,
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
	entity = d.teamApproverAuditTrailUpdate(oldRecord.Edges.HiringTeamApprovers, newRecord.Edges.HiringTeamApprovers, entity)
	result.Update = append(result.Update, entity...)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringTeamDtoImpl) recordAudit(record *ent.HiringTeam) []interface{} {
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
	entity = d.teamApproverAuditTrail(record.Edges.HiringTeamApprovers, entity)
	return entity
}

func (d hiringTeamDtoImpl) teamMemberAuditTrail(record *ent.HiringTeam, atInterface []interface{}) []interface{} {
	if len(record.Edges.UserEdges) == 0 {
		return atInterface
	}
	teamMembers := lo.Map(record.Edges.UserEdges, func(document *ent.User, index int) string {
		return document.Name
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.hiring_teams.members",
		Value: teamMembers,
	})
	return atInterface
}

func (d hiringTeamDtoImpl) teamMemberAuditTrailUpdate(oldRecord *ent.HiringTeam, newRecord *ent.HiringTeam, atInterface []interface{}) []interface{} {
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
			Field: "model.hiring_teams.members",
			Value: models.ValueChange{
				OldValue: oldUserNames,
				NewValue: newUserNames,
			},
		})
	}
	return atInterface
}

func (d hiringTeamDtoImpl) teamApproverAuditTrail(teamApprovers []*ent.HiringTeamApprover, atInterface []interface{}) []interface{} {
	teamApproversAuditTrail := lo.Map(teamApprovers, func(approver *ent.HiringTeamApprover, _ int) models.HiringTeamApproverAuditTrail {
		return models.HiringTeamApproverAuditTrail{
			Name:    approver.Edges.User.Name,
			OrderID: approver.OrderID,
		}
	})
	auditTrailJSON, _ := json.Marshal(teamApproversAuditTrail)
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.hiring_teams.approvers",
		Value: string(auditTrailJSON),
	})
	return atInterface
}

func (d hiringTeamDtoImpl) teamApproverAuditTrailUpdate(oldApprovers []*ent.HiringTeamApprover, newApprovers []*ent.HiringTeamApprover, atInterface []interface{}) []interface{} {
	oldApproversAuditTrail := lo.Map(oldApprovers, func(approver *ent.HiringTeamApprover, _ int) models.HiringTeamApproverAuditTrail {
		return models.HiringTeamApproverAuditTrail{
			Name:    approver.Edges.User.Name,
			OrderID: approver.OrderID,
		}
	})
	newApproversAuditTrail := lo.Map(newApprovers, func(approver *ent.HiringTeamApprover, _ int) models.HiringTeamApproverAuditTrail {
		return models.HiringTeamApproverAuditTrail{
			Name:    approver.Edges.User.Name,
			OrderID: approver.OrderID,
		}
	})
	if len(oldApproversAuditTrail) == len(newApproversAuditTrail) && reflect.DeepEqual(oldApproversAuditTrail, newApproversAuditTrail) {
		return atInterface
	}
	oldApproversJSON, _ := json.Marshal(oldApproversAuditTrail)
	newApproversJSON, _ := json.Marshal(newApproversAuditTrail)
	atInterface = append(atInterface, models.AuditTrailUpdate{
		Field: "model.hiring_teams.approvers",
		Value: models.ValueChange{
			OldValue: string(oldApproversJSON),
			NewValue: string(newApproversJSON),
		},
	})
	return atInterface
}

func (d hiringTeamDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.hiring_teams.name"
	case "Description":
		return "model.hiring_teams.description"
	}
	return ""
}
