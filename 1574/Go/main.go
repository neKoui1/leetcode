package main

import "fmt"

func Min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func findLengthOfShortestSubarray(arr []int) (res int) {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	res = right
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		res = Min(res, right-left-1)
	}
	return
}
func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 4, 10, 4, 2, 3, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))
	fmt.Println(findLengthOfShortestSubarray([]int{1}))
	fmt.Println(findLengthOfShortestSubarray([]int{2, 2, 2, 1, 1, 1}))
}
