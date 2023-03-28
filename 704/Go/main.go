package main

import "fmt"

/**
typical binary search
*/
func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left <= right {
		middle := (left + right) >> 1
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
