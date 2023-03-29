package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sentinal := new(ListNode)
	sentinal.Next = head
	cur := sentinal
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return sentinal.Next
}
func main() {

}
