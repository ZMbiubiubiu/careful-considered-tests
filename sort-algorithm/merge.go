package main

import "fmt"

func mergeSort(arr []int) {
	var temp = make([]int, len(arr))
	mergeSort0(temp, arr, 0, len(arr)-1)
}

func mergeSort0(temp, arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2

	mergeSort0(temp, arr, low, mid)
	mergeSort0(temp, arr, mid+1, high)

	merge(temp, arr, low, mid, high)
}

func merge(temp, arr []int, low, mid, high int) {
	if low >= high {
		return
	}
	// 复制一份
	for i := low; i <= high; i++ {
		temp[i] = arr[i]
	}

	// 排序两个子数组
	var left, right = low, mid + 1
	for pos := low; pos <= high; pos++ {
		// 左面已经完成
		if left == mid+1 {
			arr[pos] = temp[right]
			right++
			// 右面已经完成
		} else if right == high+1 {
			arr[pos] = temp[left]
			left++
		} else if temp[left] <= temp[right] {
			arr[pos] = temp[left]
			left++
		} else {
			arr[pos] = temp[right]
			right++
		}
	}
}

func main() {
	arr := []int{1, 4, 6, 2, 3, 9, 8, 5, 7}
	mergeSort(arr)
	fmt.Println("merge sort:", arr)
}
