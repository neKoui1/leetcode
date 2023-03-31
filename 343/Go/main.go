package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = Max(dp[i], Max((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}
