package main

func main() {
	var ch chan struct{}
	ch <- struct{}{} // fatal error: all goroutines are asleep - deadlock!
}
