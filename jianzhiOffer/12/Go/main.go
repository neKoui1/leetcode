package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var findPath func(i, j, k int) bool
	findPath = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		//defer func() {
		//	vis[i][j] = false
		//}()
		for g := 0; g < 4; g++ {
			newX := i + dx[g]
			newY := j + dy[g]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && !vis[newX][newY] {
				if findPath(newX, newY, k+1) {
					return true
				}
			}
		}
		vis[i][j] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findPath(i, j, 0) {
				return true
			}
		}
	}

	return false
}
func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	fmt.Println(exist([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "abcd"))
}
