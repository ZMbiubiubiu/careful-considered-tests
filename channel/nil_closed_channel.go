package main

/*
op    |  nil channel       |  closed channel                  |
read  |  block or return   |  channel elem or elem zero value |
write |  block or return   |  panic                           |
close |  panic        	   |  panic                           |
*/

func main() {
	var ch chan struct{}
	ch <- struct{}{} // fatal error: all goroutines are asleep - deadlock!
}
