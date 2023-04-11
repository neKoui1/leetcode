package main

import (
	"fmt"
	"strings"
)

func totalNQueens(n int) int {
	res := make([][]string, 0)
	diag1 := make([]bool, 2*n-1)
	diag2 := make([]bool, 2*n-1)
	col := make([]int, n)
	path := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			board := make([]string, n)
			for i, c := range col {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			res = append(res, board)
			return
		}
		for c, on := range path {
			rc := cur - c + n - 1
			if !on && !diag1[cur+c] && !diag2[rc] {
				col[cur] = c
				path[c], diag1[cur+c], diag2[rc] = true, true, true
				dfs(cur + 1)
				path[c], diag1[cur+c], diag2[rc] = false, false, false
			}
		}
	}
	dfs(0)
	return len(res)
}
func main() {
	fmt.Println(totalNQueens(4))
}
