package main

import (
	"fmt"
)

func threeSum(nums []int, target int) []int {
	m := make(map[int]int)
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			twosum := nums[i] + nums[j]
			if idx, ok := m[target-twosum]; ok {
				if idx != i && idx != j {
					return []int{i, j, idx}
				}
			}
			m[nums[j]] = j
		}
		m[nums[i]] = i
	}
	return nil
}



func main() {

	nums := []int{1, 2, 9, 20, 23, 31, 45, 63, 70, 100}

	result := threeSum(nums, 12)

	fmt.Println("result: ", result)
}
