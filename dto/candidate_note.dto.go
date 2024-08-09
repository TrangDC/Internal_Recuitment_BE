package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/samber/lo"
)

type CandidateNoteDto interface {
	AuditTrailCreate(record *ent.CandidateNote) (string, error)
	AuditTrailUpdate(oldRecord, newRecord *ent.CandidateNote) (string, error)
	AuditTrailDelete(record *ent.CandidateNote) (string, error)
}

type candidateNoteDtoImpl struct{}

func NewCandidateNoteDto() CandidateNoteDto {
	return &candidateNoteDtoImpl{}
}

func (d *candidateNoteDtoImpl) AuditTrailCreate(record *ent.CandidateNote) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateNoteI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	resultJSON, err := json.Marshal(result)
	return string(resultJSON), err
}

func (d *candidateNoteDtoImpl) AuditTrailUpdate(oldRecord, newRecord *ent.CandidateNote) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateNoteI18n,
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
			if fieldName == "" {
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
	entity = d.attachmentAuditTrailUpdate(oldRecord.Edges.AttachmentEdges, newRecord.Edges.AttachmentEdges, entity)
	result.Update = append(result.Update, entity...)
	resultJSON, err := json.Marshal(result)
	return string(resultJSON), err
}

func (d *candidateNoteDtoImpl) AuditTrailDelete(record *ent.CandidateNote) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateNoteI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	resultJSON, err := json.Marshal(result)
	return string(resultJSON), err
}

func (d *candidateNoteDtoImpl) recordAudit(record *ent.CandidateNote) []interface{} {
	entity := make([]interface{}, 0)
	value := reflect.ValueOf(interface{}(record)).Elem()
	recordType := reflect.TypeOf(record).Elem()
	for i := 1; i < value.NumField(); i++ {
		field := recordType.Field(i)
		valueField := value.Field(i).Interface()
		fieldName := d.formatFieldI18n(field.Name)
		if fieldName == "" {
			continue
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.attachmentAuditTrail(record.Edges.AttachmentEdges, entity)
	return entity
}

func (d *candidateNoteDtoImpl) attachmentAuditTrail(records []*ent.Attachment, atInterface []interface{}) []interface{} {
	if len(records) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(records, func(document *ent.Attachment, _ int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_notes.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d *candidateNoteDtoImpl) attachmentAuditTrailUpdate(oldRecords, newRecords []*ent.Attachment, atInterface []interface{}) []interface{} {
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return atInterface
	}
	oldAttachmentNames := lo.Map(oldRecords, func(document *ent.Attachment, _ int) string {
		return document.DocumentName
	})
	newAttachmentNames := lo.Map(newRecords, func(document *ent.Attachment, _ int) string {
		return document.DocumentName
	})
	if !CompareArray(oldAttachmentNames, newAttachmentNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.candidate_notes.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d *candidateNoteDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidate_notes.name"
	case "Description":
		return "model.candidate_notes.description"
	default:
		return ""
	}
}
