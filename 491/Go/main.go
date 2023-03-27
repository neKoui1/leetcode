package main

import "fmt"

/**
去重没处理好
*/
func findSubsequences(nums []int) (res [][]int) {

	n := len(nums)
	path := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {

		if len(path) >= 2 {
			res = append(res, append([]int{}, path...))
		}

		used := make(map[int]bool, n)
		for i := start; i < n; i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return
}
func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
	fmt.Println(findSubsequences([]int{1, 2, 3, 1, 1, 1}))
	fmt.Println(findSubsequences([]int{1, 1, 1, 1}))
}
