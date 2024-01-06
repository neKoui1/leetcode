package main

import (
	"fmt"
	"math"
	"reflect"
)

func findMax(row []int) (index, value int) {
	val := math.MinInt
	idx := -1
	for i, v := range row {
		if v > val {
			val = v
			idx = i
		}
	}

	return idx, val
}

func deleteGreatestValue(grid [][]int) int {
	test := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		test = append(test, []int{})
	}
	if reflect.DeepEqual(grid, test) {
		return 0
	}
	ans := 0
	max := math.MinInt
	for i, v := range grid {
		idx, val := findMax(v)
		if val > max {
			max = val
		}
		grid[i] = append(grid[i][:idx], grid[i][idx+1:]...)
	}
	ans += max
	return ans + deleteGreatestValue(grid)
}

func main() {
	fmt.Println(deleteGreatestValue([][]int{
		{1, 2, 4},
		{3, 3, 1},
	}))

	fmt.Println(deleteGreatestValue([][]int{
		{10},
	}))
}
