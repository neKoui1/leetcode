package main

import "fmt"

func permute(nums []int) (res [][]int) {
	n := len(nums)
	path := make([]int, 0)
	used := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, v := range nums {
			if !used[i] {
				used[i] = true
				path = append(path, v)
				dfs(i + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
