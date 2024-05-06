package dto

import "trec/ent/hiringjob"

func FormatHiringJobField(input string) string {
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
	}
	return ""
}

func LocationI18n(input hiringjob.Location) string {
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

func StatusI18n(input hiringjob.Status) string {
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

func SalaryTypeI18n(input hiringjob.SalaryType) string {
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

func CurrencyI18n(input hiringjob.Currency) string {
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
