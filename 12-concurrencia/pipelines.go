package main

import "fmt"

// Un Canal solamnente de escritura
func Generator(c chan<- int) {
	for i := 0; i <= 10; i++ {
		// Enviamos el valor i al canal c
		c <- i
	}
	// Cuando termina de enviar todos los valores, cerramos el canal
	// Por lo cual no recibe más valores
	close(c)
}

// In Canal de lectura
// Out Canal de escritura
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		// Enviamos el doble de value al canal out
		out <- value * 2
	}
	// Cuando termina de enviar todos los valores, cerramos el canal
	// Por lo cual no recibe más valores
	close(out)
}

func Print(c <-chan int) {
	for value := range c {
		// Imprimimos el valor del canal c
		fmt.Println(value)
	}
}

func mainCop() {
	// Creamos un canal de tipo int
	c := make(chan int)
	// Creamos un canal de tipo int
	d := make(chan int)

	// Creamos una go rutina que envie valores al canal c
	go Generator(c)
	// Creamos una go rutina que reciba valores del canal c y los doble
	go Double(c, d)
	// Creamos una go rutina que reciba valores del canal d y los imprima
	go Print(d)

	// Esperamos a que termine la go rutina
	// Esto es para que no se cierre el programa hasta que termine la go rutina
	// ya que si se cierra el programa antes de terminar la go rutina, el programa
	// se cierra y no se ejecuta la go rutina
	// Esto es porque la go rutina se ejecuta en un hilo diferente al principal
	// y el programa principal se cierra antes de que termine la go rutina
	// Por eso se cierra el programa antes de terminar la go rutina
	<-d
}

/*
	Canales de solo lectura o solo escritura
*/

func main() {
	// Creamos un canal de tipo int
	generator := make(chan int)
	// Creamos un canal de tipo int
	doubles := make(chan int)

	// Creamos una go rutina que envie valores al canal generator
	go Generator(generator)
	// Creamos una go rutina que reciba valores del canal generator y los doubles
	go Double(generator, doubles)

	// Creamos una go rutina que reciba valores del canal doubles y los imprima

	// Esperamos a que termine la go rutina
	// Esto es para que no se cierre el programa hasta que termine la go rutina
	// ya que si se cierra el programa antes de terminar la go rutina, el programa
	// se cierra y no se ejecuta la go rutina
	// Esto es porque la go rutina se ejecuta en un hilo diferente al principal
	// y el programa principal se cierra antes de que termine la go rutina
	// Por eso se cierra el programa antes de terminar la go rutina
	Print(doubles)
}
