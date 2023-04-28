package main

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	res := make([]int, n)
	res[0] = nums[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] + nums[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			res = append(res, res[i]-res[j])
		}
	}
	sum := 0
	sort.Ints(res)
	for i := left; i < right; i++ {
		sum += res[i]
	}
	return sum % (1e9 + 7)
}
func main() {

}
