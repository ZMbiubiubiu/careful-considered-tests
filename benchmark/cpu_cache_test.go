package benchmark

import (
	"math/rand"
	"testing"
)

var (
	matrix1, matrix2, matrix3 [][]int
	list1                     *Node
)

const (
	rows = 1_000
	cols = 1_000
)

func constructMatrix() [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i + j
		}
	}
	return matrix
}

type Node struct {
	val  int
	next *Node
}

func constructList() *Node {
	var hummpy = &Node{}
	var cur = hummpy
	for i := 0; i < cols*rows; i++ {
		node := Node{
			val: i,
		}
		cur.next = &node
		cur = cur.next
	}
	return hummpy.next
}

func init() {
	// construct matrix
	matrix1 = constructMatrix()
	matrix2 = constructMatrix()
	matrix3 = constructMatrix()
	list1 = constructList()
}

func BenchmarkCPUCache(b *testing.B) {
	var tmp int
	b.Run("horizontal travel", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					tmp = matrix1[i][j]
				}
			}
		}
	})

	b.Run("vertical travel", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for j := 0; j < cols; j++ {
				for i := 0; i < rows; i++ {
					//if matrix2[i][j] == 666 {
					//	tmp++
					//}
					tmp = matrix2[i][j]
				}
			}
		}
	})

	b.Run("random travel", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for i := 0; i < cols*rows; i++ {
				r := rand.Intn(rows)
				//if matrix3[r][r] == 66 {
				//	tmp++
				//}
				tmp = matrix3[r][r]
			}
		}
	})

	b.Run("linked list", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for node := list1; node != nil; node = node.next {
				//if node.val == 666 {
				//	tmp++
				//}
				tmp = node.val
			}
		}
	})

	_ = tmp
}
