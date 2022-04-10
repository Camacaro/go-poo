package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Los channel los usamos para comunicar datos entre procesos (go routines)

	Cuando sabemos que nuestras go routinas no van a compartir datos, podemos usar las waitgroups
*/
func main() {
	fmt.Println("============== Ejercicio 1  ==================")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		/*
			De esta manera main lanza las go routines pero no las va a monitorear

			Ahora le agregamos las waitgroup para que las go routines sean monitoreadas
			lo que hacemos es agregar un contador de waitgroup y se detendra el programa
			hasta que sea 0

			Cabe destacar que estas go routines se ejecutan en paralelo(concurentes)
		*/
		wg.Add(1)
		go doSomething(i, &wg)
	}

	/*
		Bloquea el Main hasta que terminen todas las go routines
		hasta que waitgroup sea 0
	*/
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	/*
		deffer es una funcion que se ejecuta al finalizar la funcion
		al finalizar termina el waitgroup, decrementa el contador del waitgroup
	*/
	defer wg.Done()
	fmt.Println("Start", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End", i)
}
