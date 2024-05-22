package util

import "time"

type Model string

var DefaultTime = time.Time{}

const (
	Coefficient     Model = "Coefficient"
	CoefficientType Model = "Coefficient Type"
	SalaryComponent Model = "Salary Component"
	EmployeeSalary  Model = "Employee's Salary"
	SalaryTemplate  Model = "Salary Template"
	PayrollTemplate Model = "Payroll Template"
)
