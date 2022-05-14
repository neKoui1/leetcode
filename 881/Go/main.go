package main

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) (cnt int) {
	sort.Ints(people)
	p := 0
	q := len(people) - 1
	for p <= q {
		if people[p]+people[q] <= limit {
			p++
		}
		q--
		cnt++
	}
	return
}
func main() {
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}
