package main

import (
	"fmt"
	"unsafe"
)

type demo struct {
	A int8
	B int64
	C byte
}

func main() {
	d := demo{}
	// it returns the number of bytes between
	//the start of the struct and the start of the field.
	fmt.Println(unsafe.Offsetof(d.A)) // 0
	fmt.Println(unsafe.Offsetof(d.B)) // 8
	fmt.Println(unsafe.Offsetof(d.C)) // 16
}
