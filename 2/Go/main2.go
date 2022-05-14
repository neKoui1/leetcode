package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Triple(condition bool, a, b int) int {
	if condition {
		return a
	}
	return b
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	OF := 0
	var p *ListNode
	var q *ListNode
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + OF
		OF = sum / 10
		if p == nil {
			p = &ListNode{sum % 10, nil}
			q = p
		} else {
			q.Next = &ListNode{sum % 10, nil}
			q = q.Next
		}
	}
	if OF > 0 {
		q.Next = &ListNode{OF, nil}
	}
	return p
}
func main() {
	var p *ListNode
	fmt.Println(p)
	fmt.Println(&p)
}
