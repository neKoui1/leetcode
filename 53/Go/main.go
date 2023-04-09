package main

import (
	"fmt"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSubArray(nums []int) (res int) {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res = dp[0]
	for i := 1; i < n; i++ {
		dp[i] = Max(nums[i], dp[i-1]+nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return
}
func main() {
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
