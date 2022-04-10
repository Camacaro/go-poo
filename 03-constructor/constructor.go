package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	fmt.Println("============== Forma 1 - Constructor vacio  ==================")
	employee := Employee{}
	fmt.Println("Employee", employee)

	fmt.Println("============== Forma 2 - Inicializar values  ==================")
	employee2 := Employee{
		id:       1,
		name:     "Jesus",
		vacation: true,
	}
	fmt.Println("2. Employee", employee2)

	fmt.Println("============== Forma 3 - Usando new  ==================")
	// De esa forma me devolvera la referencia del Employee y no sus valores como tal tengo que usar el *
	employee3 := new(Employee)
	fmt.Println("3. Employee - referencia", employee3)
	fmt.Println("3. Employee - valores", *employee3)

	employee3.id = 3
	employee3.name = "Tercer nombre"
	fmt.Println("3. Employee - valores", *employee3)

	fmt.Println("============== Forma 4 - Funcion independiente con apuntadores  ==================")
	// Recuerda que te devolvera la referencia, hay que usar el * para obtener los valores
	employee4 := NewEmployee(4, "Cuarto nombre", true)
	fmt.Println("4. Employee - referencia &", employee4)
	fmt.Println("4. Employee - valores *", *employee4)
}
