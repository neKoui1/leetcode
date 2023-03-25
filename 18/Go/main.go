package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) (res [][]int) {
	n := len(nums)
	if n < 4 {
		return [][]int{}
	}

	if n == 4 {
		if cur := nums[0] + nums[1] + nums[2] + nums[3]; cur == target {
			return [][]int{{nums[0], nums[1], nums[2], nums[3]}}
		}
	}

	sort.Ints(nums)

	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if (i > 0 && nums[i] == nums[i-1]) || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if (j > i+1 && nums[j] == nums[j-1]) || nums[j]+nums[n-3]+nums[n-2]+nums[n-1] < target {
				continue
			}

			for left, right := j+1, n-1; left < right; {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return
}
func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
