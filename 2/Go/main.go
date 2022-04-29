package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := make([]int, 0)
	temp, OF := 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp = (l2.Val + OF) % 10
			OF = (l2.Val + OF) / 10
		} else if l2 == nil {
			temp = (l1.Val + OF) % 10
			OF = (l1.Val + OF) / 10
		} else {
			temp = (l1.Val + l2.Val + OF) % 10
			OF = (l1.Val + l2.Val + OF) / 10
		}
		res = append(res, temp)

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if OF > 0 {
		res = append(res, OF)
	}

	root := new(ListNode)
	root.Val = res[0]
	p := root
	for i := 1; i < len(res); i++ {
		q := new(ListNode)
		q.Val = res[i]
		p.Next = q
		p = q
	}

	return root
}

func main() {

}
