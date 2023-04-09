package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	n := len(s)
	check := make([]int, 26)
	for i := 0; i < 26; i++ {
		check[i] = -1
	}
	for i := 0; i < n; i++ {
		temp := int(s[i] - 'a')
		if check[temp] == -1 {
			check[temp] = i
		} else {
			if i-1-check[temp] != distance[temp] {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
