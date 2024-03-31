package main

import (
	"fmt"
	"time"
)

// 延迟调用
func delayCall(f func(), d time.Duration) {
	timer := time.NewTimer(d)
	select {
	case <-timer.C:
		f()
	}
}

// 设置超时时间
func waitChannel(wait time.Duration, ch <-chan string) bool {
	timer := time.NewTimer(wait)
	select {
	case result := <-ch:
		timer.Stop()
		fmt.Println(result)
		return true
	case <-timer.C:
		fmt.Println("wait channel timeout!")
		return false
	}
}

func main() {

}
