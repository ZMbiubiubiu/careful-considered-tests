package main

import "fmt"

var (
	workerNum = 3
	jobs      = make(chan int, workerNum)

	done = make(chan struct{}, 3)
)

func work(index int, jobs <-chan int) {
	for n := range jobs {
		fmt.Printf("worker(%d): consume:%d\n", index, n)
	}
	fmt.Printf("worker(%d) done ===================================== \n", index)
	done <- struct{}{}
}

func main() {
	for i := 0; i < workerNum; i++ {
		go work(i, jobs)
	}
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	<-done
	<-done
	<-done
}
