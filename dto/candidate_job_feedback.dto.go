package dto

import (
	"encoding/json"
	"fmt"
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/samber/lo"
)

type CandidateJobFeedbackDto interface {
	AuditTrailCreate(record *ent.CandidateJobFeedback) (string, error)
	AuditTrailDelete(record *ent.CandidateJobFeedback) (string, error)
	AuditTrailUpdate(oldRecord *ent.CandidateJobFeedback, newRecord *ent.CandidateJobFeedback) (string, error)
}

type candidateJobFeedbackDtoImpl struct {
}

func NewCandidateJobFeedbackDto() CandidateJobFeedbackDto {
	return &candidateJobFeedbackDtoImpl{}
}

func (d *candidateJobFeedbackDtoImpl) AuditTrailCreate(record *ent.CandidateJobFeedback) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateJobFBAt := models.AuditTrailData{
		Module: CandidateJobFeedbackI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	fatherModule := models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", record.Edges.CandidateJobEdge.Status),
		Value:  record.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	}
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateJobFBAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobFeedbackDtoImpl) AuditTrailDelete(record *ent.CandidateJobFeedback) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateJobFBAt := models.AuditTrailData{
		Module: CandidateJobFeedbackI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	fatherModule := models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", record.Edges.CandidateJobEdge.Status),
		Value:  record.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	}
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateJobFBAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobFeedbackDtoImpl) AuditTrailUpdate(oldRecord *ent.CandidateJobFeedback, newRecord *ent.CandidateJobFeedback) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateInterviewAt := models.AuditTrailData{
		Module: CandidateJobFeedbackI18n,
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
	entity = d.attachmentAuditTrailUpdate(oldRecord, newRecord, entity)
	candidateInterviewAt.Update = append(candidateInterviewAt.Update, entity...)
	fatherModule := models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", oldRecord.Edges.CandidateJobEdge.Status),
		Value:  oldRecord.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	}
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateInterviewAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobFeedbackDtoImpl) recordAudit(record *ent.CandidateJobFeedback) []interface{} {
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
		case "model.candidate_job_feedbacks.created_by":
			valueField = record.Edges.CreatedByEdge.Name
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.attachmentAuditTrail(record, entity)
	return entity
}

func (d candidateJobFeedbackDtoImpl) attachmentAuditTrail(record *ent.CandidateJobFeedback, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_job_feedbacks.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateJobFeedbackDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.CandidateJobFeedback, newRecord *ent.CandidateJobFeedback, atInterface []interface{}) []interface{} {
	if len(oldRecord.Edges.AttachmentEdges) == 0 && len(newRecord.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	oldAttachmentNames := lo.Map(oldRecord.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	newAttachmentNames := lo.Map(newRecord.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	if !CompareArray(oldAttachmentNames, newAttachmentNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.candidate_job_feedbacks.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d candidateJobFeedbackDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "CreatedBy":
		return "model.candidate_job_feedbacks.created_by"
	case "Feedback":
		return "model.candidate_job_feedbacks.feedback"
	}
	return ""
}
