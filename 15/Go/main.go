package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (res [][]int) {
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{{nums[0], nums[1], nums[2]}}
		} else {
			return
		}
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, v := range nums {
		if v > 0 {
			return
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if v+nums[left]+nums[right] == 0 {
				res = append(res, []int{v, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if v+nums[left]+nums[right] > 0 {
				right--
			} else if v+nums[left]+nums[right] < 0 {
				left++
			}
		}
	}
	return
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
