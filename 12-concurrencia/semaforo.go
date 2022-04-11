package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("============== Ejercicio 1  ==================")

	/*
		Cuando no se pone el buffer, el canal es un canal de tipo unbuffered
		y se bloquea la función que lo llama porque no hay nadie que lea
		el valor del canal

		Cuando se pone el buffer, el canal es un canal de tipo buffered
		y se habilitan la cantidad de canales para que en cada canal se
		ejecute una go rutina y al finalizar la go rutina se pueda liberar
		el canal y se asigne otra go rutina.

		Buffered Channels como semáforos
		Estos channels puedes ser utilizados para que actuén como semáforos
		y permitan una cantidad limitada de ejecuciones para perseguir
		objetivos como evitar la ejecución de una cantidad indeterminada de
		rutinas.

		Esto nos sirve para cuando se tiene una iteracion muy larga y se
		necesita que se ejecuten varias go rutinas por bloques y no todas
		al mismo tiempo.

		Ejemplo:
		c : = [][]
		1. Primera iteracion:
		c := [go rutina 1][]
		2. Segunda iteracion:
		c := [go rutina 2][go rutina 1]
		3. Tercera iteracion:
		Se librea un canal y se asigna otra go rutina, pero
		mientras se libera alguno de los canales, se bloquea la función
		c := [go rutina 3][go rutina 2]
	*/
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// Nuestro canal envia valores de 1
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Id started", i)
	time.Sleep(4 * time.Second)
	fmt.Println("End", i)
	// Recibe valores por el canal
	// libera el canal, saca el valor del canal y lo elimina
	<-c
}
