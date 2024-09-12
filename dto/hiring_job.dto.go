package dto

import (
	"encoding/json"
	"fmt"
	"reflect"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/models"

	"github.com/samber/lo"
)

type HiringJobDto interface {
	AuditTrailCreate(record *ent.HiringJob) (string, error)
	AuditTrailDelete(record *ent.HiringJob) (string, error)
	AuditTrailUpdate(oldRecord *ent.HiringJob, newRecord *ent.HiringJob) (string, error)
	MappingEdge(records []*ent.HiringJob, hiringTeams []*ent.HiringTeam)
	MappingPriority(input int) string
	MappingLocation(input hiringjob.Location) string
	MappingStatus(input hiringjob.Status) string
	MappingSalary(input *ent.HiringJob) string
}

type hiringJobDtoImpl struct {
	hiringJobStepDto HiringJobStepDto
}

var hiringJobFieldI18n = map[string]models.I18nFormat{
	"Name": {
		AuditTrail: "model.hiring_jobs.name",
		Email:      "Name",
	},
	"Description": {
		AuditTrail: "model.hiring_jobs.description",
		Email:      "Job description",
	},
	"Amount": {
		AuditTrail: "model.hiring_jobs.amount",
		Email:      "Staff required",
	},
	"Location": {
		AuditTrail: "model.hiring_jobs.location",
		Email:      "Location",
	},
	"SalaryType": {
		AuditTrail: "model.hiring_jobs.salary_type",
		Email:      "Salary",
	},
	"SalaryFrom": {
		AuditTrail: "model.hiring_jobs.salary_from",
		Email:      "From",
	},
	"SalaryTo": {
		AuditTrail: "model.hiring_jobs.salary_to",
		Email:      "To",
	},
	"Currency": {
		AuditTrail: "model.hiring_jobs.currency",
		Email:      "Unit",
	},
	"HiringTeamID": {
		AuditTrail: "model.hiring_jobs.hiring_team",
		Email:      "Hiring team",
	},
	"CreatedBy": {
		AuditTrail: "model.hiring_jobs.created_by",
		Email:      "Requester",
	},
	"Status": {
		AuditTrail: "model.hiring_jobs.status",
		Email:      "Status",
	},
	"Priority": {
		AuditTrail: "model.hiring_jobs.priority",
		Email:      "Priority",
	},
	"JobPositionID": {
		AuditTrail: "model.hiring_jobs.job_position",
		Email:      "Job position",
	},
	"RecInChargeID": {
		AuditTrail: "model.hiring_jobs.rec_in_charge",
		Email:      "REC in charge",
	},
	"Level": {
		AuditTrail: "model.hiring_jobs.level",
		Email:      "Staff level",
	},
	"RecTeamID": {
		AuditTrail: "model.hiring_jobs.rec_team",
		Email:      "REC team",
	},
	"Note": {
		AuditTrail: "model.hiring_jobs.note",
		Email:      "Note",
	},
}

func NewHiringJobDto() HiringJobDto {
	return &hiringJobDtoImpl{
		hiringJobStepDto: NewHiringJobStepDto(),
	}
}

