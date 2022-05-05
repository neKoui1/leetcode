package main

import (
	"fmt"
	"sort"
)

func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	if max-min <= 2*k {
		return 0
	}
	return max - min - 2*k
}

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
}
