package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}
func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
