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
	t.Logf("%T slice address: %p\t arrary address: %p\n", nums, &nums, nums)
	assert.Equal(t, fmt.Sprintf("%p", nums), fmt.Sprintf("%p", &nums[0]))
}

func TestNilAndEmptySlice(t *testing.T) {
	var s1 []string         // nil
	s2 := []string{}        // empty
	s3 := make([]string, 0) // empty, equivalent to s2 := []string{}

	t.Logf("s1 is nil: %t, len: %d, cap: %d, backend array addr: %p\n", s1 == nil, len(s1), cap(s1), s1)
	t.Logf("s2 is nil: %t, len: %d, cap: %d, backend array addr: %p\n", s2 == nil, len(s2), cap(s2), s2)
	t.Logf("s3 is nil: %t, len: %d, cap: %d, backend array addr: %p\n", s3 == nil, len(s3), cap(s3), s3)
}
