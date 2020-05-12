package workerThread

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// var MaxWorker *int

func TestingMain(m *testing.M) {
	//setup
	// MaxWorker = flag.Int("mw", 1000, "maxworkers")

	code := m.Run() //调用测试用例函数

	//teardown
	os.Exit(code) //注意: os.Exit不会执行defer
}

//implements the payload
type PayloadTest struct {
	Id   int
	Name string
	Age  int
}

func (p *PayloadTest) Do() error {
	if p.Name == "" {
		return errors.New("payload id and name can not empty.")
	}
	fmt.Printf("id - %v - name - %v - age - %v\n", p.Id, p.Name, p.Age)
	return nil
}

//benchmark test
func BenchmarkWorkerThread(b *testing.B) {
	JobQueue = make(chan Job, 1000)
	dispatcher := NewDispatcher(1000)
	dispatcher.Run()

	for i := 0; i < b.N; i++ {
		job := Job{&PayloadTest{i, "testWorker", i}}
		EnJobQueue(job)
	}
}
