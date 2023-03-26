package main

import "fmt"

func combine(n int, k int) (res [][]int) {

	path := make([]int, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return
}
func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
}
