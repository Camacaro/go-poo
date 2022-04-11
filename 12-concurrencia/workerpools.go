package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker with id %d started job %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("worker with id %d finished job %d, result is %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
	Worker Pools
	Permiten la creación de múltiples trabajadores que llevarán
	a cabo determinadas tareas, en este caso Go puede explotar
	las GoRutines para alcanzar worker pools concurrentes.

	En este ejemplo, se crea un worker pool de 3 workers,
	que se encargarán de calcular el fibonacci de cada uno
	de los números en el slice tasks.

	Primero hacemos un for para los workers con esto se crea
	en memoria tres workers, que se encargarán de calcular
	el fibonacci de cada uno de los números en el slice tasks.
	Cuando se empiecen a mandar valores al canal jobs,
	Se empiezan a ocupar los worker pool que creamos en memoria.
	Cuando se complete uno queda liberado para recibir el siguiente
	valor del canal jobs.
	De esta menera se calculan 3 fibonacci al mismo tiempo.
*/

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		fmt.Println("=======", i)
		go worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for j := 0; j < len(tasks); j++ {
		// fmt.Println(<-results)
		<-results
	}
}
