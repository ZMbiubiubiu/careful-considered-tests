package channel

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func pooling() {
	ch := make(chan string)
	var wg sync.WaitGroup
	g := runtime.GOMAXPROCS(0)
	wg.Add(g)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, d)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
			defer wg.Done()
		}(c)
	}
	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("-------------------------------------------------")
}

func TestChannelType(t *testing.T) {
	t.Run("buffer or not buffer, that's a question", func(t *testing.T) {
		pooling()
	})
}
