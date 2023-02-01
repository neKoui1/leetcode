package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	temp := sum
	for index := 1; index <= len(nums)-k; index++ {
		temp = temp - nums[index-1] + nums[index+k-1]
		if temp > sum {
			sum = temp
		}
	}
	return float64(sum) / float64(k)
}
func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
	fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}
