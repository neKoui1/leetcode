package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) (res int) {
	hash := make(map[int]int)
	for i, v := range nums {
		hash[v] = i
	}
	for _, v := range nums {
		if _, ok1 := hash[v+diff]; ok1 {
			if _, ok2 := hash[v+diff+diff]; ok2 {
				res++
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return
}
func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
}
