package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("============== Func en variable  ==================")

	x := 5

	// Auto llamado
	y := func() int {
		return x * 2
	}()

	fmt.Println("Resultado", y)

	fmt.Println("============== Func Anonima  ==================")
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	// Bloqueo el programa a la espera de ese valor
	<-c
}
