package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Val = 0
	sentinel.Next = head

	p := sentinel
	q := sentinel.Next
	for q != nil {
		if q.Val == val {
			p.Next = q.Next
		}
		q = q.Next
		p = p.Next
	}

	return sentinel.Next
}
func main() {

}
