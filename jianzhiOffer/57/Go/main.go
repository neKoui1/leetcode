package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {
	for _, v := range nums {
		if idx := BinarySearch(nums, target-v); idx != -1 {
			return []int{v, nums[idx]}
		}
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{10, 26, 30, 31, 47, 60}, 40))
}
