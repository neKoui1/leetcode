package main

import (
	"fmt"
	"math"
)

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		usedShelf := 0
		maxHeight := 0
		for j := i; j > 0; j-- {
			usedShelf += books[j-1][0]
			if usedShelf > shelfWidth {
				break
			}
			maxHeight = Max(maxHeight, books[j-1][1])
			dp[i] = Min(dp[i], dp[j-1]+maxHeight)
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(minHeightShelves([][]int{
		{1, 1},
		{2, 3},
		{2, 3},
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 2},
	}, 4))
}
