package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	flag := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			flag = true
			break
		}
	}

	if flag {
		slow = head
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}
		return slow
	}
	return nil
}
func main() {

}
