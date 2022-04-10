package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("================ Ejercicio 1 - Variables ==================")

	var x int
	x = 9
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("================ Ejercicio 2 - Exepciones ==================")

	value, err := strconv.ParseInt("NaN", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(value)
	}

	fmt.Println("================ Ejercicio 3 - Maps ==================")

	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])

	fmt.Println("================ Ejercicio 4 - Slice ==================")

	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index, value)
	}
	s = append(s, 4)
	fmt.Println(s)

	fmt.Println("================ Ejercicio 5 - Apuntadores ==================")
	g := 25
	h := &g
	fmt.Println("Direccion", h)
	fmt.Println("Valor", *h)

	fmt.Println("================ Ejercicio - Gorutines ==================")

	/*
		Ejemplo 1
		Go routines: Execute a function in a new thread
		Nunca veremos ese mensaje ya que el programa termina y no espera que termine la funcion

		La crea pero no la monitorea

		Cabe mencionar que si se ejecuta la funcion pero no vemos el mensaje ya que se ejecuta en otro plano
	*/
	go doSomething()

	/*
		Ejemplo 2
		Vamos a crear un canal para que el programa espere a que termine la funcion.
		Para ello vamos a crear una variable de tipo canal y le vamos a enviar un valor
		main se va a quedar esperando hasta que el canal reciba un valor.

		Esto se debe a que necesitamos un canal que comunique las dos goroutines
	*/
	c := make(chan int)
	go doSomething2(c)
	// Recibimos un valor del canal, y al recibir un valor del canal, el programa termina
	<-c
}

func doSomething() {
	time.Sleep(time.Second * 3)
	fmt.Println("Do something")
}

func doSomething2(c chan int) {
	time.Sleep(time.Second * 3)
	fmt.Println("Do something 2")
	// Enviamos un valor al canal
	c <- 1
}
