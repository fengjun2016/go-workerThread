package workerThread

import "github.com/sirupsen/logrus"

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		logrus.Println("new a worker", i)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has received
			go func(job Job) {
				// try to obtain a worker job channel that is available
				// this will block unitl a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker jon channel
				jobChannel <- job
			}(job)
		}
	}
}
