package main

import (
	"fmt"
)


func binSearch( arr []int, target int) int {
	low  := 0
	high := len(arr) - 1;
	mid := 0
	for low < high {
		mid = low + (high - mid) / 2;
		if (arr[mid] == target){
			return mid
		} else if (arr[mid] < target) {
			low = mid + 1
		} else if (arr[mid] > target){
			high = mid - 1
		}
	}
	return -1;
}

func main() {
	items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}

	result := binSearch(items, 31);
	fmt.Println("get result ", result)

}