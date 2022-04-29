package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	flag := true
	sub := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if sub != arr[i+1]-arr[i] {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
