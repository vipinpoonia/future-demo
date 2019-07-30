package future_demo

import (
	"fmt"
	"future-demo/workerpool"
	"math/rand"
	"time"
)

// uses Examples
func addStrings(args workerpool.Args) (interface{}, error) {
	w, _ := args["str1"]
	x, _ := args["str2"]
	y, _ := args["str3"]
	z := w.(string) + x.(string) + y.(string)
	time.Sleep(1 * time.Second)
	return z, nil
}

func add(args workerpool.Args) (interface{}, error) {
	x, _ := args["num1"]
	y, _ := args["num2"]
	z := x.(int) + y.(int)
	//fmt.Println("Executing Task.....")
	time.Sleep(1 * time.Second)
	//fmt.Println("Execution Finished..........")
	return z, nil
}

//Create multiple tasks to process
func createAddTask(noOfTasks int) [] *workerpool.Task {
	tasks := make([]*workerpool.Task, noOfTasks)
	for i := 0; i < noOfTasks; i++ {
		randomNo := rand.Intn(999)
		args := workerpool.Args{
			"num1": randomNo,
			"num2": 30,
		}
		t := workerpool.NewTask(args, add)
		//for testing purpose only
		t.Id = i
		tasks[i] = t
	}
	return tasks
}

func main() {
	fmt.Println("Starting Worker Pool")
	startTime := time.Now()
	noOfTasks := 80
	tasks := createAddTask(noOfTasks)
	done := make(chan bool)
	go workerpool.ExecuteTasks(10, tasks, done)
	for i, task := range tasks {
		if i%10==0{
			isCancelled := task.Cancel()
			println("task is cancelled", task.Id, " : ", isCancelled)
		}
	}

	<-done
	for _, task := range tasks {
		res, err := task.Result()
		if err == nil {
			println("testing func", res.(int))
		} else {
			println("testing func", err.Error())
		}
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
