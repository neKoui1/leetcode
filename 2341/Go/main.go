package main

import "fmt"

func numberOfPairs(nums []int) []int {
	cnt, left := 0, 0
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
		if hash[v]%2 == 0 {
			cnt++
			hash[v] = 0
		}
	}
	for _, v := range hash {
		if v != 0 {
			left++
		}
	}
	return []int{cnt, left}

}
func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}))
	fmt.Println(numberOfPairs([]int{1, 1}))
	fmt.Println(numberOfPairs([]int{0}))
}
