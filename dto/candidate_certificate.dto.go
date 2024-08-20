package dto

import (
	"reflect"
	"trec/ent"
	"trec/models"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CandidateCertificateDto interface {
	ProcessAuditTrail(oldRecord []*ent.CandidateCertificate, newRecord []*ent.CandidateCertificate, recordAudit *models.AuditTrailData)
}

type candidateCertificateDtoImpl struct {
}

func NewCandidateCertificateDto() CandidateCertificateDto {
	return &candidateCertificateDtoImpl{}
}

// ProcessAuditTrails process audit trails for candidate certificates.
//
// oldRecords contains the old records of candidate certificates.
// newRecords contains the new records of candidate certificates.
// newCdCertificate := record not in oldRecords but in newRecords
// delCdCertificate := record not in newRecords but in oldRecords
// updCdCertificate := record in both oldRecords and newRecords
func (d candidateCertificateDtoImpl) ProcessAuditTrail(oldRecords []*ent.CandidateCertificate, newRecords []*ent.CandidateCertificate, recordAudit *models.AuditTrailData) {
	cdCertificateAudit := models.AuditTrailData{
		Module:    CandidateCertificateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	if len(oldRecords) == 0 && len(newRecords) == 0 {
		return
	}
	oldRecordCps := lo.Map(oldRecords, func(entity *ent.CandidateCertificate, _ int) interface{} {
		return entity
	})
	newRecordCps := lo.Map(newRecords, func(entity *ent.CandidateCertificate, _ int) interface{} {
		return entity
	})
	credIds, updIds, delIds := FindCUDArray(oldRecordCps, newRecordCps)
	d.recordAuditCreated(credIds, newRecords, &cdCertificateAudit)
	d.recordAuditUpdated(updIds, oldRecords, newRecords, &cdCertificateAudit)
	d.recordAuditDeleted(delIds, oldRecords, &cdCertificateAudit)
	recordAudit.SubModule = append(recordAudit.SubModule, cdCertificateAudit)
}

func (d candidateCertificateDtoImpl) recordAuditCreated(creIds []uuid.UUID, newRecords []*ent.CandidateCertificate, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(newRecords, func(entity *ent.CandidateCertificate) bool {
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

func (d candidateCertificateDtoImpl) recordAuditDeleted(creIds []uuid.UUID, oldRecords []*ent.CandidateCertificate, recordAudit *models.AuditTrailData) {
	for _, v := range creIds {
		entity := []interface{}{}
		record, _ := lo.Find(oldRecords, func(entity *ent.CandidateCertificate) bool {
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

func (d candidateCertificateDtoImpl) recordAuditUpdated(updIds []uuid.UUID, oldRecords []*ent.CandidateCertificate, newRecords []*ent.CandidateCertificate, recordAudit *models.AuditTrailData) {
	for _, v := range updIds {
		entity := []interface{}{}
		oldRecord, _ := lo.Find(oldRecords, func(entity *ent.CandidateCertificate) bool {
			return entity.ID == v
		})
		newRecord, _ := lo.Find(newRecords, func(entity *ent.CandidateCertificate) bool {
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
func (d candidateCertificateDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidate_certificates.name"
	case "Score":
		return "model.candidate_certificates.score"
	case "AchievedDate":
		return "model.candidate_certificates.achieved_date"
	}
	return ""
}

func (d candidateCertificateDtoImpl) attachmentAuditTrail(record *ent.CandidateCertificate, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_certificates.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateCertificateDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.CandidateCertificate, newRecord *ent.CandidateCertificate, atInterface []interface{}) []interface{} {
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
			Field: "model.candidate_certificates.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}
