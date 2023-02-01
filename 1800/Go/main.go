package main

import "fmt"

func maxAscendingSum(nums []int) int {
	sum, max := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= nums[j-1] {
				break
			}
			sum += nums[j]
		}
		if sum > max {
			max = sum
		}
		sum = 0
	}
	return max
}
func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 40, 50}))
	fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
}
