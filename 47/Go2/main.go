package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) (res [][]int) {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	path := make([]int, 0)
	used := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				dfs(i + 1)
				used[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
