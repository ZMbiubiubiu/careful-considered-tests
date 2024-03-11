package main

import (
	"fmt"
	"runtime"

	"testing"

	"github.com/stretchr/testify/assert"
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
			if i == 2 {
				assert.Equal(t, 30, arr[2])
			}
		}
		assert.Equal(t, 30, arr[2])
	})

	t.Run("value semantics on array", func(t *testing.T) {
		var arr = [5]int{1, 2, 3, 4, 5}
		// 会copy一份arr
		for i, num := range arr {
			if i == 0 {
				arr[2] = 30
			}
			if i == 2 {
				assert.Equal(t, 3, num)
			}
		}
		assert.Equal(t, 30, arr[2])
	})
}

func TestValueAndPointerForRangeMemory(t *testing.T) {
	t.Run("value semantic", func(t *testing.T) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("mallocs: %+v\n", m)
		var arr = [1_0000]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		runtime.ReadMemStats(&m)
		fmt.Printf("mallocs: %+v\n", m)
		for _, v := range arr {
			runtime.ReadMemStats(&m)
			fmt.Printf("mallocs: %+v\n", m)
			_ = v
			break
		}
		fmt.Printf("mallocs: %+v\n", m)
	})
}

func BenchmarkValueAndPointerForRange(b *testing.B) {
	var arr = [10]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var tmp int
	b.Run("value semantic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i, v := range arr {
				arr[i] = v
			}
		}
	})

	b.Run("pointer semantic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := range arr {
				arr[i] = arr[i]
			}
		}
	})

	_ = tmp
}

func TestForRangeOnSliceAndMap(t *testing.T) {
	type Foo struct {
		a int
		b string
	}

	foos := []*Foo{
		&Foo{1, "a"},
		&Foo{2, "b"},
		&Foo{3, "c"},
	}

	m := make(map[int]*Foo)
	for _, foo := range foos {
		m[foo.a] = foo
	}

	for k, v := range m {
		assert.Equal(t, k, v.a)
	}

	t.Run("cus", func(t *testing.T) {
		type Custom struct {
			Id   uint32 `json:"id"`
			Name string `json:"name"`
		}
		customs := []*Custom{
			&Custom{1, "one"},
			&Custom{2, "two"},
			&Custom{3, "three"},
		}
		var m = make(map[int]*Custom, len(customs))

		for i, cm := range customs {
			m[i] = cm
			fmt.Printf("%p\n", cm)
		}
		for k, v := range m {
			fmt.Printf("key:%d, val:%+v\n", k, v)
		}
	})
}
