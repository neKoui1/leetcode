package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	isOdd := make([]int, 0)
	NOdd := make([]int, 0)
	for _, v := range nums {
		if v&1 == 1 {
			isOdd = append(isOdd, v)
		} else {
			NOdd = append(NOdd, v)
		}
	}
	return append(NOdd, isOdd...)
}
func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}
