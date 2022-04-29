package main

import "fmt"

func judge(n int) bool {
	store := n
	for store != 0 {
		j := store % 10
		if j == 0 || n%j != 0 {
			return false
		}
		store /= 10
	}
	return true
}
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if judge(i) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(selfDividingNumbers(47, 85))
}
