package main

import "fmt"

/**
NO REPEAT
CAN BE CHOSEN UNLIMITEDLY
*/
func combinationSum(candidates []int, target int) (res [][]int) {

	n := len(candidates)
	path := make([]int, 0)
	sum := 0

	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}

		for i := startIndex; i < n; i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backtracking(i) // significance, not i+1
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtracking(0)

	return
}
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
