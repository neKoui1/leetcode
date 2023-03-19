package main

import (
	"fmt"
)

func Abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func helper(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 1
	} else if n == 2 {
		cnt := 0
		if Abs(nums[0], nums[1]) != k {
			cnt++
		}
		return cnt
	} else if n == 3 {
		cnt := 0
		if (Abs(nums[0], nums[1]) != k) && (Abs(nums[0], nums[2]) != k) && (Abs(nums[2], nums[1]) != k) {
			cnt++
		}
		return cnt
	} else {
		cnt := 0
		for i := 0; i < n-1; i++ {
			flag := true
			for j := i + 1; j < n; j++ {
				if Abs(nums[i], nums[j]) == k {
					flag = false
					break
				}
			}
			if flag {
				cnt++
			}
		}
		return cnt
	}
}

func beautifulSubsets(nums []int, k int) int {
	cnt := 0
	n := len(nums)
	for i := 0; i < n; i++ {

	}
	return cnt
}
func main() {
	fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2))
	fmt.Println(beautifulSubsets([]int{1}, 1))
	fmt.Println(beautifulSubsets([]int{4, 2, 5, 9, 10, 3}, 1))
}
