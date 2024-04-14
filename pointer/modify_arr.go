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
func letterCombinations(digits string) []string {
	var keyboards = map[byte][]byte{'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'}}

	var result = make([]string, 0, 4^len(digits))

	if len(digits) == 0 {
		return result
	}

	var subset = make([]byte, 0, len(digits))
	traverse(keyboards, []byte(digits), 0, subset, &result)
	return result
}

func traverse(keyboards map[byte][]byte, digits []byte, i int, subset []byte, res *[]string) {
	if len(subset) == len(digits) {
		var temp = make([]byte, len(digits))
		copy(temp, subset)
		*res = append(*res, string(temp))
	} else if i < len(digits) {
		digit := digits[i]
		for _, b := range keyboards[digit] {
			subset = append(subset, b)
			traverse(keyboards, digits, i+1, subset, res)
			subset = subset[:len(subset)-1]
		}
	}
}
