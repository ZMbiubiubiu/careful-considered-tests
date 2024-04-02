package main

import (
	"sync"
	"time"
)

/*
针对固定个数协程的缺点，另一个思路是借鉴限流器的实现，控制每个时刻最大允许协程数量也达到控制协程数量的目的。这里也提供两种实现思路
*/

func runRateLimitTask(dataChan <-chan int, ticker time.Duration) {
	// rate limiter
	var limiter = make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(ticker):
				limiter <- struct{}{}
			}
		}
	}()

	var wg sync.WaitGroup

	for data := range dataChan {
		<-limiter

		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = data
			time.Sleep(time.Second)
		}()
	}
}
