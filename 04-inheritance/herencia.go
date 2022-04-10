package main

import "fmt"

/**
Composicion sobre herencia

Go no permite la herencia, go utiliza la composicion.
la composicion, a diferencia de la herencia, no es una clase hija de… sino que contiene los metodos de las clases indicadas.

En Go no se puede aplicar el polimorfismo como en el ejemplo de GetMessages ya que este recibe es una persona
y no un FullTimeEmployee (Este tiene una propiedad de persona) pero no es un tipo de persona
*/

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// Aplicamos la composicion sobre la herencia
type FullTimeEmployee struct {
	// person Person -> usar la forma de abajo
	Person   // De esta forma crear un campo anonimo
	Employee // Es una desestructuracion de los campos de ambos struct
}

func GetMessages(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	fmt.Println("============== Forma  ==================")
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 25
	ftEmployee.id = 1
	ftEmployee.name = "Jesus Camacaro"
	// Al imprimir me secepciona o agrupa los valores de Person y Employee
	fmt.Println("Full time employee", ftEmployee)

	// Error ya que ftEmployee no es de tipo persona
	// GetMessages(ftEmployee)

	// O usar persona directamente de ftEmployee
	GetMessages(ftEmployee.Person)
}
