package main

import (
	"fmt"
	"time"
)

func main() {
	var a int

	go func() {
		a = 2
	}()

	go func() {
		a = 3
	}()

	time.Sleep(time.Second)
	fmt.Println("final a: ", a)
}
