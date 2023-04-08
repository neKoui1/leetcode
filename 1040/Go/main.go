package main

import (
	"fmt"
	"sort"
)

func numMovesStonesII(s []int) []int {
	sort.Ints(s)
	n := len(s)
	e1 := s[n-2] - s[0] - n + 2
	e2 := s[n-1] - s[1] - n + 2 // 计算空位
	maxMove := max(e1, e2)
	if e1 == 0 || e2 == 0 { // 特殊情况：没有空位
		return []int{min(2, maxMove), maxMove}
	}
	maxCnt, left := 0, 0
	for right, x := range s { // 滑动窗口：枚举右端点
		for s[left] <= x-n { // 窗口大小大于 n
			left++
		}
		maxCnt = max(maxCnt, right-left+1) // 维护窗口内的最大石子数
	}
	return []int{n - maxCnt, maxMove}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func main() {
	fmt.Println(numMovesStonesII([]int{7, 4, 9}))
	fmt.Println(numMovesStonesII([]int{6, 5, 4, 3, 10}))
	fmt.Println(numMovesStonesII([]int{100, 101, 104, 102, 103}))
}
