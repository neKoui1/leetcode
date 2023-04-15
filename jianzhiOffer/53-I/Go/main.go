package main

import "fmt"

func BinarySearch(nums []int, target int, isLeft bool) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		middle := (left + right) >> 1
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			res = middle
			if isLeft {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return res
}

func search(nums []int, target int) (res int) {
	left := BinarySearch(nums, target, true)
	if left == -1 {
		return 0
	}
	right := BinarySearch(nums, target, false)
	if left == right {
		return 1
	} else {
		return right - left + 1
	}
}
func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}
