package main

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

var (
	err  error
	pool *ants.Pool
)

func init() {
	pool, err = ants.NewPool(10)
	if err != nil {
		panic("init pool failed")
	}
}

type Task struct {
	ID int
}

func printTask(task *Task) {
	fmt.Println(task.ID)
}

func main() {
	var tasks = make([]*Task, 0, 1000)
	for i := 0; i < 1000; i++ {
		tasks = append(tasks, &Task{i})
	}

	for i := 0; i < 1000; i++ {
		_ = pool.Submit(func() {
			printTask(tasks[i])
		})
	}

}