func (d hiringJobDtoImpl) AuditTrailCreate(record *ent.HiringJob) (string, error) {
	result := models.AuditTrailData{
		Module: HiringJobI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
	d.hiringJobStepDto.ProcessAuditTrail([]*ent.HiringJobStep{}, record.Edges.ApprovalSteps, &result)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringJobDtoImpl) AuditTrailDelete(record *ent.HiringJob) (string, error) {
	result := models.AuditTrailData{
		Module: HiringJobI18n,
		Create: []interface{}{},
		Update: []interface{}{},
		Delete: d.recordAudit(record),
	}
	d.hiringJobStepDto.ProcessAuditTrail([]*ent.HiringJobStep{}, record.Edges.ApprovalSteps, &result)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringJobDtoImpl) AuditTrailUpdate(oldRecord *ent.HiringJob, newRecord *ent.HiringJob) (string, error) {
	result := models.AuditTrailData{
		Module: HiringJobI18n,
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
			case "model.hiring_jobs.status":
				oldValueField = d.statusI18n(oldRecord.Status)
				newValueField = d.statusI18n(newRecord.Status)
			case "model.hiring_jobs.location":
				oldValueField = d.locationI18n(oldRecord.Location)
				newValueField = d.locationI18n(newRecord.Location)
			case "model.hiring_jobs.salary_type":
				oldValueField = d.salaryTypeI18n(oldRecord.SalaryType)
				newValueField = d.salaryTypeI18n(newRecord.SalaryType)
			case "model.hiring_jobs.currency":
				oldValueField = d.currencyI18n(oldRecord.Currency)
				newValueField = d.currencyI18n(newRecord.Currency)
			case "model.hiring_jobs.hiring_team":
				oldValueField = ""
				if oldRecord.Edges.HiringTeamEdge != nil {
					oldValueField = oldRecord.Edges.HiringTeamEdge.Name
				}
				newValueField = ""
				if newRecord.Edges.HiringTeamEdge != nil {
					newValueField = newRecord.Edges.HiringTeamEdge.Name
				}
			case "model.hiring_jobs.created_by":
				oldValueField = ""
				if oldRecord.Edges.OwnerEdge != nil {
					oldValueField = oldRecord.Edges.OwnerEdge.Name
				}
				newValueField = ""
				if newRecord.Edges.OwnerEdge != nil {
					newValueField = newRecord.Edges.OwnerEdge.Name
				}
			case "model.hiring_jobs.priority":
				oldValueField = d.priorityI18n(oldRecord.Priority)
				newValueField = d.priorityI18n(newRecord.Priority)
			case "model.hiring_jobs.job_position":
				oldValueField = ""
				if oldRecord.Edges.JobPositionEdge != nil {
					oldValueField = oldRecord.Edges.JobPositionEdge.Name
				}
				newValueField = ""
				if newRecord.Edges.JobPositionEdge != nil {
					newValueField = newRecord.Edges.JobPositionEdge.Name
				}
			case "model.hiring_jobs.rec_team":
				oldValueField = oldRecord.Edges.RecTeamEdge.Name
				newValueField = newRecord.Edges.RecTeamEdge.Name
			case "model.hiring_jobs.level":
				oldValueField = d.levelI18n(oldRecord.Level)
				newValueField = d.levelI18n(newRecord.Level)
			case "model.hiring_jobs.rec_in_charge":
				oldValueField = ""
				if oldRecord.Edges.RecInChargeEdge != nil {
					oldValueField = oldRecord.Edges.RecInChargeEdge.Name
				}
				newValueField = newRecord.Edges.RecInChargeEdge.Name
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
	entity = d.skillAuditTrailUpdate(oldRecord, newRecord, entity)
	result.Update = append(result.Update, entity...)
	d.hiringJobStepDto.ProcessAuditTrail(oldRecord.Edges.ApprovalSteps, newRecord.Edges.ApprovalSteps, &result)
	jsonObj, err := json.Marshal(result)
	return string(jsonObj), err
}

func (d hiringJobDtoImpl) recordAudit(record *ent.HiringJob) []interface{} {
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
		case "model.hiring_jobs.status":
			valueField = d.statusI18n(record.Status)
		case "model.hiring_jobs.location":
			valueField = d.locationI18n(record.Location)
		case "model.hiring_jobs.salary_type":
			valueField = d.salaryTypeI18n(record.SalaryType)
		case "model.hiring_jobs.currency":
			valueField = d.currencyI18n(record.Currency)
		case "model.hiring_jobs.hiring_team":
			valueField = ""
			if record.Edges.HiringTeamEdge != nil {
				valueField = record.Edges.HiringTeamEdge.Name
			}
		case "model.hiring_jobs.created_by":
			valueField = ""
			if record.Edges.OwnerEdge != nil {
				valueField = record.Edges.OwnerEdge.Name
			}
		case "model.hiring_jobs.priority":
			valueField = d.priorityI18n(record.Priority)
		case "model.hiring_jobs.job_position":
			valueField = ""
			if record.Edges.JobPositionEdge != nil {
				valueField = record.Edges.JobPositionEdge.Name
			}
		case "model.hiring_jobs.rec_team":
			valueField = record.Edges.RecTeamEdge.Name
		case "model.hiring_jobs.rec_in_charge":
			valueField = ""
			if record.Edges.RecInChargeEdge != nil {
				valueField = record.Edges.RecInChargeEdge.Name
			}
		case "model.hiring_jobs.level":
			valueField = d.levelI18n(record.Level)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	entity = d.skillAuditTrail(record, entity)
	return entity
}

func (d hiringJobDtoImpl) skillAuditTrail(record *ent.HiringJob, atInterface []interface{}) []interface{} {
	if len(record.Edges.HiringJobSkillEdges) == 0 {
		return atInterface
	}
	attachmentNames := lo.Map(record.Edges.HiringJobSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	atInterface = append(atInterface, models.AuditTrailCreateDelete{
		Field: "model.hiring_jobs.skills",
		Value: attachmentNames,
	})
	return atInterface
}

func (d hiringJobDtoImpl) skillAuditTrailUpdate(oldRecord *ent.HiringJob, newRecord *ent.HiringJob, atInterface []interface{}) []interface{} {
	if len(oldRecord.Edges.HiringJobSkillEdges) == 0 && len(newRecord.Edges.HiringJobSkillEdges) == 0 {
		return atInterface
	}
	oldUserNames := lo.Map(oldRecord.Edges.HiringJobSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	newUserNames := lo.Map(newRecord.Edges.HiringJobSkillEdges, func(entity *ent.EntitySkill, index int) string {
		return entity.Edges.SkillEdge.Name
	})
	if !CompareArray(oldUserNames, newUserNames) {
		atInterface = append(atInterface, models.AuditTrailUpdate{
			Field: "model.hiring_jobs.skills",
			Value: models.ValueChange{
				OldValue: oldUserNames,
				NewValue: newUserNames,
			},
		})
	}
	return atInterface
}

func (d hiringJobDtoImpl) formatFieldI18n(input string) string {
	if v, ok := hiringJobFieldI18n[input]; ok {
		return v.AuditTrail
	}
	return ""
}

func (d hiringJobDtoImpl) locationI18n(input hiringjob.Location) string {
	switch input {
	case hiringjob.LocationHaNoi:
		return "model.hiring_jobs.location_enum.ha_noi"
	case hiringjob.LocationHoChiMinh:
		return "model.hiring_jobs.location_enum.ho_chi_minh"
	case hiringjob.LocationDaNang:
		return "model.hiring_jobs.location_enum.da_nang"
	case hiringjob.LocationJapan:
		return "model.hiring_jobs.location_enum.japan"
	case hiringjob.LocationSingapore:
		return "model.hiring_jobs.location_enum.singapore"
	}
	return ""
}

func (d hiringJobDtoImpl) statusI18n(input hiringjob.Status) string {
	switch input {
	case hiringjob.StatusOpened:
		return "model.hiring_jobs.status_enum.opened"
	case hiringjob.StatusClosed:
		return "model.hiring_jobs.status_enum.closed"
	case hiringjob.StatusPendingApprovals:
		return "model.hiring_jobs.status_enum.pending_approvals"
	case hiringjob.StatusCancelled:
		return "model.hiring_jobs.status_enum.cancelled"
	}
	return ""
}

func (d hiringJobDtoImpl) salaryTypeI18n(input hiringjob.SalaryType) string {
	switch input {
	case hiringjob.SalaryTypeRange:
		return "model.hiring_jobs.salary_type_enum.range"
	case hiringjob.SalaryTypeUpTo:
		return "model.hiring_jobs.salary_type_enum.up_to"
	case hiringjob.SalaryTypeNegotiate:
		return "model.hiring_jobs.salary_type_enum.negotiate"
	case hiringjob.SalaryTypeMinimum:
		return "model.hiring_jobs.salary_type_enum.minimum"
	}
	return ""
}

func (d hiringJobDtoImpl) currencyI18n(input hiringjob.Currency) string {
	switch input {
	case hiringjob.CurrencyVnd:
		return "model.hiring_jobs.currency_enum.vnd"
	case hiringjob.CurrencyUsd:
		return "model.hiring_jobs.currency_enum.usd"
	case hiringjob.CurrencyJpy:
		return "model.hiring_jobs.currency_enum.jpy"
	}
	return ""
}

func (d hiringJobDtoImpl) priorityI18n(input int) string {
	switch input {
	case 4:
		return "model.hiring_jobs.priority_enum.low"
	case 3:
		return "model.hiring_jobs.priority_enum.medium"
	case 2:
		return "model.hiring_jobs.priority_enum.high"
	case 1:
		return "model.hiring_jobs.priority_enum.urgent"
	}
	return ""
}

func (d hiringJobDtoImpl) levelI18n(input hiringjob.Level) string {
	switch input {
	case hiringjob.LevelIntern:
		return "model.hiring_jobs.level_enum.intern"
	case hiringjob.LevelJunior:
		return "model.hiring_jobs.level_enum.junior"
	case hiringjob.LevelSenior:
		return "model.hiring_jobs.level_enum.senior"
	case hiringjob.LevelFresher:
		return "model.hiring_jobs.level_enum.fresher"
	case hiringjob.LevelMiddle:
		return "model.hiring_jobs.level_enum.middle"
	case hiringjob.LevelManager:
		return "model.hiring_jobs.level_enum.manager"
	case hiringjob.LevelDirector:
		return "model.hiring_jobs.level_enum.director"
	}
	return ""
}

func (d hiringJobDtoImpl) MappingEdge(records []*ent.HiringJob, hiringTeams []*ent.HiringTeam) {
	for _, record := range records {
		record.Edges.HiringTeamEdge, _ = lo.Find(hiringTeams, func(team *ent.HiringTeam) bool {
			return team.ID == record.HiringTeamID
		})
	}
}

func (d hiringJobDtoImpl) MappingPriority(input int) string {
	switch input {
	case 4:
		return "LOW"
	case 3:
		return "MEDIUM"
	case 2:
		return "HIGH"
	case 1:
		return "URGENT"
	}
	return ""
}

func (d hiringJobDtoImpl) MappingLocation(input hiringjob.Location) string {
	switch input {
	case hiringjob.LocationHaNoi:
		return "Hà Nội"
	case hiringjob.LocationHoChiMinh:
		return "Hồ Chí Minh"
	case hiringjob.LocationDaNang:
		return "Đà Nẵng"
	case hiringjob.LocationJapan:
		return "Japan"
	case hiringjob.LocationSingapore:
		return "Singapore"
	}
	return ""
}

func (d hiringJobDtoImpl) MappingStatus(input hiringjob.Status) string {
	switch input {
	case hiringjob.StatusOpened:
		return "Opening"
	case hiringjob.StatusClosed:
		return "Closed"
	case hiringjob.StatusPendingApprovals:
		return "Pending approvals"
	case hiringjob.StatusCancelled:
		return "Cancelled"
	}
	return ""
}

func (d hiringJobDtoImpl) mappingCurrency(input hiringjob.Currency) string {
	switch input {
	case hiringjob.CurrencyVnd:
		return "VND"
	case hiringjob.CurrencyUsd:
		return "USD"
	case hiringjob.CurrencyJpy:
		return "JPY"
	}
	return ""
}

func (d hiringJobDtoImpl) MappingSalary(input *ent.HiringJob) string {
	switch input.SalaryType {
	case hiringjob.SalaryTypeRange:
		return fmt.Sprintf("Range: %s - %s %s", FormatCurrency(input.SalaryFrom), FormatCurrency(input.SalaryTo), d.mappingCurrency(input.Currency))
	case hiringjob.SalaryTypeUpTo:
		return fmt.Sprintf("Up To: %s %s", FormatCurrency(input.SalaryTo), d.mappingCurrency(input.Currency))
	case hiringjob.SalaryTypeNegotiate:
		return "Negotiate"
	case hiringjob.SalaryTypeMinimum:
		return fmt.Sprintf("Minimum: %s %s", FormatCurrency(input.SalaryFrom), d.mappingCurrency(input.Currency))
	}
	return ""
}
