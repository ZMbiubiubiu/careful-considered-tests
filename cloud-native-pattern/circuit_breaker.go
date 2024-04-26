package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// what is a closure
// A closure is a nested function that has access to the variables
// of its parent function, even after the parent has executed

// circuit-breaker
// 当一个功能连续失败到达阈值，进入熔断状态，直接返回错误
// 何时可以重试？达到时间要求，采用指数退避的形式

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0 // 连续失败的次数
	var lastAttempt = time.Now()
	var mu sync.RWMutex // 保护consecutiveFailures

	return func(ctx context.Context) (string, error) {
		mu.RLock()

		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d) // 指数回退的重试策略
			if time.Now().Before(shouldRetryAt) {
				mu.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		mu.RUnlock()

		response, err := circuit(ctx)
		mu.Lock()
		defer mu.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0
		return response, nil
	}
}

func circuitFunc(ctx context.Context) (string, error) {
	return "", errors.New("oh no")
}

func main() {
	threshold := 5
	breaker := Breaker(circuitFunc, uint(threshold))
	for i := 0; i < 7; i++ {
		_, err := breaker(context.Background())
		if i < threshold {
			fmt.Println(err.Error() == "oh no")
		} else {
			fmt.Println(err.Error() == "service unreachable")
		}
	}

}
