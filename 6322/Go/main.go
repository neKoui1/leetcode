package main

import (
	"fmt"
	"math"
)

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	triCoor := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			triCoor[grid[i][j]] = append(triCoor[grid[i][j]], i)
			triCoor[grid[i][j]] = append(triCoor[grid[i][j]], j)
		}
	}
	for i := 0; i <= n*n-2; i++ {
		if math.Pow(float64(triCoor[i][0]-triCoor[i+1][0]), 2)+math.Pow(float64(triCoor[i][1]-triCoor[i+1][1]), 2) != 5 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkValidGrid([][]int{
		{0, 11, 16, 5, 20},
		{17, 4, 19, 10, 15},
		{12, 1, 8, 21, 6},
		{3, 18, 23, 14, 9},
		{24, 13, 2, 7, 22},
	}))
	fmt.Println(checkValidGrid([][]int{
		{0, 3, 6},
		{5, 8, 1},
		{2, 7, 4},
	}))
}
