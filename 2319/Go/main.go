package main

import "fmt"

func checkXMatrix(grid [][]int) bool {
	length := len(grid)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j || i+j == length-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}
func main() {
	fmt.Println(checkXMatrix([][]int{
		{2, 0, 0, 1},
		{0, 3, 1, 0},
		{0, 5, 2, 0},
		{4, 0, 0, 2},
	}))
	fmt.Println(checkXMatrix([][]int{
		{5, 7, 0},
		{0, 3, 1},
		{0, 5, 0},
	}))
}
