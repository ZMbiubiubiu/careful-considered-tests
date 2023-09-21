package main

import (
	"fmt"
	"testing"
)

/*
When using value semantic iteration, two things happen.
First, the collection I’m iterating over is copied and I iterate over the copy.
In the case of an array, the copy could be expensive since the entire array is copied.
In the case of a slice, there is no real cost since only the internal slice value is copied and not the backing array.
Second, I get a copy of each element being iterated on.

When using pointer semantic iteration, I iterate over the original collection and I access each element associated
with the collection directly.
*/

func TestForRangeOnArray(t *testing.T) {
	t.Run("pointer semantics on array", func(t *testing.T) {
		var arr = [5]int{1, 2, 3, 4, 5}
		for i := range arr {
			if i == 0 {
				arr[2] = 30
			}
			fmt.Println(i, arr[i])
		}
	})

	t.Run("value semantics on array", func(t *testing.T) {
		var arr = [5]int{1, 2, 3, 4, 5}
		for i, num := range arr {
			if i == 0 {
				arr[2] = 30
			}
			fmt.Println(i, num)
		}
	})
}
