package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 2 {
		return nums[0]
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		flag := 1
		max := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[i-1] {
				flag++
			}
			if flag == 3 {
				max = nums[i]
				break
			}
			if flag < 3 && i == len(nums)-1 {
				max = nums[i]
				break
			}
		}
		return max
	}
}

func main() {
	fmt.Println(thirdMax([]int{1, 1, 2}))
}
