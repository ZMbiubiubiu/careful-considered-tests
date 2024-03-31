package main

import "fmt"

var (
	workerNum = 3
	jobs      = make(chan *task, workerNum)

	done = make(chan struct{}, 3)
)

type task struct {
	i int
}

func work(index int, jobs <-chan *task) {
	for task := range jobs {
		fmt.Printf("worker(%d): consume:%d\n", index, task.i)
	}
	fmt.Printf("worker(%d) done ===================================== \n", index)
	done <- struct{}{}
}

func main() {
	for i := 0; i < workerNum; i++ {
		go work(i, jobs)
	}
	for i := 0; i < 50; i++ {
		jobs <- &task{i: i}
	}
	close(jobs)
	<-done
	<-done
	<-done
}
