package main

import (
	"fmt"
	"math"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		max := math.MinInt
		for j := i - 1; j >= Max(0, i-k); j-- {
			max = Max(max, arr[j])
			dp[i] = Max(dp[i], dp[j]+max*(i-j))
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
	fmt.Println(maxSumAfterPartitioning([]int{1}, 1))
}
