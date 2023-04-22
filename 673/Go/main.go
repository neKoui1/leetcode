package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func findNumberOfLIS(nums []int) (ans int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	dp := make([]int, n)
	cnts := make([]int, n)
	maxLength := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnts[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnts[i] = cnts[j]
				} else if dp[i] == dp[j]+1 {
					cnts[i] += cnts[j]
				}
			}
		}
		maxLength = Max(dp[i], maxLength)
	}
	for i, v := range dp {
		if v == maxLength {
			ans += cnts[i]
		}
	}
	return
}
func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
}
