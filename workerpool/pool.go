package workerpool

import (
	"sync"
)

func Worker(wg *sync.WaitGroup, tasks chan *Task) {
	for task := range tasks {
		// Execute a task if its not processed or cancelled
		if task.state == Pending {
			println("Executing Task with Id:", task.Id)
			task.state = Running
			res, err := task.Execute()
			if err == nil {
				task.result = res
				task.state = Finished
			} else {
				task.Error = err
				task.state = FailedWithError
			}
		} else {
			println("Task Was already processed: ", task.Id, " : ", task.state)
		}
	}
	wg.Done()
}

func AddTasksToChannel(ch chan *Task, tasks []*Task) {
	println("Adding tasks to channel")
	for _, task := range tasks {
		ch <- task
	}
	close(ch)
}

func ExecuteTasks(noOfWorkers int, tasks []*Task, done chan bool) {
	//creating channel with buffer size of workers
	var ch = make(chan *Task, noOfWorkers)
	go AddTasksToChannel(ch, tasks)
	//before starting worker group adding tasks to channel
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		println("STARTING GOROUTINE: ", i)
		wg.Add(1)
		go Worker(&wg, ch)
	}
	wg.Wait()
	done <- true
}
