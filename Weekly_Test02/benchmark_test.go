package main

import "testing"

func BenchmarkAllTwoGenerator(b *testing.B) {
	b.Run("BenchmarkEmployeeTotalSalaries", BenchmarkEmployeeTotalSalaries)
	b.Run("BenchmarkEmployeeTotalSalariesWithChannel", BenchmarkEmployeeTotalSalariesWithChannel)
}

func BenchGenerateEmployees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateEmployees(200)
	}
}

func BenchGenerateEmployeesWithChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateEmployeesWithChannel(200)
	}
}

func BenchmarkEmployeeTotalSalariesWithChannel(b *testing.B) {
	employees := GenerateEmployeesWithChannel(100)

	for i := 0; i < b.N; i++ {
		EmployeeTotalSalariesWithChannel(employees)
	}
}

func BenchmarkEmployeeTotalSalaries(b *testing.B) {
	employees := GenerateEmployeesWithChannel(100)

	for i := 0; i < b.N; i++ {
		EmployeeTotalSalaries(employees)
	}
}
