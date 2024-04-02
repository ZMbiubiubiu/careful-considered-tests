package main

import (
	"log"
	"runtime"
	"time"
)

// runNumGoroutineMonitor 协程数量监控
func runNumGoroutineMonitor() {
	log.Printf("协程数量->%d\n", runtime.NumGoroutine())

	for {
		select {
		case <-time.After(time.Second):
			log.Printf("协程数量->%d\n", runtime.NumGoroutine())
		}
	}
}
