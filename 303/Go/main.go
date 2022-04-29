package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}
func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("na.SumRange(0, 2): %v\n", na.SumRange(0, 2))
	fmt.Printf("na.SumRange(2, 5): %v\n", na.SumRange(2, 5))
	fmt.Printf("na.SumRange(0, 5): %v\n", na.SumRange(0, 5))
}
