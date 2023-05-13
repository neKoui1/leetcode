package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	start, longest := 0, 1
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i < n-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			longest = 2
		}
	}

	for l := 3; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				longest = l
			}
		}
	}
	return s[start : start+longest]
}
func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
