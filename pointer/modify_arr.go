package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [3]int{1, 2, 3}
	arrPtr := &arr

	arr0ElePtr := (*int)(unsafe.Pointer(arrPtr))
	arr1ElePtr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(arrPtr)) + unsafe.Sizeof(arr[0])))

	*arr0ElePtr = 10
	*arr1ElePtr = 20

	fmt.Println(arr)
}
