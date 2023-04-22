package main

import (
	"fmt"
	"sort"
)

// time limit exceed
// N^2
// func Max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// func maxEnvelopes(envelopes [][]int) int {
// 	ans := 0
// 	n := len(envelopes)
// 	if n == 0 {
// 		return 0
// 	} else if n == 1 {
// 		return 1
// 	}
// 	sort.SliceStable(envelopes, func(i, j int) bool {
// 		if envelopes[i][0] != envelopes[j][0] {
// 			return envelopes[i][0] < envelopes[j][0]
// 		}
// 		return envelopes[i][1] > envelopes[j][1]
// 	})
// 	dp := make([]int, n)
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			if envelopes[j][1] < envelopes[i][1] {
// 				dp[i] = Max(dp[i], dp[j]+1)
// 			}
// 		}
// 		ans = Max(ans, dp[i])
// 	}
// 	return ans
// }

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	sort.SliceStable(envelopes, func(i, j int) bool {
		a := envelopes[i]
		b := envelopes[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})
	// 声明一个stored切片存放已经成立的严格递增子序列
	stored := make([]int, 0)
	for i := 0; i < n; i++ {
		height := envelopes[i][1]
		if idx := sort.SearchInts(stored, height); idx < len(stored) {
			stored[idx] = height
		} else {
			stored = append(stored, height)
		}
	}
	return len(stored)
}

func main() {
	fmt.Println(maxEnvelopes([][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}))
	fmt.Println(maxEnvelopes([][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}))
	fmt.Println(maxEnvelopes([][]int{
		{1, 1},
		{1, 1},
	}))
}
