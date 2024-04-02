package main

import (
	"fmt"
	"sync"
)

var (
	notifyChan = make(chan struct{})
	wg         sync.WaitGroup
)

func printA() {
	for i := 1; i < 100; i += 2 {
		fmt.Println(i)
		notifyChan <- struct{}{}
		<-notifyChan
	}
	wg.Done()
}

func printB() {
	for i := 2; i <= 100; i += 2 {
		<-notifyChan
		fmt.Println(i)
		notifyChan <- struct{}{}
	}
	wg.Done()
}

func main() {
	// 实时打印goroutine的数量
	//go runNumGoroutineMonitor()
	wg.Add(2)
	go printA()
	go printB()
	wg.Wait()
}
