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
	MappingLocation(input string) string
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
			case "model.candidate_interviews.location":
				oldValueField = d.MappingLocation(oldRecord.Location)
				newValueField = d.MappingLocation(newRecord.Location)
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
			if record.Edges.CreatedByEdge != nil {
				valueField = record.Edges.CreatedByEdge.Name
			} else {
				valueField = ""
			}
		case "model.candidate_interviews.status":
			valueField = d.statusI18n(record.Status)
		case "model.candidate_interviews.location":
			valueField = d.MappingLocation(record.Location)
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
	case "Location":
		return "model.candidate_interviews.location"
	case "MeetingLink":
		return "model.candidate_interviews.meeting_link"
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

func (d candidateInterviewDtoImpl) MappingLocation(input string) string {
	switch input {
	case "Hanoi":
		return "[Hanoi] TECHVIFY Office_Thanh Dong Bld 19 To Huu, Trung Van, Nam Tu Liem"
	case "HCM":
		return "[HCM] TECHVIFY Office_ H3 Building, 384 Hoang Dieu str, 6 Ward, 4 Dist, Ho Chi Minh City"
	case "ĐN":
		return "[ĐN] F3 Ricco Building, 363 Nguyen Huu Tho Str, Cam Le Dist, Da Nang City, Vietnam"
	case "Japan":
		return "[Japan] Hakata Ekimae City Building 10F, 1-9-3 Hakata Ekimae, Hakata-ku, Fukuoka-shi, Fukuoka 812-0011 Japan"
	case "Online Interview":
		return "Online interview"
	default:
		return input
	}
}
