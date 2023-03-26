package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (res [][]int) {

	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	n := len(candidates)
	path := make([]int, 0)
	used := make([]bool, n)

	sum := 0
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if sum > target || (startIndex < n && used[startIndex]) {
			return
		}
		for i := startIndex; i < n && !used[i]; i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			used[i] = true
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(i + 1)
			sum -= candidates[i]
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking(0)
	/**
	deduplication
	*/
	//deduplication := make(map[string]struct{})
	//for i := 0; i < len(res); i++ {
	//	marshal, _ := json.Marshal(res[i])
	//	deduplication[string(marshal)] = struct{}{}
	//}
	//res = nil
	//for i := range deduplication {
	//	temp := make([]int, 0)
	//	json.Unmarshal([]byte(i), &temp)
	//	res = append(res, temp)
	//}
	return
}
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
