package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch = make(chan struct{}, 2)
	wg sync.WaitGroup
)

func doSomeWork(i int) {
	fmt.Println("doSomeWork", i)
	time.Sleep(time.Second)

	<-ch
	wg.Done()
}

func main() {
	total := 20
	wg.Add(total)
	for i := 0; i < total; i++ {
		ch <- struct{}{}
		go doSomeWork(i)
	}
	close(ch)
	wg.Wait()
}
