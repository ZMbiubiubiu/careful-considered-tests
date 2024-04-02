package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch  = make(chan struct{}, 2)
	wg1 sync.WaitGroup
)

func doSomeWork(i int) {
	fmt.Println("doSomeWork", i)
	time.Sleep(time.Second)

	<-ch
	wg1.Done()
}

func main() {
	total := 20
	wg1.Add(total)
	for i := 0; i < total; i++ {
		ch <- struct{}{}
		go doSomeWork(i)
	}
	close(ch)
	wg1.Wait()
}

// runBoundedTask 起maxTaskNum个协程共同处理任务
// 实现简单，但是有可能数据太多，导致处理受到阻塞；数据太小，协程空闲
func runBoundedTask(dataChan <-chan int, maxTaskNum int) {
	var wg sync.WaitGroup
	wg.Add(maxTaskNum)

	for i := 0; i < maxTaskNum; i++ {
		go func() {
			defer wg.Done()

			for data := range dataChan {
				func(data int) {

					// do something
					time.Sleep(3 * time.Second)
				}(data)
			}
		}()
	}

	wg.Wait()
}
