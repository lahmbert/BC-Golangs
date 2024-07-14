package models

import "fmt"

type TraineeEmp struct {
	Employee

	Tunjangan
}

func NewTrainee(
	fullName string,
	salary float64) *TraineeEmp {

	tunjangan := Tunjangan{
		insurance: 0,
		overtime:  0,
		allowance: 100_000}

	return &TraineeEmp{
		Employee: *NewEmployee(
			fullName,
			salary,
			Contracts),
		Tunjangan: tunjangan,
	}
}

func (trainee *TraineeEmp) GetAllowance() float64 {
	return trainee.allowance
}

func (trainee *TraineeEmp) GetInsurance() float64 {
	return trainee.insurance
}

func (trainee *TraineeEmp) GetOvertime() float64 {
	return trainee.overtime
}

func (trainee *TraineeEmp) GetfullName() string {
	return trainee.fullName
}

func (trainee *TraineeEmp) GetSalary() float64 {
	return trainee.salary
}

func (trainee *TraineeEmp) GetStatus() Status {
	return trainee.status
}

func (trainee *TraineeEmp) ToString() string {
	return fmt.Sprintf("\nfullname : %s\nsalary :%.2f\nstatus :%s\ninsurance :%.2f\novertime :%.2f\nallowance :%.2f\ntotal_salary :%.2f",
		trainee.fullName,
		trainee.salary,
		trainee.status,
		trainee.insurance,
		trainee.overtime,
		trainee.allowance,
		trainee.GetTotalSalary())
}

func (trainee *TraineeEmp) GetTotalSalary() float64 {
	return trainee.salary + trainee.insurance + trainee.overtime + trainee.allowance
}
