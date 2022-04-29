package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	a []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	rand.Seed(time.Now().UnixNano())
	hash := make(map[int][]int)
	for i, v := range this.a {
		hash[v] = append(hash[v], i)
	}
	return hash[target][rand.Intn(len(hash[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
func main() {
	s := Solution{
		[]int{1, 2, 3, 3, 3},
	}
	fmt.Println(s.Pick(1))
	fmt.Println(s.Pick(2))
	fmt.Println(s.Pick(3))
}
