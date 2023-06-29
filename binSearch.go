package main

import "fmt"

// 找出下面代码的问题
func binSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0
	for low < high {
		mid = (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

// 实现一个快速排序的算法
func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}

	result := binSearch(items, 100)
	fmt.Println("get result ", result)

}
