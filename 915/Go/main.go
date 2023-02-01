package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func partitionDisjoint(nums []int) int {
	leftMax, max, index := nums[0], nums[0], 0
	for i, v := range nums {
		max = Max(max, v)
		if v < leftMax {
			leftMax = max
			index = i
		}
	}
	return index + 1
}
func main() {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(partitionDisjoint([]int{32, 57, 24, 19, 0, 24, 49, 67, 87, 87}))
	fmt.Println(partitionDisjoint([]int{1, 1}))
}
