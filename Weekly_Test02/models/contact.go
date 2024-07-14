package models

import "fmt"

type ContractEmp struct {
	Employee

	Tunjangan
}

func NewContract(
	fullName string,
	salary float64) *ContractEmp {

	tunjangan := Tunjangan{
		insurance: 0,
		overtime:  55_000,
		allowance: 0}

	return &ContractEmp{
		Employee: *NewEmployee(
			fullName,
			salary,
			Contracts),
		Tunjangan: tunjangan,
	}
}

func (contract *ContractEmp) GetAllowance() float64 {
	return contract.allowance
}

func (contract *ContractEmp) GetInsurance() float64 {
	return contract.insurance
}

func (contract *ContractEmp) GetOvertime() float64 {
	return contract.overtime
}

func (contract *ContractEmp) ToString() string {
	return fmt.Sprintf("\nfullname : %s\nsalary :%.2f\nstatus :%s\ninsurance :%.2f\novertime :%.2f\nallowance :%.2f\ntotal_salary :%.2f",
		contract.fullName,
		contract.salary,
		contract.status,
		contract.insurance,
		contract.overtime,
		contract.allowance,
		contract.GetTotalSalary())
}

func (contract *ContractEmp) GetTotalSalary() float64 {
	return contract.salary + contract.insurance + contract.overtime + contract.allowance
}
