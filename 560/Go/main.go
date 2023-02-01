package main

import "fmt"

func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums))
	cnt := 0
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	for _, v := range sum {
		if v == k {
			cnt++
		}
	}
	for i := 0; i < len(sum)-1; i++ {
		for j := i + 1; j < len(sum); j++ {
			if sum[j]-sum[i] == k {
				cnt++
			}
		}
	}
	return cnt
}
func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
