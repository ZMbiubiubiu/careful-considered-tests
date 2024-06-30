package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, low, high int) {
	if low >= high {
		return
	}

	p := partition(arr, low, high)
	quick(arr, low, p-1)
	quick(arr, p+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low] // low位置为空, 循环结束 pivot就到了中间位置（当然，最好是中间位置哈，这样分的才平均）
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high] // 小的去了左边

		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low] // 大的去了右边
	}
	// 此时 low == high
	arr[low] = pivot
	return low
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 9, 8, 5, 7}
	quickSort(arr)
	fmt.Println(arr)
}
