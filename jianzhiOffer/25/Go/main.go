package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//if l1 == nil && l2 == nil {
	//	return nil
	//}
	//if l1 == nil && l2 != nil {
	//	return l2
	//} else if l1 != nil && l2 == nil {
	//	return l1
	//}
	sentinel := new(ListNode)
	pre := sentinel
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			pre = pre.Next
			l2 = l2.Next
		} else {
			pre.Next = l1
			pre = pre.Next
			l1 = l1.Next
		}
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}
	return sentinel.Next
}
func main() {

}
