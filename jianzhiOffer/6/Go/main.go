package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) (res []int) {
	for p := head; p != nil; p = p.Next {
		res = append(res, p.Val)
	}
	n := len(res)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return
}
func main() {

}
