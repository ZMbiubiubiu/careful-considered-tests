package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReuseSlice(t *testing.T) {
	var s = []int{1, 2, 3}
	s1 := s[:0]
	assert.Equal(t, 0, len(s1))
	assert.Equal(t, cap(s), cap(s1))
}

func TestSliceArrayAddress(t *testing.T) {
	var nums = []int{1, 2, 3}
	fmt.Printf("%T slice address: %p\t arrary address: %p\n", nums, &nums, nums)
	assert.Equal(t, fmt.Sprintf("%p", nums), fmt.Sprintf("%p", &nums[0]))
}
