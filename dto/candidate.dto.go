package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/candidate"
	"trec/models"

	"github.com/samber/lo"
)

type CandidateDto interface {
	AuditTrailCreate(record *ent.Candidate) (string, error)
	AuditTrailDelete(record *ent.Candidate) (string, error)
	AuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate) (string, error)
}

type candidateDtoImpl struct {
}

func NewCandidateDto() CandidateDto {
	return &candidateDtoImpl{}
}

func (d *candidateDtoImpl) AuditTrailCreate(record *ent.Candidate) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateDtoImpl) AuditTrailDelete(record *ent.Candidate) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateDtoImpl) AuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate) (string, error) {
	result := models.AuditTrailData{
		Module: CandidateI18n,
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
			case "model.candidates.is_blacklist":
				oldValueField = d.isBlacklistI18n(oldRecord.IsBlacklist)
				newValueField = d.isBlacklistI18n(newRecord.IsBlacklist)
			case "model.candidates.reference_type":
				oldValueField = d.referenceTypeI18n(oldRecord.ReferenceType)
				newValueField = d.referenceTypeI18n(newRecord.ReferenceType)
			case "model.candidates.reference_value":
				oldValueField = d.referenceTypeValueI18n(oldRecord.ReferenceType, oldRecord.ReferenceValue)
				newValueField = d.referenceTypeValueI18n(newRecord.ReferenceType, newRecord.ReferenceValue)
			case "model.candidates.reference_user":
				oldValueField = oldRecord.Edges.ReferenceUserEdge.Name
				newValueField = newRecord.Edges.ReferenceUserEdge.Name
			case "model.candidates.recruit_time":
				oldValueField = ""
				newValueField = ""
				if !oldRecord.RecruitTime.IsZero() {
					oldValueField = oldRecord.RecruitTime
				}
				if !newRecord.RecruitTime.IsZero() {
					newValueField = newRecord.RecruitTime
				}
			case "model.candidates.dob":
				oldValueField = ""
				newValueField = ""
				if !oldRecord.Dob.IsZero() {
					oldValueField = oldRecord.Dob
				}
				if !newRecord.Dob.IsZero() {
					newValueField = newRecord.Dob
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
	entity = d.skillAuditTrailUpdate(oldRecord, newRecord, entity)
	result.Update = append(result.Update, entity...)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d *candidateDtoImpl) recordAudit(record *ent.Candidate) []interface{} {
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
		case "model.candidates.is_blacklist":
			valueField = d.isBlacklistI18n(record.IsBlacklist)
		case "model.candidates.reference_type":
			valueField = d.referenceTypeI18n(record.ReferenceType)
		case "model.candidates.reference_value":
			valueField = d.referenceTypeValueI18n(record.ReferenceType, record.ReferenceValue)
		case "model.candidates.reference_user":
			valueField = record.Edges.ReferenceUserEdge.Name
		case "model.candidates.recruit_time":
			if !record.RecruitTime.IsZero() {
				valueField = record.RecruitTime
			}
		case "model.candidates.dob":
			if !record.Dob.IsZero() {
				valueField = record.Dob
			}
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.attachmentAuditTrail(record, entity)
	entity = d.skillAuditTrail(record, entity)
	return entity
}

func (d candidateDtoImpl) attachmentAuditTrail(record *ent.Candidate, atInterface []interface{}) []interface{} {
	if len(record.Edges.AttachmentEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.AttachmentEdges, func(document *ent.Attachment, index int) string {
		return document.DocumentName
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidates.document",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateDtoImpl) attachmentAuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate, atInterface []interface{}) []interface{} {
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
			Field: "model.candidates.document",
			Value: models.ValueChange{
				OldValue: oldAttachmentNames,
				NewValue: newAttachmentNames,
			},
		})
	}
	return atInterface
}

func (d candidateDtoImpl) skillAuditTrail(record *ent.Candidate, atInterface []interface{}) []interface{} {
	if len(record.Edges.CandidateSkillEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.CandidateSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.candidates.skills",
		Value: attachmentNames,
	})
	return atInterface
}

func (d candidateDtoImpl) skillAuditTrailUpdate(oldRecord *ent.Candidate, newRecord *ent.Candidate, atInterface []interface{}) []interface{} {
	if len(oldRecord.Edges.CandidateSkillEdges) == 0 && len(newRecord.Edges.CandidateSkillEdges) == 0 {
		return atInterface
	}
	oldUserNames := lo.Map(oldRecord.Edges.CandidateSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	newUserNames := lo.Map(newRecord.Edges.CandidateSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	if !CompareArray(oldUserNames, newUserNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.candidates.skills",
			Value: models.ValueChange{
				OldValue: oldUserNames,
				NewValue: newUserNames,
			},
		})
	}
	return atInterface
}

