package workerThread

import (
	"github.com/sirupsen/logrus"
)

// Payload interface
type Payload interface {
	Do() error
}

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start methods starts the run loop for the worker, listening for a quit channel
// in case we need to stop it.
func (w Worker) Start() {
	// go func() {
	for {
		// register the current worker into the worker queue.
		w.WorkerPool <- w.JobChannel

		select {
		case job := <-w.JobChannel:
			// we hava received a worker request.
			if err := job.Payload.Do(); err != nil {
				logrus.Errorf("Error uploading to S3: %s", err.Error())
			}
		case <-w.quit:
			// we hava received a signal to stop
			// to do
			return
		}
	}
	// }()
}

// Stop signals the worker to stop listening for work requests
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func EnJobQueue(job Job) {
	logrus.Println("test a enqueue")
	JobQueue <- job
	logrus.Println("en queue successfully")
}
