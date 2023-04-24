package main

import "fmt"

func exchange(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		for left < right && nums[left]&1 == 1 {
			left++
		}
		for left < right && nums[right]&1 == 0 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