func (d *candidateDtoImpl) formatFieldI18n(input string) string {
	switch input {
	case "Name":
		return "model.candidates.name"
	case "Email":
		return "model.candidates.email"
	case "Phone":
		return "model.candidates.phone"
	case "Dob":
		return "model.candidates.dob"
	case "IsBlacklist":
		return "model.candidates.is_blacklist"
	case "ReferenceType":
		return "model.candidates.reference_type"
	case "ReferenceValue":
		return "model.candidates.reference_value"
	case "ReferenceUid":
		return "model.candidates.reference_user"
	case "RecruitTime":
		return "model.candidates.recruit_time"
	case "Description":
		return "model.candidates.description"
	case "Country":
		return "model.candidates.country"
	}
	return ""
}

func (d *candidateDtoImpl) isBlacklistI18n(input bool) string {
	switch input {
	case true:
		return "model.candidates.is_blacklist_enum.yes"
	default:
		return "model.candidates.is_blacklist_enum.no"
	}
}

func (d *candidateDtoImpl) referenceTypeI18n(input candidate.ReferenceType) string {
	switch input {
	case candidate.ReferenceTypeEb:
		return "model.candidates.reference_type.eb"
	case candidate.ReferenceTypeRec:
		return "model.candidates.reference_type.rec"
	case candidate.ReferenceTypeHiringPlatform:
		return "model.candidates.reference_type.hiring_platform"
	case candidate.ReferenceTypeReference:
		return "model.candidates.reference_type.reference"
	default:
		return "model.candidates.reference_type.headhunt"
	}
}

func (d *candidateDtoImpl) referenceTypeValueI18n(referenceType candidate.ReferenceType, input string) string {
	switch referenceType {
	case candidate.ReferenceTypeEb:
		switch input {
		case ent.CandidateReferenceEbTiktokTechvifyOfficial.String():
			return "model.candidates.reference_type.eb.reference_value.tiktok_techvify_official"
		case ent.CandidateReferenceEbTiktokThedevdad.String():
			return "model.candidates.reference_type.eb.reference_value.tiktok_thedevdad"
		case ent.CandidateReferenceEbLinkedinJunieTruong.String():
			return "model.candidates.reference_type.eb.reference_value.linkedin_junie_truong"
		case ent.CandidateReferenceEbOtherLinkedin.String():
			return "model.candidates.reference_type.eb.reference_value.other_linkedin"
		case ent.CandidateReferenceEbGroupSeeding.String():
			return "model.candidates.reference_type.eb.reference_value.group_seeding"
		case ent.CandidateReferenceEbFanpageTechvifyCareers.String():
			return "model.candidates.reference_type.eb.reference_value.fanpage_techvify_careers"
		case ent.CandidateReferenceEbGoogleSearch.String():
			return "model.candidates.reference_type.eb.reference_value.google_search"
		case ent.CandidateReferenceEbYoutubeTechvifyCareers.String():
			return "model.candidates.reference_type.eb.reference_value.youtube_techvify_careers"
		case ent.CandidateReferenceEbThread.String():
			return "model.candidates.reference_type.eb.reference_value.thread"
		case ent.CandidateReferenceEbInstagram.String():
			return "model.candidates.reference_type.eb.reference_value.instagram"
		case ent.CandidateReferenceEbTwitter.String():
			return "model.candidates.reference_type.eb.reference_value.twitter"
		case ent.CandidateReferenceEbOthers.String():
			return "model.candidates.reference_type.eb.reference_value.others"
		default:
			return input
		}
	case candidate.ReferenceTypeRec:
		switch input {
		case ent.CandidateReferenceRecLinkedin.String():
			return "model.candidates.reference_type.rec.reference_value.linkedin"
		case ent.CandidateReferenceRecFacebook.String():
			return "model.candidates.reference_type.rec.reference_value.facebook"
		case ent.CandidateReferenceRecInstagram.String():
			return "model.candidates.reference_type.rec.reference_value.instagram"
		case ent.CandidateReferenceRecThread.String():
			return "model.candidates.reference_type.rec.reference_value.thread"
		case ent.CandidateReferenceRecGithub.String():
			return "model.candidates.reference_type.rec.reference_value.github"
		default:
			return input
		}
	case candidate.ReferenceTypeHiringPlatform:
		switch input {
		case ent.CandidateReferenceHiringPlatformTopcv.String():
			return "model.candidates.reference_type.hiring_platform.reference_value.topcv"
		case ent.CandidateReferenceHiringPlatformVietnamWorks.String():
			return "model.candidates.reference_type.hiring_platform.reference_value.vietnam_works"
		case ent.CandidateReferenceHiringPlatformItviec.String():
			return "model.candidates.reference_type.hiring_platform.reference_value.itviec"
		default:
			return input
		}
	}
	return input
}
