package main

import "fmt"

func majorityElement(nums []int) (res int) {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}
	max := 0
	for i, v := range hash {
		if v > max {
			res = i
			max = v
		}
	}
	return
}
func main() {
	n1 := []int{3, 2, 3}
	fmt.Println(majorityElement(n1))
	n2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(n2))
}
