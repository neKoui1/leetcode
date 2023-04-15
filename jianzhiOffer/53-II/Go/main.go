package main

import "fmt"

//func BinarySearch(nums []int, target int, isLeft bool) int {
//	left, right := 0, len(nums)-1
//	res := -1
//	for left <= right {
//		middle := (left + right) >> 1
//		if nums[middle] > target {
//			right = middle - 1
//		} else if nums[middle] < target {
//			left = middle + 1
//		} else {
//			res = middle
//			if isLeft {
//				right = middle - 1
//			} else {
//				left = middle + 1
//			}
//		}
//	}
//	return res
//}

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) >> 1
		if nums[middle] == middle {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left

	//sum := 0
	//for _, v := range nums {
	//	sum += v
	//}
	//SUM := (len(nums) * (len(nums) - 1)) / 2
	//return SUM - sum
}
func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}
