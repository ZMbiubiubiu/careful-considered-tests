// 包含channel常用的使用模式

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// fan out/in semaphore pattern provides a mechanic to control the number of goroutines executing work at any given time
func fanOutSem() {
	// get num of cpu cores
	processNum := runtime.GOMAXPROCS(0)
	// we use sem to ensure that we have max processNum goroutines at the same time
	sem := make(chan struct{}, processNum)

	job := 2_000
	ch := make(chan string, job)
	for i := 0; i < job; i++ {
		go func(n int) {
			sem <- struct{}{}

			t := time.Duration(rand.Intn(200)) * time.Millisecond
			time.Sleep(t)
			ch <- "produce data"
			fmt.Println("child : sent signal :", n)

			<-sem
		}(i)
	}

	for job > 0 {
		d := <-ch
		job--
		fmt.Println(d)
		fmt.Println("parent: recv'd signal :", job)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)
	ch := make(chan string, g)
	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}
	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// retry timeout
func retryTimeout(ctx context.Context, retryInterval time.Duration,
	check func(ctx context.Context) error) {
	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}
		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}
		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)
		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

func channelCancellation(stop <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-stop:
			cancel()
		case <-ctx.Done():
		}
	}()
	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			"https://www.ardanlabs.com/blog/index.xml",
			nil,
		)
		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}(ctx)
}
