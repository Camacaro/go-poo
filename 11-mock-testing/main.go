package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM person WHERE dni = dni
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM employee WHERE id = id
	return Employee{}, nil
}

func GetFullTimeEmployee(id int, dni string) (FullTimeEmployee, error) {

	var ftEmployee FullTimeEmployee

	employee, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = employee

	person, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = person

	return ftEmployee, nil
}
