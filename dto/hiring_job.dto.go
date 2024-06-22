package dto

import (
	"encoding/json"
	"reflect"
	"trec/ent"
	"trec/ent/hiringjob"
	"trec/models"
)

type HiringJobDto interface {
	AuditTrailCreate(record *ent.HiringJob) (string, error)
	AuditTrailDelete(record *ent.HiringJob) (string, error)
	AuditTrailUpdate(oldRecord *ent.HiringJob, newRecord *ent.HiringJob) (string, error)
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
			case "model.hiring_jobs.team":
				oldValueField = oldRecord.Edges.TeamEdge.Name
				newValueField = newRecord.Edges.TeamEdge.Name
			case "model.hiring_jobs.created_by":
				oldValueField = oldRecord.Edges.OwnerEdge.Name
				newValueField = newRecord.Edges.OwnerEdge.Name
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
		case "model.hiring_jobs.team":
			valueField = record.Edges.TeamEdge.Name
		case "model.hiring_jobs.created_by":
			valueField = record.Edges.OwnerEdge.Name
		case "model.hiring_jobs.priority":
			valueField = d.priorityI18n(record.Priority)
		}
		entity = append(entity, models.AuditTrailCreateDelete{
			Field: fieldName,
			Value: valueField,
		})
	}
	return entity
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
	case "TeamID":
		return "model.hiring_jobs.team"
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
