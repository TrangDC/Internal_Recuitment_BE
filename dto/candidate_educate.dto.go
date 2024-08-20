package dto

import (
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateEducateDto interface {
	ProcessAuditTrail(oldRecord []*ent.CandidateEducate, newRecord []*ent.CandidateEducate, recordAudit *models.AuditTrailData)
}

type candidateEducateDtoImpl struct {
}

func NewCandidateEducateDto() CandidateEducateDto {
	return &candidateEducateDtoImpl{}
}

// ProcessAuditTrails process audit trails for candidate educates.
//
// oldRecords contains the old records of candidate educates.
// newRecords contains the new records of candidate educates.
// newCdEducate := record not in oldRecords but in newRecords
// delCdEducate := record not in newRecords but in oldRecords
// updCdEducate := record in both oldRecords and newRecords
func (d candidateEducateDtoImpl) ProcessAuditTrail(oldRecords []*ent.CandidateEducate, newRecords []*ent.CandidateEducate, recordAudit *models.AuditTrailData) {
	cdEducateAudit := models.AuditTrailData{
		Module:    CandidateEducateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return
	}
	oldRecordCps := lo.Map(oldRecords, func(entity *ent.CandidateEducate, _ int) interface{} {
		return entity
	})
	newRecordCps := lo.Map(newRecords, func(entity *ent.CandidateEducate, _ int) interface{} {
		return entity
	})
	credIds, updIds, delIds := FindCUDArray(oldRecordCps, newRecordCps)
	d.recordAuditCreated(credIds, newRecords, &cdEducateAudit)
	d.recordAuditUpdated(updIds, oldRecords, newRecords, &cdEducateAudit)
	d.recordAuditDeleted(delIds, oldRecords, &cdEducateAudit)
	recordAudit.SubModule = append(recordAudit.SubModule, cdEducateAudit)
}

func (d candidateEducateDtoImpl) recordAuditCreated(creIds []uuid.UUID, newRecords []*ent.CandidateEducate, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(newRecords, func(entity *ent.CandidateEducate) bool {
			return entity.ID == v
		})
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
		entity = d.attachmentAuditTrail(record, entity)
		recordAudit.Create = append(recordAudit.Create, entity)
	}
}

func (d candidateEducateDtoImpl) recordAuditDeleted(creIds []uuid.UUID, oldRecords []*ent.CandidateEducate, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(oldRecords, func(entity *ent.CandidateEducate) bool {
			return entity.ID == v
		})
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
		entity = d.attachmentAuditTrail(record, entity)
		recordAudit.Delete = append(recordAudit.Delete, entity)
	}
}

func (d candidateEducateDtoImpl) recordAuditUpdated(updIds []uuid.UUID, oldRecords []*ent.CandidateEducate, newRecords []*ent.CandidateEducate, recordAudit *models.AuditTrailData) {
	for _, v := range updIds {
		entity := []interface{}{}
		oldRecord, _ := lo.Find(oldRecords, func(entity *ent.CandidateEducate) bool {
			return entity.ID == v
		})
		newRecord, _ := lo.Find(newRecords, func(entity *ent.CandidateEducate) bool {
			return entity.ID == v
		})
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
		recordAudit.Update = append(recordAudit.Update, entity)
	}
}

// common
func (d candidateEducateDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "SchoolName":
		return "model.candidate_educates.school_name"
	case "Major":
		return "model.candidate_educates.major"
	case "GPA":
		return "model.candidate_educates.gpa"
	case "Location":
		return "model.candidate_educates.location"
	case "Description":
		return "model.candidate_educates.description"
	case "StartDate":
		return "model.candidate_educates.start_date"
	case "EndDate":
		return "model.candidate_educates.end_date"
	case "IsCurrent":
		return "model.candidate_educates.is_current"
	}
	return ""
}

func (d candidateEducateDtoImpl) attachmentAuditTrail(record *ent.CandidateEducate, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_educates.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateEducateDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.CandidateEducate, newRecord *ent.CandidateEducate, atInterface []interface{}) []interface{} {
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
			Field: "model.candidate_educates.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}
