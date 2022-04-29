package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		a := nums[i]
		b := nums[i-1]
		c := nums[i-2]
		if b+c > a {
			return a + b + c
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
}
