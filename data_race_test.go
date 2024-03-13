package main

import (
	"fmt"
	"testing"
	"time"
)

func foo() int {
	var a int

	go func() {
		a = 2
	}()

	go func() {
		a = 3
	}()

	time.Sleep(time.Second)
	return a
}

// race condition检查：go test -race data_race_test.go
func TestFoo(t *testing.T) {
	a := foo()
	fmt.Println(a)
}
