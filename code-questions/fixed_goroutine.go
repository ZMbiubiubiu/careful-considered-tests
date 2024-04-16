package main

import (
	"sync"
	"time"
)

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
