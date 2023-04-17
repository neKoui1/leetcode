package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return Max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = Max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}
func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 1, 1}))
}
