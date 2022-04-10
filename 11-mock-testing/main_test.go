package main

import (
	"reflect"
	"testing"
)

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id       int
		dni      string
		mockFunc func()
		expected FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "Developer",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "12345678",
						Name: "John Doe",
						Age:  30,
					}, nil
				}
			},
			expected: FullTimeEmployee{
				Person: Person{
					DNI:  "12345678",
					Name: "John Doe",
					Age:  30,
				},
				Employee: Employee{
					Id:       1,
					Position: "Developer",
				},
			},
		},
	}

	// Funcciones originales del archivo main.go
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, row := range table {
		// Reemplazar las funciones con las funciones mocks
		row.mockFunc()

		ftEmployee, err := GetFullTimeEmployee(row.id, row.dni)

		if err != nil {
			t.Errorf("Error getting full time employee: %s", err)
		}

		if !reflect.DeepEqual(ftEmployee.Employee, row.expected.Employee) {
			t.Errorf("Expected employee %v, got %v", row.expected.Employee, ftEmployee.Employee)
		}

		if ftEmployee.Age != row.expected.Age {
			t.Errorf("Expected age %d, got %d", row.expected.Age, ftEmployee.Age)
		}

		if ftEmployee.Name != row.expected.Name {
			t.Errorf("Expected name %s, got %s", row.expected.Name, ftEmployee.Name)
		}

		// devolver las funciones originales a sus valores originales
		// Esto es para evitar que se afecten a otros tests y no usen esttos mocks
		GetEmployeeById = originalGetEmployeeById
		GetPersonByDNI = originalGetPersonByDNI
	}
}
