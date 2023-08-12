package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data interface{}
}

var (
	taskQueue []Task
	mutex     sync.Mutex
	taskID    int
)

func scheduler(incoming <-chan Task, execChannel chan<- Task) {
	for task := range incoming {
		mutex.Lock()
		taskQueue = append(taskQueue, task)
		mutex.Unlock()
		execChannel <- task
	}
}

func executor(execChannel <-chan Task) {
	for task := range execChannel {
		fmt.Printf("Processing Task %d with data %s\n", task.ID, task.Data)
		time.Sleep(time.Second)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	task := Task{ID: taskID, Data: data}
	taskID++

	fmt.Fprintf(w, "Task enqueued with ID %d\n", task.ID)
	go scheduleTask(task)
}

func main() {
	http.HandleFunc("/enqueue", taskHandler)
	fmt.Println("Serving requests on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Unable to serve")
		os.Exit(-1)
	}
}

func scheduleTask(task Task) {
	incoming := make(chan Task)
	execChannel := make(chan Task)
	go scheduler(incoming, execChannel)
	go executor(execChannel)

	incoming <- task
}
