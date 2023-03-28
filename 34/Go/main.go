package main

import (
	"fmt"
	"reflect"
)

func BinarySearch(nums []int, target int, isLeft bool) int {
	n := len(nums)
	res := -1
	left, right := 0, n-1
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

func searchRange(nums []int, target int) []int {
	if reflect.DeepEqual(nums, []int{}) {
		return []int{-1, -1}
	}

	return []int{BinarySearch(nums, target, true), BinarySearch(nums, target, false)}
}
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}
