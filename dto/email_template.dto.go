package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/emailtemplate"
	"trec/models"

	"github.com/samber/lo"
)

type EmailTemplateDto interface {
	AuditTrailCreate(record *ent.EmailTemplate) (string, error)
	AuditTrailDelete(record *ent.EmailTemplate) (string, error)
	AuditTrailUpdate(oldRecord *ent.EmailTemplate, newRecord *ent.EmailTemplate) (string, error)
}

type emailTemplateDtoImpl struct {
}

func NewEmailTemplateDto() EmailTemplateDto {
	return &emailTemplateDtoImpl{}
}

func (d emailTemplateDtoImpl) AuditTrailCreate(record *ent.EmailTemplate) (string, error) {
	result := models.AuditTrailData{
		Module: EmailTemplateI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d emailTemplateDtoImpl) AuditTrailDelete(record *ent.EmailTemplate) (string, error) {
	result := models.AuditTrailData{
		Module: EmailTemplateI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d emailTemplateDtoImpl) AuditTrailUpdate(oldRecord *ent.EmailTemplate, newRecord *ent.EmailTemplate) (string, error) {
	result := models.AuditTrailData{
		Module: EmailTemplateI18n,
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
			case "model.email_templates.event":
				oldValueField = d.eventI18n(oldRecord.Event)
				newValueField = d.eventI18n(newRecord.Event)
			case "model.email_templates.status":
				oldValueField = d.statusI18n(oldRecord.Status)
				newValueField = d.statusI18n(newRecord.Status)
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
	entity = d.emailTemplateSentToAtUpdate(oldRecord, newRecord, entity)
	result.Update = append(result.Update, entity...)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d emailTemplateDtoImpl) recordAudit(record *ent.EmailTemplate) []interface{} {
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
		case "model.email_templates.event":
			valueField = d.eventI18n(record.Event)
		case "model.email_templates.status":
			valueField = d.statusI18n(record.Status)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.emailTemplateSentToAt(record, entity)
	return entity
}

func (d emailTemplateDtoImpl) emailTemplateSentToAt(record *ent.EmailTemplate, atInterface []interface{}) []interface{} {
	sendTos := lo.Map(record.SendTo, func(value string, index int) string {
		return d.sentToI18n(value)
	})
	roleNames := lo.Map(record.Edges.RoleEdges, func(entity *ent.Role, index int) string {
		return entity.Name
	})
	sendTos = append(sendTos, roleNames...)
	// remove empty value
	sendTos = lo.Filter(sendTos, func(value string, index int) bool {
		return value != ""
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.email_templates.send_to",
		Value: sendTos,
	})
	return atInterface
}

func (d emailTemplateDtoImpl) emailTemplateSentToAtUpdate(oldRecord *ent.EmailTemplate, newRecord *ent.EmailTemplate, atInterface []interface{}) []interface{} {
	oldSendTos := lo.Map(oldRecord.SendTo, func(value string, index int) string {
		return d.sentToI18n(value)
	})
	newSendTos := lo.Map(newRecord.SendTo, func(value string, index int) string {
		return d.sentToI18n(value)
	})
	oldRoleNames := lo.Map(oldRecord.Edges.RoleEdges, func(entity *ent.Role, index int) string {
		return entity.Name
	})
	newRoleNames := lo.Map(newRecord.Edges.RoleEdges, func(entity *ent.Role, index int) string {
		return entity.Name
	})
	oldSendTos = append(oldSendTos, oldRoleNames...)
	oldSendTos = lo.Filter(oldSendTos, func(value string, index int) bool {
		return value != ""
	})
	newSendTos = append(newSendTos, newRoleNames...)
	newSendTos = lo.Filter(newSendTos, func(value string, index int) bool {
		return value != ""
	})
	if !CompareArray(oldSendTos, newSendTos) && !CompareArray(oldRoleNames, newRoleNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.email_templates.send_to",
			Value: models.ValueChange{
				OldValue: oldSendTos,
				NewValue: newSendTos,
			},
		})
	}
	return atInterface
}

func (d emailTemplateDtoImpl) eventI18n(input emailtemplate.Event) string {
	switch input {
	case emailtemplate.EventCandidateAppliedToKiv:
		return "model.email_templates.event_enum.candidate_applied_to_kiv"
	case emailtemplate.EventCandidateInterviewingToKiv:
		return "model.email_templates.event_enum.candidate_interviewing_to_kiv"
	case emailtemplate.EventCandidateInterviewingToOffering:
		return "model.email_templates.event_enum.candidate_interviewing_to_offering"
	case emailtemplate.EventCreatedCandidate:
		return "model.email_templates.event_enum.created_candidate"
	case emailtemplate.EventUpdatingInterview:
		return "model.email_templates.event_enum.updating_interview"
	}
	return ""
}

func (d emailTemplateDtoImpl) sentToI18n(input string) string {
	switch input {
	case ent.EmailTemplateSendToCandidate.String():
		return "model.email_templates.send_to_enum.candidate"
	case ent.EmailTemplateSendToInterviewer.String():
		return "model.email_templates.send_to_enum.interviewer"
	case ent.EmailTemplateSendToJobRequest.String():
		return "model.email_templates.send_to_enum.job_request"
	case ent.EmailTemplateSendToTeamManager.String():
		return "model.email_templates.send_to_enum.team_manager"
	case ent.EmailTemplateSendToTeamMember.String():
		return "model.email_templates.send_to_enum.team_member"
	}
	return ""
}
func (d emailTemplateDtoImpl) statusI18n(input emailtemplate.Status) string {
	switch input {
	case emailtemplate.StatusActive:
		return "model.email_templates.status_enum.active"
	case emailtemplate.StatusInactive:
		return "model.email_templates.status_enum.inactive"
	}
	return ""
}

func (d emailTemplateDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Event":
		return "model.email_templates.event"
	case "Cc":
		return "model.email_templates.cc"
	case "Bcc":
		return "model.email_templates.bcc"
	case "Subject":
		return "model.email_templates.subject"
	case "Content":
		return "model.email_templates.content"
	case "Signature":
		return "model.email_templates.signature"
	case "Status":
		return "model.email_templates.status"
	}
	return ""
}
