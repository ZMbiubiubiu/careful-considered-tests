package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func produce(n int, ch chan int) <-chan int {
	for i := 0; i < n; i++ {
		time.Sleep(time.Duration(rand.Int63n(200)) * time.Millisecond)
		ch <- i
	}
	close(ch)
	return ch
}

type consumer func(ch1, ch2 <-chan int)

func TestNilChannel(t *testing.T) {
	t.Run("inefficiency way", func(t *testing.T) {
		var (
			ch1, ch2 = make(chan int, 2), make(chan int, 2)
		)

		go produce(10, ch1)
		go produce(10, ch2)

		var consumer = func(ch1, ch2 <-chan int) {
			for num := range ch1 {
				fmt.Printf("receive from ch1: %d\n", num)
			}

			for num := range ch2 {
				fmt.Printf("receive from ch2: %d\n", num)
			}
		}

		consumer(ch1, ch2)
	})

	t.Run("waste CPU way", func(t *testing.T) {
		var (
			ch1, ch2 = make(chan int, 2), make(chan int, 2)
		)

		go produce(15, ch1)
		go produce(15, ch2)

		var consumer = func(ch1, ch2 <-chan int) {
			var (
				ch1Closed, ch2Closed bool
			)
			for {
				if ch1Closed && ch2Closed {
					break
				}
				select {
				case num, ok := <-ch1:
					if ok {
						fmt.Printf("receive from ch1: %d\n", num)
					} else {
						fmt.Println("ch1 spin:##########")
						ch1Closed = true
						break
					}
				case num, ok := <-ch2:
					if ok {
						fmt.Printf("receive from ch2: %d\n", num)
					} else {
						fmt.Println("ch2 spin:          **********")
						ch2Closed = true
						break
					}
				}
			}

		}

		consumer(ch1, ch2)
	})

	t.Run("nil channel efficiency way", func(t *testing.T) {
		var (
			ch1, ch2 = make(chan int, 2), make(chan int, 2)
		)

		go produce(15, ch1)
		go produce(15, ch2)

		var consumer = func(ch1, ch2 <-chan int) {
			for {
				if ch1 == nil && ch2 == nil {
					break
				}
				select {
				case num, ok := <-ch1:
					if ok {
						fmt.Printf("receive from ch1: %d\n", num)
					} else {
						ch1 = nil
						break
					}
				case num, ok := <-ch2:
					if ok {
						fmt.Printf("receive from ch2: %d\n", num)
					} else {
						ch2 = nil
						break
					}
				}
			}

		}

		consumer(ch1, ch2)
	})
}
