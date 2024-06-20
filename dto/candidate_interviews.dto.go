package dto

import (
	"encoding/json"
	"fmt"
	"reflect"
	"trec/ent"
	"trec/ent/candidateinterview"
	"trec/models"

	"github.com/samber/lo"
)

type CandidateInterviewDto interface {
	AuditTrailCreate(record *ent.CandidateInterview) (string, error)
	AuditTrailDelete(record *ent.CandidateInterview) (string, error)
	AuditTrailUpdate(oldRecord *ent.CandidateInterview, newRecord *ent.CandidateInterview) (string, error)
}

type candidateInterviewDtoImpl struct {
}

func NewCandidateInterviewDto() CandidateInterviewDto {
	return &candidateInterviewDtoImpl{}
}

func (d *candidateInterviewDtoImpl) AuditTrailCreate(record *ent.CandidateInterview) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateInterviewAt := models.AuditTrailData{
		Module: CandidateInterviewI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	fatherModule := []models.AuditTrailFatherModule{}
	fatherModule = append(fatherModule, models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", record.Edges.CandidateJobEdge.Status),
		Value:  record.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	})
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateInterviewAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateInterviewDtoImpl) AuditTrailDelete(record *ent.CandidateInterview) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateInterviewAt := models.AuditTrailData{
		Module: CandidateInterviewI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	fatherModule := []models.AuditTrailFatherModule{}
	fatherModule = append(fatherModule, models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", record.Edges.CandidateJobEdge.Status),
		Value:  record.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	})
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateInterviewAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateInterviewDtoImpl) AuditTrailUpdate(oldRecord *ent.CandidateInterview, newRecord *ent.CandidateInterview) (string, error) {
	result := models.AuditTrailData{
		Module:    CandidateI18n,
		Create:    []interface{}{},
		Update:    []interface{}{},
		Delete:    []interface{}{},
		SubModule: []interface{}{},
	}
	candidateInterviewAt := models.AuditTrailData{
		Module: CandidateInterviewI18n,
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
			case "model.candidate_interviews.interview_date":
				oldValueField = oldRecord.InterviewDate
				newValueField = newRecord.InterviewDate
			case "model.candidate_interviews.start_from":
				oldValueField = oldRecord.StartFrom
				newValueField = newRecord.StartFrom
			case "model.candidate_interviews.end_at":
				oldValueField = oldRecord.EndAt
				newValueField = newRecord.EndAt
			case "model.candidate_interviews.status":
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
	entity = d.memberAuditTrailUpdate(oldRecord, newRecord, entity)
	candidateInterviewAt.Update = append(candidateInterviewAt.Update, entity...)
	fatherModule := []models.AuditTrailFatherModule{}
	fatherModule = append(fatherModule, models.AuditTrailFatherModule{
		Module: CandidateJobI18n + fmt.Sprintf(".%s", oldRecord.Edges.CandidateJobEdge.Status),
		Value:  oldRecord.Edges.CandidateJobEdge.Edges.HiringJobEdge.Name,
	})
	result.SubModule = append(result.SubModule, fatherModule)
	result.SubModule = append(result.SubModule, candidateInterviewAt)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateInterviewDtoImpl) recordAudit(record *ent.CandidateInterview) []interface{} {
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
		case "model.candidate_interviews.interview_date":
			valueField = record.InterviewDate
		case "model.candidate_interviews.start_from":
			valueField = record.StartFrom
		case "model.candidate_interviews.end_at":
			valueField = record.EndAt
		case "model.candidate_interviews.created_by":
			valueField = record.Edges.CreatedByEdge.Name
		case "model.candidate_interviews.status":
			valueField = d.statusI18n(record.Status)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.memberAuditTrail(record, entity)
	return entity
}

func (d candidateInterviewDtoImpl) memberAuditTrail(record *ent.CandidateInterview, atInterface []interface{}) []interface{} {
	if len(record.Edges.InterviewerEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.InterviewerEdges, func(member *ent.User, index int) string {
		return member.Name
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidate_interviews.members",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateInterviewDtoImpl) memberAuditTrailUpdate(oldRecord *ent.CandidateInterview, newRecord *ent.CandidateInterview, atInterface []interface{}) []interface{} {
	if len(oldRecord.Edges.InterviewerEdges) == 0 && len(newRecord.Edges.InterviewerEdges) == 0 {
		return atInterface
	}
	oldAttachmentNames := lo.Map(oldRecord.Edges.InterviewerEdges, func(member *ent.User, index int) string {
		return member.Name
	})
	newAttachmentNames := lo.Map(newRecord.Edges.InterviewerEdges, func(member *ent.User, index int) string {
		return member.Name
	})
	if !CompareArray(oldAttachmentNames, newAttachmentNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.candidate_interviews.members",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d candidateInterviewDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Title":
		return "model.candidate_interviews.title"
	case "Description":
		return "model.candidate_interviews.description"
	case "InterviewDate":
		return "model.candidate_interviews.interview_date"
	case "StartFrom":
		return "model.candidate_interviews.start_from"
	case "EndAt":
		return "model.candidate_interviews.end_at"
	case "CreatedBy":
		return "model.candidate_interviews.created_by"
	case "Status":
		return "model.candidate_interviews.status"
	}
	return ""
}

func (d candidateInterviewDtoImpl) statusI18n(input candidateinterview.Status) string {
	switch input {
	case candidateinterview.StatusInvitedToInterview:
		return "model.candidate_interviews.status_enum.invited_to_interview"
	case candidateinterview.StatusInterviewing:
		return "model.candidate_interviews.status_enum.interviewing"
	case candidateinterview.StatusDone:
		return "model.candidate_interviews.status_enum.done"
	case candidateinterview.StatusCancelled:
		return "model.candidate_interviews.status_enum.cancelled"
	}
	return ""
}
