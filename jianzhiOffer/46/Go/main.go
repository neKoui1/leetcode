package main

import (
	"fmt"
)

func translateNum(num int) int {

	str := fmt.Sprintf("%d", num)
	n := len(str)
	// 令下标i表示到第i个数，dp[i]表示到第i个数有多少种翻译方法
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x := (str[i-2]-'0')*10 + (str[i-1] - '0')
		if x >= 10 && x <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(translateNum(12258))
}
