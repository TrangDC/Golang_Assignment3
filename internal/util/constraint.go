package util

import "time"

type Model string

var DefaultTime, _ = time.Parse("2006-01", "0000-01")

const (
	Coefficient     Model = "Coefficient"
	CoefficientType Model = "Coefficient Type"
	SalaryComponent Model = "Salary Component"
	EmployeeSalary  Model = "Employee's Salary"
	SalaryTemplate  Model = "Salary Template"
	PayrollTemplate Model = "Payroll Template"
)
