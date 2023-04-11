package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)
	res := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		res[i] = -1
	}
	monoStack := make([]int, 0)
	for i := 0; i < 2*n; i++ {

		for len(monoStack) != 0 && nums[monoStack[len(monoStack)-1]] < nums[i] {
			x := monoStack[len(monoStack)-1]
			monoStack = monoStack[:len(monoStack)-1]
			res[x] = nums[i]
		}
		monoStack = append(monoStack, i)
	}
	return res[:n]
}
func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
