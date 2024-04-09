package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	f1, f2 := 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func main() {
	file, _ := os.Create("fib_cpu")
	defer file.Close()

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
}
