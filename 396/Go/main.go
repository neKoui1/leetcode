package main

import "fmt"

func maxRotateFunction(nums []int) int {
	sum := 0
	f := 0
	n := len(nums)
	for i, v := range nums {
		sum += v
		f += i * v
	}
	res := f
	for i := 1; i < n; i++ {
		f = f + sum - n*nums[n-i]
		if f > res {
			res = f
		}
	}
	return res
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(maxRotateFunction(nums))
}
