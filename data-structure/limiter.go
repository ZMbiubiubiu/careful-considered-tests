package main

import (
	"fmt"
	"log"
	"time"
)

type Limiter struct {
	cap    int
	ticker time.Duration
	buf    chan struct{}
	done   chan struct{}
}

func NewLimiter(cap int, ticker time.Duration) *Limiter {
	limiter := &Limiter{
		cap:    cap,                      // peek qps
		ticker: ticker,                   // 每隔ticker段时间，往令牌桶里放
		buf:    make(chan struct{}, cap), // 实际存储令牌桶
		done:   make(chan struct{}),      // 停止后台协程
	}

	go func() {
		// 无限制往里面放令牌桶
		for {
			select {
			case <-time.After(ticker):
				select {
				case limiter.buf <- struct{}{}:
				default: // 令牌桶满了，放不下
				}
			case <-limiter.done:
				log.Println("limiter background task end.")
				return
			}
		}
	}()

	return limiter
}

func (l *Limiter) GetToken(waitTime time.Duration) (got bool) {
	select {
	case <-l.buf:
		got = true
	case <-time.After(waitTime):
		got = false
	}
	return got
}

func (l *Limiter) Stop() {
	close(l.done)
}

func main() {
	limiter := NewLimiter(4, time.Second)
	time.Sleep(4 * time.Second)

	for i := 0; i < 10; i++ {
		got := limiter.GetToken(100 * time.Millisecond)
		fmt.Println("got", i, got)
	}
}
