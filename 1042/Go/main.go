package main

import "fmt"

func gardenNoAdj(n int, paths [][]int) []int {
	g := make([][]int, n)
	for _, v := range paths {
		x := v[0] - 1
		y := v[1] - 1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	color := make([]int, n)
	for i, nodes := range g {
		used := [5]bool{}
		for _, j := range nodes {
			used[color[j]] = true
		}
		for color[i]++; used[color[i]]; color[i]++ {

		}
	}
	return color
}
func main() {
	fmt.Println(gardenNoAdj(3, [][]int{
		{1, 2},
		{2, 3},
		{3, 1},
	}))

	fmt.Println(gardenNoAdj(4, [][]int{
		{1, 2},
		{3, 4},
	}))
}
