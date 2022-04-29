package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pow(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 2
	}
	return res
}
func getDecimalValue(head *ListNode) int {
	bin := make([]int, 0)
	for head != nil {
		bin = append(bin, head.Val)
		head = head.Next
	}
	res := 0
	m := len(bin)
	for i := 0; i < m; i++ {
		res += bin[i] * pow(m-i-1)
	}
	return res
}

func main() {
	head := new(ListNode)
	head.Val = 1
	p2 := new(ListNode)
	p3 := new(ListNode)
	head.Next = p2
	p2.Next = p3
	p2.Val = 0
	p3.Val = 1
	p3.Next = nil
	fmt.Println(getDecimalValue(head))
}
