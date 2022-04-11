package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
	Golang permite utilizar conceptos como los de worker pools
	y, en combinación con buffered channels, la creación de job queues
	que utilizando concurrencia nos permitirán tener un alto rendimiento
	a la hora de la creación de MUCHAS TAREAS
*/

/*
	Job represents a job to be executed, with a name
	and a number and a delay
*/
type Job struct {
	Name   string
	Delay  time.Duration // delay between each job
	Number int           // number to calculate on the fibonacci sequence
}

/*
	Worker will be our concurrency-friendly worker
*/
type Worker struct {
	Id         int
	JobQueue   chan Job      // Para recibir un channel of type Job - Jobs to be processed
	WorkerPool chan chan Job // Pool of workers - Un canal de canales de tipo Job
	QuitChan   chan bool
}

/*
	Dispatcher is a dispatcher that will dispatch jobs to workers
*/
type Dispatcher struct {
	WorkerPool chan chan Job // Pool of workers
	MaxWorkers int           // Maximum number of workers
	JobQueue   chan Job      // Jobs to be processed
}

/*
	NewWorker returns a new Worker with the provided id and workerpool
*/
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job), // create a job queue for this worker
		WorkerPool: workerPool,
		QuitChan:   make(chan bool), // Channel to end jobs
	}
}

/*
	Start method starts all workers
*/
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue // add job to pool
			// Multiplexing
			select {
			case job := <-w.JobQueue: // get job from queue
				fmt.Printf("worker %d started job %s\n", w.Id, job.Name)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("worker %d finished job %s with results %d\n", w.Id, job.Name, fib)
			case <-w.QuitChan: // quit if worker is told to do so
				fmt.Printf("worker %d stopping\n", w.Id)
			}
		}
	}()
}

/*
 Stop method stop the worker
*/
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

/*
	Fibonacci calculates the fibonacci sequence
*/
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

/*
	NewDispatcher returns a new Dispatcher with the provided maxWorkers
*/
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		WorkerPool: worker,
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}

/*
	Dispatch will dispatch jobs to workers
*/
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue: // get job from queue
			// Asign the job to a worker
			go func() {
				workerJobQueue := <-d.WorkerPool // get worker from pool
				workerJobQueue <- job            // Workers will read from this channel
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

/*
	Creando Web server para procesar jobs
	Con la libreria estandar de go y utilizando el paquete net,
	somos capaces de crear un servidor que será el que atienda las
	peticiones y asignará los nuevos workers para que lleven a cabo
	los trabajos que se está buscando conseguir.

	Con esto puedo mandar multiples request y procesar varios jobs
*/

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	// No se aceptan peticiones que no sean POST
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value", http.StatusBadRequest)
		return
	}

	// Atoi de string a number
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Bad value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	// En colamos los Job
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":4000"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:4000/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
