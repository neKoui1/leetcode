package main

import "fmt"

func BinarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) >> 1
		if nums[middle] == target {
			return true
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		if target >= matrix[i][0] && target < matrix[i+1][0] {
			return BinarySearch(matrix[i], target)
		}
	}
	return BinarySearch(matrix[len(matrix)-1], target)
}
func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 13))
}
