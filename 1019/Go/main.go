package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Pair struct {
	Index int
	Val   int
}

func nextLargerNodes(head *ListNode) []int {
	monoStack := make([]Pair, 0)
	temp := head
	n := 0
	idx := 0
	for temp != nil {
		n++
		temp = temp.Next
	}
	res := make([]int, n)
	temp = head
	for temp != nil {

		for len(monoStack) != 0 && monoStack[len(monoStack)-1].Val < temp.Val {
			x := monoStack[len(monoStack)-1]
			monoStack = monoStack[:len(monoStack)-1]
			res[x.Index] = temp.Val
		}
		pair := Pair{
			Index: idx,
			Val:   temp.Val,
		}
		monoStack = append(monoStack, pair)
		idx++
		temp = temp.Next
	}
	return res
}
func main() {
	head := new(ListNode)
	node1 := new(ListNode)
	node2 := new(ListNode)
	head.Val = 2
	head.Next = node1
	node1.Val = 1
	node1.Next = node2
	node2.Val = 5
	node2.Next = nil
	fmt.Println(nextLargerNodes(head))
}
