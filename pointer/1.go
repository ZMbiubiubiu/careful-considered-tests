package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 100
	intI := &i
	var floatI *float64
	floatI = (*float64)(unsafe.Pointer(intI))
	*floatI = *floatI * 3
	fmt.Printf("%T\n", i) // int
	// 从侧面证明了 *float64 的指针变量是指向 i 变量的内存地址的。
	fmt.Println(i)             // 300
	fmt.Printf("%T\n", intI)   // *int
	fmt.Printf("%T\n", floatI) // *float64
}
