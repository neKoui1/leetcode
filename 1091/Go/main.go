package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	if grid[0][0] == 1 {
		return -1
	}
	dir := [][]int{
		{1, 1},
		{1, 0},
		{1, -1},
		{0, 1},
		{0, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}
	queue := make([][]int, 0)
	step := 1
	grid[0][0] = 1
	queue = append(queue, []int{0, 0})
	for len(queue) != 0 {
		size := len(queue)
		for n := size; n > 0; n-- {
			x := queue[0][0]
			y := queue[0][1]
			queue = queue[1:]
			if x == len(grid)-1 && y == len(grid[0])-1 {
				return step
			}

			for _, v := range dir {
				tempX := x + v[0]
				tempY := y + v[1]
				if tempX < 0 || tempX > len(grid)-1 || tempY < 0 || tempY > len(grid[0])-1 || grid[tempX][tempY] == 1 {
					continue
				}
				queue = append(queue, []int{tempX, tempY})
				grid[tempX][tempY] = 1
			}
		}
		step++
	}
	return -1
}
func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1},
		{1, 0},
	}))
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
}
