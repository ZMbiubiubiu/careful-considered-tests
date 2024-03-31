package main

import (
	"reflect"
	"unsafe"
)

// 核心思路：共享底层的bytes数组

/*
// string 的运行时表示
type StringHeader struct {
	Data uintptr
	Len  int
}

// slice 的运行时表示
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
*/

func string2bytes(s string) []byte {
	// 先将string转换成运行时的结构
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
