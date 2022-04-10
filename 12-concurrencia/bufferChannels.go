package main

import "fmt"

/*
	Unbuffered channel: Espera una función o una rutina
	para recibir el mensaje, es bloqueada por ser llamada
	en la misma función
*/
func UnbufferedChannel() {
	fmt.Println("============== Ejercicio 1 ==================")
	c := make(chan int)

	c <- 1

	/*
		Este mensaje es bloqueado porque no hay nadie que lea
		tiene que haber una go rutina que lea ese canal, el main
		no lo lee porque no tiene un buffer
	*/
	fmt.Println("Recibiendo el valor del canal", <-c)
}

/*
	Buffered channel: Se puede llamar de manera inmediata, en el siguiente
*/
func main() {
	fmt.Println("============== Ejercicio 2 ==================")
	// Puede recibir hasta 3 valores
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	/*
		Aqui nos va a dar un error porque no hay nadie que lea
		ese ultimo valor.
		Se bloquea porque no hay nadie que lea, ya que este buffer
		tiene 3 valores
	*/
	// c <- 4

	/*
		Aqui si se muestra el valor del canal, ya que el buffer
		tiene un valor en el mismo momento que se llama a la función
	*/
	fmt.Println("Recibiendo el valor del canal", <-c)
	fmt.Println("Recibiendo el valor del canal", <-c)
}
