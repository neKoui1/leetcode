package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	sum := 0
	for i, v := range nums {
		sum += v
		if sum*2-v == totalSum {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
