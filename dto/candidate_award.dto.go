package dto

import (
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateAwardDto interface {
	ProcessAuditTrail(oldRecord []*ent.CandidateAward, newRecord []*ent.CandidateAward, recordAudit *models.AuditTrailData)
}

type candidateAwardDtoImpl struct {
}

func NewCandidateAwardDto() CandidateAwardDto {
	return &candidateAwardDtoImpl{}
}

// ProcessAuditTrails process audit trails for candidate awards.
//
// oldRecords contains the old records of candidate awards.
// newRecords contains the new records of candidate awards.
// newCdAward := record not in oldRecords but in newRecords
// delCdAward := record not in newRecords but in oldRecords
// updCdAward := record in both oldRecords and newRecords
func (d candidateAwardDtoImpl) ProcessAuditTrail(oldRecords []*ent.CandidateAward, newRecords []*ent.CandidateAward, recordAudit *models.AuditTrailData) {
	cdAwardAudit := models.AuditTrailData{
		Module:    CandidateAwardI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return
	}
	oldRecordCps := lo.Map(oldRecords, func(entity *ent.CandidateAward, _ int) interface{} {
		return entity
	})
	newRecordCps := lo.Map(newRecords, func(entity *ent.CandidateAward, _ int) interface{} {
		return entity
	})
	credIds, updIds, delIds := FindCUDArray(oldRecordCps, newRecordCps)
	d.recordAuditCreated(credIds, newRecords, &cdAwardAudit)
	d.recordAuditUpdated(updIds, oldRecords, newRecords, &cdAwardAudit)
	d.recordAuditDeleted(delIds, oldRecords, &cdAwardAudit)
	recordAudit.SubModule = append(recordAudit.SubModule, cdAwardAudit)
}

func (d candidateAwardDtoImpl) recordAuditCreated(creIds []uuid.UUID, newRecords []*ent.CandidateAward, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(newRecords, func(entity *ent.CandidateAward) bool {
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

func (d candidateAwardDtoImpl) recordAuditDeleted(creIds []uuid.UUID, oldRecords []*ent.CandidateAward, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(oldRecords, func(entity *ent.CandidateAward) bool {
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

func (d candidateAwardDtoImpl) recordAuditUpdated(updIds []uuid.UUID, oldRecords []*ent.CandidateAward, newRecords []*ent.CandidateAward, recordAudit *models.AuditTrailData) {
	for _, v := range updIds {
		entity := []interface{}{}
		oldRecord, _ := lo.Find(oldRecords, func(entity *ent.CandidateAward) bool {
			return entity.ID == v
		})
		newRecord, _ := lo.Find(newRecords, func(entity *ent.CandidateAward) bool {
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
func (d candidateAwardDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidate_awards.name"
	case "AchievedDate":
		return "model.candidate_awards.achieved_date"
	}
	return ""
}

func (d candidateAwardDtoImpl) attachmentAuditTrail(record *ent.CandidateAward, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_awards.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateAwardDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.CandidateAward, newRecord *ent.CandidateAward, atInterface []interface{}) []interface{} {
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
			Field: "model.candidate_awards.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}
