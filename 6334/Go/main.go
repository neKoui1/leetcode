package main

import (
	"fmt"
	"math"
)

func findPrefixScore(nums []int) (ans []int64) {
	n := len(nums)
	conver := make([]int, 0)
	converSum := 0
	max := math.MinInt
	for i := 0; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		conver = append(conver, nums[i]+max)
		converSum += nums[i] + max
		ans = append(ans, int64(converSum))
	}

	return
}
func main() {
	fmt.Println(findPrefixScore([]int{2, 3, 7, 5, 10}))
	fmt.Println(findPrefixScore([]int{1, 1, 2, 4, 8, 16}))
}
