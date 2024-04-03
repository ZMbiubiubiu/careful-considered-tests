package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type MyOnce struct {
	mu   sync.Mutex
	done uint32
}

func (o *MyOnce) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.mu.Lock()
		defer o.mu.Unlock()
		if o.done == 0 {
			defer atomic.StoreUint32(&o.done, 1)
			f()
		}
	}
}

func demo() {
	fmt.Println("print....")
}

func main() {
	var myOnce MyOnce
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			myOnce.Do(demo)
		}()
	}
	wg.Wait()
}
