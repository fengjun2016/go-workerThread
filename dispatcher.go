package workerThread

type Dispatched struct {
	// A pool of workers channels that are registered with the dispatched
	WorkerPool chan chan Job
}

func NewDispatched(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatched) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.pool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatched) dispatch() {
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
