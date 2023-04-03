package main

import "fmt"

func prevPermOpt1(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			for j := n - 1; j > i-1; j-- {
				if arr[j] < arr[i-1] && arr[j-1] != arr[j] {
					arr[i-1], arr[j] = arr[j], arr[i-1]
					return arr
				}
			}
		}
	}
	return arr
}
func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7}))
}
