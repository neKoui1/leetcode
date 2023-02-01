package main

import "fmt"

func productExceptSelf(nums []int) (res []int) {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	tempL := 1
	tempR := 1

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = 1
			continue
		}
		left[i] = tempL * nums[i-1]
		tempL = left[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = 1
			continue
		}
		right[i] = tempR * nums[i+1]
		tempR = right[i]

	}
	for i := 0; i < len(nums); i++ {
		res = append(res, left[i]*right[i])
	}
	return
}
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
