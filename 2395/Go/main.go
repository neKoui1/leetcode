package main

import "fmt"

func findSubarrays(nums []int) bool {
	n := len(nums)
	if n == 2 {
		return false
	}
	hash := make(map[int]int)
	for i := 1; i < n; i++ {
		hash[nums[i]+nums[i-1]]++
	}
	for _, v := range hash {
		if v > 1 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 4, 5}))
	fmt.Println(findSubarrays([]int{0, 0, 0}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 1, 0, 2}))
}
