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
}

func NewHiringJobDto() HiringJobDto {
	return &hiringJobDtoImpl{}
}

func (d hiringJobDtoImpl) AuditTrailCreate(record *ent.HiringJob) (string, error) {
	result := models.AuditTrailData{
		Module: HiringJobI18n,
		Create: d.recordAudit(record),
		Update: []interface{}{},
		Delete: []interface{}{},
	}
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
	switch input {
	case "Name":
		return "model.hiring_jobs.name"
	case "Description":
		return "model.hiring_jobs.description"
	case "Amount":
		return "model.hiring_jobs.amount"
	case "Location":
		return "model.hiring_jobs.location"
	case "SalaryType":
		return "model.hiring_jobs.salary_type"
	case "SalaryFrom":
		return "model.hiring_jobs.salary_from"
	case "SalaryTo":
		return "model.hiring_jobs.salary_to"
	case "Currency":
		return "model.hiring_jobs.currency"
	case "HiringTeamID":
		return "model.hiring_jobs.hiring_team"
	case "CreatedBy":
		return "model.hiring_jobs.created_by"
	case "Status":
		return "model.hiring_jobs.status"
	case "Priority":
		return "model.hiring_jobs.priority"
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
	case hiringjob.StatusDraft:
		return "model.hiring_jobs.status_enum.draft"
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
		return "HA NOI"
	case hiringjob.LocationHoChiMinh:
		return "HO CHI MINH"
	case hiringjob.LocationDaNang:
		return "DA NANG"
	case hiringjob.LocationJapan:
		return "JAPAN"
	case hiringjob.LocationSingapore:
		return "SINGAPORE"
	}
	return ""
}

func (d hiringJobDtoImpl) MappingStatus(input hiringjob.Status) string {
	switch input {
	case hiringjob.StatusOpened:
		return "OPENED"
	case hiringjob.StatusClosed:
		return "CLOSED"
	case hiringjob.StatusDraft:
		return "DRAFT"
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
