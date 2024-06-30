package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/candidatejob"
	"trec/models"

	"github.com/samber/lo"
)

type CandidateJobDto interface {
	AuditTrailCreate(record *ent.CandidateJob) (string, error)
	AuditTrailDelete(record *ent.CandidateJob) (string, error)
	AuditTrailUpdate(oldRecord *ent.CandidateJob, newRecord *ent.CandidateJob) (string, error)

	MappingEdge(records []*ent.CandidateJob, candidates []*ent.Candidate, interviews []*ent.CandidateInterview, hiringJobs []*ent.HiringJob)
}

type candidateJobDtoImpl struct {
}

func NewCandidateJobDto() CandidateJobDto {
	return &candidateJobDtoImpl{}
}

func (d *candidateJobDtoImpl) AuditTrailCreate(record *ent.CandidateJob) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateUpdate := models.AuditTrailData{
		Module: CandidateJobI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	result.SubModule = append(result.SubModule, candidateUpdate)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobDtoImpl) AuditTrailDelete(record *ent.CandidateJob) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateJobUpdate := models.AuditTrailData{
		Module: CandidateJobI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	result.SubModule = append(result.SubModule, candidateJobUpdate)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobDtoImpl) AuditTrailUpdate(oldRecord *ent.CandidateJob, newRecord *ent.CandidateJob) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateJobAt := models.AuditTrailData{
		Module: CandidateJobI18n,
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
			case "model.candidate_job.status":
				oldValueField = d.statusI18n(oldRecord.Status)
				newValueField = d.statusI18n(newRecord.Status)
			case "model.candidate_job.hiring_job":
				if oldRecord.Edges.HiringJobEdge != nil {
					oldValueField = oldRecord.Edges.HiringJobEdge.Name
				} else {
					oldValueField = ""
				}
				if newRecord.Edges.HiringJobEdge != nil {
					newValueField = newRecord.Edges.HiringJobEdge.Name
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
	entity = d.attachmentAuditTrailUpdate(oldRecord, newRecord, entity)
	entity = d.reasonFailAuditTrailUpdate(oldRecord, newRecord, entity)
	candidateJobAt.Update = append(candidateJobAt.Update, entity...)
	result.SubModule = append(result.SubModule, candidateJobAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateJobDtoImpl) recordAudit(record *ent.CandidateJob) []interface{} {
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
		case "model.candidate_jobs.status":
			valueField = d.statusI18n(record.Status)
		case "model.candidate_jobs.hiring_job":
			if record.Edges.HiringJobEdge != nil {
				valueField = record.Edges.HiringJobEdge.Name
			} else {
				valueField = ""
			}
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.attachmentAuditTrail(record, entity)
	entity = d.reasonFailAuditTrail(record, entity)
	return entity
}

func (d candidateJobDtoImpl) attachmentAuditTrail(record *ent.CandidateJob, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_jobs.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateJobDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.CandidateJob, newRecord *ent.CandidateJob, atInterface []interface{}) []interface{} {
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
			Field: "model.candidate_jobs.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d candidateJobDtoImpl) reasonFailAuditTrail(record *ent.CandidateJob, atInterface []interface{}) []interface{} {
	if len(record.FailedReason) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.FailedReason, func(entity string, index int) string {
		return entity
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_jobs.failed_reason",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateJobDtoImpl) reasonFailAuditTrailUpdate(oldRecord *ent.CandidateJob, newRecord *ent.CandidateJob, atInterface []interface{}) []interface{} {
	if len(oldRecord.FailedReason) == 0 && len(newRecord.FailedReason) == 0 {
		return atInterface
	}
	oldAttachmentNames := lo.Map(oldRecord.FailedReason, func(entity string, index int) string {
		return entity
	})
	newAttachmentNames := lo.Map(newRecord.FailedReason, func(entity string, index int) string {
		return entity
	})
	if !CompareArray(oldAttachmentNames, newAttachmentNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.candidate_jobs.failed_reason",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d candidateJobDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "HiringJobID":
		return "model.candidate_jobs.hiring_job"
	case "Status":
		return "model.candidate_jobs.status"
	}
	return ""
}

func (d candidateJobDtoImpl) statusI18n(input candidatejob.Status) string {
	switch input {
	case candidatejob.StatusHired:
		return "model.candidate_jobs.status_enum.hired"
	case candidatejob.StatusKiv:
		return "model.candidate_jobs.status_enum.kiv"
	case candidatejob.StatusOfferLost:
		return "model.candidate_jobs.status_enum.offer_lost"
	case candidatejob.StatusExStaff:
		return "model.candidate_jobs.status_enum.ex_staff"
	case candidatejob.StatusApplied:
		return "model.candidate_jobs.status_enum.applied"
	case candidatejob.StatusInterviewing:
		return "model.candidate_jobs.status_enum.interviewing"
	case candidatejob.StatusOffering:
		return "model.candidate_jobs.status_enum.offering"
	}
	return ""
}

func (d candidateJobDtoImpl) MappingEdge(records []*ent.CandidateJob, candidates []*ent.Candidate, interviews []*ent.CandidateInterview, hiringJobs []*ent.HiringJob) {
	for _, record := range records {
		candidateEdge, _ := lo.Find(candidates, func(candidate *ent.Candidate) bool {
			return candidate.ID == record.CandidateID
		})
		interviewEdges := lo.Filter(interviews, func(interview *ent.CandidateInterview, index int) bool {
			return interview.CandidateJobID == record.ID
		})
		hiringJobEdge, _ := lo.Find(hiringJobs, func(hiringJob *ent.HiringJob) bool {
			return hiringJob.ID == record.HiringJobID
		})
		record.Edges.CandidateEdge = candidateEdge
		record.Edges.CandidateJobInterview = interviewEdges
		record.Edges.HiringJobEdge = hiringJobEdge
	}
}
