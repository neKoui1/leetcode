package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p, q := dummy.Next, dummy
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p != nil {
		p, q = p.Next, q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}
func main() {

}
