package main

import (
	"fmt"
	"math"
)

func findColumnWidth(grid [][]int) (ans []int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		max := math.MinInt
		for j := 0; j < m; j++ {
			str := fmt.Sprintf("%d", grid[j][i])
			x := len(str)
			if x > max {
				max = x
			}
		}
		ans = append(ans, max)
	}
	return
}
func main() {
	fmt.Println(findColumnWidth([][]int{
		{1},
		{22},
		{333},
	}))

	fmt.Println(findColumnWidth([][]int{
		{-15, 1, 3},
		{15, 7, 12},
		{5, 6, -2},
	}))
}
