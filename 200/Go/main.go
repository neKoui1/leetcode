package main

import "fmt"

/*
将所有连在一起的1都置为2
*/
func Spread(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	Spread(grid, i, j+1)
	Spread(grid, i, j-1)
	Spread(grid, i+1, j)
	Spread(grid, i-1, j)
}

func numIslands(grid [][]byte) (islandcnt int) {
	for i, v := range grid {
		for j, land := range v {
			if land == '1' {
				islandcnt++
				Spread(grid, i, j)
			}
		}
	}
	return
}
func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
