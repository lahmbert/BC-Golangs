package models

import "fmt"

type PermanentEmp struct {
	Employee

	Tunjangan
}

func NewPermanent(
	fullName string,
	salary float64) *PermanentEmp {

	tunjangan := Tunjangan{
		insurance: 500_000,
		overtime:  0,
		allowance: 0}

	return &PermanentEmp{
		Employee: *NewEmployee(
			fullName,
			salary,
			Permanents),
		Tunjangan: tunjangan,
	}
}

func (emp *PermanentEmp) GetAllowance() float64 {
	return emp.allowance
}

func (emp *PermanentEmp) GetInsurance() float64 {
	return emp.insurance
}

func (emp *PermanentEmp) GetOvertime() float64 {
	return emp.overtime
}

func (permanent *PermanentEmp) GetfullName() string {
	return permanent.fullName
}

func (permanent *PermanentEmp) GetSalary() float64 {
	return permanent.salary
}

func (permanent *PermanentEmp) GetStatus() Status {
	return permanent.status
}

func (permanent *PermanentEmp) ToString() string {
	return fmt.Sprintf("\nfullname : %s\nsalary :%.2f\nstatus :%s\ninsurance :%.2f\novertime :%.2f\nallowance :%.2f\ntotal_salary :%.2f",
		permanent.fullName,
		permanent.salary,
		permanent.status,
		permanent.insurance,
		permanent.overtime,
		permanent.allowance,
		permanent.GetTotalSalary())
}

func (permanent *PermanentEmp) GetTotalSalary() float64 {
	return permanent.salary + permanent.insurance + permanent.overtime + permanent.allowance
}
