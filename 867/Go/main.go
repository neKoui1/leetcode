package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = matrix[j][i]
		}
	}

	return ans
}
func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
