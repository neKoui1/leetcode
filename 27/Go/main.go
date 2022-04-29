package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	cnt := 0
	for _, v := range nums {
		if v == val {
			cnt++
		}
	}
	sort.Ints(nums)
	res := len(nums)
	for i, v := range nums {
		if v == val {
			nums = append(nums[:i], nums[i+cnt:]...)
			break
		}
	}
	return res - cnt
}
func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}
