package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestArithSeqLength(nums []int) (ans int) {
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 1001)
	}
	for i := 1; i < n; i++ {
		for k := 0; k < i; k++ {
			j := nums[i] - nums[k] + 500
			f[i][j] = Max(f[i][j], f[k][j]+1)
			ans = Max(ans, f[i][j])
		}
	}
	return ans + 1
}
func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10}))
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8}))
}
