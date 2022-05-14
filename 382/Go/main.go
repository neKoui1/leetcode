package main

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Nums []int
}

func Constructor(head *ListNode) Solution {
	p := head
	nums := make([]int, 0)
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}
	return Solution{nums}
}

func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return this.Nums[rand.Intn(len(this.Nums))]
}
func main() {

}
