package main

import (
	"encoding/json"
	"fmt"
)

func unmarshal() {

	nums := []int{1, 2}
	bs, _ := json.Marshal(nums)

	var (
		// 指定长度
		numsWithBigLen   = make([]int, 4)
		numsWithSmallLen = make([]int, 1)

		// 只指定容量
		numsWithBigCap   = make([]int, 0, 4)
		numsWithSmallCap = make([]int, 0, 1)
	)

	fmt.Printf("numsWithBigLen:  \toriginal address:%p\n", numsWithBigLen)
	fmt.Printf("numsWithSmallLen:\toriginal address:%p\n", numsWithSmallLen)
	fmt.Printf("numsWithBigCap:  \toriginal address:%p\n", numsWithBigCap)
	fmt.Printf("numsWithSmallCap:\toriginal address:%p\n", numsWithSmallCap)

	json.Unmarshal(bs, &numsWithBigLen)
	json.Unmarshal(bs, &numsWithSmallLen)
	json.Unmarshal(bs, &numsWithBigCap)
	json.Unmarshal(bs, &numsWithSmallCap)

	fmt.Printf("numsWithBigLen:  \tnow address(%p) len(%d) cap(%d)\n", numsWithBigLen, len(numsWithBigLen), cap(numsWithBigLen))
	fmt.Printf("numsWithSmallLen:\tnow address(%p) len(%d) cap(%d)\n", numsWithSmallLen, len(numsWithSmallLen), cap(numsWithSmallLen))
	fmt.Printf("numsWithBigCap:  \tnow address(%p) len(%d) cap(%d)\n", numsWithBigCap, len(numsWithBigCap), cap(numsWithBigCap))
	fmt.Printf("numsWithSmallCap:\tnow address(%p) len(%d) cap(%d)\n", numsWithSmallCap, len(numsWithSmallCap), cap(numsWithSmallCap))
}
