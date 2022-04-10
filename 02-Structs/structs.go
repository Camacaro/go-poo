package main

import "fmt"

/*
De Esta manera de hacen las clases
*/

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {
	employee := Employee{}
	fmt.Println("employee vacio", employee)

	employee.id = 1
	employee.name = "Name"
	fmt.Println("employee con valores", employee)

	employee.setId(2)
	employee.setName("Segundo nombre")
	fmt.Println("employee con funciones", employee)

	fmt.Println("Get functiones", employee.getName(), employee.getId())
}
