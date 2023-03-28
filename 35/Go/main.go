package main

import "fmt"

//func searchInsert(nums []int, target int) int {
//	for i, v := range nums {
//		if v >= target {
//			return i
//		}
//	}
//	return len(nums)
//}

/**
binary search
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
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
	return right + 1
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
}
