package main

import (
	"fmt"
	"math"
)

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 { // 无法合并成一堆
		return -1
	}

	s := make([]int, n+1)
	for i, x := range stones {
		s[i+1] = s[i] + x // 前缀和
	}

	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for p := range memo[i][j] {
				memo[i][j][p] = -1 // -1 表示还没有计算过
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, p int) (res int) {
		ptr := &memo[i][j][p]
		if *ptr != -1 {
			return *ptr
		}
		defer func() { *ptr = res }()
		if p == 1 {
			if i == j { // 只有一堆石头，无需合并
				return
			}
			return dfs(i, j, k) + s[j+1] - s[i]
		}
		res = math.MaxInt
		for m := i; m < j; m += k - 1 { // 枚举哪些石头堆合并成第一堆
			res = min(res, dfs(i, m, 1)+dfs(m+1, j, p-1))
		}
		return
	}
	return dfs(0, n-1, 1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func main() {
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 2))
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 3))
	fmt.Println(mergeStones([]int{3, 5, 1, 2, 6}, 3))
}
