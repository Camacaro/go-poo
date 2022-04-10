package main

import "fmt"

// Values es tratado como un slice
func sum(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (int, int, int) {
	return 2 * x, 3 * x, 4 * x
}

/*
Puedo ponerle nombres a los return y usarlo
solo tengo que poner el return y Go sabrá lo que tiene que devolcer
*/
func getValuesName(x int) (double, triple, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func main() {
	fmt.Println("============== Funciones variadicas  ==================")
	fmt.Println("Suma", sum(1))
	fmt.Println("Suma", sum(1, 2))
	fmt.Println("Suma", sum(1, 2, 3))
	fmt.Println("Suma", sum(1, 2, 3, 4))
	printNames("Mujica")
	printNames("Jesus", "Alejandro", "oriana")

	fmt.Println("============== Retornos con nombre  ==================")
	fmt.Println((getValues(2)))
	fmt.Println((getValues(4)))
}
