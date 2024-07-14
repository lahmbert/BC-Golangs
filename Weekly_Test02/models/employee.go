package models

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Status string
)

const (
	Permanents Status = "Permanent" //Permanents = iota + 1
	Contracts  Status = "Contract"  //2
	Trainees   Status = "Trainee"   //3
)

type Employee struct {
	id       int
	fullName string
	salary   float64
	status   Status
}

type Tunjangan struct {
	insurance float64
	overtime  float64
	allowance float64
}

// constructor
func NewEmployee(
	fullName string,
	salary float64,
	status Status) *Employee {
	rand.Seed(time.Now().UnixNano())
	ids := rand.Intn(500)

	return &Employee{
		id:       ids,
		fullName: fullName,
		salary:   salary,
		status:   status}
}

type EmployeeInterface interface {
	// SetFullName(string)
	// SetSalary(float64)
	ToString() string
	GetTotalSalary() float64
	GetfullName() string
	GetSalary() float64
	GetStatus() Status
	GetId() int
}

type TunjanganInterface interface {
	GetAllowance() float64
	GetInsurance() float64
	GetOvertime() float64
}

type ChildInterface interface {
	EmployeeInterface
	TunjanganInterface
}

// func (emp *Employee) GetAllowance()float64{
// 	return emp.allowance
// }

// func (emp *Employee) GetInsurance()float64{
// 	return emp.insurance
// }

// func (emp *Employee) GetOvertime()float64{
// 	return emp.overtime
// }

func (emp *Employee) GetId() int {
	return emp.id
}

func (emp *Employee) GetfullName() string {
	return emp.fullName
}

func (emp *Employee) GetSalary() float64 {
	return emp.salary
}

func (emp *Employee) GetStatus() Status {
	return emp.status
}

// copy value
func (emp *Employee) ToString() string {
	return fmt.Sprintf("\nfullName : %s\nsalary :%.2f\nstatus :%s",
		emp.fullName,
		emp.salary,
		emp.status)
}

func (emp *Employee) GetTotalSalary() float64 {
	return emp.salary
}
