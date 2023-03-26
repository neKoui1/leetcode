package main

import "fmt"

func combinationSum3(k int, n int) (res [][]int) {

	path := make([]int, 0)
	sum := 0
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == k {
			if sum == n {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		if sum > n {
			return
		}

		for i := start; i <= n && i < 10; i++ {
			path = append(path, i)
			sum += i
			backtracking(i + 1)
			path = path[:len(path)-1]
			sum -= i
		}
	}
	backtracking(1)

	return
}
func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
	fmt.Println(combinationSum3(2, 18))
}
