package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	num := make([]int, 0)
	for list1 != nil {
		num = append(num, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		num = append(num, list2.Val)
		list2 = list2.Next
	}
	sort.Ints(num)
	root := new(ListNode)
	if len(num) == 0 {
		return nil
	}
	root.Val = num[0]
	p := root
	for i := 1; i < len(num); i++ {
		q := new(ListNode)
		q.Val = num[i]
		p.Next = q
		p = q
	}
	return root
}
func main() {
	nums := make([]int, 0)
	fmt.Println(len(nums))
	fmt.Println("hdwuioahforje")
}
