package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p := head
	q := head
	for q != nil {
		if q.Next != nil {
			q = q.Next.Next
		} else {
			break
		}
		p = p.Next
	}
	return p
}
