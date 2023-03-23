package main

import "fmt"

func threeSum(nums []int) (res [][]int) {
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{{nums[0], nums[1], nums[2]}}
		} else {
			return
		}
	}
	return
}
func main() {
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
