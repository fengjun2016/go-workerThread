package workerThread

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
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

//test worker thread
func TestWorkerThead(t *testing.T) {
	JobQueue = make(chan Job, 100)
	dispatcher := NewDispatcher(100)
	dispatcher.Run()

	for i := 0; i < 10; i++ {
		logrus.Println("test : ", i)
		job := Job{&PayloadTest{i, "testWorker", i}}
		logrus.Println("sdasda")
		EnJobQueue(job)
	}
}
