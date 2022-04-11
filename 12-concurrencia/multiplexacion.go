package main

import (
	"fmt"
	"time"
)

/*
	Multiplexación con Select y Case
	Cuando una rutina se está comunicando con varios
	channels es muy útil utilizar la palabra reservada select
	para poder interacturar de una manera más ordenada con todos
	los mensajes que estás siendo recibidos.
*/

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	/*
		De esta menera se bloque el programa esperando que terminen
		o se reciba el valor de c1
		fmt.Println(<-c1)

		Pero en este caso, pasa algo curioso
		el c1 tarda más que el c2 pero se imprime primero el c1
		Esto es debido a que se detiene el programa en c1
		fmt.Println(<-c1)

		Si estuviera al contrario
		fmt.Println(<-c2)
		fmt.Println(<-c1)
		Se imprime primero el c2 y luego el c1
	*/
	fmt.Println("Waiting for results")
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	/*
		Pero si quiero saber quien es primero
		independientemente en que orden se escriba
		el código ya sea:

		fmt.Println(<-c2)
		fmt.Println(<-c1)

		ó

		fmt.Println(<-c1)
		fmt.Println(<-c2)

		Puedo usar select para leer los valores
		de los diferenctes canales y saber quien es primero
		o cual fue el activado
	*/

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println("c1", channelMsg1)
		case channelMsg2 := <-c2:
			fmt.Println("c2", channelMsg2)
		}
	}
}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
